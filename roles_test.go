package gokong

import (
	"fmt"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func Test_RoleGetByIdWorkspace(t *testing.T) {

	skipEnterprise(t)

	roleRequest := &RoleRequest{
		Name: "role-" + uuid.NewV4().String(),
	}

	workspaceRequest := &WorkspaceRequest{
		Name: "workspace-roleadd-" + uuid.NewV4().String(),
	}

	client := NewClient(NewDefaultConfig())

	createdWorkspace, err := client.Workspaces().Create(workspaceRequest)

	client = NewClient(NewWorkspaceConfig(createdWorkspace.Name))

	createdRole, err := client.Roles().Create(roleRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdRole)

	result, err := client.Roles().GetById(createdRole.Id)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, createdRole, result)

}
func Test_RoleGetById(t *testing.T) {

	skipEnterprise(t)

	roleRequest := &RoleRequest{
		Name: "role-" + uuid.NewV4().String(),
	}

	client := NewClient(NewDefaultConfig())
	createdRole, err := client.Roles().Create(roleRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdRole)

	result, err := client.Roles().GetById(createdRole.Id)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, createdRole, result)

}
func Test_RoleGetByName(t *testing.T) {

	skipEnterprise(t)

	roleRequest := &RoleRequest{
		Name: "role-" + uuid.NewV4().String(),
	}

	client := NewClient(NewDefaultConfig())
	createdRole, err := client.Roles().Create(roleRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdRole)

	result, err := client.Roles().GetByName(createdRole.Name)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, createdRole, result)

}

func Test_RoleGetByIdForNonExistentRoleId(t *testing.T) {

	skipEnterprise(t)

	result, err := NewClient(NewDefaultConfig()).Roles().GetById(uuid.NewV4().String())

	assert.Nil(t, err)
	assert.Nil(t, result)
}

func Test_RoleGetByIdForNonExistentRoleByName(t *testing.T) {

	skipEnterprise(t)

	result, err := NewClient(NewDefaultConfig()).Roles().GetByName(uuid.NewV4().String())

	assert.Nil(t, err)
	assert.Nil(t, result)

}

func Test_RoleCreate(t *testing.T) {

	skipEnterprise(t)

	roleRequest := &RoleRequest{
		Name:    "role-" + uuid.NewV4().String(),
		Comment: "testing",
	}

	client := NewClient(NewDefaultConfig())
	createdRole, err := client.Roles().Create(roleRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdRole)
	assert.Equal(t, roleRequest.Name, createdRole.Name)
	assert.Equal(t, roleRequest.Comment, createdRole.Comment)
}

func Test_RoleList(t *testing.T) {

	skipEnterprise(t)

	roleRequest1 := &RoleRequest{
		Name: "role-" + uuid.NewV4().String(),
	}
	roleRequest2 := &RoleRequest{
		Name: "role-" + uuid.NewV4().String(),
	}
	roleRequest3 := &RoleRequest{
		Name: "role-" + uuid.NewV4().String(),
	}

	client := NewClient(NewDefaultConfig())
	createdRole1, err := client.Roles().Create(roleRequest1)
	createdRole2, err := client.Roles().Create(roleRequest2)
	createdRole3, err := client.Roles().Create(roleRequest3)

	assert.Nil(t, err)
	assert.NotNil(t, createdRole1)
	assert.NotNil(t, createdRole2)
	assert.NotNil(t, createdRole3)

	roles, err := client.Roles().List()
	assert.Nil(t, err)
	assert.NotNil(t, roles)
	assert.True(t, len(roles.Data) > 0)
}

func Test_RolesUpdateById(t *testing.T) {

	skipEnterprise(t)

	client := NewClient(NewDefaultConfig())

	roleRequest := &RoleRequest{
		Name:    "role-" + uuid.NewV4().String(),
		Comment: "testing",
	}

	createdRole, err := client.Roles().Create(roleRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdRole)
	assert.Equal(t, "testing", createdRole.Comment)

	roleRequest.Comment = "new comment"

	result, err := client.Roles().UpdateById(createdRole.Id, roleRequest)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "new comment", result.Comment)
}

func Test_RolesUpdateByName(t *testing.T) {

	skipEnterprise(t)

	client := NewClient(NewDefaultConfig())

	roleRequest := &RoleRequest{
		Name:    "role-" + uuid.NewV4().String(),
		Comment: "testing",
	}

	createdRole, err := client.Roles().Create(roleRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdRole)
	assert.Equal(t, "testing", createdRole.Comment)

	roleRequest.Comment = "new comment"

	result, err := client.Roles().UpdateByName(createdRole.Name, roleRequest)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "new comment", result.Comment)
}

