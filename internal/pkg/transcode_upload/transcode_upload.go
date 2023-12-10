package transcode_upload

import (
	"abdullayev13/timeup/internal/pkg/ffmpeg"
	"abdullayev13/timeup/internal/pkg/upload"
	"abdullayev13/timeup/internal/utill"
	"github.com/google/uuid"
	"mime/multipart"
	"sync"
)

const (
	tempFileName      = "temp"
	mediaTempFileName = "./media/temp"
)

func TranscodeAndUploadS3Video(file *multipart.FileHeader, callback func(string, error)) {
	filePath, err := utill.Upload(file, tempFileName)
	if err != nil {
		callback("", err)
		return
	}

	filePath = "." + filePath
	defer utill.RemoveFile(filePath)

	uuidName := uuid.New().String()
	outputFilePath := mediaTempFileName + "/" + uuidName + ".mp4"

	wg := sync.WaitGroup{}
	wg.Add(1)
	ffmpeg.TranscodeVideo(filePath, outputFilePath, func(callbackErr error) {
		err = callbackErr
		wg.Done()
	})
	wg.Wait()
	if err != nil {
		callback("", err)
		return
	}
	defer utill.RemoveFile(outputFilePath)

	fileUrl, err := upload.UploadToS3(outputFilePath)
	if err != nil {
		callback("", err)
		return
	}

	callback(fileUrl, nil)
}

func TranscodeAndUploadS3Img(file *multipart.FileHeader, callback func(string, error)) {
	filePath, err := utill.Upload(file, tempFileName)
	if err != nil {
		callback("", err)
		return
	}
	filePath = "." + filePath
	defer utill.RemoveFile(filePath)

	uuidName := uuid.New().String()
	outputFilePath := mediaTempFileName + "/" + uuidName + ".jpg"

	wg := sync.WaitGroup{}
	wg.Add(1)
	ffmpeg.TranscodeImg(filePath, outputFilePath, func(callbackErr error) {
		err = callbackErr
		wg.Done()
	})

	wg.Wait()
	if err != nil {
		callback("", err)
		return
	}
	defer utill.RemoveFile(outputFilePath)

	fileUrl, err := upload.UploadToS3(outputFilePath)
	if err != nil {
		callback("", err)
		return
	}

	callback(fileUrl, nil)
}

func ThumbnailAndUploadS3(videoFile *multipart.FileHeader, callback func(string, error)) {
	{
		cloneFile := *videoFile
		cloneFile.Filename = "thumbnail_" + cloneFile.Filename
		videoFile = &cloneFile
	}

	filePath, err := utill.Upload(videoFile, tempFileName)
	if err != nil {
		callback("", err)
		return
	}
	filePath = "." + filePath
	defer utill.RemoveFile(filePath)

	uuidName := uuid.New().String()
	outputFilePath := mediaTempFileName + "/" + uuidName + ".jpg"

	wg := sync.WaitGroup{}
	wg.Add(1)
	ffmpeg.GenerateThumbnail(filePath, outputFilePath, func(callbackErr error) {
		err = callbackErr
		wg.Done()
	})

	wg.Wait()
	if err != nil {
		callback("", err)
		return
	}
	defer utill.RemoveFile(outputFilePath)

	fileUrl, err := upload.UploadToS3(outputFilePath)
	if err != nil {
		callback("", err)
		return
	}

	callback(fileUrl, nil)
}
