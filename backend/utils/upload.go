package utils

import (
	"bytes"
	"errors"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func UploadImage(bucket, key string, data []byte) error {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})
	if err != nil {
		log.Printf("Failed to create session: %v", err)
		return errors.New("failed to create session")
	}

	svc := s3.New(sess)
	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   bytes.NewReader(data),
	})
	if err != nil {
		log.Printf("Failed to upload image: %v", err)
		return errors.New("failed to upload image")
	}

	return nil
}
