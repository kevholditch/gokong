package gokong

import (
	"encoding/json"
	"fmt"
	"strings"
)

type ApiClient struct {
	config *Config
}

type ApiRequest struct {
	Name                   *string   `json:"name"`
	Hosts                  []*string `json:"hosts"`
	Uris                   []*string `json:"uris"`
	Methods                []*string `json:"methods"`
	UpstreamUrl            *string   `json:"upstream_url"`
	StripUri               *bool     `json:"strip_uri,omitempty"`
	PreserveHost           *bool     `json:"preserve_host,omitempty"`
	Retries                *int      `json:"retries,omitempty"`
	UpstreamConnectTimeout *int      `json:"upstream_connect_timeout,omitempty"`
	UpstreamSendTimeout    *int      `json:"upstream_send_timeout,omitempty"`
	UpstreamReadTimeout    *int      `json:"upstream_read_timeout,omitempty"`
	HttpsOnly              *bool     `json:"https_only,omitempty"`
	HttpIfTerminated       *bool     `json:"http_if_terminated,omitempty"`
}

type Api struct {
	Id                     *string   `json:"id"`
	CreatedAt              *int      `json:"created_at"`
	Name                   *string   `json:"name"`
	Hosts                  []*string `json:"hosts,omitempty"`
	Uris                   []*string `json:"uris,omitempty"`
	Methods                []*string `json:"methods,omitempty"`
	UpstreamUrl            *string   `json:"upstream_url"`
	StripUri               *bool     `json:"strip_uri,omitempty"`
	PreserveHost           *bool     `json:"preserve_host,omitempty"`
	Retries                *int      `json:"retries,omitempty"`
	UpstreamConnectTimeout *int      `json:"upstream_connect_timeout,omitempty"`
	UpstreamSendTimeout    *int      `json:"upstream_send_timeout,omitempty"`
	UpstreamReadTimeout    *int      `json:"upstream_read_timeout,omitempty"`
	HttpsOnly              *bool     `json:"https_only,omitempty"`
	HttpIfTerminated       *bool     `json:"http_if_terminated,omitempty"`
}

type apiNoHosts struct {
	Id                     *string   `json:"id"`
	CreatedAt              *int      `json:"created_at"`
	Name                   *string   `json:"name"`
	Uris                   []*string `json:"uris,omitempty"`
	Methods                []*string `json:"methods,omitempty"`
	UpstreamUrl            *string   `json:"upstream_url"`
	StripUri               *bool     `json:"strip_uri,omitempty"`
	PreserveHost           *bool     `json:"preserve_host,omitempty"`
	Retries                *int      `json:"retries,omitempty"`
	UpstreamConnectTimeout *int      `json:"upstream_connect_timeout,omitempty"`
	UpstreamSendTimeout    *int      `json:"upstream_send_timeout,omitempty"`
	UpstreamReadTimeout    *int      `json:"upstream_read_timeout,omitempty"`
	HttpsOnly              *bool     `json:"https_only,omitempty"`
	HttpIfTerminated       *bool     `json:"http_if_terminated,omitempty"`
}

type Apis struct {
	Results []*Api `json:"data,omitempty"`
	Total   int    `json:"total,omitempty"`
	Next    string `json:"next,omitempty"`
	Offset  string `json:"offset,omitempty"`
}

type ApiFilter struct {
	Id          string `url:"id,omitempty"`
	Name        string `url:"name,omitempty"`
	UpstreamUrl string `url:"upstream_url,omitempty"`
	Retries     int    `url:"retries,omitempty"`
	Size        int    `url:"size,omitempty"`
	Offset      int    `url:"offset,omitempty"`
}

const ApisPath = "/apis/"

func (apiClient *ApiClient) GetByName(name string) (*Api, error) {
	return apiClient.GetById(name)
}

func (apiClient *ApiClient) GetById(id string) (*Api, error) {
	_, body, errs := newGet(apiClient.config, apiClient.config.HostAddress+ApisPath+id).End()

	if errs != nil {
		return nil, fmt.Errorf("could not get api, error: %v", errs)
	}

	api := &Api{}
	err := json.Unmarshal([]byte(body), api)
	if err != nil {

		// explicitly check for case where user has updated hosts to [] as there is bug in Kong where an empty object is returned
		// {} instead of empty array
		unMarshalTypeError, ok := err.(*json.UnmarshalTypeError)
		if ok && unMarshalTypeError.Field == "hosts" {

			apiNoHosts := &apiNoHosts{}
			err = json.Unmarshal([]byte(body), apiNoHosts)

			if err != nil {
				return nil, fmt.Errorf("could not parse api get response, error: %v", err)
			}

			api.Id = apiNoHosts.Id
			api.CreatedAt = apiNoHosts.CreatedAt
			api.Hosts = StringSlice([]string{})
			api.Name = apiNoHosts.Name
			api.Uris = apiNoHosts.Uris
			api.Methods = apiNoHosts.Methods
			api.UpstreamUrl = apiNoHosts.UpstreamUrl
			api.StripUri = apiNoHosts.StripUri
			api.PreserveHost = apiNoHosts.PreserveHost
			api.Retries = apiNoHosts.Retries
			api.UpstreamConnectTimeout = apiNoHosts.UpstreamConnectTimeout
			api.UpstreamSendTimeout = apiNoHosts.UpstreamSendTimeout
			api.UpstreamReadTimeout = apiNoHosts.UpstreamReadTimeout
			api.HttpsOnly = apiNoHosts.HttpsOnly
			api.HttpIfTerminated = apiNoHosts.HttpIfTerminated

		}

		if err != nil {
			return nil, fmt.Errorf("could not parse api get response, error: %v", err)
		}

	}

	if api.Id == nil {
		return nil, nil
	}

	return api, nil
}

