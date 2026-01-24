package service

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/plugin"
	"go.uber.org/zap"
)

var ExtAiTask = new(extAiTask)

type extAiTask struct{}

type (
	blctyResp struct {
		Data []struct {
			URL     string `json:"url"`
			B64Json string `json:"b64_json"`
		} `json:"data"`
	}
)

type (
	cozeHistory struct {
		ConnectorUID      string `json:"connector_uid"`
		Logid             string `json:"logid"`
		ErrorCode         string `json:"error_code"`
		NodeExecuteStatus struct {
		} `json:"node_execute_status"`
		ErrorMessage  string `json:"error_message"`
		ExecuteID     string `json:"execute_id"`
		BotID         string `json:"bot_id"`
		ExecuteStatus string `json:"execute_status"`
		Output        string `json:"output"`
		Usage         struct {
			TokenCount  int `json:"token_count"`
			OutputCount int `json:"output_count"`
			InputCount  int `json:"input_count"`
		} `json:"usage"`
		CreateTime      int    `json:"create_time"`
		UpdateTime      int    `json:"update_time"`
		IsOutputTrimmed bool   `json:"is_output_trimmed"`
		DebugURL        string `json:"debug_url"`
		ConnectorID     string `json:"connector_id"`
		Token           string `json:"token"`
		RunMode         int    `json:"run_mode"`
	}

	cozeOutput struct {
		Output string `json:"Output"`
	}

	cozeTaskOutput struct {
		TaskId string `json:"task_Id"`
	}

	cozeTaskResult struct {
		Code     int    `json:"code"`
		Message  string `json:"message"`
		VideoURL string `json:"video_url"`
	}
)

func (s *extAiTask) generateImage(index int, taskId string, payload *request.BLCTYImagesRequest) (filename string, err error) {
	var (
		buf []byte
		req *http.Request
		res *http.Response
	)
	if buf, err = json.Marshal(payload); err != nil {
		return
	}
	global.GVA_LOG.Info("开始生成海报!", zap.String("taskId", taskId), zap.Int("index", index))
	apiURL := plugin.Config.BltcyApiUrl + "/v1/images/generations"
	if req, err = http.NewRequest("POST", apiURL, bytes.NewReader(buf)); err != nil {
		global.GVA_LOG.Error("创建请求失败!", zap.String("taskId", taskId), zap.Int("index", index), zap.Error(err))
		return
	}
	apiKey := plugin.Config.BltcyApiKey
	req.Header.Add("Authorization", "Bearer "+apiKey)
	req.Header.Add("Content-Type", "application/json")
	if res, err = http.DefaultClient.Do(req); err != nil {
		global.GVA_LOG.Error("请求生成图片失败!", zap.String("taskId", taskId), zap.Int("index", index), zap.Error(err))
		return
	}
	defer func() {
		res.Body.Close()
	}()
	if buf, err = io.ReadAll(res.Body); err != nil {
		global.GVA_LOG.Error("读取图片数据失败!", zap.String("taskId", taskId), zap.Int("index", index), zap.Error(err))
		return
	}
	var obj blctyResp
	if err = json.Unmarshal(buf, &obj); err != nil || len(obj.Data) == 0 {
		global.GVA_LOG.Error("读取返回结果失败!", zap.String("taskId", taskId), zap.Int("index", index), zap.Error(err))
		return
	}
	raw := obj.Data[0].B64Json
	var payloadB64 string
	ext := "jpg"
	if strings.HasPrefix(raw, "data:") {
		parts := strings.SplitN(raw, ",", 2)
		if len(parts) < 2 {
			return
		}
		h := parts[0]
		payloadB64 = parts[1]
		if strings.HasPrefix(h, "data:image/") {
			t := strings.TrimPrefix(h, "data:image/")
			if i := strings.Index(t, ";"); i >= 0 {
				t = t[:i]
			}
			switch strings.ToLower(t) {
			case "jpeg", "jpg":
				ext = "jpg"
			case "png":
				ext = "png"
			case "webp":
				ext = "webp"
			case "gif":
				ext = "gif"
			}
		}
	} else {
		payloadB64 = raw
	}
	if buf, err = base64.StdEncoding.DecodeString(payloadB64); err != nil || len(buf) == 0 {
		global.GVA_LOG.Warn("解析渲染海报数据错误", zap.Int("index", index), zap.String("taskId", taskId))
		return
	}
	filename = filepath.Join(global.GVA_CONFIG.Local.Path, fmt.Sprintf("blcty_%s_%d.%s", taskId, time.Now().UnixNano(), ext))
	if err = os.WriteFile(filename, buf, 0644); err != nil {
		global.GVA_LOG.Error("写入图片数据到文件失败!", zap.String("taskId", taskId), zap.Int("index", index), zap.Error(err))
	} else {
		global.GVA_LOG.Info("图片生成成功!", zap.String("taskId", taskId), zap.Int("index", index), zap.String("filename", filename))
	}
	return
}

