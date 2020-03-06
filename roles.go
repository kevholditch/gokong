package gokong

import (
	"encoding/json"
	"fmt"
	"strings"
)

// RoleClient manages kong client constructs in the context of Role endpoint interactions
type RoleClient struct {
	config *Config
}

// RoleRequest represents an RBAC role in Kong. This feature is only available in Enterprise.
type RoleRequest struct {
	Name      string `json:"name" yaml:"name"`
	Comment   string `json:"comment,omitempty" yaml:"comment,omitempty"`
	CreatedAt *int   `json:"created_at,omitempty" yaml:"created_at,omitempty"`
}

// Role represents an RBAC role in Kong, it is generally used to marshal payloads coming back from Kong.
type Role struct {
	Id string `json:"id" yaml:"id"`
	RoleRequest
}

type Roles struct {
	Data []*Role `json:"data" yaml:"data"`
	Next string  `json:"next,omitempty" yaml:"next,omitempty"`
}

// EntityPermissionRequest provides a struct used to make requests to the admin API to update Entity Permissions
// EntityId must be the ID of an entity in Kong; if the ID of a workspace is given,
// the permission will apply to all entities in that workspace. Future entities belonging
// to that workspace will get the same permissions. A wildcard * will be interpreted
// as all entities in the system.
type EntityPermissionRequest struct {
	EntityId string `json:"entity_id,omitempty" yaml:"entity_id,omitempty"`
	Negative *bool  `json:"negative" yaml:"negative"`
	Actions  string `json:"actions" yaml:"actions"` // Comma separated string (read,create,update,delete)
	Comment  string `json:"comment,omitempty" yaml:"comment,omitempty"`
}

type EndpointPermissionRequest struct {
	WorkspaceId string `json:"workspace" yaml:"workspace"`
	Endpoint    string `json:"endpoint" yaml:"endpoint"` // Path of the associated endpoint. Can be exact matches, or contain wildcards represented by *
	Negative    *bool  `json:"negative" yaml:"negative"`
	Actions     string `json:"actions" yaml:"actions"` // Comma separated string (read,create,update,delete)
	Comment     string `json:"comment,omitempty" yaml:"comment,omitempty"`
}

type EntityPermission struct {
	EntityId   string                 `json:"entity_id" yaml:"entity_id"`
	EntityType string                 `json:"entity_type" yaml:"entity_type"`
	Actions    []string               `json:"actions" yaml:"actions"`
	CreatedAt  *int                   `json:"created_at,omitempty" yaml:"created_at,omitempty"`
	Negative   bool                   `json:"negative" yaml:"negative"`
	Role       EndpointPermissionRole `json:"role,omitempty" yaml:"role,omitempty"`
}

type EndpointPermission struct {
	WorkspaceId string                 `json:"workspace" yaml:"workspace"`
	Actions     []string               `json:"actions" yaml:"actions"`
	CreatedAt   *int                   `json:"created_at,omitempty" yaml:"created_at,omitempty"`
	Endpoint    string                 `json:"endpoint,omitempty" yaml:"endpoint,omitempty"`
	Negative    bool                   `json:"negative" yaml:"negative"`
	Role        EndpointPermissionRole `json:"role,omitempty" yaml:"role,omitempty"`
}

type EndpointPermissions struct {
	Data []*EndpointPermission `json:"data" yaml:"data"`
}
type EntityPermissions struct {
	Data []*EntityPermission `json:"data" yaml:"data"`
}

type EndpointPermissionRole struct {
	Id string `json:"id" yaml:"id"`
}

const RolesPath = "/rbac/roles/"

func (roleClient *RoleClient) getWorkspacePath() string {
	if roleClient.config.Workspace != "" {
		return "/" + roleClient.config.Workspace
	}
	return ""
}

// Role
func (roleClient *RoleClient) GetByName(name string) (*Role, error) {
	return roleClient.GetById(name)
}

