// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: editor/v1/round_question.proto

package editorv1

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

// Validate checks the field values on RoundQuestion with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *RoundQuestion) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RoundQuestion with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in RoundQuestionMultiError, or
// nil if none found.
func (m *RoundQuestion) ValidateAll() error {
	return m.validate(true)
}

func (m *RoundQuestion) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for RoundId

	// no validation rules for TopicId

	if all {
		switch v := interface{}(m.GetQuestion()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RoundQuestionValidationError{
					field:  "Question",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RoundQuestionValidationError{
					field:  "Question",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetQuestion()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RoundQuestionValidationError{
				field:  "Question",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for QuestionType

	// no validation rules for QuestionCost

	if all {
		switch v := interface{}(m.GetAnswer()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RoundQuestionValidationError{
					field:  "Answer",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RoundQuestionValidationError{
					field:  "Answer",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAnswer()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RoundQuestionValidationError{
				field:  "Answer",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for AnswerTime

	// no validation rules for HostComment

	// no validation rules for SecretTopic

	// no validation rules for SecretCost

	// no validation rules for TransferType

	// no validation rules for IsKeepable

	if len(errors) > 0 {
		return RoundQuestionMultiError(errors)
	}

	return nil
}

// RoundQuestionMultiError is an error wrapping multiple validation errors
// returned by RoundQuestion.ValidateAll() if the designated constraints
// aren't met.
type RoundQuestionMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RoundQuestionMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RoundQuestionMultiError) AllErrors() []error { return m }

// RoundQuestionValidationError is the validation error returned by
// RoundQuestion.Validate if the designated constraints aren't met.
type RoundQuestionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RoundQuestionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RoundQuestionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RoundQuestionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RoundQuestionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RoundQuestionValidationError) ErrorName() string { return "RoundQuestionValidationError" }

// Error satisfies the builtin error interface
func (e RoundQuestionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRoundQuestion.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RoundQuestionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RoundQuestionValidationError{}

// Validate checks the field values on CreateRoundQuestionRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateRoundQuestionRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateRoundQuestionRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateRoundQuestionRequestMultiError, or nil if none found.
func (m *CreateRoundQuestionRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateRoundQuestionRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for QuestionId

	// no validation rules for TopicId

	// no validation rules for RoundId

	if _, ok := _CreateRoundQuestionRequest_QuestionType_InLookup[m.GetQuestionType()]; !ok {
		err := CreateRoundQuestionRequestValidationError{
			field:  "QuestionType",
			reason: "value must be in list [STANDARD SAFE SECRET SUPER_SECRET AUCTION]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetQuestionCost() < 1 {
		err := CreateRoundQuestionRequestValidationError{
			field:  "QuestionCost",
			reason: "value must be greater than or equal to 1",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetAnswerTime() == nil {
		err := CreateRoundQuestionRequestValidationError{
			field:  "AnswerTime",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if d := m.GetAnswerTime(); d != nil {
		dur, err := d.AsDuration(), d.CheckValid()
		if err != nil {
			err = CreateRoundQuestionRequestValidationError{
				field:  "AnswerTime",
				reason: "value is not a valid duration",
				cause:  err,
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		} else {

			lte := time.Duration(60*time.Second + 0*time.Nanosecond)
			gte := time.Duration(5*time.Second + 0*time.Nanosecond)

			if dur < gte || dur > lte {
				err := CreateRoundQuestionRequestValidationError{
					field:  "AnswerTime",
					reason: "value must be inside range [5s, 1m0s]",
				}
				if !all {
					return err
				}
				errors = append(errors, err)
			}

		}
	}

	// no validation rules for HostComment

	// no validation rules for SecretTopic

	// no validation rules for SecretCost

	// no validation rules for IsKeepable

	if _, ok := _CreateRoundQuestionRequest_TransferType_InLookup[m.GetTransferType()]; !ok {
		err := CreateRoundQuestionRequestValidationError{
			field:  "TransferType",
			reason: "value must be in list [TRANSFER_TYPE_UNSPECIFIED BEFORE AFTER NEVER]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return CreateRoundQuestionRequestMultiError(errors)
	}

	return nil
}

// CreateRoundQuestionRequestMultiError is an error wrapping multiple
// validation errors returned by CreateRoundQuestionRequest.ValidateAll() if
// the designated constraints aren't met.
type CreateRoundQuestionRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateRoundQuestionRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateRoundQuestionRequestMultiError) AllErrors() []error { return m }

// CreateRoundQuestionRequestValidationError is the validation error returned
// by CreateRoundQuestionRequest.Validate if the designated constraints aren't met.
type CreateRoundQuestionRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateRoundQuestionRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateRoundQuestionRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateRoundQuestionRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateRoundQuestionRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateRoundQuestionRequestValidationError) ErrorName() string {
	return "CreateRoundQuestionRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateRoundQuestionRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateRoundQuestionRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateRoundQuestionRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateRoundQuestionRequestValidationError{}

var _CreateRoundQuestionRequest_QuestionType_InLookup = map[RoundQuestionType]struct{}{
	1: {},
	2: {},
	3: {},
	4: {},
	5: {},
}

var _CreateRoundQuestionRequest_TransferType_InLookup = map[TransferType]struct{}{
	0: {},
	1: {},
	2: {},
	3: {},
}

// Validate checks the field values on CreateRoundQuestionResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateRoundQuestionResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateRoundQuestionResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateRoundQuestionResponseMultiError, or nil if none found.
func (m *CreateRoundQuestionResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateRoundQuestionResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for RoundQuestionId

	if len(errors) > 0 {
		return CreateRoundQuestionResponseMultiError(errors)
	}

	return nil
}

// CreateRoundQuestionResponseMultiError is an error wrapping multiple
// validation errors returned by CreateRoundQuestionResponse.ValidateAll() if
// the designated constraints aren't met.
type CreateRoundQuestionResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateRoundQuestionResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateRoundQuestionResponseMultiError) AllErrors() []error { return m }

// CreateRoundQuestionResponseValidationError is the validation error returned
// by CreateRoundQuestionResponse.Validate if the designated constraints
// aren't met.
type CreateRoundQuestionResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateRoundQuestionResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateRoundQuestionResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateRoundQuestionResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateRoundQuestionResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateRoundQuestionResponseValidationError) ErrorName() string {
	return "CreateRoundQuestionResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateRoundQuestionResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateRoundQuestionResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateRoundQuestionResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateRoundQuestionResponseValidationError{}

// Validate checks the field values on GetRoundQuestionRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetRoundQuestionRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetRoundQuestionRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetRoundQuestionRequestMultiError, or nil if none found.
func (m *GetRoundQuestionRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetRoundQuestionRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for RoundQuestionId

	if len(errors) > 0 {
		return GetRoundQuestionRequestMultiError(errors)
	}

	return nil
}

// GetRoundQuestionRequestMultiError is an error wrapping multiple validation
// errors returned by GetRoundQuestionRequest.ValidateAll() if the designated
// constraints aren't met.
type GetRoundQuestionRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetRoundQuestionRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetRoundQuestionRequestMultiError) AllErrors() []error { return m }

// GetRoundQuestionRequestValidationError is the validation error returned by
// GetRoundQuestionRequest.Validate if the designated constraints aren't met.
type GetRoundQuestionRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetRoundQuestionRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetRoundQuestionRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetRoundQuestionRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetRoundQuestionRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetRoundQuestionRequestValidationError) ErrorName() string {
	return "GetRoundQuestionRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetRoundQuestionRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetRoundQuestionRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetRoundQuestionRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetRoundQuestionRequestValidationError{}

// Validate checks the field values on GetRoundQuestionResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetRoundQuestionResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetRoundQuestionResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetRoundQuestionResponseMultiError, or nil if none found.
func (m *GetRoundQuestionResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetRoundQuestionResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetRoundQuestion()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetRoundQuestionResponseValidationError{
					field:  "RoundQuestion",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetRoundQuestionResponseValidationError{
					field:  "RoundQuestion",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRoundQuestion()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetRoundQuestionResponseValidationError{
				field:  "RoundQuestion",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetRoundQuestionResponseMultiError(errors)
	}

	return nil
}

// GetRoundQuestionResponseMultiError is an error wrapping multiple validation
// errors returned by GetRoundQuestionResponse.ValidateAll() if the designated
// constraints aren't met.
type GetRoundQuestionResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetRoundQuestionResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetRoundQuestionResponseMultiError) AllErrors() []error { return m }

// GetRoundQuestionResponseValidationError is the validation error returned by
// GetRoundQuestionResponse.Validate if the designated constraints aren't met.
type GetRoundQuestionResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetRoundQuestionResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetRoundQuestionResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetRoundQuestionResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetRoundQuestionResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetRoundQuestionResponseValidationError) ErrorName() string {
	return "GetRoundQuestionResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetRoundQuestionResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetRoundQuestionResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetRoundQuestionResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetRoundQuestionResponseValidationError{}

// Validate checks the field values on RoundQuestion_Question with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RoundQuestion_Question) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RoundQuestion_Question with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RoundQuestion_QuestionMultiError, or nil if none found.
func (m *RoundQuestion_Question) ValidateAll() error {
	return m.validate(true)
}

func (m *RoundQuestion_Question) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Text

	// no validation rules for Author

	// no validation rules for MediaUrl

	if all {
		switch v := interface{}(m.GetCreateTime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RoundQuestion_QuestionValidationError{
					field:  "CreateTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RoundQuestion_QuestionValidationError{
					field:  "CreateTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreateTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RoundQuestion_QuestionValidationError{
				field:  "CreateTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return RoundQuestion_QuestionMultiError(errors)
	}

	return nil
}

// RoundQuestion_QuestionMultiError is an error wrapping multiple validation
// errors returned by RoundQuestion_Question.ValidateAll() if the designated
// constraints aren't met.
type RoundQuestion_QuestionMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RoundQuestion_QuestionMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RoundQuestion_QuestionMultiError) AllErrors() []error { return m }

// RoundQuestion_QuestionValidationError is the validation error returned by
// RoundQuestion_Question.Validate if the designated constraints aren't met.
type RoundQuestion_QuestionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RoundQuestion_QuestionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RoundQuestion_QuestionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RoundQuestion_QuestionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RoundQuestion_QuestionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RoundQuestion_QuestionValidationError) ErrorName() string {
	return "RoundQuestion_QuestionValidationError"
}

// Error satisfies the builtin error interface
func (e RoundQuestion_QuestionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRoundQuestion_Question.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RoundQuestion_QuestionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RoundQuestion_QuestionValidationError{}

// Validate checks the field values on RoundQuestion_Answer with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RoundQuestion_Answer) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RoundQuestion_Answer with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RoundQuestion_AnswerMultiError, or nil if none found.
func (m *RoundQuestion_Answer) ValidateAll() error {
	return m.validate(true)
}

func (m *RoundQuestion_Answer) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Text

	// no validation rules for Author

	// no validation rules for MediaUrl

	if all {
		switch v := interface{}(m.GetCreateTime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RoundQuestion_AnswerValidationError{
					field:  "CreateTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RoundQuestion_AnswerValidationError{
					field:  "CreateTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreateTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RoundQuestion_AnswerValidationError{
				field:  "CreateTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return RoundQuestion_AnswerMultiError(errors)
	}

	return nil
}

// RoundQuestion_AnswerMultiError is an error wrapping multiple validation
// errors returned by RoundQuestion_Answer.ValidateAll() if the designated
// constraints aren't met.
type RoundQuestion_AnswerMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RoundQuestion_AnswerMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RoundQuestion_AnswerMultiError) AllErrors() []error { return m }

// RoundQuestion_AnswerValidationError is the validation error returned by
// RoundQuestion_Answer.Validate if the designated constraints aren't met.
type RoundQuestion_AnswerValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RoundQuestion_AnswerValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RoundQuestion_AnswerValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RoundQuestion_AnswerValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RoundQuestion_AnswerValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RoundQuestion_AnswerValidationError) ErrorName() string {
	return "RoundQuestion_AnswerValidationError"
}

// Error satisfies the builtin error interface
func (e RoundQuestion_AnswerValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRoundQuestion_Answer.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RoundQuestion_AnswerValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RoundQuestion_AnswerValidationError{}
