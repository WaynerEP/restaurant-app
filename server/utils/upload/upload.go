package upload

// OSS Object Storage Interface
import (
	"mime/multipart"

	"github.com/WaynerEP/restaurant-app/server/global"
)

type OSS interface {
	UploadFile(file *multipart.FileHeader) (string, string, error)
	DeleteFile(key string) error
}

// NewOss is the instantiation method for OSS
func NewOss() OSS {
	switch global.GVA_CONFIG.System.OssType {
	case "local":
		return &Local{}
	case "aliyun-oss":
		return &AliyunOSS{}
	case "huawei-obs":
		return HuaWeiObs
	case "aws-s3":
		return &AwsS3{}
	default:
		return &Local{}
	}
}
