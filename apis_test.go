package konggo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetById(t *testing.T) {
	result, err := NewKongAdminClient().GetStatus()

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.True(t, result.Database.Reachable)
	assert.True(t, result.Server.ConnectionsAccepted >= 1)
}
