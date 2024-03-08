package upload

import (
	"errors"
	"github.com/WaynerEP/restaurant-app/server/global"
	"github.com/WaynerEP/restaurant-app/server/utils"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"
	"time"

	"go.uber.org/zap"
)

type Local struct{}

// UploadFile uploads a file
// Object: *Local
// Function: UploadFile
// Description: Upload a file
// Param: file *multipart.FileHeader
// Return: string, string, error
func (*Local) UploadFile(file *multipart.FileHeader) (string, string, error) {
	ext := path.Ext(file.Filename)
	name := strings.TrimSuffix(file.Filename, ext)
	name = utils.MD5V([]byte(name))
	filename := name + "_" + time.Now().Format("20060102150405")
	mkdirErr := os.MkdirAll(global.GVA_CONFIG.Local.StorePath, os.ModePerm)
	if mkdirErr != nil {
		global.GVA_LOG.Error("os.MkdirAll() function failed", zap.Any("err", mkdirErr.Error()))
		return "", "", errors.New("os.MkdirAll() function failed, err:" + mkdirErr.Error())
	}
	p := global.GVA_CONFIG.Local.StorePath + "/" + filename
	filepath := global.GVA_CONFIG.Local.Path + "/" + filename
	f, openError := file.Open()
	if openError != nil {
		global.GVA_LOG.Error("file.Open() function failed", zap.Any("err", openError.Error()))
		return "", "", errors.New("file.Open() function failed, err:" + openError.Error())
	}
	defer f.Close()
	out, createErr := os.Create(p)
	if createErr != nil {
		global.GVA_LOG.Error("os.Create() function failed", zap.Any("err", createErr.Error()))
		return "", "", errors.New("os.Create() function failed, err:" + createErr.Error())
	}
	defer out.Close()
	_, copyErr := io.Copy(out, f)
	if copyErr != nil {
		global.GVA_LOG.Error("io.Copy() function failed", zap.Any("err", copyErr.Error()))
		return "", "", errors.New("io.Copy() function failed, err:" + copyErr.Error())
	}
	return filepath, filename, nil
}

// DeleteFile deletes a file
// Object: *Local
// Function: DeleteFile
// Description: Delete a file
// Param: key string
// Return: error
func (*Local) DeleteFile(key string) error {
	p := global.GVA_CONFIG.Local.StorePath + "/" + key
	if strings.Contains(p, global.GVA_CONFIG.Local.StorePath) {
		if err := os.Remove(p); err != nil {
			return errors.New("Local file Error al eliminar, err:" + err.Error())
		}
	}
	return nil
}
