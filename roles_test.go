package gokong

import (
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

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
