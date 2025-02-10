package twillio

import (
	"fmt"
	"log"
	"os"

	"github.com/twilio/twilio-go"
	v2010 "github.com/twilio/twilio-go/rest/api/v2010"
)

func SendWhatsAppMessage(to, body string) error {
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: os.Getenv("TWILIO_ACCOUNT_SID"),
		Password: os.Getenv("TWILIO_AUTH_TOKEN"),
	})

	params := &v2010.CreateMessageParams{}
	params.SetTo("whatsapp:" + to)
	params.SetFrom("whatsapp:+77789802147")
	params.SetBody(body)

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		log.Printf("❌ Ошибка отправки WhatsApp-сообщения через Twilio: %v", err)
		return fmt.Errorf("ошибка Twilio: %w", err)
	}

	log.Printf("📨 WhatsApp-сообщение отправлено: SID %s\n", *resp.Sid)
	return nil
}
