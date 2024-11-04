package rabbitmq

import (
	"time"
)

type AuthSignInCode struct {
	Code       string    `json:"code"`
	ChatID     int64     `json:"chat_id"`
	UserUUID   string    `json:"user_uuid"`
	Expiration time.Time `json:"expiration"`
}
