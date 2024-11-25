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

type AuthSignUpResultMessage struct {
	ChatID  int64  `json:"chat_id"`
	Message string `json:"message"`
}

type SimpleTelegramMessage struct {
	ChatID  int64  `json:"chat_id"`
	Message string `json:"message"`
}

type PaymentStatusChangedMessage struct {
	PaymentID string `json:"payment_id"`
	Status    string `json:"status"`
}

type CreditCreatedMessage struct {
	AccountID string `json:"account_id"`
	Amount    int64  `json:"amount"`
}

type CreditPaymentMessage struct {
	PaymentScheduleID string `json:"payment_schedule_id"`
	AccountID         string `json:"account_id"`
	Amount            int64  `json:"amount"`
}

type CreditPaymentSolutionMessage struct {
	PaymentScheduleID string `json:"payment_schedule_id"`
	Status            string `json:"status"`
}
