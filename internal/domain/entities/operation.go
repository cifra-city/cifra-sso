package entities

type OperationType string

const (
	ChangeEmail    = "CHANGE_EMAIL"
	ChangePassword = "CHANGE_PASSWORD"
	ChangeUsername = "CHANGE_USERNAME"
	ConfirmEmail   = "CONFIRM_EMAIL"
	Login          = "LOGIN"
)
