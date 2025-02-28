package handlers

import (
	"github_wb/application"

	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PullRequestEvent(ctx *gin.Context) {
	eventType := ctx.GetHeader("X-GitHub-Event")
	deliveryID := ctx.GetHeader("X-GitHub-Delivery")
	signature := ctx.GetHeader("X-Hub-Signature-256")
	sendHandler := NewSendHandler()
	discordUseCase := application.NewUseCaseSendToDiscord(sendHandler)

	log.Println(signature)

	log.Printf("Webhook recibido: \nEvento=%s, \nDeliveryID=%s", eventType, deliveryID)

	payload, err := ctx.GetRawData()

	if err != nil {
		log.Printf("Error al leer el cuerpo de la solicitud: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al leer el cuerpo de la solicitud"})
		return
	}

	var statusCode int
	var message string

	switch eventType {
	case "pull_request":
		statusCode, message = application.ProcessPullRequest(payload)
	}

	switch statusCode {
	case 200:
		ctx.JSON(http.StatusOK, gin.H{"status": "Evento Pull Request recibido y procesado"})
		discordUseCase.Execute(message)
	case 500:
		log.Printf("Error al deserializar el payload del pull request: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al procesar el payload del pull request"})
	default:
		ctx.JSON(http.StatusOK, gin.H{"status": "Peticion procesada"})
	}

}
