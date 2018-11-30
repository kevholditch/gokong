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
const EnvKongApiHostAddress = "KONG_API_ADDR"
const EnvKongAdminUsername = "KONG_ADMIN_USERNAME"
const EnvKongAdminPassword = "KONG_ADMIN_PASSWORD"
const EnvKongTLSSkipVerify = "TLS_SKIP_VERIFY"
const EnvKongApiKey = "KONG_API_KEY"
const EnvKongAdminToken = "KONG_ADMIN_TOKEN"
const EnvKongMaxRetries = "KONG_MAX_RETRIES"
const EnvKongRetryInterval = "KONG_RETRY_INTERVAL"

type KongAdminClient struct {
	config *Config
}

type Config struct {
	HostAddress        string
	Username           string
	Password           string
	InsecureSkipVerify bool
	ApiKey             string
	AdminToken         string
	MaxRetries         int
	RetryInterval      int
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
		MaxRetries:         10,
		RetryInterval:      5,
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

	if val, exists := os.LookupEnv(EnvKongMaxRetries); exists {
		if maxRetries, err := strconv.Atoi(val); err == nil {
			config.MaxRetries = maxRetries
		}
	}

	if val, exists := os.LookupEnv(EnvKongRetryInterval); exists {
		if retryInterval, err := strconv.Atoi(val); err == nil {
			config.RetryInterval = retryInterval
		}
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

func (kongAdminClient *KongAdminClient) Apis() *ApiClient {
	return &ApiClient{
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

func (kongAdminClient *KongAdminClient) Routes() *RouteClient {
	return &RouteClient{
		config: kongAdminClient.config,
	}
}

func (kongAdminClient *KongAdminClient) Services() *ServiceClient {
	return &ServiceClient{
		config: kongAdminClient.config,
	}
}
