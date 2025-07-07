package component

import (
	"ETLProject/config"
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func UploadFileToS3(bucket string, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("** Error abriendo archivo: %v **", err)
	}
	defer file.Close()

	key := filepath.Base(filePath)

	_, err = config.S3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: &bucket,
		Key:    &key,
		Body:   file,
	})
	if err != nil {
		return fmt.Errorf("error subiendo archivo a S3: %v", err)
	}

	log.Printf("** ARCHIVO %s SUBIDO A S3 CON EXITO **", key)
	return nil
}
