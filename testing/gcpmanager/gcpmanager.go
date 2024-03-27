// gcpmanager.go

package gcpmanager

import (
	"context"
	"fmt"

	"google.golang.org/api/iam/v1"
)

// GCPManager is a struct representing the GCP IAM manager.
type GCPManager struct {
	IAMService *iam.Service
}

// AddRemoveServiceAccountFromGroup adds or removes a service account from an IAM group.
// If addToGroup is true, the service account will be added. If false, it will be removed.
func (m *GCPManager) AddRemoveServiceAccountFromGroup(ctx context.Context, projectID, groupName, serviceAccountEmail string, addToGroup bool) error {
	members := []string{fmt.Sprintf("serviceAccount:%s", serviceAccountEmail)}

	var modifyFunc func(*iam.MembershipsService) *iam.MembershipsService

	if addToGroup {
		modifyFunc = m.IAMService.Projects.Groups.Memberships.Add
	} else {
		modifyFunc = m.IAMService.Projects.Groups.Memberships.Remove
	}

	_, err := modifyFunc(projectID, groupName, &iam.Membership{
		Members: members,
	}).Context(ctx).Do()

	return err
}
