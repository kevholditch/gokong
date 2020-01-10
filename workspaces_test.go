// +build all enterprise

package gokong

import (
	"fmt"
	"os"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func Test_WorkspaceClient_Get(t *testing.T) {
	workspaceRequest := &WorkspaceRequest{
		Name: String(fmt.Sprintf("workspace-name-%s", uuid.NewV4().String())),
	}

	client := NewClient(NewDefaultConfig())
	createdWorkspace, err := client.Workspaces().Create(workspaceRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdWorkspace)
	assert.EqualValues(t, createdWorkspace.Name, workspaceRequest.Name)

	result, err := client.Workspaces().Get(*createdWorkspace.Id)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, createdWorkspace, result)

	err = client.Workspaces().Delete()
	assert.Nil(t, err)
}

func Test_WorkspaceClient_List(t *testing.T) {
	workspaceRequest := &WorkspaceRequest{}
	createdWorkspaces := &Workspaces{}
	client := NewClient(NewDefaultConfig())

	for i := 0; i < 5; i++ {
		workspaceRequest.Name = String(fmt.Sprintf("workspace-name-%s", uuid.NewV4().String()))
		createdWorkspace, err := client.Workspaces().Create(workspaceRequest)
		assert.Nil(t, err)
		assert.NotNil(t, createdWorkspace)

		createdWorkspaces.Data = append(createdWorkspaces.Data, createdWorkspace)
	}

	result, err := client.Workspaces().List(&WorkspaceQueryString{})
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Subset(t, result, createdWorkspaces.Data)

	for range createdWorkspaces.Data {
		err = client.Workspaces().Delete()
		assert.Nil(t, err)
	}
}

func Test_WorkspaceClient_Update(t *testing.T) {
	config := NewDefaultConfig()
	config.Workspace = "default"
	client := NewClient(config)
	workspaceRequest := &WorkspaceRequest{
		Name: String(fmt.Sprintf("workspace-name-%s", uuid.NewV4().String())),
	}
	createdWorkspace, err := client.Workspaces().Create(workspaceRequest)
	assert.Nil(t, err)
	assert.NotNil(t, createdWorkspace)

	config.Workspace = *workspaceRequest.Name
	newWorkspaceClient := NewClient(config)
	workspaceRequest.Comment = String("This is a comment")
	updatedWorkspace, err := newWorkspaceClient.Workspaces().Update(workspaceRequest)
	assert.Nil(t, err)
	assert.NotNil(t, updatedWorkspace)

	result, err := client.Workspaces().Get(*createdWorkspace.Id)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, updatedWorkspace, result)

	err = newWorkspaceClient.Workspaces().Delete()
	assert.Nil(t, err)
}

func Test_Workspaces_GetNonExistent(t *testing.T) {
	client := NewClient(NewDefaultConfig())
	workspace, err := client.Workspaces().Get(uuid.NewV4().String())

	assert.Nil(t, workspace)
	assert.Nil(t, err)
}

func Test_WorkspaceClient_ListEntities(t *testing.T) {
	workspaceRequest := &WorkspaceRequest{
		Name: String(fmt.Sprintf("workspace-name-%s", uuid.NewV4().String())),
	}

	client := NewClient(NewDefaultConfig())
	createdWorkspace, err := client.Workspaces().Create(workspaceRequest)
	assert.NotNil(t, createdWorkspace)
	assert.Nil(t, err)

	os.Setenv(EnvKongWorkspace, *workspaceRequest.Name)
	workspaceClient := NewClient(NewDefaultConfig())

	serviceRequest := &ServiceRequest{
		Name:     String(fmt.Sprintf("service-name-%s", uuid.NewV4().String())),
		Protocol: String("http"),
		Host:     String("foo.com"),
		Port:     Int(8080),
	}
	createdService, err := workspaceClient.Services().Create(serviceRequest)

	assert.NotNil(t, createdService)
	assert.Nil(t, err)

	newWorkspaceEntities, err := workspaceClient.Workspaces().ListEntities()
	assert.Len(t, newWorkspaceEntities, 2)

	for _, workspaceEntity := range newWorkspaceEntities {
		if workspaceEntity.UniqueFieldName == String("name") {
			assert.Equal(t, *workspaceEntity.UniqueFieldValue, *createdService.Name)
		}
		if workspaceEntity.UniqueFieldName == String("id") {
			assert.Equal(t, *workspaceEntity.UniqueFieldValue, *createdService.Name)
		}
	}

	err = workspaceClient.Services().DeleteServiceById(*createdService.Id)
	assert.Nil(t, err)

	err = workspaceClient.Workspaces().Delete()
	assert.Nil(t, err)
}

