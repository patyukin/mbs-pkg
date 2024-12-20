// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: logger.proto

package logger_v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on LogReportRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *LogReportRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LogReportRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// LogReportRequestMultiError, or nil if none found.
func (m *LogReportRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *LogReportRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for StartTime

	// no validation rules for EndTime

	// no validation rules for ServiceName

	if len(errors) > 0 {
		return LogReportRequestMultiError(errors)
	}

	return nil
}

// LogReportRequestMultiError is an error wrapping multiple validation errors
// returned by LogReportRequest.ValidateAll() if the designated constraints
// aren't met.
type LogReportRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LogReportRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LogReportRequestMultiError) AllErrors() []error { return m }

// LogReportRequestValidationError is the validation error returned by
// LogReportRequest.Validate if the designated constraints aren't met.
type LogReportRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LogReportRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LogReportRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LogReportRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LogReportRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LogReportRequestValidationError) ErrorName() string { return "LogReportRequestValidationError" }

// Error satisfies the builtin error interface
func (e LogReportRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLogReportRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LogReportRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LogReportRequestValidationError{}

// Validate checks the field values on LogReportResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *LogReportResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LogReportResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// LogReportResponseMultiError, or nil if none found.
func (m *LogReportResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *LogReportResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Message

	if all {
		switch v := interface{}(m.GetError()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, LogReportResponseValidationError{
					field:  "Error",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, LogReportResponseValidationError{
					field:  "Error",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetError()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LogReportResponseValidationError{
				field:  "Error",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return LogReportResponseMultiError(errors)
	}

	return nil
}

// LogReportResponseMultiError is an error wrapping multiple validation errors
// returned by LogReportResponse.ValidateAll() if the designated constraints
// aren't met.
type LogReportResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LogReportResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LogReportResponseMultiError) AllErrors() []error { return m }

// LogReportResponseValidationError is the validation error returned by
// LogReportResponse.Validate if the designated constraints aren't met.
type LogReportResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LogReportResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LogReportResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LogReportResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LogReportResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LogReportResponseValidationError) ErrorName() string {
	return "LogReportResponseValidationError"
}

// Error satisfies the builtin error interface
func (e LogReportResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLogReportResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LogReportResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LogReportResponseValidationError{}
