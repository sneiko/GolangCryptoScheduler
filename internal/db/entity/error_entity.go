package entity

type ErrorEntity struct {
	Message string
}

func NewError(message string) ErrorEntity {
	return ErrorEntity{Message: message}
}