func Test_RolesDeleteById(t *testing.T) {

	skipEnterprise(t)

	client := NewClient(NewDefaultConfig())

	roleRequest := &RoleRequest{
		Name:    "role-" + uuid.NewV4().String(),
		Comment: "testing",
	}

	createdRole, err := client.Roles().Create(roleRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdRole)

	err = client.Roles().DeleteById(createdRole.Id)
	assert.Nil(t, err)

	result, err := client.Roles().GetById(createdRole.Id)
	assert.Nil(t, err)
	assert.Nil(t, result)
}
func Test_RolesDeleteByName(t *testing.T) {

	skipEnterprise(t)

	client := NewClient(NewDefaultConfig())

	roleRequest := &RoleRequest{
		Name:    "role-" + uuid.NewV4().String(),
		Comment: "testing",
	}

	createdRole, err := client.Roles().Create(roleRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdRole)

	err = client.Roles().DeleteByName(createdRole.Name)
	assert.Nil(t, err)

	result, err := client.Roles().GetByName(createdRole.Name)
	assert.Nil(t, err)
	assert.Nil(t, result)
}

func Test_AddEndpointPermissionGetByRoleId(t *testing.T) {

	skipEnterprise(t)

	roleRequest := &RoleRequest{
		Name: "role-" + uuid.NewV4().String(),
	}

	client := NewClient(NewDefaultConfig())
	createdRole, err := client.Roles().Create(roleRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdRole)

	endpointPermissionReq := &EndpointPermissionRequest{
		WorkspaceId: "*",
		Endpoint:    "*",
		Negative:    false,
		Actions:     "read,create,update,delete",
		Comment:     "a comment",
	}
	endpoint, err := client.Roles().AddEndpointPermissionByRole(createdRole.Id, endpointPermissionReq)
	assert.Nil(t, err)
	assert.NotNil(t, endpoint)
	assert.Equal(t, endpointPermissionReq.Endpoint, endpoint.Endpoint)
	assert.Equal(t, endpointPermissionReq.Negative, endpoint.Negative)
	assert.Equal(t, createdRole.Id, endpoint.Role.Id)

}
func Test_GetEndpointPermission(t *testing.T) {

	skipEnterprise(t)

	roleRequest := &RoleRequest{
		Name: "role-" + uuid.NewV4().String(),
	}

	client := NewClient(NewDefaultConfig())
	createdRole, err := client.Roles().Create(roleRequest)

	assert.Nil(t, err)

	endpointPermissionReq := &EndpointPermissionRequest{
		WorkspaceId: "*",
		Endpoint:    "*",
		Negative:    false,
		Actions:     "read,create,update,delete",
		Comment:     "a comment",
	}
	endpoint, err := client.Roles().AddEndpointPermissionByRole(createdRole.Id, endpointPermissionReq)

	result, err := client.Roles().GetEndpointPermission(createdRole.Id, endpoint.WorkspaceId, endpoint.Endpoint)

	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.Equal(t, result.WorkspaceId, endpoint.WorkspaceId)

}

func Test_ListEndpointPermissions(t *testing.T) {

	skipEnterprise(t)

	roleRequest := &RoleRequest{
		Name: "role-" + uuid.NewV4().String(),
	}

	client := NewClient(NewDefaultConfig())
	createdRole, err := client.Roles().Create(roleRequest)

	assert.Nil(t, err)

	endpointPermissionReq1 := &EndpointPermissionRequest{
		WorkspaceId: "*",
		Endpoint:    "/foo",
		Negative:    false,
		Actions:     "read,create,update,delete",
		Comment:     "a comment",
	}
	endpointPermissionReq2 := &EndpointPermissionRequest{
		WorkspaceId: "*",
		Endpoint:    "/bar",
		Negative:    false,
		Actions:     "read,create,update,delete",
		Comment:     "a comment",
	}
	endpointPermissionReq3 := &EndpointPermissionRequest{
		WorkspaceId: "*",
		Endpoint:    "/baz",
		Negative:    false,
		Actions:     "read,create,update,delete",
		Comment:     "a comment",
	}
	_, err = client.Roles().AddEndpointPermissionByRole(createdRole.Id, endpointPermissionReq1)
	_, err = client.Roles().AddEndpointPermissionByRole(createdRole.Id, endpointPermissionReq2)
	_, err = client.Roles().AddEndpointPermissionByRole(createdRole.Id, endpointPermissionReq3)

	result, err := client.Roles().ListEndpointPermissions(createdRole.Id)

	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.True(t, len(result.Data) > 0)
}

