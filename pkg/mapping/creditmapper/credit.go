package creditmapper

import (
	"fmt"
	"github.com/patyukin/mbs-pkg/pkg/proto/credit_v1"
)

// Маппинг для CreditStatus
var creditStatusToString = map[credit_v1.CreditStatus]string{
	credit_v1.CreditStatus_UNKNOWN_CREDIT_STATUS: "UNKNOWN_CREDIT_STATUS",
	credit_v1.CreditStatus_ACTIVE:                "ACTIVE",
	credit_v1.CreditStatus_CLOSED:                "CLOSED",
}

var stringToCreditStatus = map[string]credit_v1.CreditStatus{
	"UNKNOWN_CREDIT_STATUS": credit_v1.CreditStatus_UNKNOWN_CREDIT_STATUS,
	"ACTIVE":                credit_v1.CreditStatus_ACTIVE,
	"CLOSED":                credit_v1.CreditStatus_CLOSED,
}

// EnumToStringCreditStatus converts CreditStatus enum to string
func EnumToStringCreditStatus(status credit_v1.CreditStatus) (string, error) {
	str, ok := creditStatusToString[status]
	if !ok {
		return "", fmt.Errorf("invalid CreditStatus")
	}
	return str, nil
}

// StringToEnumCreditStatus converts string to CreditStatus enum
func StringToEnumCreditStatus(status string) (credit_v1.CreditStatus, error) {
	enum, ok := stringToCreditStatus[status]
	if !ok {
		return 0, fmt.Errorf("invalid CreditStatus")
	}

	return enum, nil
}

// ValidateStringCreditStatus validates CreditStatus string
func ValidateStringCreditStatus(status string) error {
	if _, ok := stringToCreditStatus[status]; !ok {
		return fmt.Errorf("invalid CreditStatus")
	}

	return nil
}
