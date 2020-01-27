package gokong

import (
	"encoding/json"
	"fmt"
)

type WorkspaceClient struct {
	config *Config
}
type WorkspaceRequest struct {
	Name      string                 `json:"name" yaml:"name"`
	Comment   string                 `json:"comment,omitempty" yaml:"comment,omitempty"`
	Config    map[string]interface{} `json:"config,omitempty" yaml:"config,omitempty"`
	CreatedAt *int                   `json:"created_at,omitempty" yaml:"created_at,omitempty"`
	Meta      map[string]interface{} `json:"meta,omitempty" yaml:"meta,omitempty"`
}

type Workspace struct {
	Id string `json:"id" yaml:"id"`
	WorkspaceRequest
}

type Workspaces struct {
	Data []*Workspace `json:"data" yaml:"data"`
	Next string       `json:"next,omitempty" yaml:"next,omitempty"`
}

// type WorkspacesEntities struct {
// 	Data []*Workspace `json:"data" yaml:"data"`
// 	Next string       `json:"next,omitempty" yaml:"next,omitempty"`
// }

const WorkspacesPath = "/workspaces/"

func (workspaceClient *WorkspaceClient) GetByName(name string) (*Workspace, error) {
	return workspaceClient.GetById(name)
}

func (workspaceClient *WorkspaceClient) GetById(id string) (*Workspace, error) {

	r, body, errs := newGet(workspaceClient.config, workspaceClient.config.HostAddress+WorkspacesPath+id).End()
	if errs != nil {
		return nil, fmt.Errorf("could not get workspace, error: %v", errs)
	}

	if r.StatusCode == 401 || r.StatusCode == 403 {
		return nil, fmt.Errorf("not authorised, message from kong: %s", body)
	}

	workspace := &Workspace{}
	err := json.Unmarshal([]byte(body), workspace)
	if err != nil {
		return nil, fmt.Errorf("could not parse workspace get response, error: %v", err)
	}

	if workspace.Id == "" {
		return nil, nil
	}

	return workspace, nil
}

func (workspaceClient *WorkspaceClient) Create(workspaceRequest *WorkspaceRequest) (*Workspace, error) {

	r, body, errs := newPost(workspaceClient.config, workspaceClient.config.HostAddress+WorkspacesPath).Send(workspaceRequest).End()
	if errs != nil {
		return nil, fmt.Errorf("could not create new workspace, error: %v", errs)
	}

	if r.StatusCode == 401 || r.StatusCode == 403 {
		return nil, fmt.Errorf("not authorised, message from kong: %s", body)
	}

	createdWorkspace := &Workspace{}
	err := json.Unmarshal([]byte(body), createdWorkspace)
	if err != nil {
		return nil, fmt.Errorf("could not parse workspace creation response, error: %v", err)
	}

	if createdWorkspace.Id == "" {
		return nil, fmt.Errorf("could not create update, error: %v", body)
	}

	return createdWorkspace, nil
}

func (workspaceClient *WorkspaceClient) DeleteByName(name string) error {
	return workspaceClient.DeleteById(name)
}

func (workspaceClient *WorkspaceClient) DeleteById(id string) error {

	r, body, errs := newDelete(workspaceClient.config, workspaceClient.config.HostAddress+WorkspacesPath+id).End()
	if errs != nil {
		return fmt.Errorf("could not delete workspace, result: %v error: %v", r, errs)
	}

	if r.StatusCode == 401 || r.StatusCode == 403 {
		return fmt.Errorf("not authorised, message from kong: %s", body)
	}

	return nil
}

func (workspaceClient *WorkspaceClient) List() (*Workspaces, error) {

	r, body, errs := newGet(workspaceClient.config, workspaceClient.config.HostAddress+WorkspacesPath).End()
	if errs != nil {
		return nil, fmt.Errorf("could not get workspaces, error: %v", errs)
	}

	if r.StatusCode == 401 || r.StatusCode == 403 {
		return nil, fmt.Errorf("not authorised, message from kong: %s", body)
	}

	workspaces := &Workspaces{}
	err := json.Unmarshal([]byte(body), workspaces)
	if err != nil {
		return nil, fmt.Errorf("could not parse workspaces list response, error: %v", err)
	}

	return workspaces, nil
}

func (workspaceClient *WorkspaceClient) UpdateByName(name string, workspaceRequest *WorkspaceRequest) (*Workspace, error) {
	return workspaceClient.UpdateById(name, workspaceRequest)
}

func (workspaceClient *WorkspaceClient) UpdateById(id string, workspaceRequest *WorkspaceRequest) (*Workspace, error) {

	r, body, errs := newPatch(workspaceClient.config, workspaceClient.config.HostAddress+WorkspacesPath+id).Send(workspaceRequest).End()
	if errs != nil {
		return nil, fmt.Errorf("could not update workspace, error: %v", errs)
	}

	if r.StatusCode == 401 || r.StatusCode == 403 {
		return nil, fmt.Errorf("not authorised, message from kong: %s", body)
	}

	updatedWorkspace := &Workspace{}
	err := json.Unmarshal([]byte(body), updatedWorkspace)
	if err != nil {
		return nil, fmt.Errorf("could not parse workspace update response, error: %v", err)
	}

	if updatedWorkspace.Id == "" {
		return nil, fmt.Errorf("could not update workspace, error: %v", body)
	}

	return updatedWorkspace, nil
}
