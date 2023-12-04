package utill

import (
	"abdullayev13/timeup/internal/pkg/upload"
	"github.com/google/uuid"
	"mime/multipart"
)

const (
	tempFileName      = "temp"
	mediaTempFileName = "./media/temp"
)

func TranscodeAndUploadS3Video(file *multipart.FileHeader) (string, error) {
	filePath, err := Upload(file, tempFileName)
	if err != nil {
		return "", err
	}
	filePath = "." + filePath
	defer RemoveFile(filePath)

	uuidName := uuid.New().String()
	outputFilePath := mediaTempFileName + "/" + uuidName + ".mp4"

	err = TranscodeVideo(filePath, outputFilePath)
	if err != nil {
		return "", err
	}
	defer RemoveFile(outputFilePath)

	fileUrl, err := upload.UploadToS3(outputFilePath)
	if err != nil {
		return "", err
	}

	return fileUrl, nil
}

func TranscodeAndUploadS3Img(file *multipart.FileHeader) (string, error) {
	filePath, err := Upload(file, tempFileName)
	if err != nil {
		return "", err
	}
	filePath = "." + filePath
	defer RemoveFile(filePath)

	uuidName := uuid.New().String()
	outputFilePath := mediaTempFileName + "/" + uuidName + ".jpg"

	err = TranscodeImg(filePath, outputFilePath)
	if err != nil {
		return "", err
	}
	defer RemoveFile(outputFilePath)

	fileUrl, err := upload.UploadToS3(outputFilePath)
	if err != nil {
		return "", err
	}

	return fileUrl, nil
}
