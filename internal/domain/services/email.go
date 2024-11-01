package services

type Email interface {
	AddToEmailList(username string, code string)
	CheckInEmailList(username string, code string)
	SendConfirmationEmail(to string, code string)
	GenerateConfirmationCode() (string, error)
}
