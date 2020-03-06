package gokong

import (
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func Test_AdminInvite(t *testing.T) {

	skipEnterprise(t)
	rnd := uuid.NewV4().String()
	inviteAdminRequest := &InviteAdminRequest{
		Email:            "admin-" + rnd + "@example.com",
		Username:         "admin-" + rnd + "@example.com",
		CustomId:         rnd,
		RBACTokenEnabled: Bool(true),
	}

	client := NewClient(NewDefaultConfig())
	createdAdmin, err := client.Admins().Invite(inviteAdminRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdAdmin)
	assert.Equal(t, createdAdmin.Admin.Status, 4)
}
func Test_AdminGetById(t *testing.T) {

	skipEnterprise(t)

	rnd := uuid.NewV4().String()
	inviteAdminRequest := &InviteAdminRequest{
		Email:            "admin-" + rnd + "@example.com",
		Username:         "admin-" + rnd + "@example.com",
		CustomId:         rnd,
		RBACTokenEnabled: Bool(true),
	}

	client := NewClient(NewDefaultConfig())
	createdAdmin, err := client.Admins().Invite(inviteAdminRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdAdmin)

	result, err := client.Admins().Get(createdAdmin.Admin.Id)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, createdAdmin.Admin.Email, result.Email)

}

func Test_AdminGetByIdForNonExistentAdminId(t *testing.T) {

	skipEnterprise(t)

	result, err := NewClient(NewDefaultConfig()).Admins().Get(uuid.NewV4().String())

	assert.NotNil(t, err)
	assert.Nil(t, result)
}

func Test_AdminList(t *testing.T) {

	skipEnterprise(t)

	rnd1 := uuid.NewV4().String()
	inviteAdminRequest1 := &InviteAdminRequest{
		Email:            "admin-" + rnd1 + "@example.com",
		Username:         "admin-" + rnd1 + "@example.com",
		CustomId:         rnd1,
		RBACTokenEnabled: Bool(true),
	}
	rnd2 := uuid.NewV4().String()
	inviteAdminRequest2 := &InviteAdminRequest{
		Email:            "admin-" + rnd2 + "@example.com",
		Username:         "admin-" + rnd2 + "@example.com",
		CustomId:         rnd2,
		RBACTokenEnabled: Bool(true),
	}
	rnd3 := uuid.NewV4().String()
	inviteAdminRequest3 := &InviteAdminRequest{
		Email:            "admin-" + rnd3 + "@example.com",
		Username:         "admin-" + rnd3 + "@example.com",
		CustomId:         rnd3,
		RBACTokenEnabled: Bool(true),
	}

	client := NewClient(NewDefaultConfig())
	createdAdmin1, err := client.Admins().Invite(inviteAdminRequest1)
	createdAdmin2, err := client.Admins().Invite(inviteAdminRequest2)
	createdAdmin3, err := client.Admins().Invite(inviteAdminRequest3)

	assert.Nil(t, err)
	assert.NotNil(t, createdAdmin1)
	assert.NotNil(t, createdAdmin2)
	assert.NotNil(t, createdAdmin3)

	admins, err := client.Admins().List()
	assert.Nil(t, err)
	assert.NotNil(t, admins)
	assert.True(t, len(admins.Data) > 0)
}

func Test_AdminsRegisterCredentials(t *testing.T) {

	skipEnterprise(t)

	rnd := uuid.NewV4().String()
	inviteAdminRequest := &InviteAdminRequest{
		Email:            "admin-" + rnd + "@example.com",
		Username:         "admin-" + rnd + "@example.com",
		CustomId:         rnd,
		RBACTokenEnabled: Bool(true),
	}

	client := NewClient(NewDefaultConfig())
	createdAdmin, err := client.Admins().Invite(inviteAdminRequest)

	registerCredsRequest := &RegisterAdminCredentialsRequest{
		Email:    createdAdmin.Admin.Email,
		Username: createdAdmin.Admin.Username,
		Token:    "my-token",
		Password: "p@ssw0rd",
	}

	err = client.Admins().RegisterAdminCredentials(registerCredsRequest)

	assert.Nil(t, err)
}

