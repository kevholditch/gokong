package gokong

import (
	"net/url"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/google/go-querystring/query"
)

const EnvKongAdminHostAddress = "KONG_ADMIN_ADDR"
const EnvKongAdminUsername = "KONG_ADMIN_USERNAME"
const EnvKongAdminPassword = "KONG_ADMIN_PASSWORD"
const EnvKongTLSSkipVerify = "TLS_SKIP_VERIFY"
const EnvKongApiKey = "KONG_API_KEY"
const EnvKongAdminToken = "KONG_ADMIN_TOKEN"

type KongAdminClient interface {
	Status() StatusClient
	Consumers() ConsumerClient
	Plugins() PluginClient
	Certificates() CertificateClient
	Snis() SnisClient
	Upstreams() UpstreamClient
	Routes() RouteClient
	Services() ServiceClient
	Targets() TargetClient
}

type kongAdminClient struct {
	config *Config
}

type Config struct {
	HostAddress        string
	Username           string
	Password           string
	InsecureSkipVerify bool
	ApiKey             string
	AdminToken         string
}

func addQueryString(currentUrl string, filter interface{}) (string, error) {
	v := reflect.ValueOf(filter)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return currentUrl, nil
	}

	u, err := url.Parse(currentUrl)
	if err != nil {
		return currentUrl, err
	}

	qs, err := query.Values(filter)
	if err != nil {
		return currentUrl, err
	}

	u.RawQuery = qs.Encode()
	return u.String(), nil
}

func NewDefaultConfig() *Config {
	config := &Config{
		HostAddress:        "http://localhost:8001",
		Username:           "",
		Password:           "",
		InsecureSkipVerify: false,
	}

	if os.Getenv(EnvKongAdminHostAddress) != "" {
		config.HostAddress = strings.TrimRight(os.Getenv(EnvKongAdminHostAddress), "/")
	}
	if os.Getenv(EnvKongAdminHostAddress) != "" {
		config.Username = os.Getenv(EnvKongAdminUsername)
	}
	if os.Getenv(EnvKongAdminPassword) != "" {
		config.Password = os.Getenv(EnvKongAdminPassword)
	}
	if os.Getenv(EnvKongTLSSkipVerify) != "" {
		skip, err := strconv.ParseBool(os.Getenv(EnvKongTLSSkipVerify))
		if err == nil {
			config.InsecureSkipVerify = skip
		}
	}
	if os.Getenv(EnvKongApiKey) != "" {
		config.ApiKey = os.Getenv(EnvKongApiKey)
	}
	if os.Getenv(EnvKongAdminToken) != "" {
		config.AdminToken = os.Getenv(EnvKongAdminToken)
	}

	return config
}

func NewClient(config *Config) *kongAdminClient {
	return &kongAdminClient{
		config: config,
	}
}

func (kongAdminClient *kongAdminClient) Status() StatusClient {
	return &statusClient{
		config: kongAdminClient.config,
	}
}

func (kongAdminClient *kongAdminClient) Consumers() ConsumerClient {
	return &consumerClient{
		config: kongAdminClient.config,
	}
}

func (kongAdminClient *kongAdminClient) Plugins() PluginClient {
	return &pluginClient{
		config: kongAdminClient.config,
	}
}

func (kongAdminClient *kongAdminClient) Certificates() CertificateClient {
	return &certificateClient{
		config: kongAdminClient.config,
	}
}

func (kongAdminClient *kongAdminClient) Snis() SnisClient {
	return &snisClient{
		config: kongAdminClient.config,
	}
}

func (kongAdminClient *kongAdminClient) Upstreams() UpstreamClient {
	return &upstreamClient{
		config: kongAdminClient.config,
	}
}

func (kongAdminClient *kongAdminClient) Routes() RouteClient {
	return &routeClient{
		config: kongAdminClient.config,
	}
}

func (kongAdminClient *kongAdminClient) Services() ServiceClient {
	return &serviceClient{
		config: kongAdminClient.config,
	}
}

func (kongAdminClient *kongAdminClient) Targets() TargetClient {
	return &targetClient{
		config: kongAdminClient.config,
	}
}
