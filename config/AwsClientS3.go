package config

import (
	"ETLProject/util"
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var S3Client *s3.Client

func InitAWS() {
	log.Println("** INICIAZION DEL CLIENT S3 **")
	accessKey := util.LoadProperty("aws_access_key_id")
	secretKey := util.LoadProperty("aws_secret_access_key")
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")))
	if err != nil {
		log.Fatalf("Error cargando configuración AWS: %v", err)
	}

	S3Client = s3.NewFromConfig(cfg)
	log.Println("** CLIENTE S3 INICIADO CORRECTAMENTE **")
}
