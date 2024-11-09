package errors

import (
	"fmt"
	"google.golang.org/grpc/codes"
)

var (
	ErrUserNotFound   = fmt.Errorf("user not found")
	ErrInvalidRequest = fmt.Errorf("invalid request")
	ErrDatabaseError  = fmt.Errorf("database connection error")
)

// NewError - Функция для создания детализированной ошибки с кодом и сообщением
func NewError(code codes.Code, shortMessage, fullDescription string) *CustomError {
	return &CustomError{
		Code:        code,
		Message:     shortMessage,
		Description: fullDescription,
	}
}

// CustomError - Кастомная структура ошибки
type CustomError struct {
	Code        codes.Code
	Message     string
	Description string
}

func (e *CustomError) Error() string {
	return e.Description
}
