package utils

import (
	"fmt"
	"time"
)

func ValidateDate(dateStr string) (bool, error) {
	layout := "2006-01-02"
	_, err := time.Parse(layout, dateStr)
	if err != nil {
		return false, fmt.Errorf("некорректная дата: %v", err)
	}

	return true, nil
}
