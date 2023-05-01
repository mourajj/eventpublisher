package main

import (
	"eventpublisher/aws"
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
		aws.UploadFileToS3(session, "testfile.txt", "jm-events-bucket")
	}
}
