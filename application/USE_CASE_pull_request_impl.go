package application

import (
	"encoding/json"
	"fmt"
	"log"

	domain "github_wb/domain/value_objects"
)

func ProcessPullRequest(payload []byte) (int, string) {
	var eventPayload domain.PullRequestEventPayload

	if err := json.Unmarshal(payload, &eventPayload); err != nil {
		errorMsg := fmt.Sprintf("Error al procesar payload: %v", err)
		log.Println(errorMsg)
		return 500, errorMsg
	}

	if eventPayload.Action == "ready_for_review" {
		title := eventPayload.PullRequest.Title
		content := eventPayload.PullRequest.Body
		user := eventPayload.PullRequest.User.Login

		successMsg := fmt.Sprintf("Pull Request aprobado.\nTítulo: %s\nContenido: %s\nUsuario que aprobó: %s", title, content, user)
		log.Println(successMsg)
		return 200, successMsg
	}

	infoMsg := fmt.Sprintf("Pull Request no aprobado o acción distinta: %s", eventPayload.Action)
	log.Println(infoMsg)
	return 400, infoMsg
}
