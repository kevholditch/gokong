package konggo

import "github.com/parnurzeal/gorequest"

type ApiClient struct {
	hostAddress string
	client      *gorequest.SuperAgent
}

type ApiRequest struct {
}

func (apis *ApiClient) GetById(id string) (*ApiRequest, error) {

	return &ApiRequest{}, nil
}

func (apiClient *ApiClient) Create(newApi *ApiRequest) (*ApiRequest, error) {

	return &ApiRequest{}, nil
}
