package kafka

import (
	"time"
)

const (
	maxRetries     = 5
	maxBackoff     = time.Second * 5
	maxGoroutines  = 100
	initialBackoff = 100 * time.Millisecond
)

const (
	DeadLetterTopic             = "dead_letter"
	LogReportTopic              = "log_report"
	TransactionReportTopic      = "transaction_report"
	CreditCreatedTopic          = "credit_created"
	CreditPaymentsTopic         = "credit_payments"
	PaymentRequestTopic         = "payment_request"
	CreditPaymentsSolutionTopic = "credit_payments_solution"
)
