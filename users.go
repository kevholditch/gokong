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

const UsersPath = "/rbac/users/"

func (userClient *UserClient) GetByName(name string) (*User, error) {
	return userClient.GetById(name)
}

func (userClient *UserClient) GetById(id string) (*User, error) {

	r, body, errs := newGet(userClient.config, userClient.config.HostAddress+UsersPath+id).End()
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
	r, body, errs := newPost(userClient.config, userClient.config.HostAddress+UsersPath).Send(userRequest).End()
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

	r, body, errs := newDelete(userClient.config, userClient.config.HostAddress+UsersPath+id).End()
	if errs != nil {
		return fmt.Errorf("could not delete user, result: %v error: %v", r, errs)
	}

	if r.StatusCode == 401 || r.StatusCode == 403 {
		return fmt.Errorf("not authorised, message from kong: %s", body)
	}

	return nil
}

func (userClient *UserClient) List() (*Users, error) {

	r, body, errs := newGet(userClient.config, userClient.config.HostAddress+UsersPath).End()
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

	r, body, errs := newPatch(userClient.config, userClient.config.HostAddress+UsersPath+id).Send(userRequest).End()
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
