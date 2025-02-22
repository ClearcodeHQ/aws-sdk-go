// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemakerruntime

import (
	"fmt"

	"github.com/ClearcodeHQ/aws-sdk-go/aws"
	"github.com/ClearcodeHQ/aws-sdk-go/aws/awsutil"
	"github.com/ClearcodeHQ/aws-sdk-go/aws/request"
	"github.com/ClearcodeHQ/aws-sdk-go/private/protocol"
)

const opInvokeEndpoint = "InvokeEndpoint"

// InvokeEndpointRequest generates a "aws/request.Request" representing the
// client's request for the InvokeEndpoint operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See InvokeEndpoint for more information on using the InvokeEndpoint
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the InvokeEndpointRequest method.
//    req, resp := client.InvokeEndpointRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/runtime.sagemaker-2017-05-13/InvokeEndpoint
func (c *SageMakerRuntime) InvokeEndpointRequest(input *InvokeEndpointInput) (req *request.Request, output *InvokeEndpointOutput) {
	op := &request.Operation{
		Name:       opInvokeEndpoint,
		HTTPMethod: "POST",
		HTTPPath:   "/endpoints/{EndpointName}/invocations",
	}

	if input == nil {
		input = &InvokeEndpointInput{}
	}

	output = &InvokeEndpointOutput{}
	req = c.newRequest(op, input, output)
	return
}

// InvokeEndpoint API operation for Amazon SageMaker Runtime.
//
// After you deploy a model into production using Amazon SageMaker hosting services,
// your client applications use this API to get inferences from the model hosted
// at the specified endpoint.
//
// For an overview of Amazon SageMaker, see How It Works (https://docs.aws.amazon.com/sagemaker/latest/dg/how-it-works.html).
//
// Amazon SageMaker strips all POST headers except those supported by the API.
// Amazon SageMaker might add additional headers. You should not rely on the
// behavior of headers outside those enumerated in the request syntax.
//
// Calls to InvokeEndpoint are authenticated by using AWS Signature Version
// 4. For information, see Authenticating Requests (AWS Signature Version 4)
// (https://docs.aws.amazon.com/AmazonS3/latest/API/sig-v4-authenticating-requests.html)
// in the Amazon S3 API Reference.
//
// A customer's model containers must respond to requests within 60 seconds.
// The model itself can have a maximum processing time of 60 seconds before
// responding to invocations. If your model is going to take 50-60 seconds of
// processing time, the SDK socket timeout should be set to be 70 seconds.
//
// Endpoints are scoped to an individual account, and are not public. The URL
// does not contain the account ID, but Amazon SageMaker determines the account
// ID from the authentication token that is supplied by the caller.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon SageMaker Runtime's
// API operation InvokeEndpoint for usage and error information.
//
// Returned Error Types:
//   * InternalFailure
//   An internal failure occurred.
//
//   * ServiceUnavailable
//   The service is unavailable. Try your call again.
//
//   * ValidationError
//   Inspect your request and try again.
//
//   * ModelError
//   Model (owned by the customer in the container) returned 4xx or 5xx error
//   code.
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/runtime.sagemaker-2017-05-13/InvokeEndpoint
func (c *SageMakerRuntime) InvokeEndpoint(input *InvokeEndpointInput) (*InvokeEndpointOutput, error) {
	req, out := c.InvokeEndpointRequest(input)
	return out, req.Send()
}

// InvokeEndpointWithContext is the same as InvokeEndpoint with the addition of
// the ability to pass a context and additional request options.
//
// See InvokeEndpoint for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SageMakerRuntime) InvokeEndpointWithContext(ctx aws.Context, input *InvokeEndpointInput, opts ...request.Option) (*InvokeEndpointOutput, error) {
	req, out := c.InvokeEndpointRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

// An internal failure occurred.
type InternalFailure struct {
	_            struct{}                  `type:"structure"`
	RespMetadata protocol.ResponseMetadata `json:"-" xml:"-"`

	Message_ *string `locationName:"Message" type:"string"`
}

// String returns the string representation
func (s InternalFailure) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s InternalFailure) GoString() string {
	return s.String()
}

func newErrorInternalFailure(v protocol.ResponseMetadata) error {
	return &InternalFailure{
		RespMetadata: v,
	}
}

// Code returns the exception type name.
func (s *InternalFailure) Code() string {
	return "InternalFailure"
}

// Message returns the exception's message.
func (s *InternalFailure) Message() string {
	if s.Message_ != nil {
		return *s.Message_
	}
	return ""
}

