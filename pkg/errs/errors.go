package errs

import (
	"errors"
	"fmt"
	"github.com/patyukin/mbs-pkg/pkg/proto/error_v1"
	"net/http"
)

var (
	ErrUserNotFound   = fmt.Errorf("user not found")
	ErrInvalidRequest = fmt.Errorf("invalid request")
	ErrDatabaseError  = fmt.Errorf("database connection error")
)

// ToErrorResponse преобразует предопределенные ошибки в error_v1.ErrorResponse
func ToErrorResponse(err error) *error_v1.ErrorResponse {
	switch {
	case errors.Is(err, ErrUserNotFound):
		return &error_v1.ErrorResponse{
			Code:        http.StatusNotFound,
			Message:     "User not found",
			Description: err.Error(),
		}
	case errors.Is(err, ErrInvalidRequest):
		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     "Invalid request",
			Description: err.Error(),
		}
	case errors.Is(err, ErrDatabaseError):
		return &error_v1.ErrorResponse{
			Code:        http.StatusInternalServerError,
			Message:     "Internal Server Error",
			Description: err.Error(),
		}
	default:
		return &error_v1.ErrorResponse{
			Code:        http.StatusInternalServerError,
			Message:     "Internal Server Error",
			Description: err.Error(),
		}
	}
}
