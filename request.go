package gokong

import (
	"crypto/tls"
	"time"

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
	r := newKongAgent(config).Get(address)
	return configureRequest(r, config)
}

func newPost(config *Config, address string) *gorequest.SuperAgent {
	r := newKongAgent(config).Post(address)
	return configureRequest(r, config)
}

func newPatch(config *Config, address string) *gorequest.SuperAgent {
	r := newKongAgent(config).Patch(address)
	return configureRequest(r, config)
}

func newDelete(config *Config, address string) *gorequest.SuperAgent {
	r := newKongAgent(config).Delete(address)
	return configureRequest(r, config)
}

type KongAgent struct {
	gorequest.SuperAgent
	maxRetries    int
	retryInterval time.Duration
}

// Added this function because the retry settings in gorquest needs exact status to retry where this really wants to retry
// on everything except for certain conditions.
func (ka *KongAgent) End(callback ...func(response gorequest.Response, body string, errs []error)) (gorequest.Response, string, []error) {
	var errors []error
	for count := 0; count < ka.maxRetries; count++ {
		r, body, errs := ka.SuperAgent.End()

		if errs != nil || (r.StatusCode > 400 && r.StatusCode < 500) || (r.StatusCode >= 200 || r.StatusCode < 300) {
			return r, body, errs
		}

		errors = append(errors, errs...)
		time.Sleep(ka.retryInterval)
	}
	return nil, "", errors
}

func newKongAgent(config *Config) *KongAgent {
	maxRetries := config.MaxRetries
	if maxRetries <= 0 {
		maxRetries = 1
	}
	retryInterval := config.RetryInterval
	if retryInterval <= 0 {
		retryInterval = 5
	}
	ka := KongAgent{*gorequest.New(), maxRetries, time.Duration(retryInterval) * time.Second}
	return &ka
}
