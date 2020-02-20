package gokong

import (
	"encoding/json"
	"fmt"
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

const RolesPath = "/rbac/roles/"

func (roleClient *RoleClient) GetByName(name string) (*Role, error) {
	return roleClient.GetById(name)
}

func (roleClient *RoleClient) GetById(id string) (*Role, error) {

	r, body, errs := newGet(roleClient.config, roleClient.config.HostAddress+RolesPath+id).End()
	if errs != nil {
		return nil, fmt.Errorf("could not get role, error: %v", errs)
	}

	if r.StatusCode == 401 || r.StatusCode == 403 {
		return nil, fmt.Errorf("not authorised, message from kong: %s", body)
	}

	role := &Role{}
	err := json.Unmarshal([]byte(body), role)
	if err != nil {
		return nil, fmt.Errorf("could not parse role get response, error: %v", err)
	}

	if role.Id == "" {
		return nil, nil
	}

	return role, nil
}

func (roleClient *RoleClient) Create(roleRequest *RoleRequest) (*Role, error) {
	r, body, errs := newPost(roleClient.config, roleClient.config.HostAddress+RolesPath).Send(roleRequest).End()
	if errs != nil {
		return nil, fmt.Errorf("could not create new role, error: %v", errs)
	}

	if r.StatusCode == 401 || r.StatusCode == 403 {
		return nil, fmt.Errorf("not authorised, message from kong: %s", body)
	}

	createdRole := &Role{}
	err := json.Unmarshal([]byte(body), createdRole)
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

	r, body, errs := newDelete(roleClient.config, roleClient.config.HostAddress+RolesPath+id).End()
	if errs != nil {
		return fmt.Errorf("could not delete role, result: %v error: %v", r, errs)
	}

	if r.StatusCode == 401 || r.StatusCode == 403 {
		return fmt.Errorf("not authorised, message from kong: %s", body)
	}

	return nil
}

func (roleClient *RoleClient) List() (*Roles, error) {

	r, body, errs := newGet(roleClient.config, roleClient.config.HostAddress+RolesPath).End()
	if errs != nil {
		return nil, fmt.Errorf("could not get roles, error: %v", errs)
	}

	if r.StatusCode == 401 || r.StatusCode == 403 {
		return nil, fmt.Errorf("not authorised, message from kong: %s", body)
	}

	roles := &Roles{}
	err := json.Unmarshal([]byte(body), roles)
	if err != nil {
		return nil, fmt.Errorf("could not parse role list response, error: %v", err)
	}

	return roles, nil
}

func (roleClient *RoleClient) UpdateByName(name string, roleRequest *RoleRequest) (*Role, error) {
	return roleClient.UpdateById(name, roleRequest)
}

func (roleClient *RoleClient) UpdateById(id string, roleRequest *RoleRequest) (*Role, error) {

	r, body, errs := newPatch(roleClient.config, roleClient.config.HostAddress+RolesPath+id).Send(roleRequest).End()
	if errs != nil {
		return nil, fmt.Errorf("could not update role, error: %v", errs)
	}

	if r.StatusCode == 401 || r.StatusCode == 403 {
		return nil, fmt.Errorf("not authorised, message from kong: %s", body)
	}

	updatedRole := &Role{}
	err := json.Unmarshal([]byte(body), updatedRole)
	if err != nil {
		return nil, fmt.Errorf("could not parse role update response, error: %v", err)
	}

	if updatedRole.Id == "" {
		return nil, fmt.Errorf("could not update role, error: %v", body)
	}

	return updatedRole, nil
}
