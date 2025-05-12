package external_api

import (
	"net/http"
	"strings"
)

func SendToAli(text string) {
	go func() {
		_, _ = http.Post(
			"https://example.com/endpoint",
			"text/plain",
			strings.NewReader(text),
		)
	}()
}
