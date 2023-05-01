package main

import (
	"eventpublisher/aws"
	"fmt"
	"log"
	"time"
)

func main() {

	awsSession, err := aws.CreateSession()
	if err != nil {
		log.Fatal(err)
	}
	aws.UploadFileToS3(awsSession, "testfile.txt", "testbucket")

	for {
		time.Sleep(3 * time.Second)
		fmt.Println("Running on docker")
	}

}
