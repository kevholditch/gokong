package gokong

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/parnurzeal/gorequest"
)

type APIClient struct {
	config *Config
}

type APIRequest struct {
	ID                     string   `json:"id,omitempty"`
	Name                   string   `json:"name"`
	CreatedAt              int      `json:"created_at,omitempty"`
	Hosts                  []string `json:"hosts,omitempty"`
	URIs                   []string `json:"uris,omitempty"`
	Methods                []string `json:"methods,omitempty"`
	UpstreamURL            string   `json:"upstream_url"`
	StripURI               bool     `json:"strip_uri"`
	PreserveHost           bool     `json:"preserve_host"`
	Retries                int      `json:"retries,omitempty"`
	UpstreamConnectTimeout int      `json:"upstream_connect_timeout,omitempty"`
	UpstreamSendTimeout    int      `json:"upstream_send_timeout,omitempty"`
	UpstreamReadTimeout    int      `json:"upstream_read_timeout,omitempty"`
	HTTPSOnly              bool     `json:"https_only"`
	HTTPIfTerminated       bool     `json:"http_if_terminated"`
}

type API struct {
	ID                     string   `json:"id"`
	CreatedAt              int      `json:"created_at"`
	Name                   string   `json:"name"`
	Hosts                  []string `json:"hosts,omitempty"`
	URIs                   []string `json:"uris,omitempty"`
	Methods                []string `json:"methods,omitempty"`
	UpstreamURL            string   `json:"upstream_url"`
	StripURI               bool     `json:"strip_uri,omitempty"`
	PreserveHost           bool     `json:"preserve_host,omitempty"`
	Retries                int      `json:"retries,omitempty"`
	UpstreamConnectTimeout int      `json:"upstream_connect_timeout,omitempty"`
	UpstreamSendTimeout    int      `json:"upstream_send_timeout,omitempty"`
	UpstreamReadTimeout    int      `json:"upstream_read_timeout,omitempty"`
	HTTPSOnly              bool     `json:"https_only,omitempty"`
	HTTPIfTerminated       bool     `json:"http_if_terminated,omitempty"`
}

type APIs struct {
	Results []*API `json:"data,omitempty"`
	Total   int    `json:"total,omitempty"`
	Next    string `json:"next,omitempty"`
	Offset  string `json:"offset,omitempty"`
}

type APIFilter struct {
	ID          string `url:"id,omitempty"`
	Name        string `url:"name,omitempty"`
	UpstreamURL string `url:"upstream_url,omitempty"`
	Retries     int    `url:"retries,omitempty"`
	Size        int    `url:"size,omitempty"`
	Offset      int    `url:"offset,omitempty"`
}

const APIsPath = "/apis/"

func (apiClient *APIClient) GetByName(name string) (*API, error) {
	return apiClient.GetByID(name)
}

func (apiClient *APIClient) GetByID(id string) (*API, error) {
	address := apiClient.config.HostAddress + APIsPath + id
	_, body, errs := gorequest.New().Get(address).End()
	if errs != nil {
		return nil, fmt.Errorf("could not get api, error: %v", errs)
	}

	log.Printf("DEBUG: GetByID address: %s \n body: %s \n", address, body)

	api := &API{}
	err := json.Unmarshal([]byte(body), api)
	if err != nil {
		return nil, fmt.Errorf("could not parse api get response, error: %v", err)
	}

	if api.ID == "" {
		return nil, nil
	}

	return api, nil
}

func (apiClient *APIClient) List() (*APIs, error) {
	return apiClient.ListFiltered(nil)
}

func (apiClient *APIClient) ListFiltered(filter *APIFilter) (*APIs, error) {

	address, err := addQueryString(apiClient.config.HostAddress+APIsPath, filter)

	if err != nil {
		return nil, fmt.Errorf("could not build query string for apis filter, error: %v", err)
	}

	_, body, errs := gorequest.New().Get(address).End()
	if errs != nil {
		return nil, fmt.Errorf("could not get apis, error: %v", errs)
	}

	apis := &APIs{}
	err = json.Unmarshal([]byte(body), apis)
	if err != nil {
		return nil, fmt.Errorf("could not parse apis list response, error: %v", err)
	}

	return apis, nil
}

func (apiClient *APIClient) Create(newAPI *APIRequest) (*API, error) {

	_, body, errs := gorequest.New().Post(apiClient.config.HostAddress + APIsPath).Send(newAPI).End()
	if errs != nil {
		return nil, fmt.Errorf("could not create new api, error: %v", errs)
	}

	createdAPI := &API{}
	err := json.Unmarshal([]byte(body), createdAPI)
	if err != nil {
		return nil, fmt.Errorf("could not parse api creation response, error: %v %s", err, body)
	}

	if createdAPI.ID == "" {
		return nil, fmt.Errorf("could not create api, error: %v", body)
	}

	return createdAPI, nil
}

func (apiClient *APIClient) DeleteByName(name string) error {
	return apiClient.DeleteByID(name)
}

func (apiClient *APIClient) DeleteByID(id string) error {

	res, _, errs := gorequest.New().Delete(apiClient.config.HostAddress + APIsPath + id).End()
	if errs != nil {
		return fmt.Errorf("could not delete api, result: %v error: %v", res, errs)
	}

	return nil
}

func (apiClient *APIClient) UpdateByName(name string, apiRequest *APIRequest) (*API, error) {
	return apiClient.UpdateByID(name, apiRequest)
}

func (apiClient *APIClient) UpdateByID(id string, apiRequest *APIRequest) (*API, error) {

	_, body, errs := gorequest.New().Patch(apiClient.config.HostAddress + APIsPath + id).Send(apiRequest).End()
	if errs != nil {
		return nil, fmt.Errorf("could not update api, error: %v", errs)
	}

	updatedAPI := &API{}
	err := json.Unmarshal([]byte(body), updatedAPI)
	if err != nil {
		return nil, fmt.Errorf("could not parse api update response, error: %v", err)
	}

	if updatedAPI.ID == "" {
		return nil, fmt.Errorf("could not update certificate, error: %v", body)
	}

	return updatedAPI, nil
}

func (apiClient *APIClient) CreateOrUpdate(apiRequest *APIRequest) (*API, error) {

	foundAPI := false
	api, err := apiClient.GetByName(apiRequest.Name)
	if err != nil {
		return nil, fmt.Errorf("WORDZ: could get api, error: %v", err)
	}

	if api != nil {
		apiRequest.ID = api.ID
		apiRequest.CreatedAt = api.CreatedAt
		foundAPI = true
	}

	_, body, errs := gorequest.New().Put(apiClient.config.HostAddress + APIsPath).Send(apiRequest).End()
	if errs != nil {
		return nil, fmt.Errorf("could not update api, error: %v", errs)
	}

	updatedAPI := &API{}
	err = json.Unmarshal([]byte(body), updatedAPI)
	if err != nil {
		return nil, fmt.Errorf("could not parse api update response, error: %v %s %+v found: %v", err, body, apiRequest, foundAPI)
	}

	if updatedAPI.ID == "" {
		return nil, fmt.Errorf("could not update certificate, error: %v", body)
	}

	return updatedAPI, nil
}
