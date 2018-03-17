package gokong

import "reflect"
import (
	"github.com/parnurzeal/gorequest"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRequest(t *testing.T) {
	config := &Config{
		HostAddress: "http://localhost:8001",
		Username:    "",
		Password:    "",
	}

	current := reflect.TypeOf(NewRequest(config))
	expected := reflect.TypeOf(gorequest.New())
	assert.IsType(t, expected, current)
}

func TestNewRequestBasicAuth(t *testing.T) {
	config := &Config{
		HostAddress: "http://localhost:8001",
		Username:    "AnTestUser",
		Password:    "AnyPassword",
	}

	current := NewRequest(config)
	assert.EqualValues(t, current.BasicAuth.Username, config.Username)
	assert.EqualValues(t, current.BasicAuth.Password, config.Password)
}
