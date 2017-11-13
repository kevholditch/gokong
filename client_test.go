package konggo

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_GetStatus(t *testing.T) {
	result, err := NewClient().GetStatus()

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.True(t, result.Database.Reachable)
	assert.True(t, result.Server.ConnectionsAccepted >= 1)
}

func TestMain(m *testing.M) {

	testContext := StartTestContainers()

	code := m.Run()

	StopTestContainers(testContext)

	os.Exit(code)

}
