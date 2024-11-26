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

type PaymentRequest struct {
	PaymentID string `json:"payment_id"`
}

type Transaction struct {
	ID          string `json:"id"`
	PaymentID   string `json:"payment_id"`
	AccountID   string `json:"account_id"`
	Type        string `json:"type"`
	Amount      int64  `json:"amount"`
	Currency    string `json:"currency"`
	Description string `json:"description"`
	Status      string `json:"status"`
	SendStatus  string `json:"send_status"`
	CreatedAt   string `json:"created_at"`
}

type TransactionSendStatus struct {
	ID         string `json:"id"`
	SendStatus string `json:"send_status"`
}

type CreditCreated struct {
	AccountID string `json:"account_id"`
	Amount    int64  `json:"amount"`
}

type CreditPayment struct {
	PaymentScheduleID string `json:"payment_schedule_id"`
	AccountID         string `json:"account_id"`
	Amount            int64  `json:"amount"`
}

type CreditPaymentSolution struct {
	PaymentScheduleID string `json:"payment_schedule_id"`
	Status            string `json:"status"`
}
