package creditmapper

import (
	"fmt"

	"github.com/patyukin/mbs-pkg/pkg/proto/credit_v1"
)

const (
	CreditApplicationStatusDraft      = "DRAFT"
	CreditApplicationStatusPending    = "PENDING"
	CreditApplicationStatusApproved   = "APPROVED"
	CreditApplicationStatusRejected   = "REJECTED"
	CreditApplicationStatusArchived   = "ARCHIVED"
	CreditApplicationStatusProcessing = "PROCESSING"
)

// Маппинг для CreditApplicationStatus
var creditApplicationStatusToString = map[credit_v1.CreditApplicationStatus]string{
	credit_v1.CreditApplicationStatus_UNKNOWN_STATUS: "UNKNOWN_STATUS",
	credit_v1.CreditApplicationStatus_PENDING:        "PENDING",
	credit_v1.CreditApplicationStatus_APPROVED:       "APPROVED",
	credit_v1.CreditApplicationStatus_REJECTED:       "REJECTED",
	credit_v1.CreditApplicationStatus_ARCHIVED:       "ARCHIVED",
	credit_v1.CreditApplicationStatus_PROCESSING:     "PROCESSING",
	credit_v1.CreditApplicationStatus_DRAFT:          "DRAFT",
}

var stringToCreditApplicationStatus = map[string]credit_v1.CreditApplicationStatus{
	"UNKNOWN_STATUS": credit_v1.CreditApplicationStatus_UNKNOWN_STATUS,
	"PENDING":        credit_v1.CreditApplicationStatus_PENDING,
	"APPROVED":       credit_v1.CreditApplicationStatus_APPROVED,
	"REJECTED":       credit_v1.CreditApplicationStatus_REJECTED,
	"ARCHIVED":       credit_v1.CreditApplicationStatus_ARCHIVED,
	"PROCESSING":     credit_v1.CreditApplicationStatus_PROCESSING,
	"DRAFT":          credit_v1.CreditApplicationStatus_DRAFT,
}

// EnumToStringCreditApplicationStatus converts CreditApplicationStatus enum to string
func EnumToStringCreditApplicationStatus(status credit_v1.CreditApplicationStatus) (string, error) {
	str, ok := creditApplicationStatusToString[status]
	if !ok {
		return "", fmt.Errorf("invalid CreditApplicationStatus")
	}

	return str, nil
}

// StringToEnumCreditApplicationStatus converts string to CreditApplicationStatus enum
func StringToEnumCreditApplicationStatus(status string) (credit_v1.CreditApplicationStatus, error) {
	enum, ok := stringToCreditApplicationStatus[status]
	if !ok {
		return 0, fmt.Errorf("invalid CreditApplicationStatus")
	}

	return enum, nil
}

// ValidateStringCreditApplicationStatus validates CreditApplicationStatus string
func ValidateStringCreditApplicationStatus(status string) error {
	if _, ok := stringToCreditApplicationStatus[status]; !ok {
		return fmt.Errorf("invalid CreditApplicationStatus")
	}

	return nil
}
