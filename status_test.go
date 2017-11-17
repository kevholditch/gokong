package gokong

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetStatus(t *testing.T) {
	result, err := NewClient(NewDefaultConfig()).Status().Get()

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.True(t, result.Database.Reachable)
	assert.True(t, result.Server.ConnectionsAccepted >= 1)
}
