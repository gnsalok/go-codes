// gcpmanager_test.go

package gcpmanager

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/api/iam/v1"
)

// MockIAMService is a mock implementation of the IAMService interface for testing purposes.
type MockIAMService struct{}

func (m *MockIAMService) Projects() *iam.ProjectsService {
	return nil
}

func (m *MockIAMService) Groups() *iam.GroupsService {
	return nil
}

func (m *MockIAMService) Memberships() *iam.MembershipsService {
	return nil
}

func TestAddRemoveServiceAccountFromGroup(t *testing.T) {
	mockIAMService := & {} // You may use a mocking library for more complex cases

	manager := GCPManager{
		IAMService: mockIAMService,
	}

	// Test case 1: Add service account to group
	err := manager.AddRemoveServiceAccountFromGroup(context.Background(), "my-project", "my-group", "user@example.com", true)
	assert.NoError(t, err, "Error should be nil for adding service account to group")

	// Test case 2: Remove service account from group
	err = manager.AddRemoveServiceAccountFromGroup(context.Background(), "my-project", "my-group", "user@example.com", false)
	assert.NoError(t, err, "Error should be nil for removing service account from group")

	// Test case 3: Simulate an error
	mockIAMService = &MockIAMService{} // Reset the mock
	manager.IAMService = mockIAMService

	mockIAMService.Memberships = func() *iam.MembershipsService {
		return &iam.MembershipsService{
			Add: func(project, group string, membership *iam.Membership) *iam.MembershipsService {
				return &iam.MembershipsService{}
			},
			Remove: func(project, group string, membership *iam.Membership) *iam.MembershipsService {
				return &iam.MembershipsService{}
			},
		}
	}

	// Trigger an error condition by setting an empty project ID
	err = manager.AddRemoveServiceAccountFromGroup(context.Background(), "", "my-group", "user@example.com", true)
	assert.Error(t, err, "Error should occur for an empty project ID")
	assert.True(t, errors.Is(err, iam.ErrEmpty), "Error should be of type iam.ErrEmpty")
}