// OrigErr always returns nil, satisfies awserr.Error interface.
func (s *InternalFailure) OrigErr() error {
	return nil
}

func (s *InternalFailure) Error() string {
	return fmt.Sprintf("%s: %s", s.Code(), s.Message())
}

// Status code returns the HTTP status code for the request's response error.
func (s *InternalFailure) StatusCode() int {
	return s.RespMetadata.StatusCode
}

// RequestID returns the service's response RequestID for request.
func (s *InternalFailure) RequestID() string {
	return s.RespMetadata.RequestID
}

type InvokeEndpointInput struct {
	_ struct{} `type:"structure" payload:"Body"`

	// The desired MIME type of the inference in the response.
	Accept *string `location:"header" locationName:"Accept" type:"string"`

	// Provides input data, in the format specified in the ContentType request header.
	// Amazon SageMaker passes all of the data in the body to the model.
	//
	// For information about the format of the request body, see Common Data Formats-Inference
	// (https://docs.aws.amazon.com/sagemaker/latest/dg/cdf-inference.html).
	//
	// Body is a required field
	Body []byte `type:"blob" required:"true" sensitive:"true"`

	// The MIME type of the input data in the request body.
	ContentType *string `location:"header" locationName:"Content-Type" type:"string"`

	// Provides additional information about a request for an inference submitted
	// to a model hosted at an Amazon SageMaker endpoint. The information is an
	// opaque value that is forwarded verbatim. You could use this value, for example,
	// to provide an ID that you can use to track a request or to provide other
	// metadata that a service endpoint was programmed to process. The value must
	// consist of no more than 1024 visible US-ASCII characters as specified in
	// Section 3.3.6. Field Value Components (https://tools.ietf.org/html/rfc7230#section-3.2.6)
	// of the Hypertext Transfer Protocol (HTTP/1.1).
	//
	// The code in your model is responsible for setting or updating any custom
	// attributes in the response. If your code does not set this value in the response,
	// an empty value is returned. For example, if a custom attribute represents
	// the trace ID, your model can prepend the custom attribute with Trace ID:
	// in your post-processing function.
	//
	// This feature is currently supported in the AWS SDKs but not in the Amazon
	// SageMaker Python SDK.
	CustomAttributes *string `location:"header" locationName:"X-Amzn-SageMaker-Custom-Attributes" type:"string" sensitive:"true"`

	// The name of the endpoint that you specified when you created the endpoint
	// using the CreateEndpoint (https://docs.aws.amazon.com/sagemaker/latest/dg/API_CreateEndpoint.html)
	// API.
	//
	// EndpointName is a required field
	EndpointName *string `location:"uri" locationName:"EndpointName" type:"string" required:"true"`

	// If you provide a value, it is added to the captured data when you enable
	// data capture on the endpoint. For information about data capture, see Capture
	// Data (https://docs.aws.amazon.com/sagemaker/latest/dg/model-monitor-data-capture.html).
	InferenceId *string `location:"header" locationName:"X-Amzn-SageMaker-Inference-Id" min:"1" type:"string"`

	// If the endpoint hosts multiple containers and is configured to use direct
	// invocation, this parameter specifies the host name of the container to invoke.
	TargetContainerHostname *string `location:"header" locationName:"X-Amzn-SageMaker-Target-Container-Hostname" type:"string"`

	// The model to request for inference when invoking a multi-model endpoint.
	TargetModel *string `location:"header" locationName:"X-Amzn-SageMaker-Target-Model" min:"1" type:"string"`

	// Specify the production variant to send the inference request to when invoking
	// an endpoint that is running two or more variants. Note that this parameter
	// overrides the default behavior for the endpoint, which is to distribute the
	// invocation traffic based on the variant weights.
	//
	// For information about how to use variant targeting to perform a/b testing,
	// see Test models in production (https://docs.aws.amazon.com/sagemaker/latest/dg/model-ab-testing.html)
	TargetVariant *string `location:"header" locationName:"X-Amzn-SageMaker-Target-Variant" type:"string"`
}

