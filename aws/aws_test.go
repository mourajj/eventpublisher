package aws

import (
	"eventpublisher/constants"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

func TestSessionCreation(t *testing.T) {

	_, err := session.NewSession(&aws.Config{
		Region: aws.String(constants.AWS_REGION_STRING),
	})

	if err != nil {
		t.Errorf("Failed to create AWS session: %s", err)
	}

}
