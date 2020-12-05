package gokong

import (
	"crypto/tls"
	"fmt"

	"github.com/parnurzeal/gorequest"
)

func configureRequest(r *gorequest.SuperAgent, config *Config) *gorequest.SuperAgent {
	r.TLSClientConfig(&tls.Config{InsecureSkipVerify: config.InsecureSkipVerify})
	if config.Username != "" || config.Password != "" {
		r.SetBasicAuth(config.Username, config.Password)
	}

	if config.ApiKey != "" {
		r.Set("apikey", config.ApiKey)
	}

	if config.AdminToken != "" {
		r.Set("kong-admin-token", config.AdminToken)
	}

	return r
}

func buildRequestUri(config *Config, path string) string {
	if config.Workspace == "" {
		return config.HostAddress + path
	}
	return fmt.Sprintf("%s/%s%s", config.HostAddress, config.Workspace, path)
}

func newRawGet(config *Config, address string) *gorequest.SuperAgent {
	r := gorequest.New().Get(address)
	return configureRequest(r, config)
}

func newRawPost(config *Config, address string) *gorequest.SuperAgent {
	r := gorequest.New().Post(address)
	return configureRequest(r, config)
}

func newRawPatch(config *Config, address string) *gorequest.SuperAgent {
	r := gorequest.New().Patch(address)
	return configureRequest(r, config)
}

func newRawDelete(config *Config, address string) *gorequest.SuperAgent {
	r := gorequest.New().Delete(address)
	return configureRequest(r, config)
}

func newGet(config *Config, path string) *gorequest.SuperAgent {
	r := gorequest.New().Get(buildRequestUri(config, path))
	return configureRequest(r, config)
}

func newPost(config *Config, path string) *gorequest.SuperAgent {
	r := gorequest.New().Post(buildRequestUri(config, path))
	return configureRequest(r, config)
}

func newPatch(config *Config, path string) *gorequest.SuperAgent {
	r := gorequest.New().Patch(buildRequestUri(config, path))
	return configureRequest(r, config)
}

func newDelete(config *Config, path string) *gorequest.SuperAgent {
	r := gorequest.New().Delete(buildRequestUri(config, path))
	return configureRequest(r, config)
}
