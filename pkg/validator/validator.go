package validator

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

func ValidateDate(dateStr string) (bool, error) {
	layout := "2006-01-02"
	_, err := time.Parse(layout, dateStr)
	if err != nil {
		return false, fmt.Errorf("invalid date format: %s, expected format is YYYY-MM-DD", dateStr)
	}

	return true, nil
}

func ValidateUUID(uuidStr string) (bool, error) {
	_, err := uuid.Parse(uuidStr)
	if err != nil {
		return false, fmt.Errorf("invalid format: %s", uuidStr)
	}

	return true, nil
}
