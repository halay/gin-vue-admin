import service from '@/utils/request'
import { ElMessage, ElNotification } from 'element-plus'
import { defineStore } from 'pinia'
import { ref } from 'vue'

// 上传图片资源接口
const uploadCozeFile = (body) => {
  let formData = new FormData()
  if (!(body instanceof FormData)) {
    Object.keys(body).forEach((key) => {
      formData.append(key, body[key])
    })
  } else {
    formData = body
  }
  return service({
    url: '/yApi/uploadCozeFile',
    method: 'post',
    data: formData,
    donNotShowLoading: true
  })
}
// 执行一个工作流
const runCozeWorkflow = (body) => {
  return service({
    url: '/yApi/executeCozeTask',
    method: 'post',
    data: body,
    donNotShowLoading: true
  })
}
// 获取执行结果
const getCozeWorkflowResult = (id) => {
  return service({
    url: '/yApi/getCozeTaskResult',
    method: 'post',
    data: { id },
    donNotShowLoading: true
  })
}
// 重新执行任务
const retryCozeWorkflow = (id) => {
  return service({
    url: '/yApi/reExecuteCozeTask',
    method: 'post',
    data: { id },
    donNotShowLoading: true
  })
}

export const defineCozeWorkflowStore = (storeId) => {

  // 工作流 ID 配置
  const consig = {
    tip: {
      priduct: 'AI 商品宣传片智能生成完成',
      vlog: 'AI 探店视频智能生成完成',
    },
    type: {
      priduct: 'video1',
      vlog: 'video2',
    },
    workflowIds: '7579514463738167323',
    priduct: '7589196415885983779',
    vlog: '7579514432432472127',
    suApiKey: '2af09cae-106a-49db-9361-00b7b0b20ca8',
  }
  
  return defineStore(`coze-${storeId}`, () => {
    const isGenerating = ref(false) // 是否正在生成中
    const taskId = ref(0) // 生成任务 ID
    const taskStatus = ref('') // 任务状态
    const imageList = ref([
      {
        id: '',
        url: '',
        status: '',
      },
      {
        id: '',
        url: '',
        status: '',
      },
      {
        id: '',
        url: '',
        status: '',
      },
      {
        id: '',
        url: '',
        status: '',
      },
    ]) // 上传图片列表
    const taskTime = ref('') // 任务创建时间
    const videoUrl = ref('') // 视频 URL
    const text = ref('') // 提示内容

    // 删除图片
    const deleteImage = (index) => {
      imageList.value[index] = {
        id: '',
        url: '',
        status: '',
      }
    }

    // 上传图片
    const uploadImages = async (index, params) => {
      imageList.value[index].status = 'loading'
      try {
        const { code, data, msg } = await uploadCozeFile(params);
        if (code !== 0) {
          ElMessage.error(msg || '图片上传失败')
          throw new Error(msg || '图片上传失败')
        }
        const { upload, data: cozeData } = data;
        const image = {
          id: cozeData.id,
          url: upload.url,
          status: 'success',
        };
        imageList.value[index] = image
        return image
      } catch (error) {
        ElMessage.error(error.msg || error || '图片上传失败')
        imageList.value[index].status = 'error'
        throw error
      }
    }

    // 发起生成请求的 Action
    const createTask = async (params) => {
      isGenerating.value = true;
      try {
        const { code, data } = await runCozeWorkflow({
          type: consig.type[storeId],
          workflows: [
            {
              workflow_id: consig[storeId],
              parameters: {
                api_token: consig.suApiKey,
                ...params
              }
            },
            {
              workflow_id: consig.workflowIds,
              parameters: {
                api: consig.suApiKey
              }
            }
          ]
        });
        // const { code, data } = { code: 0, data: { id: 9 } }
        if (code !== 0) {
          ElMessage.error(data.msg)
          throw new Error(data.msg)
        }
        taskId.value = data.id;
        // 开启轮询
        pollTaskStatus(data.id);
        return taskId;
      } catch (error) {
        ElMessage.error('智能海报生成失败')
        imageList.value = []
        throw error;
      }
    };

    // 轮询逻辑 (递归 setTimeout)
    const pollTaskStatus = async (taskId) => {
      isGenerating.value = true;
      try {
        const { data } = await getCozeWorkflowResult(taskId);
        const { status, data: cozeData } = data;
        console.log(status, cozeData)
        taskStatus.value = status
        if (status === 'completed') {
          videoUrl.value = cozeData.url
          isGenerating.value = false
          ElNotification({
            title: consig.tip[storeId],
            message: `视频已生成，请查看`,
            type: 'success',
            duration: 5000
          });
        } else if (status === 'failed' || status === 'canceled') {
          isGenerating.value = false
          ElNotification({
            title: consig.tip[storeId],
            message: `视频生成失败，${cozeData.error || '请重试'}` ,
            type: 'error',
            duration: 5000
          });
          // throw new Error('视频生成失败')
        } else {
          // 未完成：3秒后继续下一次轮询
          setTimeout(() => pollTaskStatus(taskId), 20000);
        }
      } catch (error) {
        console.error('获取结果失败，5秒后重试:', error);
        // 遇到网络报错不中断，延长间隔继续重试
        setTimeout(() => pollTaskStatus(taskId), 5000);
      }
    }

    // 重新执行任务
    const retryTask = async (id) => {
      isGenerating.value = true;
      taskStatus.value = ''
      try {
        const taskId = id || taskId.value
        const { code, data, msg } = await retryCozeWorkflow(taskId);
        if (code !== 0) {
          ElMessage.error(msg || data.msg)
          throw new Error(msg || data.msg)
        }
        ElMessage.success('任务已重新执行')
        // 开启轮询
        setTimeout(() => pollTaskStatus(taskId), 10000);
      } catch (error) {
        ElMessage.error(error.msg || error || '任务重新执行失败')
        throw error
      }
    }

    // 重置，重新生成
    const onResetCreate = () => {
      isGenerating.value = false
      taskId.value = 0
      taskStatus.value = ''
      videoUrl.value = ''
    }

    return {
      isGenerating,
      taskId,
      taskStatus,
      imageList,
      taskTime,
      text,
      videoUrl,
      deleteImage,
      uploadImages,
      createTask,
      retryTask,
      onResetCreate
    }
  })()
}
