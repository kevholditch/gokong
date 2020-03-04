package gokong

import (
	"encoding/json"
	"fmt"
)

// UserClient manages kong client constructs in the context of User endpoint interactions
type UserClient struct {
	config *Config
}

// UserRequest represents an RBAC user in Kong. This feature is only available in Enterprise.
type UserRequest struct {
	Name      string `json:"name" yaml:"name"`
	UserToken string `json:"user_token" yaml:"user_token"`
	Enabled   bool   `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	Comment   string `json:"comment,omitempty" yaml:"comment,omitempty"`
	CreatedAt *int   `json:"created_at,omitempty" yaml:"created_at,omitempty"`
}

// User represents an RBAC user in Kong, it is generally used to marshal payloads coming back from Kong.
type User struct {
	Id string `json:"id" yaml:"id"`
	UserRequest
}

type Users struct {
	Data []*User `json:"data" yaml:"data"`
	Next string  `json:"next,omitempty" yaml:"next,omitempty"`
}

type UserRoleRequest struct {
	Roles string `json:"roles" yaml:"roles"`
}

type UserRoles struct {
	Roles *[]Role `json:"roles" yaml:"roles"`
	User  *User   `json:"user" yaml:"user"`
}

type UserPermissions struct {
	Endpoints map[string]interface{} `json:"endpoints" yaml:"endpoints"`
	Entities  map[string]interface{} `json:"entitites" yaml:"entities"`
}

const UsersPath = "/rbac/users/"

func (userClient *UserClient) getWorkspacePath() string {
	if userClient.config.Workspace != "" {
		return "/" + userClient.config.Workspace
	}
	return ""
}

func (userClient *UserClient) GetByName(name string) (*User, error) {
	return userClient.GetById(name)
}

func (userClient *UserClient) GetById(id string) (*User, error) {

	r, body, errs := newGet(userClient.config, userClient.config.HostAddress+userClient.getWorkspacePath()+UsersPath+id).End()
	if errs != nil {
		return nil, fmt.Errorf("could not get user, error: %v", errs)
	}

	if r.StatusCode == 401 || r.StatusCode == 403 {
		return nil, fmt.Errorf("not authorised, message from kong: %s", body)
	}

	user := &User{}
	err := json.Unmarshal([]byte(body), user)
	if err != nil {
		return nil, fmt.Errorf("could not parse user get response, error: %v", err)
	}

	if user.Id == "" {
		return nil, nil
	}

	return user, nil
}

func (userClient *UserClient) Create(userRequest *UserRequest) (*User, error) {
	r, body, errs := newPost(userClient.config, userClient.config.HostAddress+userClient.getWorkspacePath()+UsersPath).Send(userRequest).End()
	if errs != nil {
		return nil, fmt.Errorf("could not create new user, error: %v", errs)
	}

	if r.StatusCode == 401 || r.StatusCode == 403 {
		return nil, fmt.Errorf("not authorised, message from kong: %s", body)
	}

	createdUser := &User{}
	err := json.Unmarshal([]byte(body), createdUser)
	if err != nil {
		return nil, fmt.Errorf("could not parse user creation response, error: %v", err)
	}

	if createdUser.Id == "" {
		return nil, fmt.Errorf("could not create update, error: %v", body)
	}

	return createdUser, nil
}

func (userClient *UserClient) DeleteByName(name string) error {
	return userClient.DeleteById(name)
}

func (userClient *UserClient) DeleteById(id string) error {

	r, body, errs := newDelete(userClient.config, userClient.config.HostAddress+userClient.getWorkspacePath()+UsersPath+id).End()
	if errs != nil {
		return fmt.Errorf("could not delete user, result: %v error: %v", r, errs)
	}

	if r.StatusCode == 401 || r.StatusCode == 403 {
		return fmt.Errorf("not authorised, message from kong: %s", body)
	}

	return nil
}

func (userClient *UserClient) List() (*Users, error) {

	r, body, errs := newGet(userClient.config, userClient.config.HostAddress+userClient.getWorkspacePath()+UsersPath).End()
	if errs != nil {
		return nil, fmt.Errorf("could not get users, error: %v", errs)
	}

	if r.StatusCode == 401 || r.StatusCode == 403 {
		return nil, fmt.Errorf("not authorised, message from kong: %s", body)
	}

	users := &Users{}
	err := json.Unmarshal([]byte(body), users)
	if err != nil {
		return nil, fmt.Errorf("could not parse user list response, error: %v", err)
	}

	return users, nil
}

func (userClient *UserClient) UpdateByName(name string, userRequest *UserRequest) (*User, error) {
	return userClient.UpdateById(name, userRequest)
}

func (userClient *UserClient) UpdateById(id string, userRequest *UserRequest) (*User, error) {

	r, body, errs := newPatch(userClient.config, userClient.config.HostAddress+userClient.getWorkspacePath()+UsersPath+id).Send(userRequest).End()
	if errs != nil {
		return nil, fmt.Errorf("could not update user, error: %v", errs)
	}

	if r.StatusCode == 401 || r.StatusCode == 403 {
		return nil, fmt.Errorf("not authorised, message from kong: %s", body)
	}

	updatedUser := &User{}
	err := json.Unmarshal([]byte(body), updatedUser)
	if err != nil {
		return nil, fmt.Errorf("could not parse user update response, error: %v", err)
	}

	if updatedUser.Id == "" {
		return nil, fmt.Errorf("could not update user, error: %v", body)
	}

	return updatedUser, nil
}

func (userClient *UserClient) AddUserToRole(userId string, userRoleRequest *UserRoleRequest) (*UserRoles, error) {
	r, body, errs := newPost(userClient.config, userClient.config.HostAddress+userClient.getWorkspacePath()+UsersPath+userId+"/roles").Send(userRoleRequest).End()
	if errs != nil {
		return nil, fmt.Errorf("could not add user to role, error: %v", errs)
	}

	if r.StatusCode == 401 || r.StatusCode == 403 {
		return nil, fmt.Errorf("not authorised, message from kong: %v", body)
	}

	userRoles := &UserRoles{}

	err := json.Unmarshal([]byte(body), userRoles)
	if err != nil {
		return nil, fmt.Errorf("could not parse the add user roles response, error: %v", err)
	}

	return userRoles, nil
}

func (userClient *UserClient) ListUserRoles(userId string) (*UserRoles, error) {
	r, body, errs := newGet(userClient.config, userClient.config.HostAddress+userClient.getWorkspacePath()+UsersPath+userId+"/roles").End()
	if errs != nil {
		return nil, fmt.Errorf("could not list user roles, error: %v", errs)
	}

	if r.StatusCode == 401 || r.StatusCode == 403 {
		return nil, fmt.Errorf("not authorised, message from kong: %v", body)
	}

	userRoles := &UserRoles{}

	err := json.Unmarshal([]byte(body), userRoles)
	if err != nil {
		return nil, fmt.Errorf("could not parse the list user roles response, error: %v", err)
	}

	return userRoles, nil
}

func (userClient *UserClient) DeleteRoleFromUser(userId string, userRoleRequest *UserRoleRequest) error {
	r, body, errs := newDelete(userClient.config, userClient.config.HostAddress+userClient.getWorkspacePath()+UsersPath+userId+"/roles").Send(userRoleRequest).End()
	if errs != nil {
		return fmt.Errorf("could not list user roles, error: %v", body)
	}

	if r.StatusCode == 401 || r.StatusCode == 403 {
		return fmt.Errorf("not authorised, message from kong: %v", body)
	}

	return nil
}

func (userClient *UserClient) ListUserPermissions(userId string) (*UserPermissions, error) {
	r, body, errs := newGet(userClient.config, userClient.config.HostAddress+userClient.getWorkspacePath()+UsersPath+userId+"/permissions").End()
	if errs != nil {
		return nil, fmt.Errorf("could not list user permissions, error: %v", errs)
	}

	if r.StatusCode == 401 || r.StatusCode == 403 {
		return nil, fmt.Errorf("not authorised, message from kong: %v", body)
	}

	userPermissions := &UserPermissions{}

	err := json.Unmarshal([]byte(body), userPermissions)
	if err != nil {
		return nil, fmt.Errorf("could not parse the list user permissions response, error: %v", err)
	}

	return userPermissions, nil
}
