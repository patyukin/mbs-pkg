package kafka

import (
	"time"
)

const (
	maxRetries = 3
	maxBackoff = time.Second * 2
)

const (
	DeadLetterTopic = "dead_letter_topic"
)

const (
	LogReportTopic = "log_report_topic"
)
