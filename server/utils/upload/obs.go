package upload

import (
	"mime/multipart"

	"github.com/WaynerEP/restaurant-app/server/global"

	"github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
	"github.com/pkg/errors"
)

type _obs struct{}

// HuaWeiObs represents the Huawei OBS client.
var HuaWeiObs = new(_obs)

// NewHuaWeiObsClient creates a new Huawei OBS client.
func NewHuaWeiObsClient() (client *obs.ObsClient, err error) {
	return obs.New(global.GVA_CONFIG.HuaWeiObs.AccessKey, global.GVA_CONFIG.HuaWeiObs.SecretKey, global.GVA_CONFIG.HuaWeiObs.Endpoint)
}

// UploadFile uploads a file to Huawei OBS.
func (o *_obs) UploadFile(file *multipart.FileHeader) (filename string, filepath string, err error) {
	var open multipart.File
	open, err = file.Open()
	if err != nil {
		return filename, filepath, err
	}
	filename = file.Filename
	input := &obs.PutObjectInput{
		PutObjectBasicInput: obs.PutObjectBasicInput{
			ObjectOperationInput: obs.ObjectOperationInput{
				Bucket: global.GVA_CONFIG.HuaWeiObs.Bucket,
				Key:    filename,
			},
			//HttpHeader: file.Header.Get("content-type"),
		},
		Body: open,
	}

	var client *obs.ObsClient
	client, err = NewHuaWeiObsClient()
	if err != nil {
		return filepath, filename, errors.Wrap(err, "Failed to get Huawei OBS client!")
	}

	_, err = client.PutObject(input)
	if err != nil {
		return filepath, filename, errors.Wrap(err, "File upload failed!")
	}
	filepath = global.GVA_CONFIG.HuaWeiObs.Path + "/" + filename
	return filepath, filename, err
}

// DeleteFile deletes a file from Huawei OBS.
func (o *_obs) DeleteFile(key string) error {
	client, err := NewHuaWeiObsClient()
	if err != nil {
		return errors.Wrap(err, "Failed to get Huawei OBS client!")
	}
	input := &obs.DeleteObjectInput{
		Bucket: global.GVA_CONFIG.HuaWeiObs.Bucket,
		Key:    key,
	}
	var output *obs.DeleteObjectOutput
	output, err = client.DeleteObject(input)
	if err != nil {
		return errors.Wrapf(err, "Deletion failed object (%s)!, output: %v", key, output)
	}
	return nil
}
