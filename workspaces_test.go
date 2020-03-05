package gokong

import (
	"fmt"
	"os"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func skipEnterprise(t *testing.T) {
	if os.Getenv("KONG_LICENSE") == "" {
		t.Skip("Skipping enterprise feature test")
	}
}

func Test_WorkspaceGetById(t *testing.T) {

	skipEnterprise(t)

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

	skipEnterprise(t)

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

func Test_WorkspaceGetByIdForNonExistentWorkspaceId(t *testing.T) {

	skipEnterprise(t)

	result, err := NewClient(NewDefaultConfig()).Workspaces().GetById(uuid.NewV4().String())

	assert.Nil(t, err)
	assert.Nil(t, result)
}

func Test_WorkspaceGetByIdForNonExistentWorkspaceByName(t *testing.T) {

	skipEnterprise(t)

	result, err := NewClient(NewDefaultConfig()).Workspaces().GetByName(uuid.NewV4().String())

	assert.Nil(t, err)
	assert.Nil(t, result)

}

func Test_WorkspaceCreate(t *testing.T) {

	skipEnterprise(t)

	workspaceRequest := &WorkspaceRequest{
		Name:    "workspace-" + uuid.NewV4().String(),
		Comment: "testing",
		Meta: map[string]interface{}{
			"color":     nil,
			"thumbnail": nil,
		},
		Config: map[string]interface{}{
			"meta":                         nil,
			"portal":                       false,
			"portal_access_request_email":  nil,
			"portal_approved_email":        nil,
			"portal_auth":                  nil,
			"portal_auth_conf":             nil,
			"portal_auto_approve":          nil,
			"portal_cors_origins":          nil,
			"portal_developer_meta_fields": "[{\"label\":\"Full Name\",\"title\":\"full_name\",\"validator\":{\"required\":true,\"type\":\"string\"}}]",
			"portal_emails_from":           nil,
			"portal_emails_reply_to":       nil,
			"portal_invite_email":          nil,
			"portal_reset_email":           nil,
			"portal_reset_success_email":   nil,
			"portal_token_exp":             nil,
		},
	}

	client := NewClient(NewDefaultConfig())
	createdWorkspace, err := client.Workspaces().Create(workspaceRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdWorkspace)
	assert.Equal(t, workspaceRequest.Name, createdWorkspace.Name)
	assert.Equal(t, workspaceRequest.Comment, createdWorkspace.Comment)
}
func Test_WorkspaceCreateWithPortal(t *testing.T) {

	skipEnterprise(t)

	workspaceRequest := &WorkspaceRequest{
		Name:    "workspace-withportal" + uuid.NewV4().String(),
		Comment: "testing",
		Meta: map[string]interface{}{
			"color":     nil,
			"thumbnail": nil,
		},
		Config: map[string]interface{}{
			"meta":                         nil,
			"portal":                       true,
			"portal_access_request_email":  nil,
			"portal_approved_email":        nil,
			"portal_auth":                  nil,
			"portal_auth_conf":             nil,
			"portal_auto_approve":          nil,
			"portal_cors_origins":          nil,
			"portal_developer_meta_fields": "[{\"label\":\"Full Name\",\"title\":\"full_name\",\"validator\":{\"required\":true,\"type\":\"string\"}}]",
			"portal_emails_from":           nil,
			"portal_emails_reply_to":       nil,
			"portal_invite_email":          nil,
			"portal_reset_email":           nil,
			"portal_reset_success_email":   nil,
			"portal_token_exp":             nil,
		},
	}

	client := NewClient(NewDefaultConfig())
	createdWorkspace, err := client.Workspaces().Create(workspaceRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdWorkspace)
	assert.Equal(t, workspaceRequest.Name, createdWorkspace.Name)
	assert.Equal(t, workspaceRequest.Comment, createdWorkspace.Comment)
}

func Test_WorkspaceList(t *testing.T) {

	skipEnterprise(t)

	workspaceRequest1 := &WorkspaceRequest{
		Name:    "workspace-" + uuid.NewV4().String(),
		Comment: "testing",
	}
	workspaceRequest2 := &WorkspaceRequest{
		Name:    "workspace-" + uuid.NewV4().String(),
		Comment: "testing 2",
	}
	workspaceRequest3 := &WorkspaceRequest{
		Name:    "workspace-" + uuid.NewV4().String(),
		Comment: "testing 3",
	}

	client := NewClient(NewDefaultConfig())
	createdWorkspace1, err := client.Workspaces().Create(workspaceRequest1)
	createdWorkspace2, err := client.Workspaces().Create(workspaceRequest2)
	createdWorkspace3, err := client.Workspaces().Create(workspaceRequest3)

	assert.Nil(t, err)
	assert.NotNil(t, createdWorkspace1)
	assert.NotNil(t, createdWorkspace2)
	assert.NotNil(t, createdWorkspace3)

	workspaces, err := client.Workspaces().List()
	assert.Nil(t, err)
	assert.NotNil(t, workspaces)
	assert.True(t, len(workspaces.Data) > 0)
}

func Test_WorkspacesDeleteById(t *testing.T) {

	skipEnterprise(t)

	client := NewClient(NewDefaultConfig())

	workspaceRequest := &WorkspaceRequest{
		Name: "workspace-" + uuid.NewV4().String(),
	}

	createdWorkspace, err := client.Workspaces().Create(workspaceRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdWorkspace)

	err = client.Workspaces().DeleteById(createdWorkspace.Id)
	assert.Nil(t, err)

	result, err := client.Workspaces().GetById(createdWorkspace.Id)
	assert.Nil(t, err)
	assert.Nil(t, result)
}
func Test_WorkspacesDeleteByName(t *testing.T) {

	skipEnterprise(t)

	client := NewClient(NewDefaultConfig())

	workspaceRequest := &WorkspaceRequest{
		Name: "workspace-" + uuid.NewV4().String(),
	}

	createdWorkspace, err := client.Workspaces().Create(workspaceRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdWorkspace)

	err = client.Workspaces().DeleteByName(createdWorkspace.Name)
	assert.Nil(t, err)

	result, err := client.Workspaces().GetByName(createdWorkspace.Name)
	assert.Nil(t, err)
	assert.Nil(t, result)
}

func Test_WorkspacesUpdateById(t *testing.T) {

	skipEnterprise(t)

	client := NewClient(NewDefaultConfig())

	workspaceRequest := &WorkspaceRequest{
		Name:    "workspace-" + uuid.NewV4().String(),
		Comment: "testing",
		Meta: map[string]interface{}{
			"color":     nil,
			"thumbnail": nil,
		},
		Config: map[string]interface{}{
			"meta":                         nil,
			"portal":                       false,
			"portal_access_request_email":  nil,
			"portal_approved_email":        nil,
			"portal_auth":                  nil,
			"portal_auth_conf":             nil,
			"portal_auto_approve":          nil,
			"portal_cors_origins":          nil,
			"portal_developer_meta_fields": "[{\"label\":\"Full Name\",\"title\":\"full_name\",\"validator\":{\"required\":true,\"type\":\"string\"}}]",
			"portal_emails_from":           nil,
			"portal_emails_reply_to":       nil,
			"portal_invite_email":          nil,
			"portal_reset_email":           nil,
			"portal_reset_success_email":   nil,
			"portal_token_exp":             nil,
		},
	}

	createdWorkspace, err := client.Workspaces().Create(workspaceRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdWorkspace)
	assert.Equal(t, "testing", createdWorkspace.Comment)

	workspaceRequest.Comment = "new comment"

	result, err := client.Workspaces().UpdateById(createdWorkspace.Id, workspaceRequest)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "new comment", result.Comment)
}
func Test_WorkspacesUpdateByName(t *testing.T) {

	skipEnterprise(t)

	client := NewClient(NewDefaultConfig())

	workspaceRequest := &WorkspaceRequest{
		Name:    "workspace-" + uuid.NewV4().String(),
		Comment: "testing",
		Meta: map[string]interface{}{
			"color":     nil,
			"thumbnail": nil,
		},
		Config: map[string]interface{}{
			"meta":                         nil,
			"portal":                       false,
			"portal_access_request_email":  nil,
			"portal_approved_email":        nil,
			"portal_auth":                  nil,
			"portal_auth_conf":             nil,
			"portal_auto_approve":          nil,
			"portal_cors_origins":          nil,
			"portal_developer_meta_fields": "[{\"label\":\"Full Name\",\"title\":\"full_name\",\"validator\":{\"required\":true,\"type\":\"string\"}}]",
			"portal_emails_from":           nil,
			"portal_emails_reply_to":       nil,
			"portal_invite_email":          nil,
			"portal_reset_email":           nil,
			"portal_reset_success_email":   nil,
			"portal_token_exp":             nil,
		},
	}

	createdWorkspace, err := client.Workspaces().Create(workspaceRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdWorkspace)
	assert.Equal(t, "testing", createdWorkspace.Comment)

	workspaceRequest.Comment = "new comment"

	result, err := client.Workspaces().UpdateByName(createdWorkspace.Name, workspaceRequest)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "new comment", result.Comment)
}

