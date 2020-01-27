package gokong

import (
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func Test_WorkspaceGetById(t *testing.T) {

	workspaceRequest := &WorkspaceRequest{
		Name: "workspace-" + uuid.NewV4().String(),
	}

	client := NewClient(NewDefaultConfig())
	createdWorkspace, err := client.Workspaces().Create(workspaceRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdWorkspace)

	result, err := client.Workspaces().GetById(createdWorkspace.Id)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, createdWorkspace, result)

}

func Test_WorkspaceGetByName(t *testing.T) {

	workspaceRequest := &WorkspaceRequest{
		Name: "workspace-" + uuid.NewV4().String(),
	}

	client := NewClient(NewDefaultConfig())
	createdWorkspace, err := client.Workspaces().Create(workspaceRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdWorkspace)

	result, err := client.Workspaces().GetByName(createdWorkspace.Name)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, createdWorkspace, result)

}