// String returns the string representation
func (s InvokeEndpointInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s InvokeEndpointInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *InvokeEndpointInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "InvokeEndpointInput"}
	if s.Body == nil {
		invalidParams.Add(request.NewErrParamRequired("Body"))
	}
	if s.EndpointName == nil {
		invalidParams.Add(request.NewErrParamRequired("EndpointName"))
	}
	if s.EndpointName != nil && len(*s.EndpointName) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("EndpointName", 1))
	}
	if s.InferenceId != nil && len(*s.InferenceId) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("InferenceId", 1))
	}
	if s.TargetModel != nil && len(*s.TargetModel) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("TargetModel", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAccept sets the Accept field's value.
func (s *InvokeEndpointInput) SetAccept(v string) *InvokeEndpointInput {
	s.Accept = &v
	return s
}

// SetBody sets the Body field's value.
func (s *InvokeEndpointInput) SetBody(v []byte) *InvokeEndpointInput {
	s.Body = v
	return s
}

// SetContentType sets the ContentType field's value.
func (s *InvokeEndpointInput) SetContentType(v string) *InvokeEndpointInput {
	s.ContentType = &v
	return s
}

// SetCustomAttributes sets the CustomAttributes field's value.
func (s *InvokeEndpointInput) SetCustomAttributes(v string) *InvokeEndpointInput {
	s.CustomAttributes = &v
	return s
}

// SetEndpointName sets the EndpointName field's value.
func (s *InvokeEndpointInput) SetEndpointName(v string) *InvokeEndpointInput {
	s.EndpointName = &v
	return s
}

// SetInferenceId sets the InferenceId field's value.
func (s *InvokeEndpointInput) SetInferenceId(v string) *InvokeEndpointInput {
	s.InferenceId = &v
	return s
}

// SetTargetContainerHostname sets the TargetContainerHostname field's value.
func (s *InvokeEndpointInput) SetTargetContainerHostname(v string) *InvokeEndpointInput {
	s.TargetContainerHostname = &v
	return s
}

// SetTargetModel sets the TargetModel field's value.
func (s *InvokeEndpointInput) SetTargetModel(v string) *InvokeEndpointInput {
	s.TargetModel = &v
	return s
}

// SetTargetVariant sets the TargetVariant field's value.
func (s *InvokeEndpointInput) SetTargetVariant(v string) *InvokeEndpointInput {
	s.TargetVariant = &v
	return s
}

type InvokeEndpointOutput struct {
	_ struct{} `type:"structure" payload:"Body"`

	// Includes the inference provided by the model.
	//
	// For information about the format of the response body, see Common Data Formats-Inference
	// (https://docs.aws.amazon.com/sagemaker/latest/dg/cdf-inference.html).
	//
	// Body is a required field
	Body []byte `type:"blob" required:"true" sensitive:"true"`

	// The MIME type of the inference returned in the response body.
	ContentType *string `location:"header" locationName:"Content-Type" type:"string"`

	// Provides additional information in the response about the inference returned
	// by a model hosted at an Amazon SageMaker endpoint. The information is an
	// opaque value that is forwarded verbatim. You could use this value, for example,
	// to return an ID received in the CustomAttributes header of a request or other
	// metadata that a service endpoint was programmed to produce. The value must
	// consist of no more than 1024 visible US-ASCII characters as specified in
	// Section 3.3.6. Field Value Components (https://tools.ietf.org/html/rfc7230#section-3.2.6)
	// of the Hypertext Transfer Protocol (HTTP/1.1). If the customer wants the
	// custom attribute returned, the model must set the custom attribute to be
	// included on the way back.
	//
	// The code in your model is responsible for setting or updating any custom
	// attributes in the response. If your code does not set this value in the response,
	// an empty value is returned. For example, if a custom attribute represents
	// the trace ID, your model can prepend the custom attribute with Trace ID:
	// in your post-processing function.
	//
	// This feature is currently supported in the AWS SDKs but not in the Amazon
	// SageMaker Python SDK.
	CustomAttributes *string `location:"header" locationName:"X-Amzn-SageMaker-Custom-Attributes" type:"string" sensitive:"true"`

	// Identifies the production variant that was invoked.
	InvokedProductionVariant *string `location:"header" locationName:"x-Amzn-Invoked-Production-Variant" type:"string"`
}

// String returns the string representation
func (s InvokeEndpointOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s InvokeEndpointOutput) GoString() string {
	return s.String()
}

// SetBody sets the Body field's value.
func (s *InvokeEndpointOutput) SetBody(v []byte) *InvokeEndpointOutput {
	s.Body = v
	return s
}

// SetContentType sets the ContentType field's value.
func (s *InvokeEndpointOutput) SetContentType(v string) *InvokeEndpointOutput {
	s.ContentType = &v
	return s
}

// SetCustomAttributes sets the CustomAttributes field's value.
func (s *InvokeEndpointOutput) SetCustomAttributes(v string) *InvokeEndpointOutput {
	s.CustomAttributes = &v
	return s
}

// SetInvokedProductionVariant sets the InvokedProductionVariant field's value.
func (s *InvokeEndpointOutput) SetInvokedProductionVariant(v string) *InvokeEndpointOutput {
	s.InvokedProductionVariant = &v
	return s
}

// Model (owned by the customer in the container) returned 4xx or 5xx error
// code.
type ModelError struct {
	_            struct{}                  `type:"structure"`
	RespMetadata protocol.ResponseMetadata `json:"-" xml:"-"`

	// The Amazon Resource Name (ARN) of the log stream.
	LogStreamArn *string `type:"string"`

	Message_ *string `locationName:"Message" type:"string"`

	// Original message.
	OriginalMessage *string `type:"string"`

	// Original status code.
	OriginalStatusCode *int64 `type:"integer"`
}

// String returns the string representation
func (s ModelError) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ModelError) GoString() string {
	return s.String()
}

