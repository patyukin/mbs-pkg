package rabbitmq

import (
	"time"
)

type AuthSignInCode struct {
	Code       string    `json:"code"`
	UserUUID   string    `json:"user_uuid"`
	Expiration time.Time `json:"expiration"`
}