func Test_WorkspaceClient_ShouldNotAllowToDeleteWorkspaceWithEntity(t *testing.T) {
	workspaceRequest := &WorkspaceRequest{
		Name: String("testingworkspace"),
	}

	client := NewClient(NewDefaultConfig())
	createdWorkspace, err := client.Workspaces().Create(workspaceRequest)
	assert.NotNil(t, createdWorkspace)
	assert.Nil(t, err)

	os.Setenv(EnvKongWorkspace, *workspaceRequest.Name)
	workspaceClient := NewClient(NewDefaultConfig())

	serviceRequest := &ServiceRequest{
		Name:     String(fmt.Sprintf("service-name-%s", uuid.NewV4().String())),
		Protocol: String("http"),
		Host:     String("foo.com"),
		Port:     Int(8080),
	}
	createdService, err := workspaceClient.Services().Create(serviceRequest)
	assert.NotNil(t, createdService)
	assert.Nil(t, err)

	newWorkspaceEntities, err := workspaceClient.Workspaces().ListEntities()
	assert.Len(t, newWorkspaceEntities, 2)

	err = workspaceClient.Workspaces().Delete()
	errorMessage := "bad request, message from kong: {\"message\":\"Workspace is not empty\"}"
	assert.Equal(t, err.Error(), errorMessage)
}

func Test_WorkspaceClient_Delete(t *testing.T) {
	workspaceRequest := &WorkspaceRequest{
		Name: String(fmt.Sprintf("workspace-name-%s", uuid.NewV4().String())),
	}

	config := NewDefaultConfig()
	config.Workspace = *workspaceRequest.Name
	client := NewClient(config)
	createdWorkspace, err := client.Workspaces().Create(workspaceRequest)
	assert.Nil(t, err)
	assert.NotNil(t, createdWorkspace)
	assert.EqualValues(t, createdWorkspace.Name, workspaceRequest.Name)

	result, err := client.Workspaces().Get(*createdWorkspace.Id)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, createdWorkspace, result)

	err = client.Workspaces().Delete()
	assert.Nil(t, err)
}

func Test_WorkspaceClient_DeleteMultipleEntitiesFromWorkspaceByIds(t *testing.T) {
	workspaceRequest := &WorkspaceRequest{
		Name: String(fmt.Sprintf("workspace-name-%s", uuid.NewV4().String())),
	}

	client := NewClient(NewDefaultConfig())
	createdWorkspace, err := client.Workspaces().Create(workspaceRequest)
	assert.NotNil(t, createdWorkspace)
	assert.Nil(t, err)

	os.Setenv(EnvKongWorkspace, *workspaceRequest.Name)
	workspaceClient := NewClient(NewDefaultConfig())

	serviceRequest := &ServiceRequest{
		Name:     String(fmt.Sprintf("service-name-%s", uuid.NewV4().String())),
		Protocol: String("http"),
		Host:     String("foo.com"),
		Port:     Int(8080),
	}
	createdService, err := workspaceClient.Services().Create(serviceRequest)

	assert.NotNil(t, createdService)
	assert.Nil(t, err)

	newWorkspaceEntities, err := workspaceClient.Workspaces().ListEntities()
	assert.Len(t, newWorkspaceEntities, 2)

	entityIds := []string{}
	for _, workspaceEntity := range newWorkspaceEntities {
		entityIds = append(entityIds, *workspaceEntity.EntityId)
		if workspaceEntity.UniqueFieldName == String("name") {
			assert.Equal(t, *workspaceEntity.UniqueFieldValue, *createdService.Name)
		}
		if workspaceEntity.UniqueFieldName == String("id") {
			assert.Equal(t, *workspaceEntity.UniqueFieldValue, *createdService.Name)
		}
	}

	err = workspaceClient.Workspaces().DeleteMultipleEntitiesFromWorkspace(
		entityIds,
	)
	assert.Nil(t, err)

	newWorkspaceEntities, err = workspaceClient.Workspaces().ListEntities()
	assert.Nil(t, err)
	assert.Len(t, newWorkspaceEntities, 0)

	err = workspaceClient.Workspaces().Delete()
	assert.Nil(t, err)
}