// Unable to test this without auth configured and migrations run. Endpoint returns 404.
// func Test_AdminsSendAdminPasswordResetEmail(t *testing.T) {

// 	skipEnterprise(t)

// 	rnd := uuid.NewV4().String()
// 	inviteAdminRequest := &InviteAdminRequest{
// 		Email:            "admin-" + rnd + "@example.com",
// 		Username:         "admin-" + rnd + "@example.com",
// 		CustomId:         rnd,
// 		RBACTokenEnabled: Bool(true),
// 	}

// 	client := NewClient(NewDefaultConfig())
// 	createdAdmin, err := client.Admins().Invite(inviteAdminRequest)

// 	sendResetRequest := &SendPasswordResetRequest{
// 		Email: createdAdmin.Admin.Email,
// 	}

// 	err = client.Admins().SendAdminPasswordResetEmail(sendResetRequest)

// 	assert.Nil(t, err)
// }

// Unable to test this without auth configured and migrations run. Endpoint returns 404.
// func Test_AdminsResetAdminPassword(t *testing.T) {

// 	skipEnterprise(t)

// 	rnd := uuid.NewV4().String()
// 	inviteAdminRequest := &InviteAdminRequest{
// 		Email:            "admin-" + rnd + "@example.com",
// 		Username:         "admin-" + rnd + "@example.com",
// 		CustomId:         rnd,
// 		RBACTokenEnabled: Bool(true),
// 	}

// 	client := NewClient(NewDefaultConfig())
// 	createdAdmin, err := client.Admins().Invite(inviteAdminRequest)

// 	resetRequest := &PasswordResetRequest{
// 		Email:    createdAdmin.Admin.Email,
// 		Password: "new-pass",
// 	}

// 	err = client.Admins().ResetAdminPassword(resetRequest)

// 	assert.Nil(t, err)
// }

func Test_AdminsUpdate(t *testing.T) {

	skipEnterprise(t)

	rnd := uuid.NewV4().String()
	inviteAdminRequest := &InviteAdminRequest{
		Email:            "admin-" + rnd + "@example.com",
		Username:         "admin-" + rnd + "@example.com",
		CustomId:         rnd,
		RBACTokenEnabled: Bool(true),
	}

	client := NewClient(NewDefaultConfig())

	createdAdmin, err := client.Admins().Invite(inviteAdminRequest)

	assert.Nil(t, err)

	updateAdminRequest := &AdminRequest{
		Email: "test@example.com",
	}

	updatedAdmin, err := client.Admins().Update(createdAdmin.Admin.Id, updateAdminRequest)

	assert.Nil(t, err)
	assert.Equal(t, createdAdmin.Admin.Username, updatedAdmin.Username)
	assert.Equal(t, "test@example.com", updatedAdmin.Email)
}

func Test_AdminsDelete(t *testing.T) {

	skipEnterprise(t)

	rnd := uuid.NewV4().String()
	inviteAdminRequest := &InviteAdminRequest{
		Email:            "admin-" + rnd + "@example.com",
		Username:         "admin-" + rnd + "@example.com",
		CustomId:         rnd,
		RBACTokenEnabled: Bool(true),
	}

	client := NewClient(NewDefaultConfig())

	createdAdmin, err := client.Admins().Invite(inviteAdminRequest)

	assert.Nil(t, err)

	err = client.Admins().Delete(createdAdmin.Admin.Id)

	assert.Nil(t, err)

	result, err := client.Admins().Get(createdAdmin.Admin.Id)

	assert.Nil(t, result)
	assert.NotNil(t, err)

}

