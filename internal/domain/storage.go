package domain

type Storage interface {
	GetSentences(userID string) string
}
