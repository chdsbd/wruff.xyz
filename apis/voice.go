package apis

import (
	"os"

	"github.com/wyc/utwil"
)

func SendCall(toNumber string) error {
	fromNumber := os.Getenv("Twilio_Number")
	accountSid := os.Getenv("Twilio_AccoutSid")
	authToken := os.Getenv("Twilio_AuthToken")
	client := utwil.NewClient(accountSid, authToken)
	_, err := client.Call(fromNumber, toNumber, "https://raw.githubusercontent.com/chdsbd/wruff.xyz/master/apis/voice.xml")
	if err != nil {
		return err
	}
	return nil
}
