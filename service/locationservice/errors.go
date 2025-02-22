// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package locationservice

import (
	"github.com/ClearcodeHQ/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeAccessDeniedException for service response error code
	// "AccessDeniedException".
	//
	// The request was denied due to insufficient access or permission. Check with
	// an administrator to verify your permissions.
	ErrCodeAccessDeniedException = "AccessDeniedException"

	// ErrCodeConflictException for service response error code
	// "ConflictException".
	//
	// The request was unsuccessful due to a conflict.
	ErrCodeConflictException = "ConflictException"

	// ErrCodeInternalServerException for service response error code
	// "InternalServerException".
	//
	// The request has failed to process because of an unknown server error, exception,
	// or failure.
	ErrCodeInternalServerException = "InternalServerException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// The resource that you've entered was not found in your AWS account.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeThrottlingException for service response error code
	// "ThrottlingException".
	//
	// The request was denied due to request throttling.
	ErrCodeThrottlingException = "ThrottlingException"

	// ErrCodeValidationException for service response error code
	// "ValidationException".
	//
	// The input failed to meet the constraints specified by the AWS service.
	ErrCodeValidationException = "ValidationException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"AccessDeniedException":     newErrorAccessDeniedException,
	"ConflictException":         newErrorConflictException,
	"InternalServerException":   newErrorInternalServerException,
	"ResourceNotFoundException": newErrorResourceNotFoundException,
	"ThrottlingException":       newErrorThrottlingException,
	"ValidationException":       newErrorValidationException,
}