func Test_UpdateEndpointPermissions(t *testing.T) {

	skipEnterprise(t)

	roleRequest := &RoleRequest{
		Name: "role-" + uuid.NewV4().String(),
	}

	client := NewClient(NewDefaultConfig())
	createdRole, err := client.Roles().Create(roleRequest)

	assert.Nil(t, err)

	endpointPermissionReq := &EndpointPermissionRequest{
		WorkspaceId: "*",
		Endpoint:    "/foo",
		Negative:    false,
		Actions:     "read,create,update,delete",
		Comment:     "a comment",
	}
	_, err = client.Roles().AddEndpointPermissionByRole(createdRole.Id, endpointPermissionReq)

	updatePermissionReq := &EndpointPermissionRequest{
		WorkspaceId: "*",
		Endpoint:    "/foo",
		Negative:    false,
		Actions:     "read,create,update",
	}
	result, err := client.Roles().UpdateEndpointPermissions(
		createdRole.Id,
		endpointPermissionReq.WorkspaceId,
		endpointPermissionReq.Endpoint,
		updatePermissionReq,
	)

	assert.NotNil(t, result)
	assert.Nil(t, err)
	// Confirm we've removed the delete action from the permission
	assert.Equal(t, []string{"create", "update", "read"}, result.Actions)
}

func Test_DeleteEndpointPermissions(t *testing.T) {

	skipEnterprise(t)

	roleRequest := &RoleRequest{
		Name: "role-" + uuid.NewV4().String(),
	}

	client := NewClient(NewDefaultConfig())
	createdRole, err := client.Roles().Create(roleRequest)

	assert.Nil(t, err)

	endpointPermissionReq := &EndpointPermissionRequest{
		WorkspaceId: "*",
		Endpoint:    "/foo",
		Negative:    false,
		Actions:     "read,create,update,delete",
		Comment:     "a comment",
	}
	endpoint, err := client.Roles().AddEndpointPermissionByRole(createdRole.Id, endpointPermissionReq)

	err = client.Roles().DeleteRoleEndpointPermission(createdRole.Id, endpoint.WorkspaceId, endpoint.Endpoint)

	assert.Nil(t, err)
}

func Test_AddEntityPermissionGetByRoleId(t *testing.T) {

	skipEnterprise(t)

	client := NewClient(NewDefaultConfig())

	// Create a service (an entity)
	serviceRequest := &ServiceRequest{
		Name:     String(fmt.Sprintf("service-name-%s", uuid.NewV4().String())),
		Protocol: String("http"),
		Host:     String("foo.com"),
	}

	createdService, err := client.Services().Create(serviceRequest)

	assert.Nil(t, err)

	roleRequest := &RoleRequest{
		Name: "role-entity-" + uuid.NewV4().String(),
	}

	createdRole, err := client.Roles().Create(roleRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdRole)

	entityPermissionReq := &EntityPermissionRequest{
		EntityId: *createdService.Id,
		Negative: false,
		Actions:  "read,create,update,delete",
		Comment:  "a comment",
	}

	entity, err := client.Roles().AddEntityPermissionByRole(createdRole.Id, entityPermissionReq)
	assert.Nil(t, err)
	assert.NotNil(t, entity)
	assert.Equal(t, createdService.Id, &entity.EntityId)
	assert.Equal(t, entityPermissionReq.Negative, entity.Negative)
	assert.Equal(t, createdRole.Id, entity.Role.Id)

	err = client.Services().DeleteServiceById(*createdService.Id)
	assert.Nil(t, err)

}

func Test_GetEntityPermission(t *testing.T) {

	skipEnterprise(t)

	client := NewClient(NewDefaultConfig())

	// Create a service (an entity)
	serviceRequest := &ServiceRequest{
		Name:     String(fmt.Sprintf("service-name-%s", uuid.NewV4().String())),
		Protocol: String("http"),
		Host:     String("foo.com"),
	}

	createdService, err := client.Services().Create(serviceRequest)

	assert.Nil(t, err)

	roleRequest := &RoleRequest{
		Name: "role-entity-" + uuid.NewV4().String(),
	}

	createdRole, err := client.Roles().Create(roleRequest)

	assert.Nil(t, err)

	entityPermissionReq := &EntityPermissionRequest{
		EntityId: *createdService.Id,
		Negative: false,
		Actions:  "read,create,update,delete",
		Comment:  "a comment",
	}
	entity, err := client.Roles().AddEntityPermissionByRole(createdRole.Id, entityPermissionReq)

	result, err := client.Roles().GetEntityPermission(createdRole.Id, entity.EntityId)

	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.Equal(t, result.EntityId, entity.EntityId)

	err = client.Services().DeleteServiceById(*createdService.Id)
	assert.Nil(t, err)
}