func Test_WorkspacesUpdateByIdInvalid(t *testing.T) {

	skipEnterprise(t)

	client := NewClient(NewDefaultConfig())

	workspaceRequest := &WorkspaceRequest{
		Name:    "workspace-" + uuid.NewV4().String(),
		Comment: "testing",
		Meta: map[string]interface{}{
			"color":     nil,
			"thumbnail": nil,
		},
		Config: map[string]interface{}{
			"meta":                         nil,
			"portal":                       false,
			"portal_access_request_email":  nil,
			"portal_approved_email":        nil,
			"portal_auth":                  nil,
			"portal_auth_conf":             nil,
			"portal_auto_approve":          nil,
			"portal_cors_origins":          nil,
			"portal_developer_meta_fields": "[{\"label\":\"Full Name\",\"title\":\"full_name\",\"validator\":{\"required\":true,\"type\":\"string\"}}]",
			"portal_emails_from":           nil,
			"portal_emails_reply_to":       nil,
			"portal_invite_email":          nil,
			"portal_reset_email":           nil,
			"portal_reset_success_email":   nil,
			"portal_token_exp":             nil,
		},
	}

	createdWorkspace, err := client.Workspaces().Create(workspaceRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdWorkspace)
	assert.Equal(t, "testing", createdWorkspace.Comment)

	workspaceRequest.Name = "test-workspace"

	result, err := client.Workspaces().UpdateById(createdWorkspace.Id, workspaceRequest)

	// Updating workspace names is not allowed.
	assert.NotNil(t, err)
	assert.Nil(t, result)
}
func Test_WorkspacesUpdateByNameInvalid(t *testing.T) {

	skipEnterprise(t)

	client := NewClient(NewDefaultConfig())

	workspaceRequest := &WorkspaceRequest{
		Name:    "workspace-" + uuid.NewV4().String(),
		Comment: "testing",
		Meta: map[string]interface{}{
			"color":     nil,
			"thumbnail": nil,
		},
		Config: map[string]interface{}{
			"meta":                         nil,
			"portal":                       false,
			"portal_access_request_email":  nil,
			"portal_approved_email":        nil,
			"portal_auth":                  nil,
			"portal_auth_conf":             nil,
			"portal_auto_approve":          nil,
			"portal_cors_origins":          nil,
			"portal_developer_meta_fields": "[{\"label\":\"Full Name\",\"title\":\"full_name\",\"validator\":{\"required\":true,\"type\":\"string\"}}]",
			"portal_emails_from":           nil,
			"portal_emails_reply_to":       nil,
			"portal_invite_email":          nil,
			"portal_reset_email":           nil,
			"portal_reset_success_email":   nil,
			"portal_token_exp":             nil,
		},
	}

	createdWorkspace, err := client.Workspaces().Create(workspaceRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdWorkspace)
	assert.Equal(t, "testing", createdWorkspace.Comment)

	workspaceRequest.Name = "test-workspace"

	result, err := client.Workspaces().UpdateByName(createdWorkspace.Name, workspaceRequest)

	// Updating workspace names is not allowed.
	assert.NotNil(t, err)
	assert.Nil(t, result)
}

