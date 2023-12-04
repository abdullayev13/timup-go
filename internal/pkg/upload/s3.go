package upload

import (
	"abdullayev13/timeup/internal/config"
	"fmt"
	"mime"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func UploadToS3(filename string) (string, error) {
	// The session the S3 Uploader will use
	sess := session.Must(session.NewSession())

	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)

	f, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("failed to open file %q, %v", filename, err)
	}

	// Upload the file to S3.
	name := filepath.Base(f.Name())
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(config.S3BucketName),
		Key:         aws.String(name),
		Body:        f,
		ContentType: aws.String(mime.TypeByExtension(filepath.Ext(name))),
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload file, %v", err)
	}
	fmt.Printf("file uploaded to, %s\n", result.Location)
	return result.Location, nil
}
