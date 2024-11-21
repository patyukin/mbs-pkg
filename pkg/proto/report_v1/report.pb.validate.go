// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: report.proto

package report_v1

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

// Validate checks the field values on GetUserReportRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetUserReportRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetUserReportRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetUserReportRequestMultiError, or nil if none found.
func (m *GetUserReportRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetUserReportRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for UserId

	// no validation rules for StartDate

	// no validation rules for EndDate

	if len(errors) > 0 {
		return GetUserReportRequestMultiError(errors)
	}

	return nil
}

// GetUserReportRequestMultiError is an error wrapping multiple validation
// errors returned by GetUserReportRequest.ValidateAll() if the designated
// constraints aren't met.
type GetUserReportRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetUserReportRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetUserReportRequestMultiError) AllErrors() []error { return m }

// GetUserReportRequestValidationError is the validation error returned by
// GetUserReportRequest.Validate if the designated constraints aren't met.
type GetUserReportRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUserReportRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUserReportRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUserReportRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUserReportRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUserReportRequestValidationError) ErrorName() string {
	return "GetUserReportRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetUserReportRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUserReportRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUserReportRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUserReportRequestValidationError{}

// Validate checks the field values on GetUserReportResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetUserReportResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetUserReportResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetUserReportResponseMultiError, or nil if none found.
func (m *GetUserReportResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetUserReportResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Message

	if all {
		switch v := interface{}(m.GetError()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetUserReportResponseValidationError{
					field:  "Error",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetUserReportResponseValidationError{
					field:  "Error",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetError()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetUserReportResponseValidationError{
				field:  "Error",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetUserReportResponseMultiError(errors)
	}

	return nil
}

// GetUserReportResponseMultiError is an error wrapping multiple validation
// errors returned by GetUserReportResponse.ValidateAll() if the designated
// constraints aren't met.
type GetUserReportResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetUserReportResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetUserReportResponseMultiError) AllErrors() []error { return m }

// GetUserReportResponseValidationError is the validation error returned by
// GetUserReportResponse.Validate if the designated constraints aren't met.
type GetUserReportResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUserReportResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUserReportResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUserReportResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUserReportResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUserReportResponseValidationError) ErrorName() string {
	return "GetUserReportResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetUserReportResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUserReportResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUserReportResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUserReportResponseValidationError{}