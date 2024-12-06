package paymentmapper

import (
	"fmt"

	"github.com/patyukin/mbs-pkg/pkg/proto/payment_v1"
)

const (
	TransactionTypeDebit  = "DEBIT"
	TransactionTypeCredit = "CREDIT"
)

// Маппинг для TransactionType
var paymentTransactionTypeToString = map[payment_v1.TransactionType]string{
	payment_v1.TransactionType_UNKNOWN_TRANSACTION_TYPE: "UNKNOWN_TRANSACTION_TYPE",
	payment_v1.TransactionType_DEBIT:                    "DEBIT",
	payment_v1.TransactionType_CREDIT:                   "CREDIT",
}

var stringToPaymentTransactionType = map[string]payment_v1.TransactionType{
	"UNKNOWN_TRANSACTION_TYPE": payment_v1.TransactionType_UNKNOWN_TRANSACTION_TYPE,
	"DEBIT":                    payment_v1.TransactionType_DEBIT,
	"CREDIT":                   payment_v1.TransactionType_CREDIT,
}

// EnumToStringTransactionType converts TransactionType enum to string
func EnumToStringTransactionType(status payment_v1.TransactionType) (string, error) {
	str, ok := paymentTransactionTypeToString[status]
	if !ok {
		return "", fmt.Errorf("invalid TransactionType")
	}

	return str, nil
}

// StringToEnumTransactionType converts string to TransactionType enum
func StringToEnumTransactionType(status string) (payment_v1.TransactionType, error) {
	enum, ok := stringToPaymentTransactionType[status]
	if !ok {
		return 0, fmt.Errorf("invalid TransactionType")
	}

	return enum, nil
}

// ValidateStringTransactionType validates TransactionType string
func ValidateStringTransactionType(status string) error {
	if _, ok := stringToPaymentTransactionType[status]; !ok {
		return fmt.Errorf("invalid TransactionType")
	}

	return nil
}
