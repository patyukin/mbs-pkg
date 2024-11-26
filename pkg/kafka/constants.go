package kafka

import (
	"time"
)

const (
	maxRetries = 3
	maxBackoff = time.Second * 2
)

const (
	DeadLetterTopic     = "dead_letter"
	LogReportTopic      = "log_report"
	TransactionsTopic   = "transactions"
	CreditCreatedTopic  = "credit_created"
	CreditPaymentsTopic = "credit_payments"
	PaymentRequestTopic = "payment_request"
)
