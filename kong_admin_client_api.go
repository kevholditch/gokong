package konggo

import "github.com/parnurzeal/gorequest"

type ApiClient struct {
	hostAddress string
	client      *gorequest.SuperAgent
}

type Api struct {
}

func (apiClient *ApiClient) Get(id string) (*Api, error) {

	return &Api{}, nil
}

func (apiClient *ApiClient) Create(api Api) (*Api, error) {

	return &Api{}, nil
}