func (roleClient *RoleClient) GetById(id string) (*Role, error) {

	r, body, errs := newGet(roleClient.config, roleClient.config.HostAddress+roleClient.getWorkspacePath()+RolesPath+id).End()
	if errs != nil {
		return nil, fmt.Errorf("could not get role, error: %v", errs)
	}

	err := checkResponse(r, body, errs)
	if err != nil {
		return nil, err
	}

	role := &Role{}
	err = json.Unmarshal([]byte(body), role)
	if err != nil {
		return nil, fmt.Errorf("could not parse role get response, error: %v", err)
	}

	if role.Id == "" {
		return nil, nil
	}

	return role, nil
}

func (roleClient *RoleClient) Create(roleRequest *RoleRequest) (*Role, error) {
	r, body, errs := newPost(roleClient.config, roleClient.config.HostAddress+roleClient.getWorkspacePath()+RolesPath).Send(roleRequest).End()
	if errs != nil {
		return nil, fmt.Errorf("could not create new role, error: %v", errs)
	}

	err := checkResponse(r, body, errs)
	if err != nil {
		return nil, err
	}

	createdRole := &Role{}
	err = json.Unmarshal([]byte(body), createdRole)
	if err != nil {
		return nil, fmt.Errorf("could not parse role creation response, error: %v", err)
	}

	if createdRole.Id == "" {
		return nil, fmt.Errorf("could not create update, error: %v", body)
	}

	return createdRole, nil
}

func (roleClient *RoleClient) DeleteByName(name string) error {
	return roleClient.DeleteById(name)
}

func (roleClient *RoleClient) DeleteById(id string) error {

	r, body, errs := newDelete(roleClient.config, roleClient.config.HostAddress+roleClient.getWorkspacePath()+RolesPath+id).End()
	if errs != nil {
		return fmt.Errorf("could not delete role, result: %v error: %v", r, errs)
	}

	err := checkResponse(r, body, errs)
	if err != nil {
		return err
	}

	return nil
}

func (roleClient *RoleClient) List() (*Roles, error) {

	r, body, errs := newGet(roleClient.config, roleClient.config.HostAddress+roleClient.getWorkspacePath()+RolesPath).End()
	if errs != nil {
		return nil, fmt.Errorf("could not get roles, error: %v", errs)
	}

	err := checkResponse(r, body, errs)
	if err != nil {
		return nil, err
	}

	roles := &Roles{}
	err = json.Unmarshal([]byte(body), roles)
	if err != nil {
		return nil, fmt.Errorf("could not parse role list response, error: %v", err)
	}

	return roles, nil
}

func (roleClient *RoleClient) UpdateByName(name string, roleRequest *RoleRequest) (*Role, error) {
	return roleClient.UpdateById(name, roleRequest)
}

func (roleClient *RoleClient) UpdateById(id string, roleRequest *RoleRequest) (*Role, error) {

	r, body, errs := newPatch(roleClient.config, roleClient.config.HostAddress+roleClient.getWorkspacePath()+RolesPath+id).Send(roleRequest).End()
	if errs != nil {
		return nil, fmt.Errorf("could not update role, error: %v", errs)
	}

	err := checkResponse(r, body, errs)
	if err != nil {
		return nil, err
	}

	updatedRole := &Role{}
	err = json.Unmarshal([]byte(body), updatedRole)
	if err != nil {
		return nil, fmt.Errorf("could not parse role update response, error: %v", err)
	}

	if updatedRole.Id == "" {
		return nil, fmt.Errorf("could not update role, error: %v", body)
	}

	return updatedRole, nil
}

