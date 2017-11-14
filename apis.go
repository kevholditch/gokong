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

func (apiClient *ApiClient) GetById(id string) (*NewApi, error) {

	return &NewApi{}, nil
}

func (apiClient *ApiClient) Create(newApi *NewApi) (*Api, error) {

	_, body, errs := apiClient.client.Post(NewUrlBuilder(apiClient.hostAddress).Apis().Build()).
		Send(newApi).
		End()
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
