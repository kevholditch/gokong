package gokong

import (
	"crypto/tls"
	"fmt"
	"net/http"

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

func newGet(config *Config, address string) *gorequest.SuperAgent {
	r := gorequest.New().Get(address)
	return configureRequest(r, config)
}

func newPost(config *Config, address string) *gorequest.SuperAgent {
	r := gorequest.New().Post(address)
	return configureRequest(r, config)
}

func newPatch(config *Config, address string) *gorequest.SuperAgent {
	r := gorequest.New().Patch(address)
	return configureRequest(r, config)
}

func newDelete(config *Config, address string) *gorequest.SuperAgent {
	r := gorequest.New().Delete(address)
	return configureRequest(r, config)
}

func checkResponse(r gorequest.Response, body string, errs []error) error {
	var responseError error
	switch r.StatusCode {
	case http.StatusBadRequest:
		responseError = fmt.Errorf("bad request, message from kong: %s", body)
	case http.StatusForbidden, http.StatusUnauthorized:
		responseError = fmt.Errorf("not authorised, message from kong: %s", body)
	case http.StatusConflict:
		responseError = fmt.Errorf("there was a conflict, message from kong: %v", body)
	case http.StatusNotFound:
		responseError = fmt.Errorf("not found: %v", body)
	default:
		return nil
	}
	return responseError
}