func (apiClient *ApiClient) List() (*Apis, error) {
	return apiClient.ListFiltered(nil)
}

func (apiClient *ApiClient) ListFiltered(filter *ApiFilter) (*Apis, error) {

	address, err := addQueryString(apiClient.config.HostAddress+ApisPath, filter)

	if err != nil {
		return nil, fmt.Errorf("could not build query string for apis filter, error: %v", err)
	}

	_, body, errs := newGet(apiClient.config, address).End()
	if errs != nil {
		return nil, fmt.Errorf("could not get apis, error: %v", errs)
	}

	apis := &Apis{}
	err = json.Unmarshal([]byte(body), apis)
	if err != nil {
		return nil, fmt.Errorf("could not parse apis list response, error: %v", err)
	}

	return apis, nil
}

func (apiClient *ApiClient) Create(newApi *ApiRequest) (*Api, error) {

	_, body, errs := newPost(apiClient.config, apiClient.config.HostAddress+ApisPath).Send(newApi).End()
	if errs != nil {
		return nil, fmt.Errorf("could not create new api, error: %v", errs)
	}

	createdApi := &Api{}
	err := json.Unmarshal([]byte(body), createdApi)
	if err != nil {
		return nil, fmt.Errorf("could not parse api creation response, error: %v %s", err, body)
	}

	if createdApi.Id == nil {
		return nil, fmt.Errorf("could not create api, error: %v", body)
	}

	return createdApi, nil
}

func (apiClient *ApiClient) DeleteByName(name string) error {
	return apiClient.DeleteById(name)
}

func (apiClient *ApiClient) DeleteById(id string) error {

	res, _, errs := newDelete(apiClient.config, apiClient.config.HostAddress+ApisPath+id).End()
	if errs != nil {
		return fmt.Errorf("could not delete api, result: %v error: %v", res, errs)
	}

	return nil
}

func (apiClient *ApiClient) UpdateByName(name string, apiRequest *ApiRequest) (*Api, error) {
	return apiClient.UpdateById(name, apiRequest)
}

func (apiClient *ApiClient) UpdateById(id string, apiRequest *ApiRequest) (*Api, error) {

	j, _ := json.Marshal(apiRequest)
	js := string(j)
	fmt.Sprintf("%s", js)

	_, body, errs := newPatch(apiClient.config, apiClient.config.HostAddress+ApisPath+id).Send(apiRequest).End()
	if errs != nil {
		return nil, fmt.Errorf("could not update api, error: %v", errs)
	}

	updatedApi := &Api{}
	err := json.Unmarshal([]byte(body), updatedApi)
	if err != nil {
		return nil, fmt.Errorf("could not parse api update response, error: %v", err)
	}

	if updatedApi.Id == nil {
		return nil, fmt.Errorf("could not update api, error: %v", body)
	}

	return updatedApi, nil
}

func (a *ApiRequest) MarshalJSON() ([]byte, error) {

	uris := a.Uris
	if uris == nil {
		uris = make([]*string, 0)
	}

	hosts := a.Hosts
	if hosts == nil {
		hosts = make([]*string, 0)
	}

	methods := a.Methods
	if methods == nil {
		methods = make([]*string, 0)
	}

	type Alias ApiRequest
	return json.Marshal(&struct {
		Uris    []*string `json:"uris"`
		Hosts   []*string `json:"hosts"`
		Methods []*string `json:"methods"`
		*Alias
	}{
		Uris:    uris,
		Hosts:   hosts,
		Methods: methods,
		Alias:   (*Alias)(a),
	})
}

func (a *Api) UnmarshalJSON(data []byte) error {

	fixedJson := strings.Replace(string(data), `"hosts":{}`, `"hosts":[]`, -1)
	fixedJson = strings.Replace(fixedJson, `"uris":{}`, `"uris":[]`, -1)
	fixedJson = strings.Replace(fixedJson, `"methods":{}`, `"methods":[]`, -1)

	type Alias Api
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(a),
	}

	return json.Unmarshal([]byte(fixedJson), &aux)
}
