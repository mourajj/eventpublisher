package aws

import (
	"eventpublisher/constants"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func CreateSession() (*session.Session, error) {
	//Create a new AWS session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(constants.AWS_REGION_STRING),
	})
	if err != nil {
		panic(err)
	}

	return sess, err
}

func UploadFileToS3(awsSession *session.Session, filename, bucket string) {

	//Creating a s3 client
	service := s3.New(awsSession)

	//Opening the file to upload
	file, err := os.Open("events/" + filename)
	if err != nil {
		fmt.Println("Failed to open file: ", err)
		os.Exit(1)
	}
	defer file.Close()

	// Upload the file to S3
	_, err = service.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),   // Replace with your S3 bucket's name
		Key:    aws.String(filename), // Replace with the name you want to give to the file in S3
		Body:   file,
	})
	if err != nil {
		log.Fatal("Failed to upload fileL: ", err)
		os.Exit(1)
	}

	log.Printf("File %s uploaded to S3 bucket %s", filename, bucket)
}
