package gokong

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/go-querystring/query"
	"github.com/parnurzeal/gorequest"
	"net/url"
)

type ApiClient struct {
	config *Config
	client *gorequest.SuperAgent
}

type ApiRequest struct {
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

type Apis struct {
	Results []*Api `json:"data,omitempty"`
	Total   int    `json:"total,omitempty"`
	Next    string `json:"next,omitempty"`
	Offset  string `json:"offset,omitempty"`
}

type GetAllFilter struct {
	Id          string `url:"id"`
	Name        string `url:"name"`
	UpstreamUrl string `url:"upstream_url"`
	Retries     string `url:"retries"`
	Size        string `url:"size"`
	Offset      string `url:"offset"`
}

const ApisPath = "/apis/"

func (apiClient *ApiClient) GetById(id string) (*Api, error) {

	_, body, errs := apiClient.client.Get(apiClient.config.HostAddress + ApisPath + id).End()
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

func (apiClient *ApiClient) GetAll() (*Apis, error) {
	return apiClient.GetAllFiltered(nil)
}

func buildQueryString(filter *GetAllFilter) string {
	v, _ := query.Values(filter)
	queryStringValues := make(url.Values)

	for key, values := range v {
		if len(values) > 0 && values[0] != "" {
			queryStringValues[key] = values
		}
	}

	return fmt.Sprintf(queryStringValues.Encode())
}

func (apiClient *ApiClient) GetAllFiltered(getAllFilter *GetAllFilter) (*Apis, error) {

	address := apiClient.config.HostAddress + ApisPath
	queryString := buildQueryString(getAllFilter)
	if queryString != "" {
		address += "?" + queryString
	}

	_, body, errs := apiClient.client.Get(address).End()
	if errs != nil {
		return nil, errors.New(fmt.Sprintf("Could not get apis, error: %v", errs))
	}

	apis := &Apis{}
	err := json.Unmarshal([]byte(body), apis)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Could not parse apis get response, error: %v", err))
	}

	return apis, nil
}

func (apiClient *ApiClient) Create(newApi *ApiRequest) (*Api, error) {

	_, body, errs := apiClient.client.Post(apiClient.config.HostAddress + ApisPath).Send(newApi).End()
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
