package konggo

import (
	"github.com/parnurzeal/gorequest"
	"strings"
)

const EnvKongAdminHostAddress = "KONG_ADMIN_ADDR"

type KongAdminClient struct {
	hostAddress string
	client      *gorequest.SuperAgent
}

func NewClient() *KongAdminClient {
	return &KongAdminClient{
		hostAddress: strings.TrimRight(GetEnvOrDefault("KONG_ADMIN_ADDR", "http://localhost:8001"), "/"),
		client:      gorequest.New(),
	}
}

func (kongAdminClient *KongAdminClient) Status() *StatusClient {
	return &StatusClient{
		hostAddress: kongAdminClient.hostAddress,
		client:      kongAdminClient.client,
	}

}

func (kongAdminClient *KongAdminClient) Apis() *ApiClient {
	return &ApiClient{
		hostAddress: kongAdminClient.hostAddress,
		client:      kongAdminClient.client,
	}
}
