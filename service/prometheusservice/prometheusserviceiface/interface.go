// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package prometheusserviceiface provides an interface to enable mocking the Amazon Prometheus Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package prometheusserviceiface

import (
	"github.com/ClearcodeHQ/aws-sdk-go/aws"
	"github.com/ClearcodeHQ/aws-sdk-go/aws/request"
	"github.com/ClearcodeHQ/aws-sdk-go/service/prometheusservice"
)

// PrometheusServiceAPI provides an interface to enable mocking the
// prometheusservice.PrometheusService service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Prometheus Service.
//    func myFunc(svc prometheusserviceiface.PrometheusServiceAPI) bool {
//        // Make svc.CreateWorkspace request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := prometheusservice.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockPrometheusServiceClient struct {
//        prometheusserviceiface.PrometheusServiceAPI
//    }
//    func (m *mockPrometheusServiceClient) CreateWorkspace(input *prometheusservice.CreateWorkspaceInput) (*prometheusservice.CreateWorkspaceOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockPrometheusServiceClient{}
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
type PrometheusServiceAPI interface {
	CreateWorkspace(*prometheusservice.CreateWorkspaceInput) (*prometheusservice.CreateWorkspaceOutput, error)
	CreateWorkspaceWithContext(aws.Context, *prometheusservice.CreateWorkspaceInput, ...request.Option) (*prometheusservice.CreateWorkspaceOutput, error)
	CreateWorkspaceRequest(*prometheusservice.CreateWorkspaceInput) (*request.Request, *prometheusservice.CreateWorkspaceOutput)

	DeleteWorkspace(*prometheusservice.DeleteWorkspaceInput) (*prometheusservice.DeleteWorkspaceOutput, error)
	DeleteWorkspaceWithContext(aws.Context, *prometheusservice.DeleteWorkspaceInput, ...request.Option) (*prometheusservice.DeleteWorkspaceOutput, error)
	DeleteWorkspaceRequest(*prometheusservice.DeleteWorkspaceInput) (*request.Request, *prometheusservice.DeleteWorkspaceOutput)

	DescribeWorkspace(*prometheusservice.DescribeWorkspaceInput) (*prometheusservice.DescribeWorkspaceOutput, error)
	DescribeWorkspaceWithContext(aws.Context, *prometheusservice.DescribeWorkspaceInput, ...request.Option) (*prometheusservice.DescribeWorkspaceOutput, error)
	DescribeWorkspaceRequest(*prometheusservice.DescribeWorkspaceInput) (*request.Request, *prometheusservice.DescribeWorkspaceOutput)

	ListWorkspaces(*prometheusservice.ListWorkspacesInput) (*prometheusservice.ListWorkspacesOutput, error)
	ListWorkspacesWithContext(aws.Context, *prometheusservice.ListWorkspacesInput, ...request.Option) (*prometheusservice.ListWorkspacesOutput, error)
	ListWorkspacesRequest(*prometheusservice.ListWorkspacesInput) (*request.Request, *prometheusservice.ListWorkspacesOutput)

	ListWorkspacesPages(*prometheusservice.ListWorkspacesInput, func(*prometheusservice.ListWorkspacesOutput, bool) bool) error
	ListWorkspacesPagesWithContext(aws.Context, *prometheusservice.ListWorkspacesInput, func(*prometheusservice.ListWorkspacesOutput, bool) bool, ...request.Option) error

	UpdateWorkspaceAlias(*prometheusservice.UpdateWorkspaceAliasInput) (*prometheusservice.UpdateWorkspaceAliasOutput, error)
	UpdateWorkspaceAliasWithContext(aws.Context, *prometheusservice.UpdateWorkspaceAliasInput, ...request.Option) (*prometheusservice.UpdateWorkspaceAliasOutput, error)
	UpdateWorkspaceAliasRequest(*prometheusservice.UpdateWorkspaceAliasInput) (*request.Request, *prometheusservice.UpdateWorkspaceAliasOutput)
}

var _ PrometheusServiceAPI = (*prometheusservice.PrometheusService)(nil)