func Test_AllWorkspaceEndpointsShouldReturnErrorWhenRequestUnauthorised(t *testing.T) {

	skipEnterprise(t)

	unauthorisedClient := NewClient(&Config{HostAddress: kong401Server})

	workspace, err := unauthorisedClient.Workspaces().GetByName("foo")
	assert.NotNil(t, err)
	assert.Nil(t, workspace)

	workspace, err = unauthorisedClient.Workspaces().GetById(uuid.NewV4().String())
	assert.NotNil(t, err)
	assert.Nil(t, workspace)

	results, err := unauthorisedClient.Workspaces().List()
	assert.NotNil(t, err)
	assert.Nil(t, results)

	err = unauthorisedClient.Workspaces().DeleteByName("bar")
	assert.NotNil(t, err)

	err = unauthorisedClient.Workspaces().DeleteById(uuid.NewV4().String())
	assert.NotNil(t, err)

	workspaceResult, err := unauthorisedClient.Workspaces().Create(&WorkspaceRequest{
		Name: "workspace-" + uuid.NewV4().String(),
	})
	assert.Nil(t, workspaceResult)
	assert.NotNil(t, err)

	updatedWorkspace, err := unauthorisedClient.Workspaces().UpdateByName("foo", &WorkspaceRequest{
		Name: "workspace-" + uuid.NewV4().String(),
	})
	assert.Nil(t, updatedWorkspace)
	assert.NotNil(t, err)

	updatedWorkspace, err = unauthorisedClient.Workspaces().UpdateById(uuid.NewV4().String(), &WorkspaceRequest{
		Name: "workspace-" + uuid.NewV4().String(),
	})
	assert.Nil(t, updatedWorkspace)
	assert.NotNil(t, err)

}

