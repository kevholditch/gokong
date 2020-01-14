// +build all community

package gokong

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetStatus(t *testing.T) {
	result, err := NewClient(NewDefaultConfig()).Status().Get()

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.True(t, result.Database.Reachable)
	assert.True(t, result.Server.ConnectionsAccepted >= 1)
}
