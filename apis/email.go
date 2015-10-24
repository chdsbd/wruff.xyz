package apis

import (
	"os"

	"github.com/mailgun/mailgun-go"
)

func SendEmail(email, message string) error {
	publicAPIKey := os.Getenv("mg_publicApiKey")
	apiKey := os.Getenv("mg_apiKey")
	mg := mailgun.NewMailgun("wruff.xyz", apiKey, publicAPIKey)
	m := mg.NewMessage(
		"Wruff Wruff <grr@wruff.xyz>",
		"Hello",
		message,
		email,
	)
	_, _, err := mg.Send(m)
	if err != nil {
		return err
	}
	return nil
}
