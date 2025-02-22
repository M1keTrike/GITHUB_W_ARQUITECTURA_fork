package application

import (
	"encoding/json"
	domain "github_wb/domain/value_objects"
	"log"
)

func ProcessPullRequest(payload []byte) int {
	var eventPayload domain.PullRequestReviewEventPayload

	log.Printf("Payload recibido: %s", payload)


	if err := json.Unmarshal(payload, &eventPayload); err != nil {
		log.Printf("Error al procesar payload: %v", err)
		return 500
	}

	if eventPayload.Action == "submitted" && eventPayload.Review.State == "approved" {
		title := eventPayload.PullRequest.Title
		content := eventPayload.Review.Body
		user := eventPayload.Review.User.Login

		log.Printf("Pull Request Reviewed:\nTítulo: %s\nContenido: %s\nUsuario que aprobó: %s", title, content, user)
	} else {
		log.Printf("Pull Request Review no aprobado o acción distinta: %s", eventPayload.Action)
	}

	return 200
}