func TestWorkspaceAddEntitiesById(t *testing.T) {
	skipEnterprise(t)

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

	serviceRequest := &ServiceRequest{
		Name:     String(fmt.Sprintf("service-name-%s", uuid.NewV4().String())),
		Protocol: String("http"),
		Host:     String("foo.com"),
		Port:     Int(8080),
	}

	createdService, err := client.Services().Create(serviceRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdService)

	entityRequest := &EntityRequest{
		Entities: createdService.Id,
	}

	updatedEntities, err := client.Workspaces().AddEntitiesByWorkspaceId(result.Id, entityRequest)

	assert.Nil(t, err)
	assert.NotNil(t, updatedEntities)
}
func TestWorkspaceAddEntitiesByName(t *testing.T) {
	skipEnterprise(t)

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

	serviceRequest := &ServiceRequest{
		Name:     String(fmt.Sprintf("service-name-%s", uuid.NewV4().String())),
		Protocol: String("http"),
		Host:     String("foo.com"),
		Port:     Int(8080),
	}

	createdService, err := client.Services().Create(serviceRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdService)

	entityRequest := &EntityRequest{
		Entities: createdService.Id,
	}

	updatedEntities, err := client.Workspaces().AddEntitiesByWorkspaceName(result.Name, entityRequest)

	assert.Nil(t, err)
	assert.NotNil(t, updatedEntities)
}

func TestWorkspaceListEntitiesById(t *testing.T) {
	skipEnterprise(t)

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

	serviceRequest := &ServiceRequest{
		Name:     String(fmt.Sprintf("service-name-%s", uuid.NewV4().String())),
		Protocol: String("http"),
		Host:     String("foo.com"),
		Port:     Int(8080),
	}

	createdService, err := client.Services().Create(serviceRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdService)

	entityRequest := &EntityRequest{
		Entities: createdService.Id,
	}

	updatedEntities, err := client.Workspaces().AddEntitiesByWorkspaceName(result.Name, entityRequest)

	assert.Nil(t, err)
	assert.NotNil(t, updatedEntities)

	entities, err := client.Workspaces().ListEntitiesById(createdWorkspace.Id)

	assert.Nil(t, err)
	assert.NotNil(t, entities)
	assert.True(t, len(entities.Data) > 0)

}

