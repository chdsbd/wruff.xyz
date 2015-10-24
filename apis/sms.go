package apis

import (
	"os"

	"github.com/missionMeteora/twilio"
)

func SendSMS(toNumber, message string) error {
	accountSid := os.Getenv("Twilio_AccoutSid")
	authToken := os.Getenv("Twilio_AuthToken")
	fromNumber := os.Getenv("Twilio_Number")
	c := twilio.New(accountSid, authToken, fromNumber)
	err := c.Send(toNumber, message)
	if err != nil {
		return err
	}
	return nil
}