// Role Endpoint Permission
func (roleClient *RoleClient) AddEndpointPermissionByRole(id string, roleEndpointPermissionRequest *EndpointPermissionRequest) (*EndpointPermission, error) {
	r, body, errs := newPost(roleClient.config, roleClient.config.HostAddress+roleClient.getWorkspacePath()+RolesPath+id+"/endpoints").Send(roleEndpointPermissionRequest).End()
	if errs != nil {
		return nil, fmt.Errorf("could not update role entities, error: %v", errs)
	}

	err := checkResponse(r, body, errs)
	if err != nil {
		return nil, err
	}

	roleEndpointPermission := &EndpointPermission{}

	err = json.Unmarshal([]byte(body), roleEndpointPermission)
	if err != nil {
		return nil, fmt.Errorf("could not parse the entity update response, error: %v", err)
	}

	return roleEndpointPermission, nil
}

func (roleClient *RoleClient) GetEndpointPermission(roleId string, workspaceId string, endpoint string) (*EndpointPermission, error) {

	r, body, errs := newGet(roleClient.config, roleClient.config.HostAddress+roleClient.getWorkspacePath()+RolesPath+roleId+"/endpoints/"+workspaceId+"/"+strings.TrimLeft(endpoint, "/")).End()
	if errs != nil {
		return nil, fmt.Errorf("could not get role endpoint permission, error: %v", errs)
	}

	err := checkResponse(r, body, errs)
	if err != nil {
		return nil, err
	}

	roleEndpointPermission := &EndpointPermission{}
	err = json.Unmarshal([]byte(body), roleEndpointPermission)
	if err != nil {
		return nil, fmt.Errorf("could not parse role endpoint permission get response, error: %v", err)
	}

	return roleEndpointPermission, nil
}

func (roleClient *RoleClient) ListEndpointPermissions(roleId string) (*EndpointPermissions, error) {

	r, body, errs := newGet(roleClient.config, roleClient.config.HostAddress+roleClient.getWorkspacePath()+RolesPath+roleId+"/endpoints").End()
	if errs != nil {
		return nil, fmt.Errorf("could not get role endpoint permission, error: %v", errs)
	}

	err := checkResponse(r, body, errs)
	if err != nil {
		return nil, err
	}

	roleEndpointPermissions := &EndpointPermissions{}
	err = json.Unmarshal([]byte(body), roleEndpointPermissions)
	if err != nil {
		return nil, fmt.Errorf("could not parse role endpoint permission list response, error: %v", err)
	}

	return roleEndpointPermissions, nil
}

func (roleClient *RoleClient) UpdateEndpointPermissions(roleId string, workspaceId string, endpoint string, roleEpRequest *EndpointPermissionRequest) (*EndpointPermission, error) {

	r, body, errs := newPatch(roleClient.config, roleClient.config.HostAddress+roleClient.getWorkspacePath()+RolesPath+roleId+"/endpoints/"+workspaceId+"/"+strings.TrimLeft(endpoint, "/")).Send(roleEpRequest).End()
	if errs != nil {
		return nil, fmt.Errorf("could not update role endpoint permission, error: %v", errs)
	}

	err := checkResponse(r, body, errs)
	if err != nil {
		return nil, err
	}

	updatedPermission := &EndpointPermission{}
	err = json.Unmarshal([]byte(body), updatedPermission)
	if err != nil {
		return nil, fmt.Errorf("could not parse role endpoint permission update response, error: %v", err)
	}

	if updatedPermission.Actions == nil {
		return nil, nil
	}

	return updatedPermission, nil
}

func (roleClient *RoleClient) DeleteRoleEndpointPermission(roleId string, workspaceId string, endpoint string) error {

	r, body, errs := newDelete(roleClient.config, roleClient.config.HostAddress+roleClient.getWorkspacePath()+RolesPath+roleId+"/endpoints/"+workspaceId+"/"+strings.TrimLeft(endpoint, "/")).End()
	if errs != nil {
		return fmt.Errorf("could not delete role endpoint permission, result: %v error: %v", r, errs)
	}

	err := checkResponse(r, body, errs)
	if err != nil {
		return err
	}

	return nil
}

