package kafka

import (
	"time"
)

const (
	maxRetries = 3
	maxBackoff = time.Second * 2
)

const (
	DeadLetterTopic      = "dead_letter_topic"
	LogReportTopic       = "log_report_topic"
	ReportTopic          = "report_topic"
	CreditCreatedTopic   = "credit_created_topic"
	CreditPaymentsTopic  = "credit_payments_topic"
	PaymentRequestsTopic = "payment_requests_topic"
)