func newErrorModelError(v protocol.ResponseMetadata) error {
	return &ModelError{
		RespMetadata: v,
	}
}

// Code returns the exception type name.
func (s *ModelError) Code() string {
	return "ModelError"
}

// Message returns the exception's message.
func (s *ModelError) Message() string {
	if s.Message_ != nil {
		return *s.Message_
	}
	return ""
}

// OrigErr always returns nil, satisfies awserr.Error interface.
func (s *ModelError) OrigErr() error {
	return nil
}

func (s *ModelError) Error() string {
	return fmt.Sprintf("%s: %s\n%s", s.Code(), s.Message(), s.String())
}

// Status code returns the HTTP status code for the request's response error.
func (s *ModelError) StatusCode() int {
	return s.RespMetadata.StatusCode
}

// RequestID returns the service's response RequestID for request.
func (s *ModelError) RequestID() string {
	return s.RespMetadata.RequestID
}

// The service is unavailable. Try your call again.
type ServiceUnavailable struct {
	_            struct{}                  `type:"structure"`
	RespMetadata protocol.ResponseMetadata `json:"-" xml:"-"`

	Message_ *string `locationName:"Message" type:"string"`
}

// String returns the string representation
func (s ServiceUnavailable) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ServiceUnavailable) GoString() string {
	return s.String()
}

func newErrorServiceUnavailable(v protocol.ResponseMetadata) error {
	return &ServiceUnavailable{
		RespMetadata: v,
	}
}

// Code returns the exception type name.
func (s *ServiceUnavailable) Code() string {
	return "ServiceUnavailable"
}

// Message returns the exception's message.
func (s *ServiceUnavailable) Message() string {
	if s.Message_ != nil {
		return *s.Message_
	}
	return ""
}

// OrigErr always returns nil, satisfies awserr.Error interface.
func (s *ServiceUnavailable) OrigErr() error {
	return nil
}

func (s *ServiceUnavailable) Error() string {
	return fmt.Sprintf("%s: %s", s.Code(), s.Message())
}

// Status code returns the HTTP status code for the request's response error.
func (s *ServiceUnavailable) StatusCode() int {
	return s.RespMetadata.StatusCode
}

// RequestID returns the service's response RequestID for request.
func (s *ServiceUnavailable) RequestID() string {
	return s.RespMetadata.RequestID
}

// Inspect your request and try again.
type ValidationError struct {
	_            struct{}                  `type:"structure"`
	RespMetadata protocol.ResponseMetadata `json:"-" xml:"-"`

	Message_ *string `locationName:"Message" type:"string"`
}

// String returns the string representation
func (s ValidationError) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ValidationError) GoString() string {
	return s.String()
}

func newErrorValidationError(v protocol.ResponseMetadata) error {
	return &ValidationError{
		RespMetadata: v,
	}
}

// Code returns the exception type name.
func (s *ValidationError) Code() string {
	return "ValidationError"
}

// Message returns the exception's message.
func (s *ValidationError) Message() string {
	if s.Message_ != nil {
		return *s.Message_
	}
	return ""
}

// OrigErr always returns nil, satisfies awserr.Error interface.
func (s *ValidationError) OrigErr() error {
	return nil
}

func (s *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", s.Code(), s.Message())
}

// Status code returns the HTTP status code for the request's response error.
func (s *ValidationError) StatusCode() int {
	return s.RespMetadata.StatusCode
}

// RequestID returns the service's response RequestID for request.
func (s *ValidationError) RequestID() string {
	return s.RespMetadata.RequestID
}
