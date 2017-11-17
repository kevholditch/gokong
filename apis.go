package konggo

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/parnurzeal/gorequest"
)

type ApiClient struct {
	hostAddress string
	client      *gorequest.SuperAgent
}

type NewApi struct {
	Name                   string   `json:"name"`
	Hosts                  []string `json:"hosts,omitempty"`
	Uris                   []string `json:"uris,omitempty"`
	Methods                []string `json:"methods,omitempty"`
	UpstreamUrl            string   `json:"upstream_url"`
	StripUri               bool     `json:"strip_uri,omitempty"`
	PreserveHost           bool     `json:"preserve_host,omitempty"`
	Retries                int      `json:"retries,omitempty"`
	UpstreamConnectTimeout int      `json:"upstream_connect_timeout,omitempty"`
	UpstreamSendTimeout    int      `json:"upstream_send_timeout,omitempty"`
	UpstreamReadTimeout    int      `json:"upstream_read_timeout,omitempty"`
	HttpsOnly              bool     `json:"https_only,omitempty"`
	HttpIfTerminated       bool     `json:"http_if_terminated,omitempty"`
}

type Api struct {
	Id                     string   `json:"id"`
	CreatedAt              int      `json:"created_at"`
	Name                   string   `json:"name"`
	Hosts                  []string `json:"hosts,omitempty"`
	Uris                   []string `json:"uris,omitempty"`
	Methods                []string `json:"methods,omitempty"`
	UpstreamUrl            string   `json:"upstream_url"`
	StripUri               bool     `json:"strip_uri,omitempty"`
	PreserveHost           bool     `json:"preserve_host,omitempty"`
	Retries                int      `json:"retries,omitempty"`
	UpstreamConnectTimeout int      `json:"upstream_connect_timeout,omitempty"`
	UpstreamSendTimeout    int      `json:"upstream_send_timeout,omitempty"`
	UpstreamReadTimeout    int      `json:"upstream_read_timeout,omitempty"`
	HttpsOnly              bool     `json:"https_only,omitempty"`
	HttpIfTerminated       bool     `json:"http_if_terminated,omitempty"`
}

const ApisPath = "/apis/"

func (apiClient *ApiClient) GetById(id string) (*Api, error) {

	_, body, errs := apiClient.client.Get(apiClient.hostAddress + ApisPath + id).End()
	if errs != nil {
		return nil, errors.New(fmt.Sprintf("Could not get api, error: %v", errs))
	}

	api := &Api{}
	err := json.Unmarshal([]byte(body), api)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Could not parse api get by id response, error: %v", err))
	}

	return api, nil
}

func (apiClient *ApiClient) GetAll() ([] *Api, error) {

	return nil, nil
}

func (apiClient *ApiClient) Create(newApi *NewApi) (*Api, error) {

	_, body, errs := apiClient.client.Post(apiClient.hostAddress + ApisPath).Send(newApi).End()
	if errs != nil {
		return nil, errors.New(fmt.Sprintf("Could not create new api, error: %v", errs))
	}

	createdApi := &Api{}
	err := json.Unmarshal([]byte(body), createdApi)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Could not parse api creation response, error: %v", err))
	}

	return createdApi, nil
}
