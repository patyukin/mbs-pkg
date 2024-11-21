package mapping

import (
	"fmt"
	"github.com/patyukin/mbs-pkg/pkg/proto/credit_v1"
)

// Маппинг для PaymentStatus
var paymentStatusToString = map[credit_v1.PaymentStatus]string{
	credit_v1.PaymentStatus_UNKNOWN_PAYMENT_STATUS: "UNKNOWN_PAYMENT_STATUS",
	credit_v1.PaymentStatus_SCHEDULED:              "SCHEDULED",
	credit_v1.PaymentStatus_PAID:                   "PAID",
	credit_v1.PaymentStatus_MISSED:                 "MISSED",
	credit_v1.PaymentStatus_OVERPAID:               "OVERPAID",
	credit_v1.PaymentStatus_REFUNDED:               "REFUNDED",
}

var stringToPaymentStatus = map[string]credit_v1.PaymentStatus{
	"UNKNOWN_PAYMENT_STATUS": credit_v1.PaymentStatus_UNKNOWN_PAYMENT_STATUS,
	"SCHEDULED":              credit_v1.PaymentStatus_SCHEDULED,
	"PAID":                   credit_v1.PaymentStatus_PAID,
	"MISSED":                 credit_v1.PaymentStatus_MISSED,
	"OVERPAID":               credit_v1.PaymentStatus_OVERPAID,
	"REFUNDED":               credit_v1.PaymentStatus_REFUNDED,
}

// EnumToStringPaymentStatus converts PaymentStatus enum to string
func EnumToStringPaymentStatus(status credit_v1.PaymentStatus) (string, error) {
	str, ok := paymentStatusToString[status]
	if !ok {
		return "", fmt.Errorf("invalid PaymentStatus enum value")
	}

	return str, nil
}

// StringToEnumPaymentStatus converts string to PaymentStatus enum
func StringToEnumPaymentStatus(status string) (credit_v1.PaymentStatus, error) {
	enum, ok := stringToPaymentStatus[status]
	if !ok {
		return 0, fmt.Errorf("invalid PaymentStatus string value")
	}

	return enum, nil
}

// ValidateStringPaymentStatus validates PaymentStatus string
func ValidateStringPaymentStatus(status string) error {
	if _, ok := stringToPaymentStatus[status]; !ok {
		return fmt.Errorf("invalid PaymentStatus value")
	}

	return nil
}
