package gokong

import (
	"net/url"
	"os"
	"reflect"
	"strings"

	"github.com/google/go-querystring/query"
)

const EnvKongAdminHostAddress = "KONG_ADMIN_ADDR"

type KongAdminClient struct {
	config *Config
}

type Config struct {
	HostAddress string
}

func addQueryString(currentURL string, filter interface{}) (string, error) {
	v := reflect.ValueOf(filter)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return currentURL, nil
	}

	u, err := url.Parse(currentURL)
	if err != nil {
		return currentURL, err
	}

	qs, err := query.Values(filter)
	if err != nil {
		return currentURL, err
	}

	u.RawQuery = qs.Encode()
	return u.String(), nil
}

func NewDefaultConfig() *Config {
	config := &Config{
		HostAddress: "http://localhost:8001",
	}

	if os.Getenv(EnvKongAdminHostAddress) != "" {
		config.HostAddress = strings.TrimRight(os.Getenv(EnvKongAdminHostAddress), "/")
	}

	return config
}

func NewClient(config *Config) *KongAdminClient {
	return &KongAdminClient{
		config: config,
	}
}

func (kongAdminClient *KongAdminClient) Status() *StatusClient {
	return &StatusClient{
		config: kongAdminClient.config,
	}

}

func (kongAdminClient *KongAdminClient) APIs() *APIClient {
	return &APIClient{
		config: kongAdminClient.config,
	}
}

func (kongAdminClient *KongAdminClient) Consumers() *ConsumerClient {
	return &ConsumerClient{
		config: kongAdminClient.config,
	}
}

func (kongAdminClient *KongAdminClient) Plugins() *PluginClient {
	return &PluginClient{
		config: kongAdminClient.config,
	}
}

func (kongAdminClient *KongAdminClient) Certificates() *CertificateClient {
	return &CertificateClient{
		config: kongAdminClient.config,
	}
}

func (kongAdminClient *KongAdminClient) Snis() *SnisClient {
	return &SnisClient{
		config: kongAdminClient.config,
	}
}

func (kongAdminClient *KongAdminClient) Upstreams() *UpstreamClient {
	return &UpstreamClient{
		config: kongAdminClient.config,
	}
}