// Role Entity Permission
func (roleClient *RoleClient) AddEntityPermissionByRole(id string, roleEntityPermissionRequest *EntityPermissionRequest) (*EntityPermission, error) {
	r, body, errs := newPost(roleClient.config, roleClient.config.HostAddress+roleClient.getWorkspacePath()+RolesPath+id+"/entities").Send(roleEntityPermissionRequest).End()
	if errs != nil {
		return nil, fmt.Errorf("could not update role entities, error: %v", errs)
	}

	err := checkResponse(r, body, errs)
	if err != nil {
		return nil, err
	}

	roleEntityPermission := &EntityPermission{}

	err = json.Unmarshal([]byte(body), roleEntityPermission)
	if err != nil {
		return nil, fmt.Errorf("could not parse the entity update response, error: %v", err)
	}
	if roleEntityPermission.EntityId == "" {
		return nil, nil
	}

	return roleEntityPermission, nil
}

func (roleClient *RoleClient) GetEntityPermission(roleId string, entityId string) (*EntityPermission, error) {

	r, body, errs := newGet(roleClient.config, roleClient.config.HostAddress+roleClient.getWorkspacePath()+RolesPath+roleId+"/entities/"+entityId).End()
	if errs != nil {
		return nil, fmt.Errorf("could not get role entity permission, error: %v", errs)
	}

	err := checkResponse(r, body, errs)
	if err != nil {
		return nil, err
	}

	roleEntityPermission := &EntityPermission{}
	err = json.Unmarshal([]byte(body), roleEntityPermission)
	if err != nil {
		return nil, fmt.Errorf("could not parse role entity permission get response, error: %v", err)
	}

	return roleEntityPermission, nil
}

func (roleClient *RoleClient) ListEntityPermissions(roleId string) (*EntityPermissions, error) {

	r, body, errs := newGet(roleClient.config, roleClient.config.HostAddress+roleClient.getWorkspacePath()+RolesPath+roleId+"/entities").End()
	if errs != nil {
		return nil, fmt.Errorf("could not get role entity permission, error: %v", errs)
	}

	err := checkResponse(r, body, errs)
	if err != nil {
		return nil, err
	}

	roleEntityPermissions := &EntityPermissions{}
	err = json.Unmarshal([]byte(body), roleEntityPermissions)
	if err != nil {
		return nil, fmt.Errorf("could not parse role entity permission list response, error: %v", err)
	}

	return roleEntityPermissions, nil
}

func (roleClient *RoleClient) UpdateEntityPermissions(roleId string, entityId string, roleEpRequest *EntityPermissionRequest) (*EntityPermission, error) {

	r, body, errs := newPatch(roleClient.config, roleClient.config.HostAddress+roleClient.getWorkspacePath()+RolesPath+roleId+"/entities/"+entityId).Send(roleEpRequest).End()
	if errs != nil {
		return nil, fmt.Errorf("could not update role entity permission, error: %v", errs)
	}

	err := checkResponse(r, body, errs)
	if err != nil {
		return nil, err
	}

	updatedPermission := &EntityPermission{}
	err = json.Unmarshal([]byte(body), updatedPermission)
	if err != nil {
		return nil, fmt.Errorf("could not parse role entity permission update response, error: %v", err)
	}

	if updatedPermission.Actions == nil {
		return nil, nil
	}

	return updatedPermission, nil
}

func (roleClient *RoleClient) DeleteRoleEntityPermission(roleId string, entityId string) error {

	r, body, errs := newDelete(roleClient.config, roleClient.config.HostAddress+roleClient.getWorkspacePath()+RolesPath+roleId+"/entities/"+entityId).End()
	if errs != nil {
		return fmt.Errorf("could not delete role entity permission, result: %v error: %v", r, errs)
	}

	err := checkResponse(r, body, errs)
	if err != nil {
		return err
	}

	return nil
}
