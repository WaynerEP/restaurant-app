package upload

import (
	"errors"
	"mime/multipart"
	"time"

	"github.com/WaynerEP/restaurant-app/server/global"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"

	"go.uber.org/zap"
)

type AliyunOSS struct{}

func (*AliyunOSS) UploadFile(file *multipart.FileHeader) (string, string, error) {
	bucket, err := NewBucket()
	if err != nil {
		global.GVA_LOG.Error("function AliyunOSS.NewBucket() Failed", zap.Any("err", err.Error()))
		return "", "", errors.New("function AliyunOSS.NewBucket() Failed, err:" + err.Error())
	}

	// Read the local file.
	f, openError := file.Open()
	if openError != nil {
		global.GVA_LOG.Error("function file.Open() Failed", zap.Any("err", openError.Error()))
		return "", "", errors.New("function file.Open() Failed, err:" + openError.Error())
	}
	defer f.Close() // Create file defer close.
	// Upload path on Aliyun, you can change the file name format as needed for uniqueness.
	// yunFileTmpPath := filepath.Join("uploads", time.Now().Format("2006-01-02")) + "/" + file.Filename
	yunFileTmpPath := global.GVA_CONFIG.AliyunOSS.BasePath + "/" + "uploads" + "/" + time.Now().Format("2006-01-02") + "/" + file.Filename

	// Upload the file stream.
	err = bucket.PutObject(yunFileTmpPath, f)
	if err != nil {
		global.GVA_LOG.Error("function formUploader.Put() Failed", zap.Any("err", err.Error()))
		return "", "", errors.New("function formUploader.Put() Failed, err:" + err.Error())
	}

	return global.GVA_CONFIG.AliyunOSS.BucketUrl + "/" + yunFileTmpPath, yunFileTmpPath, nil
}

func (*AliyunOSS) DeleteFile(key string) error {
	bucket, err := NewBucket()
	if err != nil {
		global.GVA_LOG.Error("function AliyunOSS.NewBucket() Failed", zap.Any("err", err.Error()))
		return errors.New("function AliyunOSS.NewBucket() Failed, err:" + err.Error())
	}

	// Delete a single file. objectName represents the complete path to the OSS file, including the file extension, e.g., abc/efg/123.jpg.
	// If you want to delete a folder, set objectName to the folder name. If the folder is not empty, you need to delete all objects within the folder before deleting the folder itself.
	err = bucket.DeleteObject(key)
	if err != nil {
		global.GVA_LOG.Error("function bucketManager.Delete() Filed", zap.Any("err", err.Error()))
		return errors.New("function bucketManager.Delete() Filed, err:" + err.Error())
	}

	return nil
}

func NewBucket() (*oss.Bucket, error) {
	// Create an OSSClient instance.
	client, err := oss.New(global.GVA_CONFIG.AliyunOSS.Endpoint, global.GVA_CONFIG.AliyunOSS.AccessKeyId, global.GVA_CONFIG.AliyunOSS.AccessKeySecret)
	if err != nil {
		return nil, err
	}

	// Get the repository bucket.
	bucket, err := client.Bucket(global.GVA_CONFIG.AliyunOSS.BucketName)
	if err != nil {
		return nil, err
	}

	return bucket, nil
}