func (s *extAiTask) ExecuteImageTask(taskId string) {
	var (
		err         error
		buf         []byte
		mutex       sync.Mutex
		filename    string
		isCompleted bool
		payload     *request.BLCTYImagesRequest
		results     []*response.ImgResult
	)
	payload = &request.BLCTYImagesRequest{}
	results = make([]*response.ImgResult, 0, 4)
	modelValue := &model.ExtAiTask{}
	global.GVA_LOG.Info("开始生成宣传海报!", zap.String("taskId", taskId))
	if err = global.GVA_DB.First(modelValue, "task_id = ?", taskId).Error; err != nil {
		global.GVA_LOG.Error("找不到任务数据!", zap.String("taskId", taskId), zap.Error(err))
		return
	}
	if err = json.Unmarshal([]byte(modelValue.Options), payload); err != nil {
		global.GVA_LOG.Error("配置参数无效!", zap.String("taskId", taskId), zap.Error(err))
		modelValue.Status = "failed"
		modelValue.Description = "无效的配置参数"
		global.GVA_DB.Save(modelValue)
		return
	}
	for i, img := range payload.Image {
		if !strings.HasPrefix(img, "http") {
			if buf, err = os.ReadFile(img); err == nil {
				payload.Image[i] = base64.StdEncoding.EncodeToString(buf)
			} else {
				global.GVA_LOG.Warn("读取本地模板海报信息失败!", zap.String("taskId", taskId), zap.Error(err))
			}
		}
	}
	defer func() {
		if r := recover(); r != nil {
			global.GVA_LOG.Error("生成宣传海报失败!", zap.String("taskId", taskId), zap.Error(fmt.Errorf("%v", r)))
		}
		if !isCompleted {
			modelValue.Status = "failed"
			modelValue.Description = "生成异常"
			global.GVA_DB.Save(modelValue)
		}
	}()
	var wg sync.WaitGroup
	imgBaseUrl := plugin.Config.HrefUrl
	for i := range 4 {
		wg.Add(1)
		go func(idx int) {
			stm := time.Now()
			result := &response.ImgResult{
				Index: idx,
			}
			filename, err = s.generateImage(idx, taskId, payload)
			result.Duration = time.Since(stm)
			if err == nil {
				result.Url = imgBaseUrl + filename
			} else {
				result.Error = err.Error()
			}
			mutex.Lock()
			results = append(results, result)
			if buf, err = json.Marshal(results); err == nil {
				modelValue.Result = string(buf)
				global.GVA_DB.Save(modelValue)
			}
			mutex.Unlock()
			wg.Done()
		}(i)
	}
	wg.Wait()
	modelValue.Status = "completed"
	modelValue.Description = "成功"
	modelValue.CompleteAt = time.Now().Unix()
	global.GVA_DB.Save(modelValue)
	global.GVA_LOG.Info("任务执行成功!", zap.String("taskId", taskId))
	isCompleted = true
}

