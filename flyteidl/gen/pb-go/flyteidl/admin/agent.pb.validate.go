// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: flyteidl/admin/agent.proto

package admin

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _agent_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on TaskExecutionMetadata with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *TaskExecutionMetadata) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetTaskExecutionId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TaskExecutionMetadataValidationError{
				field:  "TaskExecutionId",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Namespace

	// no validation rules for Labels

	// no validation rules for Annotations

	// no validation rules for K8SServiceAccount

	// no validation rules for EnvironmentVariables

	return nil
}

// TaskExecutionMetadataValidationError is the validation error returned by
// TaskExecutionMetadata.Validate if the designated constraints aren't met.
type TaskExecutionMetadataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TaskExecutionMetadataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TaskExecutionMetadataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TaskExecutionMetadataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TaskExecutionMetadataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TaskExecutionMetadataValidationError) ErrorName() string {
	return "TaskExecutionMetadataValidationError"
}

// Error satisfies the builtin error interface
func (e TaskExecutionMetadataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTaskExecutionMetadata.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TaskExecutionMetadataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TaskExecutionMetadataValidationError{}

// Validate checks the field values on CreateTaskRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *CreateTaskRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetInputs()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateTaskRequestValidationError{
				field:  "Inputs",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetTemplate()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateTaskRequestValidationError{
				field:  "Template",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for OutputPrefix

	if v, ok := interface{}(m.GetTaskExecutionMetadata()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateTaskRequestValidationError{
				field:  "TaskExecutionMetadata",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// CreateTaskRequestValidationError is the validation error returned by
// CreateTaskRequest.Validate if the designated constraints aren't met.
type CreateTaskRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateTaskRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateTaskRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateTaskRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateTaskRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateTaskRequestValidationError) ErrorName() string {
	return "CreateTaskRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateTaskRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateTaskRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateTaskRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateTaskRequestValidationError{}

// Validate checks the field values on CreateTaskResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateTaskResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for ResourceMeta

	return nil
}

// CreateTaskResponseValidationError is the validation error returned by
// CreateTaskResponse.Validate if the designated constraints aren't met.
type CreateTaskResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateTaskResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateTaskResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateTaskResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateTaskResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateTaskResponseValidationError) ErrorName() string {
	return "CreateTaskResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateTaskResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateTaskResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateTaskResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateTaskResponseValidationError{}

// Validate checks the field values on GetTaskRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *GetTaskRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for TaskType

	// no validation rules for ResourceMeta

	return nil
}

// GetTaskRequestValidationError is the validation error returned by
// GetTaskRequest.Validate if the designated constraints aren't met.
type GetTaskRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTaskRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTaskRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTaskRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTaskRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTaskRequestValidationError) ErrorName() string { return "GetTaskRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetTaskRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTaskRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTaskRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTaskRequestValidationError{}

// Validate checks the field values on GetTaskResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GetTaskResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetResource()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetTaskResponseValidationError{
				field:  "Resource",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// GetTaskResponseValidationError is the validation error returned by
// GetTaskResponse.Validate if the designated constraints aren't met.
type GetTaskResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTaskResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTaskResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTaskResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTaskResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTaskResponseValidationError) ErrorName() string { return "GetTaskResponseValidationError" }

// Error satisfies the builtin error interface
func (e GetTaskResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTaskResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTaskResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTaskResponseValidationError{}

// Validate checks the field values on Resource with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Resource) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for State

	if v, ok := interface{}(m.GetOutputs()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ResourceValidationError{
				field:  "Outputs",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Message

	return nil
}

// ResourceValidationError is the validation error returned by
// Resource.Validate if the designated constraints aren't met.
type ResourceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ResourceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ResourceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ResourceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ResourceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ResourceValidationError) ErrorName() string { return "ResourceValidationError" }

// Error satisfies the builtin error interface
func (e ResourceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResource.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ResourceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ResourceValidationError{}

// Validate checks the field values on DeleteTaskRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *DeleteTaskRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for TaskType

	// no validation rules for ResourceMeta

	return nil
}

// DeleteTaskRequestValidationError is the validation error returned by
// DeleteTaskRequest.Validate if the designated constraints aren't met.
type DeleteTaskRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteTaskRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteTaskRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteTaskRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteTaskRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteTaskRequestValidationError) ErrorName() string {
	return "DeleteTaskRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteTaskRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteTaskRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteTaskRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteTaskRequestValidationError{}

// Validate checks the field values on DeleteTaskResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DeleteTaskResponse) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// DeleteTaskResponseValidationError is the validation error returned by
// DeleteTaskResponse.Validate if the designated constraints aren't met.
type DeleteTaskResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteTaskResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteTaskResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteTaskResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteTaskResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteTaskResponseValidationError) ErrorName() string {
	return "DeleteTaskResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteTaskResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteTaskResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteTaskResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteTaskResponseValidationError{}
