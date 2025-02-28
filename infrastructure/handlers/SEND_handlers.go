package handlers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"

	domain "github_wb/domain/value_objects"
)

type SendHandler struct {
	WebHookURL string
}

func NewSendHandler() *SendHandler {
	webhookURL := os.Getenv("WEB_HOOK_DISCORD_URL")
	if webhookURL == "" {
		log.Println("Error: WEB_HOOK_DISCORD_URL no está configurado")
	}
	return &SendHandler{
		WebHookURL: webhookURL,
	}
}

func (s *SendHandler) Send(message string) int {

	if s.WebHookURL == "" {
		log.Println("Error: No se puede enviar el mensaje porque la URL del webhook no está configurada")
		return http.StatusInternalServerError
	}

	payload := domain.Message{
		Content: message,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Println("Error al convertir el mensaje a JSON:", err)
		return http.StatusInternalServerError
	}

	resp, err := http.Post(s.WebHookURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("Error al enviar el mensaje a Discord:", err)
		return http.StatusInternalServerError
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		log.Println("Error: Discord respondió con estado", resp.Status)
		return resp.StatusCode
	}

	log.Println("Mensaje enviado correctamente a Discord")
	return resp.StatusCode
}