func TestWorkspaceListEntitiesByName(t *testing.T) {
	skipEnterprise(t)

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

	serviceRequest := &ServiceRequest{
		Name:     String(fmt.Sprintf("service-name-%s", uuid.NewV4().String())),
		Protocol: String("http"),
		Host:     String("foo.com"),
		Port:     Int(8080),
	}

	createdService, err := client.Services().Create(serviceRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdService)

	entityRequest := &EntityRequest{
		Entities: createdService.Id,
	}

	updatedEntities, err := client.Workspaces().AddEntitiesByWorkspaceName(result.Name, entityRequest)

	assert.Nil(t, err)
	assert.NotNil(t, updatedEntities)

	entities, err := client.Workspaces().ListEntitiesByName(createdWorkspace.Id)

	assert.Nil(t, err)
	assert.NotNil(t, entities)
	assert.True(t, len(entities.Data) > 0)

}

func TestWorkspaceGetEntityById(t *testing.T) {
	skipEnterprise(t)

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

	serviceRequest := &ServiceRequest{
		Name:     String(fmt.Sprintf("service-name-%s", uuid.NewV4().String())),
		Protocol: String("http"),
		Host:     String("foo.com"),
		Port:     Int(8080),
	}

	createdService, err := client.Services().Create(serviceRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdService)

	entityRequest := &EntityRequest{
		Entities: createdService.Id,
	}

	updatedEntities, err := client.Workspaces().AddEntitiesByWorkspaceName(result.Name, entityRequest)

	assert.Nil(t, err)
	assert.NotNil(t, updatedEntities)

	entity, err := client.Workspaces().GetEntityById(createdWorkspace.Id, updatedEntities.Entities[0]["id"].(string))

	assert.Nil(t, err)
	assert.NotNil(t, entity)

}

func TestWorkspaceGetEntityByName(t *testing.T) {
	skipEnterprise(t)

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

	serviceRequest := &ServiceRequest{
		Name:     String(fmt.Sprintf("service-name-%s", uuid.NewV4().String())),
		Protocol: String("http"),
		Host:     String("foo.com"),
		Port:     Int(8080),
	}

	createdService, err := client.Services().Create(serviceRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdService)

	entityRequest := &EntityRequest{
		Entities: createdService.Id,
	}

	updatedEntities, err := client.Workspaces().AddEntitiesByWorkspaceName(result.Name, entityRequest)

	assert.Nil(t, err)
	assert.NotNil(t, updatedEntities)

	entity, err := client.Workspaces().GetEntityByName(createdWorkspace.Id, updatedEntities.Entities[0]["id"].(string))

	assert.Nil(t, err)
	assert.NotNil(t, entity)

}

func TestWorkspaceDeleteEntityById(t *testing.T) {
	skipEnterprise(t)

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

	serviceRequest := &ServiceRequest{
		Name:     String(fmt.Sprintf("service-name-%s", uuid.NewV4().String())),
		Protocol: String("http"),
		Host:     String("foo.com"),
		Port:     Int(8080),
	}

	createdService, err := client.Services().Create(serviceRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdService)

	entityRequest := &EntityRequest{
		Entities: createdService.Id,
	}

	updatedEntities, err := client.Workspaces().AddEntitiesByWorkspaceName(result.Name, entityRequest)

	assert.Nil(t, err)
	assert.NotNil(t, updatedEntities)

	err = client.Workspaces().DeleteEntityById(createdWorkspace.Id, updatedEntities.Entities[0]["id"].(string))

	assert.Nil(t, err)

}

func TestWorkspaceDeleteEntityByName(t *testing.T) {
	skipEnterprise(t)

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

	serviceRequest := &ServiceRequest{
		Name:     String(fmt.Sprintf("service-name-%s", uuid.NewV4().String())),
		Protocol: String("http"),
		Host:     String("foo.com"),
		Port:     Int(8080),
	}

	createdService, err := client.Services().Create(serviceRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdService)

	entityRequest := &EntityRequest{
		Entities: createdService.Id,
	}

	updatedEntities, err := client.Workspaces().AddEntitiesByWorkspaceName(result.Name, entityRequest)

	assert.Nil(t, err)
	assert.NotNil(t, updatedEntities)

	err = client.Workspaces().DeleteEntityByName(createdWorkspace.Id, updatedEntities.Entities[0]["id"].(string))

	assert.Nil(t, err)

}
