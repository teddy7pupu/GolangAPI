// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package migrationhubiface provides an interface to enable mocking the AWS Migration Hub service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package migrationhubiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/migrationhub"
)

// MigrationHubAPI provides an interface to enable mocking the
// migrationhub.MigrationHub service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Migration Hub.
//    func myFunc(svc migrationhubiface.MigrationHubAPI) bool {
//        // Make svc.AssociateCreatedArtifact request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := migrationhub.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockMigrationHubClient struct {
//        migrationhubiface.MigrationHubAPI
//    }
//    func (m *mockMigrationHubClient) AssociateCreatedArtifact(input *migrationhub.AssociateCreatedArtifactInput) (*migrationhub.AssociateCreatedArtifactOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockMigrationHubClient{}
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
type MigrationHubAPI interface {
	AssociateCreatedArtifact(*migrationhub.AssociateCreatedArtifactInput) (*migrationhub.AssociateCreatedArtifactOutput, error)
	AssociateCreatedArtifactWithContext(aws.Context, *migrationhub.AssociateCreatedArtifactInput, ...request.Option) (*migrationhub.AssociateCreatedArtifactOutput, error)
	AssociateCreatedArtifactRequest(*migrationhub.AssociateCreatedArtifactInput) (*request.Request, *migrationhub.AssociateCreatedArtifactOutput)

	AssociateDiscoveredResource(*migrationhub.AssociateDiscoveredResourceInput) (*migrationhub.AssociateDiscoveredResourceOutput, error)
	AssociateDiscoveredResourceWithContext(aws.Context, *migrationhub.AssociateDiscoveredResourceInput, ...request.Option) (*migrationhub.AssociateDiscoveredResourceOutput, error)
	AssociateDiscoveredResourceRequest(*migrationhub.AssociateDiscoveredResourceInput) (*request.Request, *migrationhub.AssociateDiscoveredResourceOutput)

	CreateProgressUpdateStream(*migrationhub.CreateProgressUpdateStreamInput) (*migrationhub.CreateProgressUpdateStreamOutput, error)
	CreateProgressUpdateStreamWithContext(aws.Context, *migrationhub.CreateProgressUpdateStreamInput, ...request.Option) (*migrationhub.CreateProgressUpdateStreamOutput, error)
	CreateProgressUpdateStreamRequest(*migrationhub.CreateProgressUpdateStreamInput) (*request.Request, *migrationhub.CreateProgressUpdateStreamOutput)

	DeleteProgressUpdateStream(*migrationhub.DeleteProgressUpdateStreamInput) (*migrationhub.DeleteProgressUpdateStreamOutput, error)
	DeleteProgressUpdateStreamWithContext(aws.Context, *migrationhub.DeleteProgressUpdateStreamInput, ...request.Option) (*migrationhub.DeleteProgressUpdateStreamOutput, error)
	DeleteProgressUpdateStreamRequest(*migrationhub.DeleteProgressUpdateStreamInput) (*request.Request, *migrationhub.DeleteProgressUpdateStreamOutput)

	DescribeApplicationState(*migrationhub.DescribeApplicationStateInput) (*migrationhub.DescribeApplicationStateOutput, error)
	DescribeApplicationStateWithContext(aws.Context, *migrationhub.DescribeApplicationStateInput, ...request.Option) (*migrationhub.DescribeApplicationStateOutput, error)
	DescribeApplicationStateRequest(*migrationhub.DescribeApplicationStateInput) (*request.Request, *migrationhub.DescribeApplicationStateOutput)

	DescribeMigrationTask(*migrationhub.DescribeMigrationTaskInput) (*migrationhub.DescribeMigrationTaskOutput, error)
	DescribeMigrationTaskWithContext(aws.Context, *migrationhub.DescribeMigrationTaskInput, ...request.Option) (*migrationhub.DescribeMigrationTaskOutput, error)
	DescribeMigrationTaskRequest(*migrationhub.DescribeMigrationTaskInput) (*request.Request, *migrationhub.DescribeMigrationTaskOutput)

	DisassociateCreatedArtifact(*migrationhub.DisassociateCreatedArtifactInput) (*migrationhub.DisassociateCreatedArtifactOutput, error)
	DisassociateCreatedArtifactWithContext(aws.Context, *migrationhub.DisassociateCreatedArtifactInput, ...request.Option) (*migrationhub.DisassociateCreatedArtifactOutput, error)
	DisassociateCreatedArtifactRequest(*migrationhub.DisassociateCreatedArtifactInput) (*request.Request, *migrationhub.DisassociateCreatedArtifactOutput)

	DisassociateDiscoveredResource(*migrationhub.DisassociateDiscoveredResourceInput) (*migrationhub.DisassociateDiscoveredResourceOutput, error)
	DisassociateDiscoveredResourceWithContext(aws.Context, *migrationhub.DisassociateDiscoveredResourceInput, ...request.Option) (*migrationhub.DisassociateDiscoveredResourceOutput, error)
	DisassociateDiscoveredResourceRequest(*migrationhub.DisassociateDiscoveredResourceInput) (*request.Request, *migrationhub.DisassociateDiscoveredResourceOutput)

	ImportMigrationTask(*migrationhub.ImportMigrationTaskInput) (*migrationhub.ImportMigrationTaskOutput, error)
	ImportMigrationTaskWithContext(aws.Context, *migrationhub.ImportMigrationTaskInput, ...request.Option) (*migrationhub.ImportMigrationTaskOutput, error)
	ImportMigrationTaskRequest(*migrationhub.ImportMigrationTaskInput) (*request.Request, *migrationhub.ImportMigrationTaskOutput)

	ListApplicationStates(*migrationhub.ListApplicationStatesInput) (*migrationhub.ListApplicationStatesOutput, error)
	ListApplicationStatesWithContext(aws.Context, *migrationhub.ListApplicationStatesInput, ...request.Option) (*migrationhub.ListApplicationStatesOutput, error)
	ListApplicationStatesRequest(*migrationhub.ListApplicationStatesInput) (*request.Request, *migrationhub.ListApplicationStatesOutput)

	ListApplicationStatesPages(*migrationhub.ListApplicationStatesInput, func(*migrationhub.ListApplicationStatesOutput, bool) bool) error
	ListApplicationStatesPagesWithContext(aws.Context, *migrationhub.ListApplicationStatesInput, func(*migrationhub.ListApplicationStatesOutput, bool) bool, ...request.Option) error

	ListCreatedArtifacts(*migrationhub.ListCreatedArtifactsInput) (*migrationhub.ListCreatedArtifactsOutput, error)
	ListCreatedArtifactsWithContext(aws.Context, *migrationhub.ListCreatedArtifactsInput, ...request.Option) (*migrationhub.ListCreatedArtifactsOutput, error)
	ListCreatedArtifactsRequest(*migrationhub.ListCreatedArtifactsInput) (*request.Request, *migrationhub.ListCreatedArtifactsOutput)

	ListCreatedArtifactsPages(*migrationhub.ListCreatedArtifactsInput, func(*migrationhub.ListCreatedArtifactsOutput, bool) bool) error
	ListCreatedArtifactsPagesWithContext(aws.Context, *migrationhub.ListCreatedArtifactsInput, func(*migrationhub.ListCreatedArtifactsOutput, bool) bool, ...request.Option) error

	ListDiscoveredResources(*migrationhub.ListDiscoveredResourcesInput) (*migrationhub.ListDiscoveredResourcesOutput, error)
	ListDiscoveredResourcesWithContext(aws.Context, *migrationhub.ListDiscoveredResourcesInput, ...request.Option) (*migrationhub.ListDiscoveredResourcesOutput, error)
	ListDiscoveredResourcesRequest(*migrationhub.ListDiscoveredResourcesInput) (*request.Request, *migrationhub.ListDiscoveredResourcesOutput)

	ListDiscoveredResourcesPages(*migrationhub.ListDiscoveredResourcesInput, func(*migrationhub.ListDiscoveredResourcesOutput, bool) bool) error
	ListDiscoveredResourcesPagesWithContext(aws.Context, *migrationhub.ListDiscoveredResourcesInput, func(*migrationhub.ListDiscoveredResourcesOutput, bool) bool, ...request.Option) error

	ListMigrationTasks(*migrationhub.ListMigrationTasksInput) (*migrationhub.ListMigrationTasksOutput, error)
	ListMigrationTasksWithContext(aws.Context, *migrationhub.ListMigrationTasksInput, ...request.Option) (*migrationhub.ListMigrationTasksOutput, error)
	ListMigrationTasksRequest(*migrationhub.ListMigrationTasksInput) (*request.Request, *migrationhub.ListMigrationTasksOutput)

	ListMigrationTasksPages(*migrationhub.ListMigrationTasksInput, func(*migrationhub.ListMigrationTasksOutput, bool) bool) error
	ListMigrationTasksPagesWithContext(aws.Context, *migrationhub.ListMigrationTasksInput, func(*migrationhub.ListMigrationTasksOutput, bool) bool, ...request.Option) error

	ListProgressUpdateStreams(*migrationhub.ListProgressUpdateStreamsInput) (*migrationhub.ListProgressUpdateStreamsOutput, error)
	ListProgressUpdateStreamsWithContext(aws.Context, *migrationhub.ListProgressUpdateStreamsInput, ...request.Option) (*migrationhub.ListProgressUpdateStreamsOutput, error)
	ListProgressUpdateStreamsRequest(*migrationhub.ListProgressUpdateStreamsInput) (*request.Request, *migrationhub.ListProgressUpdateStreamsOutput)

	ListProgressUpdateStreamsPages(*migrationhub.ListProgressUpdateStreamsInput, func(*migrationhub.ListProgressUpdateStreamsOutput, bool) bool) error
	ListProgressUpdateStreamsPagesWithContext(aws.Context, *migrationhub.ListProgressUpdateStreamsInput, func(*migrationhub.ListProgressUpdateStreamsOutput, bool) bool, ...request.Option) error

	NotifyApplicationState(*migrationhub.NotifyApplicationStateInput) (*migrationhub.NotifyApplicationStateOutput, error)
	NotifyApplicationStateWithContext(aws.Context, *migrationhub.NotifyApplicationStateInput, ...request.Option) (*migrationhub.NotifyApplicationStateOutput, error)
	NotifyApplicationStateRequest(*migrationhub.NotifyApplicationStateInput) (*request.Request, *migrationhub.NotifyApplicationStateOutput)

	NotifyMigrationTaskState(*migrationhub.NotifyMigrationTaskStateInput) (*migrationhub.NotifyMigrationTaskStateOutput, error)
	NotifyMigrationTaskStateWithContext(aws.Context, *migrationhub.NotifyMigrationTaskStateInput, ...request.Option) (*migrationhub.NotifyMigrationTaskStateOutput, error)
	NotifyMigrationTaskStateRequest(*migrationhub.NotifyMigrationTaskStateInput) (*request.Request, *migrationhub.NotifyMigrationTaskStateOutput)

	PutResourceAttributes(*migrationhub.PutResourceAttributesInput) (*migrationhub.PutResourceAttributesOutput, error)
	PutResourceAttributesWithContext(aws.Context, *migrationhub.PutResourceAttributesInput, ...request.Option) (*migrationhub.PutResourceAttributesOutput, error)
	PutResourceAttributesRequest(*migrationhub.PutResourceAttributesInput) (*request.Request, *migrationhub.PutResourceAttributesOutput)
}

var _ MigrationHubAPI = (*migrationhub.MigrationHub)(nil)
