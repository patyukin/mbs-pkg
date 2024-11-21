package paymentmapper

import (
	"fmt"
	"github.com/patyukin/mbs-pkg/pkg/proto/payment_v1"
)

// Маппинг для PaymentStatus
var paymentStatusToString = map[payment_v1.PaymentStatus]string{
	payment_v1.PaymentStatus_UNKNOWN_PAYMENT_STATUS: "UNKNOWN_PAYMENT_STATUS",
	payment_v1.PaymentStatus_DRAFT:                  "DRAFT",
	payment_v1.PaymentStatus_PENDING:                "PENDING",
	payment_v1.PaymentStatus_COMPLETED:              "COMPLETED",
	payment_v1.PaymentStatus_FAILED:                 "FAILED",
}

var stringToPaymentStatus = map[string]payment_v1.PaymentStatus{
	"UNKNOWN_PAYMENT_STATUS": payment_v1.PaymentStatus_UNKNOWN_PAYMENT_STATUS,
	"DRAFT":                  payment_v1.PaymentStatus_DRAFT,
	"PENDING":                payment_v1.PaymentStatus_PENDING,
	"COMPLETED":              payment_v1.PaymentStatus_COMPLETED,
	"FAILED":                 payment_v1.PaymentStatus_FAILED,
}

// EnumToStringPaymentStatus converts PaymentStatus enum to string
func EnumToStringPaymentStatus(status payment_v1.PaymentStatus) (string, error) {
	str, ok := paymentStatusToString[status]
	if !ok {
		return "", fmt.Errorf("invalid PaymentStatus")
	}

	return str, nil
}

// StringToEnumPaymentStatus converts string to PaymentStatus enum
func StringToEnumPaymentStatus(status string) (payment_v1.PaymentStatus, error) {
	enum, ok := stringToPaymentStatus[status]
	if !ok {
		return 0, fmt.Errorf("invalid PaymentStatus")
	}

	return enum, nil
}

// ValidateStringPaymentStatus validates PaymentStatus string
func ValidateStringPaymentStatus(status string) error {
	if _, ok := stringToPaymentStatus[status]; !ok {
		return fmt.Errorf("invalid PaymentStatus")
	}

	return nil
}
