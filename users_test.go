package gokong

import (
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func Test_UserGetById(t *testing.T) {

	skipEnterprise(t)

	userRequest := &UserRequest{
		Name:      "user-" + uuid.NewV4().String(),
		UserToken: "testToken" + uuid.NewV4().String(),
	}

	client := NewClient(NewDefaultConfig())
	createdUser, err := client.Users().Create(userRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdUser)

	result, err := client.Users().GetById(createdUser.Id)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, createdUser, result)

}
func Test_UserGetByName(t *testing.T) {

	skipEnterprise(t)

	userRequest := &UserRequest{
		Name:      "user-" + uuid.NewV4().String(),
		UserToken: "testToken" + uuid.NewV4().String(),
	}

	client := NewClient(NewDefaultConfig())
	createdUser, err := client.Users().Create(userRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdUser)

	result, err := client.Users().GetByName(createdUser.Name)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, createdUser, result)

}

func Test_UserGetByIdForNonExistentUserId(t *testing.T) {

	skipEnterprise(t)

	result, err := NewClient(NewDefaultConfig()).Users().GetById(uuid.NewV4().String())

	assert.Nil(t, err)
	assert.Nil(t, result)
}

func Test_UserGetByIdForNonExistentUserByName(t *testing.T) {

	skipEnterprise(t)

	result, err := NewClient(NewDefaultConfig()).Users().GetByName(uuid.NewV4().String())

	assert.Nil(t, err)
	assert.Nil(t, result)

}

func Test_UserCreate(t *testing.T) {

	skipEnterprise(t)

	userRequest := &UserRequest{
		Name:      "user-" + uuid.NewV4().String(),
		UserToken: "test-token" + uuid.NewV4().String(),
		Enabled:   true,
		Comment:   "testing",
	}

	client := NewClient(NewDefaultConfig())
	createdUser, err := client.Users().Create(userRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdUser)
	assert.Equal(t, userRequest.Name, createdUser.Name)
	assert.Equal(t, userRequest.Comment, createdUser.Comment)
}

func Test_UserList(t *testing.T) {

	skipEnterprise(t)

	userRequest1 := &UserRequest{
		Name:      "user-" + uuid.NewV4().String(),
		UserToken: "testToken" + uuid.NewV4().String(),
	}
	userRequest2 := &UserRequest{
		Name:      "user-" + uuid.NewV4().String(),
		UserToken: "testToken2" + uuid.NewV4().String(),
	}
	userRequest3 := &UserRequest{
		Name:      "user-" + uuid.NewV4().String(),
		UserToken: "testToken3" + uuid.NewV4().String(),
	}

	client := NewClient(NewDefaultConfig())
	createdUser1, err := client.Users().Create(userRequest1)
	createdUser2, err := client.Users().Create(userRequest2)
	createdUser3, err := client.Users().Create(userRequest3)

	assert.Nil(t, err)
	assert.NotNil(t, createdUser1)
	assert.NotNil(t, createdUser2)
	assert.NotNil(t, createdUser3)

	users, err := client.Users().List()
	assert.Nil(t, err)
	assert.NotNil(t, users)
	assert.True(t, len(users.Data) > 0)
}

func Test_UsersUpdateById(t *testing.T) {

	skipEnterprise(t)

	client := NewClient(NewDefaultConfig())

	userRequest := &UserRequest{
		Name:      "user-" + uuid.NewV4().String(),
		Comment:   "testing",
		UserToken: uuid.NewV4().String(),
	}

	createdUser, err := client.Users().Create(userRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdUser)
	assert.Equal(t, "testing", createdUser.Comment)

	userRequest.Comment = "new comment"

	result, err := client.Users().UpdateById(createdUser.Id, userRequest)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "new comment", result.Comment)
}

func Test_UsersUpdateByName(t *testing.T) {

	skipEnterprise(t)

	client := NewClient(NewDefaultConfig())

	userRequest := &UserRequest{
		Name:      "user-" + uuid.NewV4().String(),
		Comment:   "testing",
		UserToken: uuid.NewV4().String(),
	}

	createdUser, err := client.Users().Create(userRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdUser)
	assert.Equal(t, "testing", createdUser.Comment)

	userRequest.Comment = "new comment"

	result, err := client.Users().UpdateByName(createdUser.Name, userRequest)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "new comment", result.Comment)
}

