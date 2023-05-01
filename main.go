package main

import (
	"eventpublisher/aws"
	"eventpublisher/constants"
	"log"
	"time"
)

func main() {

	session, err := aws.CreateSession()
	if err != nil {
		log.Fatal(err)
	}

	for {
		time.Sleep(10 * time.Second)
		aws.UploadFileToS3(session, "testfile.txt", constants.EVENT_BUCKET)
	}
}