func Test_ListEntityPermissions(t *testing.T) {

	skipEnterprise(t)

	client := NewClient(NewDefaultConfig())

	// Create a service (an entity)
	serviceRequest1 := &ServiceRequest{
		Name:     String(fmt.Sprintf("service-name-%s", uuid.NewV4().String())),
		Protocol: String("http"),
		Host:     String("foo.com"),
	}
	serviceRequest2 := &ServiceRequest{
		Name:     String(fmt.Sprintf("service-name-%s", uuid.NewV4().String())),
		Protocol: String("http"),
		Host:     String("foo.com"),
	}

	createdService1, err := client.Services().Create(serviceRequest1)
	assert.Nil(t, err)
	createdService2, err := client.Services().Create(serviceRequest2)
	assert.Nil(t, err)

	roleRequest := &RoleRequest{
		Name: "role-entity-" + uuid.NewV4().String(),
	}

	createdRole, err := client.Roles().Create(roleRequest)

	assert.Nil(t, err)

	entityPermissionReq1 := &EntityPermissionRequest{
		EntityId: *createdService1.Id,
		Negative: false,
		Actions:  "read,create,update,delete",
		Comment:  "a comment",
	}
	entityPermissionReq2 := &EntityPermissionRequest{
		EntityId: *createdService2.Id,
		Negative: false,
		Actions:  "read,create,update,delete",
		Comment:  "a comment",
	}
	_, err = client.Roles().AddEntityPermissionByRole(createdRole.Id, entityPermissionReq1)
	_, err = client.Roles().AddEntityPermissionByRole(createdRole.Id, entityPermissionReq2)

	result, err := client.Roles().ListEntityPermissions(createdRole.Id)

	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.True(t, len(result.Data) > 0)

	err = client.Services().DeleteServiceById(*createdService1.Id)
	assert.Nil(t, err)
	err = client.Services().DeleteServiceById(*createdService2.Id)
	assert.Nil(t, err)
}

func Test_UpdateEntityPermissions(t *testing.T) {

	skipEnterprise(t)

	client := NewClient(NewDefaultConfig())

	// Create a service (an entity)
	serviceRequest := &ServiceRequest{
		Name:     String(fmt.Sprintf("service-name-%s", uuid.NewV4().String())),
		Protocol: String("http"),
		Host:     String("foo.com"),
	}

	createdService, err := client.Services().Create(serviceRequest)

	assert.Nil(t, err)

	roleRequest := &RoleRequest{
		Name: "role-entity-update-" + uuid.NewV4().String(),
	}

	createdRole, err := client.Roles().Create(roleRequest)

	assert.Nil(t, err)

	entityPermissionReq := &EntityPermissionRequest{
		EntityId: *createdService.Id,
		Negative: false,
		Actions:  "read,create,update,delete",
		Comment:  "a comment",
	}
	entity, err := client.Roles().AddEntityPermissionByRole(createdRole.Id, entityPermissionReq)

	assert.NotNil(t, entity)

	updatePermissionReq := &EntityPermissionRequest{
		EntityId: *createdService.Id,
		Negative: false,
		Actions:  "read,create,update",
	}
	result, err := client.Roles().UpdateEntityPermissions(
		createdRole.Id,
		*createdService.Id,
		updatePermissionReq,
	)

	assert.NotNil(t, result)
	assert.Nil(t, err)
	// Confirm we've removed the delete action from the permission
	assert.Equal(t, []string{"create", "update", "read"}, result.Actions)

	err = client.Services().DeleteServiceById(*createdService.Id)
	assert.Nil(t, err)
}

func Test_DeleteEntityPermissions(t *testing.T) {

	skipEnterprise(t)

	client := NewClient(NewDefaultConfig())

	// Create a service (an entity)
	serviceRequest := &ServiceRequest{
		Name:     String(fmt.Sprintf("service-name-%s", uuid.NewV4().String())),
		Protocol: String("http"),
		Host:     String("foo.com"),
	}

	createdService, err := client.Services().Create(serviceRequest)

	assert.Nil(t, err)

	roleRequest := &RoleRequest{
		Name: "role-entity-deleted-" + uuid.NewV4().String(),
	}

	createdRole, err := client.Roles().Create(roleRequest)

	assert.Nil(t, err)

	entityPermissionReq := &EntityPermissionRequest{
		EntityId: *createdService.Id,
		Negative: false,
		Actions:  "read,create,update,delete",
		Comment:  "a comment",
	}

	entity, err := client.Roles().AddEntityPermissionByRole(createdRole.Id, entityPermissionReq)

	err = client.Roles().DeleteRoleEntityPermission(createdRole.Id, entity.EntityId)
	assert.Nil(t, err)

	err = client.Services().DeleteServiceById(*createdService.Id)
	assert.Nil(t, err)
}
