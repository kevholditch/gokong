package gokong

import (
	"encoding/json"
	"fmt"
)

// AdminClient manages kong client constructs in the context of Admin endpoint interactions
type AdminClient struct {
	config *Config
}

// InviteAdminRequest represents an RBAC admin request in Kong. This feature is only available in Enterprise.
type InviteAdminRequest struct {
	Email            string `json:"email" yaml:"email"`
	Username         string `json:"username" yaml:"username"`
	CustomId         string `json:"custom_id,omitempty" yaml:"custom_id,omitempty"`
	RBACTokenEnabled bool   `json:"rbac_token_enabled" yaml:"rbac_token_enabled"`
}

type InviteAdminResponse struct {
	Admin *AdminResponse `json:"admin" yaml:"admin"`
}

type RegisterAdminCredentialsRequest struct {
	Email    string `json:"email" yaml:"email"`
	Username string `json:"username" yaml:"username"`
	Token    string `json:"token" yaml:"token"`
	Password string `json"password" yaml:"password"`
}

type SendPasswordResetRequest struct {
	Email string `json:"email" yaml:"email"`
}

type PasswordResetRequest struct {
	Email    string `json:"email" yaml:"email"`
	Password string `json:"password" yaml:"password"`
	Token    string `json:"token" yaml:"token"`
}

type AdminRequest struct {
	Name             string `json:"name,omitempty" yaml:"name,omitempty"`
	Id               string `json:"id,omitempty" yaml:"id,omitempty"`
	Email            string `json:"email,omitempty" yaml:"email,omitempty"`
	Username         string `json:"username,omitempty" yaml:"username,omitempty"`
	CustomId         string `json:"custom_id,omitempty" yaml:"custom_id,omitempty"`
	RBACTokenEnabled bool   `json"rbac_token_enabled" yaml:"rbac_token_enabled"`
}
type AdminResponse struct {
	CreatedAt        *int   `json:"created_at" yaml:"created_at"`
	UpdatedAt        *int   `json:"updated_at" yaml:"updated_at"`
	Status           int    `json:"status" yaml:"status"`
	Email            string `json:"email" yaml:"email"`
	Username         string `json:"username" yaml:"username"`
	Id               string `json:"id" yaml:"id"`
	RBACTokenEnabled bool   `json"rbac_token_enabled" yaml:"rbac_token_enabled"`
}

type Admins struct {
	Data []*AdminResponse `json:"data" yaml:"data"`
	Next string           `json:"next,omitempty" yaml:"next,omitempty"`
}

type AdminRole struct {
	Comment   string `json:"comment" yaml:"comment"`
	CreatedAt *int   `json:"created_at yaml:"created_at"`
	Id        string `json:"id" yaml:"id"`
	Name      string `json:"name" yaml:"name"`
	IsDefault bool   `json:"is_default" yaml:"is_default"`
}

type AdminRoles struct {
	Roles []*AdminRole `json:"roles" yaml:"roles"`
}

type AdminRoleRequest struct {
	Name  string `json:"name,omitempty" yaml:"name,omitempty"`
	Id    string `json:"id,omitempty" yaml:"name,omitempty"`
	Roles string `json:"roles" yaml:"roles"`
}

const AdminsPath = "/admins/"

func (adminClient *AdminClient) List() (*Admins, error) {

	r, body, errs := newGet(adminClient.config, adminClient.config.HostAddress+AdminsPath).End()
	if errs != nil {
		return nil, fmt.Errorf("could not get admins, error: %v", errs)
	}

	if r.StatusCode == 401 || r.StatusCode == 403 {
		return nil, fmt.Errorf("not authorised, message from kong: %s", body)
	}

	admins := &Admins{}
	err := json.Unmarshal([]byte(body), admins)
	if err != nil {
		return nil, fmt.Errorf("could not parse admin list response, error: %v", err)
	}

	return admins, nil
}

func (adminClient *AdminClient) Invite(inviteAdminRequest *InviteAdminRequest) (*InviteAdminResponse, error) {
	r, body, errs := newPost(adminClient.config, adminClient.config.HostAddress+AdminsPath).Send(inviteAdminRequest).End()
	if errs != nil {
		return nil, fmt.Errorf("could not invite new admin, error: %v", errs)
	}

	if r.StatusCode == 401 || r.StatusCode == 403 {
		return nil, fmt.Errorf("not authorised, message from kong: %s", body)
	}

	createdAdmin := &InviteAdminResponse{}
	err := json.Unmarshal([]byte(body), createdAdmin)
	if err != nil {
		return nil, fmt.Errorf("could not parse admin invite response, error: %v", err)
	}

	if createdAdmin.Admin.Id == "" {
		return nil, fmt.Errorf("could not invite admin, error: %v", body)
	}

	return createdAdmin, nil
}

func (adminClient *AdminClient) RegisterAdminCredentials(registerAdminCreds *RegisterAdminCredentialsRequest) error {
	r, body, errs := newPost(adminClient.config, adminClient.config.HostAddress+AdminsPath+"register").Send(registerAdminCreds).End()
	if errs != nil {
		return fmt.Errorf("could not register admin credentials, error: %v", errs)
	}

	if r.StatusCode == 401 || r.StatusCode == 403 {
		return fmt.Errorf("not authorised, message from kong: %s", body)
	}

	return nil
}

func (adminClient *AdminClient) SendAdminPasswordResetEmail(sendResetRequest *SendPasswordResetRequest) error {
	r, body, errs := newPost(adminClient.config, adminClient.config.HostAddress+AdminsPath+"password_resets").Send(sendResetRequest).End()
	if errs != nil {
		return fmt.Errorf("could not send password-reset email, error: %v", errs)
	}

	if r.StatusCode == 401 || r.StatusCode == 403 {
		return fmt.Errorf("not authorised, message from kong: %s", body)
	}

	return nil
}

