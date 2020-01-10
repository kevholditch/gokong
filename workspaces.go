package gokong

import (
	"encoding/json"
	"fmt"
	"strings"
)

type WorkspaceClient struct {
	config *Config
}

type WorkspaceRequest struct {
	Name    *string `json:"name" yaml:"name"`
	Comment *string `json:"comment" yaml:"comment"`
}

type WorkspaceEntitiesRequest struct {
	Entities *string `json:"entities" yaml:"entities"`
}

type Workspace struct {
	Id      *string     `json:"id" yaml:"id"`
	Name    *string     `json:"name" yaml:"name"`
	Comment *string     `json:"comment" yaml:"comment"`
	Config  interface{} `json:"config,omitempty" yaml:"config,omitempty"`
	Meta    interface{} `json:"meta,omitempty" yaml:"meta,omitempty"`
}

type Workspaces struct {
	Data   []*Workspace `json:"data" yaml:"data,omitempty"`
	Next   *string      `json:"next" yaml:"next,omitempty"`
	Offset string       `json:"offset,omitempty" yaml:"offset,omitempty"`
}

type WorkspaceQueryString struct {
	Offset *string `json:"offset,omitempty" yaml:"offset,omitempty"`
	Size   int     `json:"size" yaml:"size,omitempty"`
}

type WorkspaceEntity struct {
	WorkspaceId      *string `json:"workspace_id" yaml:"workspace_id"`
	WorkspaceName    *string `json:"workspace_name" yaml:"workspace_name"`
	EntityId         *string `json:"entity_id" yaml:"entity_id"`
	EntityType       *string `json:"entity_type" yaml:"entity_type"`
	UniqueFieldName  *string `json:"unique_field_name" yaml:"unique_field_name"`
	UniqueFieldValue *string `json:"unique_field_value" yaml:"unique_field_value"`
}

type WorkspaceEntities struct {
	Data  []*WorkspaceEntity `json:"data" yaml:"data,omitempty"`
	Total int                `json:"total,omitempty" yaml:"total,omitempty"`
}

const WorkspacesPath = "/workspaces/"

func (workspaceClient *WorkspaceClient) GetByName(name string) (*Workspace, error) {
	return workspaceClient.Get(name)
}

func (workspaceClient *WorkspaceClient) Get(id string) (*Workspace, error) {
	r, body, errs := newRawGet(workspaceClient.config, workspaceClient.config.HostAddress+WorkspacesPath+id).End()
	if errs != nil {
		return nil, fmt.Errorf("could not get workspace, error: %v", errs)
	}

	if r.StatusCode == 400 {
		return nil, fmt.Errorf("bad request, message from kong: %s", body)
	}

	if r.StatusCode == 401 || r.StatusCode == 403 {
		return nil, fmt.Errorf("not authorised, message from kong: %s", body)
	}

	workspace := &Workspace{}
	err := json.Unmarshal([]byte(body), workspace)
	if err != nil {
		return nil, fmt.Errorf("could not parse workspace workspace response, error: %v", err)
	}

	if workspace.Id == nil {
		return nil, nil
	}

	return workspace, nil
}

func (workspaceClient *WorkspaceClient) List(query *WorkspaceQueryString) ([]*Workspace, error) {
	workspaces := make([]*Workspace, 0)

	if query.Size < 100 {
		query.Size = 100
	}

	if query.Size > 1000 {
		query.Size = 1000
	}

	for {
		data := &Workspaces{}

		r, body, errs := newRawGet(workspaceClient.config, workspaceClient.config.HostAddress+WorkspacesPath).Query(*query).End()
		if errs != nil {
			return nil, fmt.Errorf("could not get workspaces, error: %v", errs)
		}

		if r.StatusCode == 400 {
			return nil, fmt.Errorf("bad request, message from kong: %s", body)
		}

		if r.StatusCode == 401 || r.StatusCode == 403 {
			return nil, fmt.Errorf("not authorised, message from kong: %s", body)
		}

		err := json.Unmarshal([]byte(body), data)
		if err != nil {
			return nil, fmt.Errorf("could not parse workspaces list response, error: %v", err)
		}

		workspaces = append(workspaces, data.Data...)

		if data.Next == nil || data.Next == String("") {
			break
		}

		query.Offset = &data.Offset
	}

	return workspaces, nil
}

