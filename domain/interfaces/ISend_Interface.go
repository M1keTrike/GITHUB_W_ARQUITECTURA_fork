package domain

type ISend interface {
	Send(message string) int
}