func (s *extAiTask) ExecuteCozeTask(taskId string) (err error) {
	var (
		buf         []byte
		executeID   string
		isCompleted bool
		cozeTaskId  string
		uri         string
		result      *response.CozeResult
		ctx         context.Context
	)
	ctx = context.Background()
	payload := &request.CozeWorkflowRequest{}
	modelValue := &model.ExtAiTask{}
	defer func() {
		if !isCompleted {
			if err != nil {
				result.State = "failed"
				modelValue.Status = "failed"
			} else {
				modelValue.Status = "canceled"
			}
			if result != nil {
				if buf, err = json.Marshal(result); err == nil {
					modelValue.Result = string(buf)
				}
			}
			global.GVA_DB.Save(modelValue)
		}
	}()
	tm := time.Now()
	result = &response.CozeResult{
		State: "pending",
	}
	if err = global.GVA_DB.First(modelValue, "task_id = ?", taskId).Error; err != nil {
		global.GVA_LOG.Error("找不到任务数据!", zap.String("taskId", taskId), zap.Error(err))
		modelValue.Description = "任务参数解码失败"
		return
	}
	if err = json.Unmarshal([]byte(modelValue.Options), payload); err != nil {
		global.GVA_LOG.Error("任务参数解码失败!", zap.String("taskId", taskId), zap.Error(err))
		modelValue.Description = "任务参数解码失败"
		return
	}
	global.GVA_LOG.Info("开始执行扣子工作流!", zap.String("taskId", taskId), zap.String("WorkflowId", payload.WorkflowId))
	var wkResult *request.CozeResult
	result.State = "executing"
	if plugin.Config.CozeApiToken != "" {
		payload.Parameters["api_token"] = plugin.Config.CozeApiToken
	}
	if wkResult, err = s.executeCozeWorkflow(payload); err != nil {
		global.GVA_LOG.Error("执行任务失败!", zap.String("taskId", taskId), zap.Error(err))
		modelValue.Description = err.Error()
		return
	}
	executeID = wkResult.ExecuteId
	result.Execute.ID = executeID
	var str string
	if err = wkResult.Decode(&str); err == nil {
		tk := cozeTaskOutput{}
		if err = json.Unmarshal([]byte(str), &tk); err != nil {
			global.GVA_LOG.Error("解析扣子任务结果失败!", zap.String("taskId", taskId), zap.String("output", string(wkResult.Data)), zap.Error(err))
			return
		}
		cozeTaskId = tk.TaskId
		result.Execute.Output = str
	} else {
		global.GVA_LOG.Error("解析扣子任务结果失败!", zap.String("taskId", taskId), zap.String("output", string(wkResult.Data)), zap.Error(err))
		return
	}
	// global.GVA_LOG.Info("执行口子工作流成功!", zap.String("taskId", taskId), zap.String("executeID", executeID), zap.String("payload", modelValue.Options), zap.Error(err))
	// cozeTaskId, err = s.waitCozeWorflowOutput(ctx, payload.WorkflowId, executeID)
	result.Execute.Duration = time.Since(tm)
	result.Execute.ID = cozeTaskId
	if err != nil {
		global.GVA_LOG.Error("获取扣子工作流输出失败!", zap.String("taskId", taskId), zap.String("WorkflowId", payload.WorkflowId), zap.String("executeID", executeID), zap.Error(err))
		result.Execute.Error = err.Error()
		modelValue.Description = err.Error()
		return
	} else {
		global.GVA_LOG.Info("获取扣子工作流任务信息成功!", zap.String("taskId", taskId), zap.String("WorkflowId", payload.WorkflowId), zap.String("executeID", executeID), zap.String("cozeTaskID", cozeTaskId), zap.Error(err))
	}
	result.State = "tasking"

	uri, err = s.waitCozeTaskResult(ctx, payload.WorkflowId, cozeTaskId)
	result.Task.Duration = time.Since(tm)
	if err != nil {
		global.GVA_LOG.Error("执行扣子任务失败!", zap.String("taskId", taskId), zap.String("WorkflowId", payload.WorkflowId), zap.String("executeID", executeID), zap.String("cozeTaskID", cozeTaskId), zap.Error(err))
		modelValue.Description = err.Error()
		result.Task.Error = err.Error()
		return
	} else {
		global.GVA_LOG.Error("执行扣子任务成功!", zap.String("taskId", taskId), zap.String("WorkflowId", payload.WorkflowId), zap.String("executeID", executeID), zap.String("cozeTaskID", cozeTaskId), zap.String("uri", uri), zap.Error(err))
		modelValue.Description = "执行成功"
		modelValue.Status = "success"
		modelValue.CompleteAt = time.Now().Unix()
		result.State = "completed"
		result.Url = uri
		if buf, err = json.Marshal(result); err == nil {
			modelValue.Result = string(buf)
		}
		global.GVA_DB.Save(modelValue)
		isCompleted = true
	}
	return
}

func (s *extAiTask) DeleteImageResult(ctx context.Context, id int, indexs []int) (err error) {
	modelValue := model.ExtAiTask{}
	if modelValue, err = s.GetExtAiTask(ctx, strconv.Itoa(id)); err != nil {
		return
	}
	results := make([]*response.ImgResult, 0)
	if err = json.Unmarshal([]byte(modelValue.Result), &results); err != nil {
		return
	}
	imgBaseUrl := plugin.Config.HrefUrl
	for i := 0; i < len(results); i++ {
		if slices.Contains(indexs, results[i].Index) {
			if results[i].Url != "" {
				if err = os.Remove(strings.TrimPrefix(results[i].Url, imgBaseUrl)); err == nil {
					results[i].Url = ""
					results[i].Error = "已删除"
				}
			}
		}
	}
	var (
		buf []byte
	)
	if buf, err = json.Marshal(results); err == nil {
		modelValue.Result = string(buf)
		err = global.GVA_DB.Save(modelValue).Error
	}
	return
}

