package application

import domain "github_wb/domain/interfaces"

type USE_CASE_SendToDiscord struct {
	sender domain.ISend
}

func NewUseCaseSendToDiscord(sender domain.ISend) *USE_CASE_SendToDiscord {
	return &USE_CASE_SendToDiscord{sender: sender}
}

func (u *USE_CASE_SendToDiscord) Execute(message string) int {
	if message == "" {
		return 400
	}

	return u.sender.Send(message)
}
