package model

type AuthSignInCode struct {
	Code   string `json:"code"`
	ChatID int64  `json:"chat_id"`
}
