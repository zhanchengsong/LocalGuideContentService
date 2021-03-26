package model

type ServiceError struct {
	Source   string
	Message  string
	HttpCode int
}
