package errs

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/patyukin/mbs-pkg/pkg/proto/error_v1"
)

var (
	ErrUserNotFound           = fmt.Errorf("user not found")
	ErrUserExists             = fmt.Errorf("user exists")
	ErrInvalidRequest         = fmt.Errorf("invalid request")
	ErrDatabaseError          = fmt.Errorf("database connection error")
	ErrTelegramChatIDNotFound = errors.New("telegram chat id not found")
	ErrInvalidCode            = errors.New("invalid code")
)

// ToErrorResponse преобразует предопределенные ошибки в error_v1.ErrorResponse
func ToErrorResponse(err error) *error_v1.ErrorResponse {
	switch {
	case errors.Is(err, ErrUserExists):
		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     "User exists",
			Description: err.Error(),
		}
	case errors.Is(err, ErrUserNotFound):
		return &error_v1.ErrorResponse{
			Code:        http.StatusNotFound,
			Message:     "User not found",
			Description: err.Error(),
		}
	case errors.Is(err, sql.ErrNoRows):
		return &error_v1.ErrorResponse{
			Code:        http.StatusNotFound,
			Message:     "Not found",
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
