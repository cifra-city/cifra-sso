package entities

type OperationType string

const (
	ChangeEmail    = "CHANGE_EMAIL"
	ChangePassword = "CHANGE_PASS"
	ChangeUsername = "CHANGE_USERNAME"
	ConfirmEmail   = "CONFIRM_EMAIL"
	Login          = "LOGIN"
)
