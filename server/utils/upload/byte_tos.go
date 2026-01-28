package upload

import (
	"context"
	"errors"
	"io"
	"mime/multipart"
	"path"
	"time"

	"github.com/volcengine/ve-tos-golang-sdk/v2/tos"
	"go.uber.org/zap"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type ByteTOS struct{}

func (*ByteTOS) PutFile(ctx context.Context, filename string, reader io.Reader) (string, string, error) {
	client, err := NewByteBucket()
	if err != nil {
		global.GVA_LOG.Error("function ByteTOS.NewClient() Failed", zap.Any("err", err.Error()))
		return "", "", errors.New("function ByteTOS.NewClient() Failed, err:" + err.Error())
	}
	basename := path.Base(filename)
	// 读取本地文件。
	// 上传阿里云路径 文件名格式 自己可以改 建议保证唯一性
	// yunFileTmpPath := filepath.Join("uploads", time.Now().Format("2006-01-02")) + "/" + file.Filename
	yunFileTmpPath := global.GVA_CONFIG.ByteTos.BasePath + "/" + "uploads" + "/" + time.Now().Format("2006-01-02") + "/" + basename

	// 上传文件流。
	output, err := client.PutObjectV2(ctx, &tos.PutObjectV2Input{
		PutObjectBasicInput: tos.PutObjectBasicInput{
			Bucket: global.GVA_CONFIG.ByteTos.BucketName,
			Key:    yunFileTmpPath,
		},
		Content: reader,
	})
	if err != nil {
		global.GVA_LOG.Error("function formUploader.Put() Failed", zap.Any("err", err.Error()))
		return "", "", errors.New("function formUploader.Put() Failed, err:" + err.Error())
	}
	global.GVA_LOG.Info("file.UploadTos success", zap.Any("RequestID", output.RequestID))
	return global.GVA_CONFIG.ByteTos.BucketUrl + "/" + yunFileTmpPath, yunFileTmpPath, nil
}

func (*ByteTOS) UploadFile(file *multipart.FileHeader) (string, string, error) {
	client, err := NewByteBucket()
	if err != nil {
		global.GVA_LOG.Error("function ByteTOS.NewClient() Failed", zap.Any("err", err.Error()))
		return "", "", errors.New("function ByteTOS.NewClient() Failed, err:" + err.Error())
	}

	// 读取本地文件。
	f, openError := file.Open()
	if openError != nil {
		global.GVA_LOG.Error("function file.Open() Failed", zap.Any("err", openError.Error()))
		return "", "", errors.New("function file.Open() Failed, err:" + openError.Error())
	}
	defer f.Close() // 创建文件 defer 关闭
	// 上传阿里云路径 文件名格式 自己可以改 建议保证唯一性
	// yunFileTmpPath := filepath.Join("uploads", time.Now().Format("2006-01-02")) + "/" + file.Filename
	yunFileTmpPath := global.GVA_CONFIG.ByteTos.BasePath + "/" + "uploads" + "/" + time.Now().Format("2006-01-02") + "/" + file.Filename

	// 上传文件流。
	output, err := client.PutObjectV2(context.Background(), &tos.PutObjectV2Input{
		PutObjectBasicInput: tos.PutObjectBasicInput{
			Bucket: global.GVA_CONFIG.ByteTos.BucketName,
			Key:    yunFileTmpPath,
		},
		Content: f,
	})
	if err != nil {
		global.GVA_LOG.Error("function formUploader.Put() Failed", zap.Any("err", err.Error()))
		return "", "", errors.New("function formUploader.Put() Failed, err:" + err.Error())
	}
	global.GVA_LOG.Info("file.UploadTos success", zap.Any("RequestID", output.RequestID))
	return global.GVA_CONFIG.ByteTos.BucketUrl + "/" + yunFileTmpPath, yunFileTmpPath, nil
}

func (*ByteTOS) DeleteFile(key string) error {
	client, err := NewByteBucket()
	if err != nil {
		global.GVA_LOG.Error("function ByteTOS.NewBucket() Failed", zap.Any("err", err.Error()))
		return errors.New("function ByteTOS.NewBucket() Failed, err:" + err.Error())
	}

	output, err := client.DeleteObjectV2(context.Background(), &tos.DeleteObjectV2Input{
		Bucket: global.GVA_CONFIG.ByteTos.BucketName,
		Key:    key,
	})
	if err != nil {
		global.GVA_LOG.Error("function bucketManager.Delete() failed", zap.Any("err", err.Error()))
		return errors.New("function bucketManager.Delete() failed, err:" + err.Error())
	}
	global.GVA_LOG.Info("file.DeleteTos success", zap.Any("RequestID", output.RequestID))
	return nil
}

func NewByteBucket() (*tos.ClientV2, error) {
	client, err := tos.NewClientV2(global.GVA_CONFIG.ByteTos.Endpoint, tos.WithRegion(global.GVA_CONFIG.ByteTos.Region),
		tos.WithCredentials(tos.NewStaticCredentials(global.GVA_CONFIG.ByteTos.AccessKeyId, global.GVA_CONFIG.ByteTos.AccessKeySecret)))
	// 创建Client实例。
	if err != nil {
		return nil, err
	}

	return client, nil
}
