package services

type AuthQueue interface {
	AddToQueue(username string, eve string)
	GetEvent(username string) (string, error)
	DeleteEve(username string)
}
