// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: admin/v1/bot.proto

package v1

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

// Validate checks the field values on Auth with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Auth) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Auth with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in AuthMultiError, or nil if none found.
func (m *Auth) ValidateAll() error {
	return m.validate(true)
}

func (m *Auth) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return AuthMultiError(errors)
	}

	return nil
}

// AuthMultiError is an error wrapping multiple validation errors returned by
// Auth.ValidateAll() if the designated constraints aren't met.
type AuthMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AuthMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AuthMultiError) AllErrors() []error { return m }

// AuthValidationError is the validation error returned by Auth.Validate if the
// designated constraints aren't met.
type AuthValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AuthValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AuthValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AuthValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AuthValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AuthValidationError) ErrorName() string { return "AuthValidationError" }

// Error satisfies the builtin error interface
func (e AuthValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAuth.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AuthValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AuthValidationError{}

// Validate checks the field values on GetSentence with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetSentence) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetSentence with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetSentenceMultiError, or
// nil if none found.
func (m *GetSentence) ValidateAll() error {
	return m.validate(true)
}

func (m *GetSentence) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return GetSentenceMultiError(errors)
	}

	return nil
}

// GetSentenceMultiError is an error wrapping multiple validation errors
// returned by GetSentence.ValidateAll() if the designated constraints aren't met.
type GetSentenceMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetSentenceMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetSentenceMultiError) AllErrors() []error { return m }

// GetSentenceValidationError is the validation error returned by
// GetSentence.Validate if the designated constraints aren't met.
type GetSentenceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetSentenceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetSentenceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetSentenceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetSentenceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetSentenceValidationError) ErrorName() string { return "GetSentenceValidationError" }

// Error satisfies the builtin error interface
func (e GetSentenceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetSentence.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetSentenceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetSentenceValidationError{}

// Validate checks the field values on User with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *User) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on User with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in UserMultiError, or nil if none found.
func (m *User) ValidateAll() error {
	return m.validate(true)
}

func (m *User) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for TelegramId

	if len(errors) > 0 {
		return UserMultiError(errors)
	}

	return nil
}

// UserMultiError is an error wrapping multiple validation errors returned by
// User.ValidateAll() if the designated constraints aren't met.
type UserMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserMultiError) AllErrors() []error { return m }

// UserValidationError is the validation error returned by User.Validate if the
// designated constraints aren't met.
type UserValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserValidationError) ErrorName() string { return "UserValidationError" }

// Error satisfies the builtin error interface
func (e UserValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUser.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserValidationError{}

// Validate checks the field values on Auth_Request with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Auth_Request) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Auth_Request with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in Auth_RequestMultiError, or
// nil if none found.
func (m *Auth_Request) ValidateAll() error {
	return m.validate(true)
}

func (m *Auth_Request) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetUser()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, Auth_RequestValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, Auth_RequestValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUser()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Auth_RequestValidationError{
				field:  "User",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return Auth_RequestMultiError(errors)
	}

	return nil
}

// Auth_RequestMultiError is an error wrapping multiple validation errors
// returned by Auth_Request.ValidateAll() if the designated constraints aren't met.
type Auth_RequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Auth_RequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Auth_RequestMultiError) AllErrors() []error { return m }

// Auth_RequestValidationError is the validation error returned by
// Auth_Request.Validate if the designated constraints aren't met.
type Auth_RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Auth_RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Auth_RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Auth_RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Auth_RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Auth_RequestValidationError) ErrorName() string { return "Auth_RequestValidationError" }

// Error satisfies the builtin error interface
func (e Auth_RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAuth_Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Auth_RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Auth_RequestValidationError{}

// Validate checks the field values on Auth_Response with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Auth_Response) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Auth_Response with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in Auth_ResponseMultiError, or
// nil if none found.
func (m *Auth_Response) ValidateAll() error {
	return m.validate(true)
}

func (m *Auth_Response) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return Auth_ResponseMultiError(errors)
	}

	return nil
}

// Auth_ResponseMultiError is an error wrapping multiple validation errors
// returned by Auth_Response.ValidateAll() if the designated constraints
// aren't met.
type Auth_ResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Auth_ResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Auth_ResponseMultiError) AllErrors() []error { return m }

// Auth_ResponseValidationError is the validation error returned by
// Auth_Response.Validate if the designated constraints aren't met.
type Auth_ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Auth_ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Auth_ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Auth_ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Auth_ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Auth_ResponseValidationError) ErrorName() string { return "Auth_ResponseValidationError" }

// Error satisfies the builtin error interface
func (e Auth_ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAuth_Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Auth_ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Auth_ResponseValidationError{}

// Validate checks the field values on GetSentence_Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetSentence_Request) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetSentence_Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetSentence_RequestMultiError, or nil if none found.
func (m *GetSentence_Request) ValidateAll() error {
	return m.validate(true)
}

func (m *GetSentence_Request) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return GetSentence_RequestMultiError(errors)
	}

	return nil
}

// GetSentence_RequestMultiError is an error wrapping multiple validation
// errors returned by GetSentence_Request.ValidateAll() if the designated
// constraints aren't met.
type GetSentence_RequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetSentence_RequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetSentence_RequestMultiError) AllErrors() []error { return m }

// GetSentence_RequestValidationError is the validation error returned by
// GetSentence_Request.Validate if the designated constraints aren't met.
type GetSentence_RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetSentence_RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetSentence_RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetSentence_RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetSentence_RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetSentence_RequestValidationError) ErrorName() string {
	return "GetSentence_RequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetSentence_RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetSentence_Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetSentence_RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetSentence_RequestValidationError{}

// Validate checks the field values on GetSentence_Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetSentence_Response) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetSentence_Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetSentence_ResponseMultiError, or nil if none found.
func (m *GetSentence_Response) ValidateAll() error {
	return m.validate(true)
}

func (m *GetSentence_Response) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Sentence

	if len(errors) > 0 {
		return GetSentence_ResponseMultiError(errors)
	}

	return nil
}

// GetSentence_ResponseMultiError is an error wrapping multiple validation
// errors returned by GetSentence_Response.ValidateAll() if the designated
// constraints aren't met.
type GetSentence_ResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetSentence_ResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetSentence_ResponseMultiError) AllErrors() []error { return m }

// GetSentence_ResponseValidationError is the validation error returned by
// GetSentence_Response.Validate if the designated constraints aren't met.
type GetSentence_ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetSentence_ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetSentence_ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetSentence_ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetSentence_ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetSentence_ResponseValidationError) ErrorName() string {
	return "GetSentence_ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetSentence_ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetSentence_Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetSentence_ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetSentence_ResponseValidationError{}
