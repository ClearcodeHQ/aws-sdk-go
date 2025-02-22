// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package lakeformationiface provides an interface to enable mocking the AWS Lake Formation service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package lakeformationiface

import (
	"github.com/ClearcodeHQ/aws-sdk-go/aws"
	"github.com/ClearcodeHQ/aws-sdk-go/aws/request"
	"github.com/ClearcodeHQ/aws-sdk-go/service/lakeformation"
)

// LakeFormationAPI provides an interface to enable mocking the
// lakeformation.LakeFormation service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Lake Formation.
//    func myFunc(svc lakeformationiface.LakeFormationAPI) bool {
//        // Make svc.BatchGrantPermissions request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := lakeformation.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockLakeFormationClient struct {
//        lakeformationiface.LakeFormationAPI
//    }
//    func (m *mockLakeFormationClient) BatchGrantPermissions(input *lakeformation.BatchGrantPermissionsInput) (*lakeformation.BatchGrantPermissionsOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockLakeFormationClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type LakeFormationAPI interface {
	BatchGrantPermissions(*lakeformation.BatchGrantPermissionsInput) (*lakeformation.BatchGrantPermissionsOutput, error)
	BatchGrantPermissionsWithContext(aws.Context, *lakeformation.BatchGrantPermissionsInput, ...request.Option) (*lakeformation.BatchGrantPermissionsOutput, error)
	BatchGrantPermissionsRequest(*lakeformation.BatchGrantPermissionsInput) (*request.Request, *lakeformation.BatchGrantPermissionsOutput)

	BatchRevokePermissions(*lakeformation.BatchRevokePermissionsInput) (*lakeformation.BatchRevokePermissionsOutput, error)
	BatchRevokePermissionsWithContext(aws.Context, *lakeformation.BatchRevokePermissionsInput, ...request.Option) (*lakeformation.BatchRevokePermissionsOutput, error)
	BatchRevokePermissionsRequest(*lakeformation.BatchRevokePermissionsInput) (*request.Request, *lakeformation.BatchRevokePermissionsOutput)

	DeregisterResource(*lakeformation.DeregisterResourceInput) (*lakeformation.DeregisterResourceOutput, error)
	DeregisterResourceWithContext(aws.Context, *lakeformation.DeregisterResourceInput, ...request.Option) (*lakeformation.DeregisterResourceOutput, error)
	DeregisterResourceRequest(*lakeformation.DeregisterResourceInput) (*request.Request, *lakeformation.DeregisterResourceOutput)

	DescribeResource(*lakeformation.DescribeResourceInput) (*lakeformation.DescribeResourceOutput, error)
	DescribeResourceWithContext(aws.Context, *lakeformation.DescribeResourceInput, ...request.Option) (*lakeformation.DescribeResourceOutput, error)
	DescribeResourceRequest(*lakeformation.DescribeResourceInput) (*request.Request, *lakeformation.DescribeResourceOutput)

	GetDataLakeSettings(*lakeformation.GetDataLakeSettingsInput) (*lakeformation.GetDataLakeSettingsOutput, error)
	GetDataLakeSettingsWithContext(aws.Context, *lakeformation.GetDataLakeSettingsInput, ...request.Option) (*lakeformation.GetDataLakeSettingsOutput, error)
	GetDataLakeSettingsRequest(*lakeformation.GetDataLakeSettingsInput) (*request.Request, *lakeformation.GetDataLakeSettingsOutput)

	GetEffectivePermissionsForPath(*lakeformation.GetEffectivePermissionsForPathInput) (*lakeformation.GetEffectivePermissionsForPathOutput, error)
	GetEffectivePermissionsForPathWithContext(aws.Context, *lakeformation.GetEffectivePermissionsForPathInput, ...request.Option) (*lakeformation.GetEffectivePermissionsForPathOutput, error)
	GetEffectivePermissionsForPathRequest(*lakeformation.GetEffectivePermissionsForPathInput) (*request.Request, *lakeformation.GetEffectivePermissionsForPathOutput)

	GetEffectivePermissionsForPathPages(*lakeformation.GetEffectivePermissionsForPathInput, func(*lakeformation.GetEffectivePermissionsForPathOutput, bool) bool) error
	GetEffectivePermissionsForPathPagesWithContext(aws.Context, *lakeformation.GetEffectivePermissionsForPathInput, func(*lakeformation.GetEffectivePermissionsForPathOutput, bool) bool, ...request.Option) error

	GrantPermissions(*lakeformation.GrantPermissionsInput) (*lakeformation.GrantPermissionsOutput, error)
	GrantPermissionsWithContext(aws.Context, *lakeformation.GrantPermissionsInput, ...request.Option) (*lakeformation.GrantPermissionsOutput, error)
	GrantPermissionsRequest(*lakeformation.GrantPermissionsInput) (*request.Request, *lakeformation.GrantPermissionsOutput)

	ListPermissions(*lakeformation.ListPermissionsInput) (*lakeformation.ListPermissionsOutput, error)
	ListPermissionsWithContext(aws.Context, *lakeformation.ListPermissionsInput, ...request.Option) (*lakeformation.ListPermissionsOutput, error)
	ListPermissionsRequest(*lakeformation.ListPermissionsInput) (*request.Request, *lakeformation.ListPermissionsOutput)

	ListPermissionsPages(*lakeformation.ListPermissionsInput, func(*lakeformation.ListPermissionsOutput, bool) bool) error
	ListPermissionsPagesWithContext(aws.Context, *lakeformation.ListPermissionsInput, func(*lakeformation.ListPermissionsOutput, bool) bool, ...request.Option) error

	ListResources(*lakeformation.ListResourcesInput) (*lakeformation.ListResourcesOutput, error)
	ListResourcesWithContext(aws.Context, *lakeformation.ListResourcesInput, ...request.Option) (*lakeformation.ListResourcesOutput, error)
	ListResourcesRequest(*lakeformation.ListResourcesInput) (*request.Request, *lakeformation.ListResourcesOutput)

	ListResourcesPages(*lakeformation.ListResourcesInput, func(*lakeformation.ListResourcesOutput, bool) bool) error
	ListResourcesPagesWithContext(aws.Context, *lakeformation.ListResourcesInput, func(*lakeformation.ListResourcesOutput, bool) bool, ...request.Option) error

	PutDataLakeSettings(*lakeformation.PutDataLakeSettingsInput) (*lakeformation.PutDataLakeSettingsOutput, error)
	PutDataLakeSettingsWithContext(aws.Context, *lakeformation.PutDataLakeSettingsInput, ...request.Option) (*lakeformation.PutDataLakeSettingsOutput, error)
	PutDataLakeSettingsRequest(*lakeformation.PutDataLakeSettingsInput) (*request.Request, *lakeformation.PutDataLakeSettingsOutput)

	RegisterResource(*lakeformation.RegisterResourceInput) (*lakeformation.RegisterResourceOutput, error)
	RegisterResourceWithContext(aws.Context, *lakeformation.RegisterResourceInput, ...request.Option) (*lakeformation.RegisterResourceOutput, error)
	RegisterResourceRequest(*lakeformation.RegisterResourceInput) (*request.Request, *lakeformation.RegisterResourceOutput)

	RevokePermissions(*lakeformation.RevokePermissionsInput) (*lakeformation.RevokePermissionsOutput, error)
	RevokePermissionsWithContext(aws.Context, *lakeformation.RevokePermissionsInput, ...request.Option) (*lakeformation.RevokePermissionsOutput, error)
	RevokePermissionsRequest(*lakeformation.RevokePermissionsInput) (*request.Request, *lakeformation.RevokePermissionsOutput)

	UpdateResource(*lakeformation.UpdateResourceInput) (*lakeformation.UpdateResourceOutput, error)
	UpdateResourceWithContext(aws.Context, *lakeformation.UpdateResourceInput, ...request.Option) (*lakeformation.UpdateResourceOutput, error)
	UpdateResourceRequest(*lakeformation.UpdateResourceInput) (*request.Request, *lakeformation.UpdateResourceOutput)
}

var _ LakeFormationAPI = (*lakeformation.LakeFormation)(nil)
