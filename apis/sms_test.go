package apis

import (
	"os"
	"testing"
)

func TestSendSMS(t *testing.T) {
	toNumber := os.Getenv("Twilio_TestPhone")
	err := SendSMS(toNumber, "Testing123")
	if err != nil {
		t.Error(err)
	}
}