func Test_AddRoleToAdmin(t *testing.T) {

	skipEnterprise(t)

	client := NewClient(NewDefaultConfig())

	rnd := uuid.NewV4().String()
	inviteAdminRequest := &InviteAdminRequest{
		Email:            "admin-" + rnd + "@example.com",
		Username:         "admin-" + rnd + "@example.com",
		CustomId:         rnd,
		RBACTokenEnabled: Bool(true),
	}
	createdAdmin, err := client.Admins().Invite(inviteAdminRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdAdmin)

	roleRequest := &RoleRequest{
		Name: "role-adminadd-" + uuid.NewV4().String(),
	}

	createdRole, err := client.Roles().Create(roleRequest)

	assert.Nil(t, err)

	adminRoleRequest := &AdminRoleRequest{
		Roles: createdRole.Name,
	}

	adminRoles, err := client.Admins().AddOrUpdateRoles(createdAdmin.Admin.Id, adminRoleRequest)

	assert.Nil(t, err)
	assert.True(t, len(adminRoles.Roles) > 0)

}

func Test_ListAdminRoles(t *testing.T) {

	skipEnterprise(t)

	client := NewClient(NewDefaultConfig())

	rnd := uuid.NewV4().String()
	inviteAdminRequest := &InviteAdminRequest{
		Email:            "admin-" + rnd + "@example.com",
		Username:         "admin-" + rnd + "@example.com",
		CustomId:         rnd,
		RBACTokenEnabled: Bool(true),
	}
	createdAdmin, err := client.Admins().Invite(inviteAdminRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdAdmin)

	roleRequest := &RoleRequest{
		Name: "role-adminadd-" + uuid.NewV4().String(),
	}

	createdRole, err := client.Roles().Create(roleRequest)

	assert.Nil(t, err)

	adminRoleRequest := &AdminRoleRequest{
		Roles: createdRole.Name,
	}

	adminRoles, err := client.Admins().AddOrUpdateRoles(createdAdmin.Admin.Id, adminRoleRequest)

	adminRolesList, err := client.Admins().ListRoles(createdAdmin.Admin.Id)

	assert.Nil(t, err)
	assert.True(t, len(adminRolesList.Roles) > 0)
	assert.Equal(t, adminRoles.Roles, adminRolesList.Roles)
}

func Test_DeleteRoleFromAdmin(t *testing.T) {

	skipEnterprise(t)

	client := NewClient(NewDefaultConfig())

	rnd := uuid.NewV4().String()
	inviteAdminRequest := &InviteAdminRequest{
		Email:            "admin-" + rnd + "@example.com",
		Username:         "admin-" + rnd + "@example.com",
		CustomId:         rnd,
		RBACTokenEnabled: Bool(true),
	}
	createdAdmin, err := client.Admins().Invite(inviteAdminRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdAdmin)

	roleRequest := &RoleRequest{
		Name: "role-adminadd-" + uuid.NewV4().String(),
	}

	createdRole, err := client.Roles().Create(roleRequest)

	assert.Nil(t, err)

	adminRoleRequest := &AdminRoleRequest{
		Roles: createdRole.Name,
	}

	_, err = client.Admins().AddOrUpdateRoles(createdAdmin.Admin.Id, adminRoleRequest)

	assert.Nil(t, err)

	err = client.Admins().DeleteRoles(createdAdmin.Admin.Id, adminRoleRequest)

	assert.Nil(t, err)
}

func Test_ListAdminWorkspaces(t *testing.T) {

	skipEnterprise(t)

	client := NewClient(NewDefaultConfig())
	rnd := uuid.NewV4().String()
	workspaceRequest := &WorkspaceRequest{
		Name: "test-workspace" + rnd,
	}

	result, err := client.Workspaces().Create(workspaceRequest)

	assert.Nil(t, err)
	assert.NotNil(t, result)

	client = NewClient(NewWorkspaceConfig("test-workspace" + rnd))

	inviteAdminRequest := &InviteAdminRequest{
		Email:            "workspace-admin-" + rnd + "@example.com",
		Username:         "workspace-admin-" + rnd + "@example.com",
		CustomId:         rnd,
		RBACTokenEnabled: Bool(true),
	}
	createdAdmin, err := client.Admins().Invite(inviteAdminRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdAdmin)

	adminWorkspaces, err := client.Admins().ListWorkspaces(createdAdmin.Admin.Id)
	assert.Nil(t, err)
	assert.True(t, len(adminWorkspaces) > 0)
}
