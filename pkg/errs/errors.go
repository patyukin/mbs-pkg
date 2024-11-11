package errs

import (
	"errors"
	"fmt"
	"github.com/patyukin/mbs-pkg/pkg/proto/error_v1"
	"google.golang.org/grpc/codes"
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
			Code:        int32(codes.NotFound),
			Message:     "User not found",
			Description: "The user with the given ID was not found",
		}
	case errors.Is(err, ErrInvalidRequest):
		return &error_v1.ErrorResponse{
			Code:        int32(codes.InvalidArgument),
			Message:     "Invalid request",
			Description: "The request parameters are invalid",
		}
	case errors.Is(err, ErrDatabaseError):
		return &error_v1.ErrorResponse{
			Code:        int32(codes.Internal),
			Message:     "Database error",
			Description: "There was an error connecting to the database",
		}
	default:
		return &error_v1.ErrorResponse{
			Code:        int32(codes.Unknown),
			Message:     "Unknown error",
			Description: err.Error(),
		}
	}
}
