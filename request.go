package gokong

import (
	"crypto/tls"

	"github.com/parnurzeal/gorequest"
)

func NewRequest(adminConfig *Config) *gorequest.SuperAgent {
	request := gorequest.New().TLSClientConfig(&tls.Config{InsecureSkipVerify: adminConfig.InsecureSkipVerify})
	if adminConfig.Username != "" || adminConfig.Password != "" {
		request.SetBasicAuth(adminConfig.Username, adminConfig.Password)
	}

	if adminConfig.ApiKeyHeaderName != "" && adminConfig.ApiKeyHeaderValue != "" {
		request.Set(adminConfig.ApiKeyHeaderName, adminConfig.ApiKeyHeaderValue)
	}

	return request
}