func (adminClient *AdminClient) ResetAdminPassword(resetPasswordRequest *PasswordResetRequest) error {

	r, body, errs := newPatch(adminClient.config, adminClient.config.HostAddress+AdminsPath+"password_resets").Send(resetPasswordRequest).End()
	if errs != nil {
		return fmt.Errorf("could not reset admin password, error: %v", errs)
	}

	if r.StatusCode == 401 || r.StatusCode == 403 {
		return fmt.Errorf("not authorised, message from kong: %s", body)
	}

	return nil
}

// Get will retrieve details about an Admin when provided the name or id of an Admin entity
func (adminClient *AdminClient) Get(id string) (*AdminResponse, error) {
	r, body, errs := newGet(adminClient.config, adminClient.config.HostAddress+AdminsPath+id).End()
	if errs != nil {
		return nil, fmt.Errorf("could not get admin, error: %v", errs)
	}

	if r.StatusCode == 401 || r.StatusCode == 403 {
		return nil, fmt.Errorf("not authorised, message from kong: %s", body)
	}

	admin := &AdminResponse{}
	err := json.Unmarshal([]byte(body), admin)
	if err != nil {
		return nil, fmt.Errorf("could not parse admin get response, error: %v", err)
	}

	if admin.Id == "" {
		return nil, nil
	}

	return admin, nil
}

func (adminClient *AdminClient) Update(id string, adminRequest *AdminRequest) (*AdminResponse, error) {

	r, body, errs := newPatch(adminClient.config, adminClient.config.HostAddress+AdminsPath+id).Send(adminRequest).End()
	if errs != nil {
		return nil, fmt.Errorf("could not get admin, error: %v", errs)
	}

	if r.StatusCode == 401 || r.StatusCode == 403 {
		return nil, fmt.Errorf("not authorised, message from kong: %s", body)
	}

	admin := &AdminResponse{}
	err := json.Unmarshal([]byte(body), admin)
	if err != nil {
		return nil, fmt.Errorf("could not parse admin get response, error: %v", err)
	}

	if admin.Id == "" {
		return nil, nil
	}

	return admin, nil
}

func (adminClient *AdminClient) Delete(id string) error {

	r, body, errs := newDelete(adminClient.config, adminClient.config.HostAddress+AdminsPath+id).End()
	if errs != nil {
		return fmt.Errorf("could not delete admin, result: %v error: %v", r, errs)
	}

	if r.StatusCode == 401 || r.StatusCode == 403 {
		return fmt.Errorf("not authorised, message from kong: %s", body)
	}

	return nil
}

func (adminClient *AdminClient) ListRoles(id string) (*AdminRoles, error) {

	r, body, errs := newGet(adminClient.config, adminClient.config.HostAddress+AdminsPath+id+"/roles").End()
	if errs != nil {
		return nil, fmt.Errorf("could not list admin roles, error: %v", errs)
	}

	if r.StatusCode == 401 || r.StatusCode == 403 {
		return nil, fmt.Errorf("not authorised, message from kong: %v", body)
	}

	adminRoles := &AdminRoles{}

	err := json.Unmarshal([]byte(body), adminRoles)
	if err != nil {
		return nil, fmt.Errorf("could not parse the list admin roles response, error: %v", err)
	}

	return adminRoles, nil
}

func (adminClient *AdminClient) AddOrUpdateRoles(id string, adminRoleRequest *AdminRoleRequest) (*AdminRoles, error) {

	r, body, errs := newPost(adminClient.config, adminClient.config.HostAddress+AdminsPath+id+"/roles").Send(adminRoleRequest).End()
	if errs != nil {
		return nil, fmt.Errorf("could not add or update admin roles, error: %v", body)
	}

	if r.StatusCode == 401 || r.StatusCode == 403 {
		return nil, fmt.Errorf("not authorised, message from kong: %v", body)
	}

	adminRoles := &AdminRoles{}

	err := json.Unmarshal([]byte(body), adminRoles)
	if err != nil {
		return nil, fmt.Errorf("could not parse the add or update admin roles response, error: %v", err)
	}

	return adminRoles, nil
}

func (adminClient *AdminClient) DeleteRoles(id string, adminRoleRequest *AdminRoleRequest) error {

	r, body, errs := newDelete(adminClient.config, adminClient.config.HostAddress+AdminsPath+id+"/roles").Send(adminRoleRequest).End()
	if errs != nil {
		return fmt.Errorf("could not delete admin roles, error: %v", body)
	}

	if r.StatusCode == 401 || r.StatusCode == 403 {
		return fmt.Errorf("not authorised, message from kong: %v", body)
	}

	return nil
}

func (adminClient *AdminClient) ListWorkspaces(adminId string) ([]*WorkspaceRequest, error) {
	r, body, errs := newGet(adminClient.config, adminClient.config.HostAddress+AdminsPath+adminId+"/workspaces").End()
	if errs != nil {
		return nil, fmt.Errorf("could not list admin workspaces, error: %v", errs)
	}

	if r.StatusCode == 401 || r.StatusCode == 403 {
		return nil, fmt.Errorf("not authorised, message from kong: %v", body)
	}

	adminWorkspaces := []*WorkspaceRequest{}

	err := json.Unmarshal([]byte(body), &adminWorkspaces)
	if err != nil {
		return nil, fmt.Errorf("could not parse the list admin workspaces response, error: %v", err)
	}

	return adminWorkspaces, nil
}
