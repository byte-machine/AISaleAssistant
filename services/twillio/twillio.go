package twillio

import (
	"fmt"
	"log"
	"os"

	"github.com/twilio/twilio-go"
	v2010 "github.com/twilio/twilio-go/rest/api/v2010"
)

func SendTwilioMessage(to, body string) error {
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: os.Getenv("TWILIO_ACCOUNT_SID"),
		Password: os.Getenv("TWILIO_AUTH_TOKEN"),
	})

	params := &v2010.CreateMessageParams{}
	params.SetFrom(os.Getenv("TWILIO_PHONE_NUMBER"))
	params.SetTo(to)
	params.SetBody(body)

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		return fmt.Errorf("ошибка отправки через Twilio: %w", err)
	}

	log.Printf("📨 Сообщение отправлено через Twilio: SID %s\n", *resp.Sid)
	return nil
}
