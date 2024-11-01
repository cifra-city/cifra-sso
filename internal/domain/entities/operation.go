package entities

type OperationType string

const (
	ChangeEmail    OperationType = "CHANGE_EMAIL"
	ChangePassword OperationType = "CHANGE_PASS"
	ChangeUsername OperationType = "CHANGE_USERNAME"
	Login          OperationType = "LOGIN"
)
