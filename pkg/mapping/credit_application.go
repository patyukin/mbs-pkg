package mapping

import (
	"errors"
	"github.com/patyukin/mbs-pkg/pkg/proto/credit_v1"
)

// Маппинг для CreditApplicationStatus
var creditApplicationStatusToString = map[credit_v1.CreditApplicationStatus]string{
	credit_v1.CreditApplicationStatus_UNKNOWN_STATUS: "UNKNOWN_STATUS",
	credit_v1.CreditApplicationStatus_PENDING:        "PENDING",
	credit_v1.CreditApplicationStatus_APPROVED:       "APPROVED",
	credit_v1.CreditApplicationStatus_REJECTED:       "REJECTED",
}

var stringToCreditApplicationStatus = map[string]credit_v1.CreditApplicationStatus{
	"UNKNOWN_STATUS": credit_v1.CreditApplicationStatus_UNKNOWN_STATUS,
	"PENDING":        credit_v1.CreditApplicationStatus_PENDING,
	"APPROVED":       credit_v1.CreditApplicationStatus_APPROVED,
	"REJECTED":       credit_v1.CreditApplicationStatus_REJECTED,
}

// EnumToStringCreditApplicationStatus converts CreditApplicationStatus enum to string
func EnumToStringCreditApplicationStatus(status credit_v1.CreditApplicationStatus) (string, error) {
	str, ok := creditApplicationStatusToString[status]
	if !ok {
		return "", errors.New("invalid CreditApplicationStatus enum value")
	}

	return str, nil
}

// StringToEnumCreditApplicationStatus converts string to CreditApplicationStatus enum
func StringToEnumCreditApplicationStatus(status string) (credit_v1.CreditApplicationStatus, error) {
	enum, ok := stringToCreditApplicationStatus[status]
	if !ok {
		return 0, errors.New("invalid CreditApplicationStatus string value")
	}

	return enum, nil
}
