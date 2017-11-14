package konggo

import "github.com/parnurzeal/gorequest"

type ApiClient struct {
	hostAddress string
	client      *gorequest.SuperAgent
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

func (apiClient *ApiClient) GetById(id string) (*ApiRequest, error) {

	return &ApiRequest{}, nil
}

func (apiClient *ApiClient) Create(newApi *ApiRequest) (*ApiRequest, error) {

	return &ApiRequest{}, nil
}