func (s *extAiTask) UploadCozeFile(ctx context.Context, header *multipart.FileHeader, fp multipart.File) (data request.CoreUploadData, err error) {
	var (
		partWriter io.Writer
	)
	body := &bytes.Buffer{}
	formWriter := multipart.NewWriter(body)
	partWriter, err = formWriter.CreateFormFile("file", header.Filename)
	if err != nil {
		return
	}
	if _, err = io.Copy(partWriter, fp); err != nil {
		return
	}
	formWriter.Close()
	var (
		req *http.Request
		res *http.Response
	)
	if req, err = http.NewRequestWithContext(ctx, "POST", "https://api.coze.cn/v1/files/upload", body); err != nil {
		return
	}
	req.Header.Set("Content-Type", formWriter.FormDataContentType())
	req.Header.Set("Authorization", "Bearer "+plugin.Config.CozeAccessToken)
	if res, err = http.DefaultClient.Do(req); err != nil {
		return
	}
	defer res.Body.Close()
	result := request.CozeResult{}
	if err = json.NewDecoder(res.Body).Decode(&result); err != nil {
		return
	}
	if result.Code != 0 {
		err = errors.New(result.Message)
		return
	} else {
		data = request.CoreUploadData{}
		err = result.Decode(&data)
		return
	}
}

func (s *extAiTask) executeCozeWorkflow(payload *request.CozeWorkflowRequest) (result *request.CozeResult, err error) {
	var (
		buf []byte
		req *http.Request
		res *http.Response
	)
	if payload.Parameters == nil {
		payload.Parameters = make(map[string]string)
	}
	if buf, err = json.Marshal(payload); err != nil {
		return
	}
	if req, err = http.NewRequest(http.MethodPost, "https://api.coze.cn/v1/workflow/run", bytes.NewReader(buf)); err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+plugin.Config.CozeAccessToken)
	if res, err = http.DefaultClient.Do(req); err != nil {
		return
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("unexpected status code: %d", res.StatusCode)
		return
	}
	result = &request.CozeResult{}
	if err = json.NewDecoder(res.Body).Decode(result); err != nil {
		return
	}
	if result.Code != 0 {
		err = fmt.Errorf("coze error: %s", result.Message)
		return
	}
	return
}

func (s *extAiTask) getCozeWorkflowHistories(ctx context.Context, workflowId, executeID string) (histories []*cozeHistory, err error) {
	var (
		buf []byte
		req *http.Request
		res *http.Response
	)
	uri := fmt.Sprintf("https://api.coze.cn/v1/workflows/%s/run_histories/%s", workflowId, executeID)
	if req, err = http.NewRequestWithContext(ctx, http.MethodPost, uri, bytes.NewReader(buf)); err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+plugin.Config.CozeAccessToken)
	if res, err = http.DefaultClient.Do(req); err != nil {
		return
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("unexpected status code: %d", res.StatusCode)
		return
	}
	result := &request.CozeResult{}
	if err = json.NewDecoder(res.Body).Decode(result); err != nil {
		return
	}
	histories = make([]*cozeHistory, 0)
	if err = result.Decode(&histories); err != nil {
		return
	}
	return
}

func (s *extAiTask) getCozeWorkflowTaskResult(ctx context.Context, payload *request.CozeWorkflowRequest) (uri string, err error) {
	var (
		result *request.CozeResult
	)
	if plugin.Config.CozeApiToken != "" {
		payload.Parameters["api"] = plugin.Config.CozeApiToken
	}
	if result, err = s.executeCozeWorkflow(payload); err != nil {
		return
	}
	var str string
	if err = result.Decode(&str); err != nil {
		return
	}
	videoResult := &cozeTaskResult{}
	if err = json.Unmarshal([]byte(str), videoResult); err != nil {
		return
	}
	if videoResult.Code == 1 {
		return videoResult.VideoURL, nil
	} else {
		return "", io.ErrNoProgress
	}
}

