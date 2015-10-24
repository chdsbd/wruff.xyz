package apis

import (
	"os"
	"testing"
)

func TestSendCall(t *testing.T) {
	toNumber := os.Getenv("Twilio_TestPhone")
	err := SendCall(toNumber)
	if err != nil {
		t.Error(err)
	}
}