func Test_UsersDeleteById(t *testing.T) {

	skipEnterprise(t)

	client := NewClient(NewDefaultConfig())

	userRequest := &UserRequest{
		Name:      "user-" + uuid.NewV4().String(),
		Comment:   "testing",
		UserToken: uuid.NewV4().String(),
	}

	createdUser, err := client.Users().Create(userRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdUser)

	err = client.Users().DeleteById(createdUser.Id)
	assert.Nil(t, err)

	result, err := client.Users().GetById(createdUser.Id)
	assert.Nil(t, err)
	assert.Nil(t, result)
}
func Test_UsersDeleteByName(t *testing.T) {

	skipEnterprise(t)

	client := NewClient(NewDefaultConfig())

	userRequest := &UserRequest{
		Name:      "user-" + uuid.NewV4().String(),
		Comment:   "testing",
		UserToken: uuid.NewV4().String(),
	}

	createdUser, err := client.Users().Create(userRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdUser)

	err = client.Users().DeleteByName(createdUser.Name)
	assert.Nil(t, err)

	result, err := client.Users().GetByName(createdUser.Name)
	assert.Nil(t, err)
	assert.Nil(t, result)
}

func Test_AddRoleToUser(t *testing.T) {

	skipEnterprise(t)

	client := NewClient(NewDefaultConfig())

	userRequest := &UserRequest{
		Name:      "user-" + uuid.NewV4().String(),
		Comment:   "testing",
		UserToken: uuid.NewV4().String(),
	}

	createdUser, err := client.Users().Create(userRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdUser)

	roleRequest := &RoleRequest{
		Name: "role-useradd-" + uuid.NewV4().String(),
	}

	createdRole, err := client.Roles().Create(roleRequest)

	assert.Nil(t, err)

	userRoleRequest := &UserRoleRequest{
		Roles: createdRole.Name,
	}

	userRoles, err := client.Users().AddUserToRole(createdUser.Id, userRoleRequest)

	assert.Nil(t, err)
	assert.True(t, len(*userRoles.Roles) > 0)

}

func Test_ListUserRoles(t *testing.T) {

	skipEnterprise(t)

	client := NewClient(NewDefaultConfig())

	userRequest := &UserRequest{
		Name:      "user-" + uuid.NewV4().String(),
		Comment:   "testing",
		UserToken: uuid.NewV4().String(),
	}

	createdUser, err := client.Users().Create(userRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdUser)

	roleRequest := &RoleRequest{
		Name: "role-useradd-" + uuid.NewV4().String(),
	}

	createdRole, err := client.Roles().Create(roleRequest)

	assert.Nil(t, err)

	userRoleRequest := &UserRoleRequest{
		Roles: createdRole.Name,
	}

	_, err = client.Users().AddUserToRole(createdUser.Id, userRoleRequest)

	assert.Nil(t, err)

	userRoles, err := client.Users().ListUserRoles(createdUser.Id)

	assert.Nil(t, err)
	assert.True(t, len(*userRoles.Roles) > 0)
}

func Test_DeleteRoleFromUser(t *testing.T) {

	skipEnterprise(t)

	client := NewClient(NewDefaultConfig())

	userRequest := &UserRequest{
		Name:      "user-" + uuid.NewV4().String(),
		Comment:   "testing",
		UserToken: uuid.NewV4().String(),
	}

	createdUser, err := client.Users().Create(userRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdUser)

	roleRequest := &RoleRequest{
		Name: "role-useradd-" + uuid.NewV4().String(),
	}

	createdRole, err := client.Roles().Create(roleRequest)

	assert.Nil(t, err)

	userRoleRequest := &UserRoleRequest{
		Roles: createdRole.Name,
	}

	_, err = client.Users().AddUserToRole(createdUser.Id, userRoleRequest)

	assert.Nil(t, err)

	deleteRoleRequest := &UserRoleRequest{
		Roles: createdRole.Name,
	}

	err = client.Users().DeleteRoleFromUser(createdRole.Id, deleteRoleRequest)

	assert.Nil(t, err)
}

func Test_ListUserPermissions(t *testing.T) {

	skipEnterprise(t)

	client := NewClient(NewDefaultConfig())

	userRequest := &UserRequest{
		Name:      "user-" + uuid.NewV4().String(),
		Comment:   "testing",
		UserToken: uuid.NewV4().String(),
	}

	createdUser, err := client.Users().Create(userRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdUser)

	roleRequest := &RoleRequest{
		Name: "role-useradd-" + uuid.NewV4().String(),
	}

	createdRole, err := client.Roles().Create(roleRequest)

	assert.Nil(t, err)

	userRoleRequest := &UserRoleRequest{
		Roles: createdRole.Name,
	}

	endpointPermissionReq := &EndpointPermissionRequest{
		WorkspaceId: "*",
		Endpoint:    "/foo",
		Negative:    false,
		Actions:     "read,create,update,delete",
		Comment:     "a comment",
	}
	_, err = client.Roles().AddEndpointPermissionByRole(createdRole.Id, endpointPermissionReq)

	assert.Nil(t, err)

	_, err = client.Users().AddUserToRole(createdUser.Id, userRoleRequest)

	assert.Nil(t, err)

	listUserPermissions, err := client.Users().ListUserPermissions(createdUser.Id)

	assert.Nil(t, err)
	assert.NotNil(t, listUserPermissions.Endpoints)
}
