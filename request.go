package gokong

import (
	"crypto/tls"
	"fmt"
	"log"
	"time"

	"github.com/parnurzeal/gorequest"
)

func configureRequest(r *KongAgent, config *Config) *KongAgent {
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

func newGet(config *Config, address string) *KongAgent {
	r := newKongAgent(config).Get(address)
	return configureRequest(r, config)
}

func newPost(config *Config, address string) *KongAgent {
	r := newKongAgent(config).Post(address)
	return configureRequest(r, config)
}

func newPatch(config *Config, address string) *KongAgent {
	r := newKongAgent(config).Patch(address)
	return configureRequest(r, config)
}

func newDelete(config *Config, address string) *KongAgent {
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

		if errs != nil || (r.StatusCode >= 400 && r.StatusCode < 500) || (r.StatusCode >= 200 && r.StatusCode < 300) {
			log.Printf("Success on %v: %v=%v", count, r.StatusCode, r)
			return r, body, errs
		}

		log.Printf("Attempt %v unsuccessful: %v\n", count, r)
		errors = append(errors, fmt.Errorf("retry attempt %v", count))
		errors = append(errors, errs...)
		if count < (ka.maxRetries - 1) { // Don't sleep if it is leaving the loop
			time.Sleep(ka.retryInterval)
		}
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

func (ka *KongAgent) Get(targetUrl string) *KongAgent {
	ka.SuperAgent.Get(targetUrl)
	return ka
}

func (ka *KongAgent) Post(targetUrl string) *KongAgent {
	ka.SuperAgent.Post(targetUrl)
	return ka
}

func (ka *KongAgent) Head(targetUrl string) *KongAgent {
	ka.SuperAgent.Head(targetUrl)
	return ka
}

func (ka *KongAgent) Put(targetUrl string) *KongAgent {
	ka.SuperAgent.Put(targetUrl)
	return ka
}

func (ka *KongAgent) Delete(targetUrl string) *KongAgent {
	ka.SuperAgent.Delete(targetUrl)
	return ka
}

func (ka *KongAgent) Patch(targetUrl string) *KongAgent {
	ka.SuperAgent.Patch(targetUrl)
	return ka
}

func (ka *KongAgent) Options(targetUrl string) *KongAgent {
	ka.SuperAgent.Options(targetUrl)
	return ka
}