func (s *extAiTask) waitCozeWorflowOutput(ctx context.Context, workflowId, executeID string) (taskID string, err error) {
	var (
		errcount  int
		histories []*cozeHistory
	)
	ticker := time.NewTicker(time.Second * 5)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			err = context.Cause(ctx)
			return
		case <-ticker.C:
			if histories, err = s.getCozeWorkflowHistories(ctx, workflowId, executeID); err != nil {
				errcount++
				if errcount > 3 {
					return "", err
				}
			} else {
				if len(histories) > 0 {
					row := histories[0]
					if row.ErrorCode != "0" {
						return "", errors.New(row.ErrorMessage)
					}
					if row.ExecuteStatus == "Fail" {
						return "", fmt.Errorf("workflow %s execute failed(%s)", workflowId, row.Output)
					}
					if row.ExecuteStatus == "Success" {
						oret := &cozeOutput{}
						if err = json.Unmarshal([]byte(row.Output), oret); err == nil {
							tret := &cozeTaskOutput{}
							if err = json.Unmarshal([]byte(oret.Output), tret); err == nil {
								global.GVA_LOG.Info("获取口子工作流执行任务ID!", zap.String("executeID", executeID), zap.String("taskID", tret.TaskId))
								return tret.TaskId, nil
							} else {
								global.GVA_LOG.Error("解码扣子任务结果失败!", zap.String("executeID", executeID), zap.String("output", oret.Output), zap.Error(err))
							}
						} else {
							global.GVA_LOG.Error("解码扣子输出结果失败!", zap.String("executeID", executeID), zap.String("output", row.Output), zap.Error(err))
						}
					}
				}
			}
		}
	}
}

func (s *extAiTask) waitCozeTaskResult(ctx context.Context, workflowId, taskID string) (uri string, err error) {
	var (
		errcount int
	)
	payload := &request.CozeWorkflowRequest{
		WorkflowId: workflowId,
		IsAsync:    false,
		Parameters: map[string]string{
			"task_id": taskID,
		},
	}
	ticker := time.NewTicker(time.Second * 5)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			err = context.Cause(ctx)
			return
		case <-ticker.C:
			if uri, err = s.getCozeWorkflowTaskResult(ctx, payload); err == nil {
				return
			} else {
				if !errors.Is(err, io.ErrNoProgress) {
					errcount++
				}
				if errcount > 3 {
					return "", err
				}
			}
		}
	}
}

// CreateExtAiTask 创建extAiTask表记录
// Author [yourname](https://github.com/yourname)
func (s *extAiTask) CreateExtAiTask(ctx context.Context, extAiTask *model.ExtAiTask) (err error) {
	err = global.GVA_DB.Create(extAiTask).Error
	return err
}

// DeleteExtAiTask 删除extAiTask表记录
// Author [yourname](https://github.com/yourname)
func (s *extAiTask) DeleteExtAiTask(ctx context.Context, id string) (err error) {
	err = global.GVA_DB.Delete(&model.ExtAiTask{}, "id = ?", id).Error
	return err
}

// DeleteExtAiTaskByIds 批量删除extAiTask表记录
// Author [yourname](https://github.com/yourname)
func (s *extAiTask) DeleteExtAiTaskByIds(ctx context.Context, ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]model.ExtAiTask{}, "id in ?", ids).Error
	return err
}

// UpdateExtAiTask 更新extAiTask表记录
// Author [yourname](https://github.com/yourname)
func (s *extAiTask) UpdateExtAiTask(ctx context.Context, extAiTask model.ExtAiTask) (err error) {
	err = global.GVA_DB.Model(&model.ExtAiTask{}).Where("id = ?", extAiTask.ID).Updates(&extAiTask).Error
	return err
}

// GetExtAiTask 根据id获取extAiTask表记录
// Author [yourname](https://github.com/yourname)
func (s *extAiTask) GetExtAiTask(ctx context.Context, id string) (extAiTask model.ExtAiTask, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&extAiTask).Error
	return
}

// GetExtAiTaskInfoList 分页获取extAiTask表记录
// Author [yourname](https://github.com/yourname)
func (s *extAiTask) GetExtAiTaskInfoList(ctx context.Context, info request.ExtAiTaskSearch) (list []model.ExtAiTask, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.ExtAiTask{})
	var extAiTasks []model.ExtAiTask
	// 如果有条件搜索 下方会自动创建搜索语句

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}
	err = db.Find(&extAiTasks).Error
	return extAiTasks, total, err
}

func (s *extAiTask) GetExtAiTaskPublic(ctx context.Context) {

}
