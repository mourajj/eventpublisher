package aws

import (
	"bytes"
	"eventpublisher/constants"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/golang/mock/gomock"
	"github.com/xitongsys/parquet-go-source/s3/mocks"
)

type mockS3Client struct {
	s3iface.S3API
}

func TestSessionCreation(t *testing.T) {

	_, err := session.NewSession(&aws.Config{
		Region: aws.String(constants.AWS_REGION_STRING),
	})

	if err != nil {
		t.Errorf("Failed to create AWS session: %s", err)
	}

}

func TestUploadFileToS3(t *testing.T) {
	//Creating an instance of the gomock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	//Creating a mock of S3 client
	mockS3Client := mocks.NewMockS3API(ctrl)

	//Test Parameters
	file := []byte("assuming this is a file")
	bucket := "test-bucket"
	key := "test-key"

	//Creating an expected output for PutObject call
	expectedOutput := &s3.PutObjectOutput{}

	//Setting up the mock s3 client to expect a Putobject call
	mockS3Client.EXPECT().PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   bytes.NewReader(file),
	}).Return(expectedOutput, nil)

	//Running the putobject call
	_, err := mockS3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   bytes.NewReader(file),
	})
	if err != nil {
		t.Errorf("Failed to upload a file to S3.  (%s)", err)
		os.Exit(1)
	}
}