func (workspaceClient *WorkspaceClient) Create(workspaceRequest *WorkspaceRequest) (*Workspace, error) {
	r, body, errs := newRawPost(workspaceClient.config, workspaceClient.config.HostAddress+WorkspacesPath).Send(workspaceRequest).End()
	if errs != nil {
		return nil, fmt.Errorf("could not create new workspace, error: %v", errs)
	}

	if r.StatusCode == 400 {
		return nil, fmt.Errorf("bad request, message from kong: %s", body)
	}

	if r.StatusCode == 401 || r.StatusCode == 403 {
		return nil, fmt.Errorf("not authorised, message from kong: %s", body)
	}

	workspace := &Workspace{}
	err := json.Unmarshal([]byte(body), workspace)
	if err != nil {
		return nil, fmt.Errorf("could not parse workspace creation response, error: %v kong response: %s", err, body)
	}

	if workspace.Id == nil {
		return nil, fmt.Errorf("could not create workspace, err: %v", body)
	}

	return workspace, nil
}

func (workspaceClient *WorkspaceClient) Update(workspaceRequest *WorkspaceRequest) (*Workspace, error) {
	requestPath := fmt.Sprintf(
		"%s%s%s",
		workspaceClient.config.HostAddress,
		WorkspacesPath,
		workspaceClient.config.Workspace,
	)
	r, body, errs := newRawPatch(workspaceClient.config, requestPath).Send(workspaceRequest).End()
	if errs != nil {
		return nil, fmt.Errorf("could not update workspace, error: %v", errs)
	}

	if r.StatusCode == 400 {
		return nil, fmt.Errorf("bad request, message from kong: %s", body)
	}

	if r.StatusCode == 401 || r.StatusCode == 403 {
		return nil, fmt.Errorf("not authorised, message from kong: %s", body)
	}

	workspace := &Workspace{}
	err := json.Unmarshal([]byte(body), workspace)
	if err != nil {
		return nil, fmt.Errorf("could not parse workspace update response, error: %v kong response: %s", err, body)
	}

	if workspace.Id == nil {
		return nil, fmt.Errorf("could not update workspace, error: %v", body)
	}

	return workspace, nil
}

func (workspaceClient *WorkspaceClient) Delete() error {
	requestPath := fmt.Sprintf(
		"%s%s%s",
		workspaceClient.config.HostAddress,
		WorkspacesPath,
		workspaceClient.config.Workspace,
	)
	r, body, errs := newRawDelete(workspaceClient.config, requestPath).End()
	if errs != nil {
		return fmt.Errorf("could not delete workspace, result: %v error: %v", r, errs)
	}

	if r.StatusCode == 400 {
		return fmt.Errorf("bad request, message from kong: %s", body)
	}

	if r.StatusCode == 401 || r.StatusCode == 403 {
		return fmt.Errorf("not authorised, message from kong: %s", body)
	}

	return nil
}

func (workspaceClient *WorkspaceClient) ListEntities() ([]*WorkspaceEntity, error) {
	requestPath := fmt.Sprintf(
		"%s%s/entities",
		workspaceClient.config.HostAddress+WorkspacesPath,
		workspaceClient.config.Workspace,
	)
	r, body, errs := newRawGet(workspaceClient.config, requestPath).End()
	if errs != nil {
		return nil, fmt.Errorf("could not get workspaces, error: %v", errs)
	}

	if r.StatusCode == 400 {
		return nil, fmt.Errorf("bad request, message from kong: %s", body)
	}

	if r.StatusCode == 401 || r.StatusCode == 403 {
		return nil, fmt.Errorf("not authorised, message from kong: %s", body)
	}

	workspaceEntities := &WorkspaceEntities{}
	err := json.Unmarshal([]byte(body), workspaceEntities)
	if err != nil {
		return nil, fmt.Errorf("could not parse workspaces list response, error: %v", err)
	}

	return workspaceEntities.Data, nil
}

func (workspaceClient *WorkspaceClient) DeleteMultipleEntitiesFromWorkspace(entityIds []string) error {
	requestPath := fmt.Sprintf(
		"%s%s/entities",
		workspaceClient.config.HostAddress+WorkspacesPath,
		workspaceClient.config.Workspace,
	)
	workspaceEntitiesRequest := &WorkspaceEntitiesRequest{
		Entities: String(strings.Join(entityIds, ",")),
	}

	r, body, errs := newRawDelete(workspaceClient.config, requestPath).Send(workspaceEntitiesRequest).End()
	if errs != nil {
		return fmt.Errorf("could not delete workspace entities, error: %v", errs)
	}

	if r.StatusCode == 400 {
		return fmt.Errorf("bad request, message from kong: %s", body)
	}

	if r.StatusCode == 401 || r.StatusCode == 403 {
		return fmt.Errorf("not authorised, message from kong: %s", body)
	}

	return nil
}
