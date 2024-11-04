package model

type AuthSignInCode struct {
	Code   string `json:"code"`
	ChatID int64  `json:"chat_id"`
}

type AuthSignUpConfirmCode struct {
	Code              string `json:"code"`
	ChatID            int64  `json:"chat_id"`
	UserTelegramLogin string `json:"user_telegram_login"`
	UserTelegramID    int64  `json:"user_telegram_id"`
}
