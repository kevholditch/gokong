package gokong

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_Newclient(t *testing.T) {
	result := NewClient(NewDefaultConfig())

	assert.NotNil(t, result)
	assert.Equal(t, os.Getenv(EnvKongAdminHostAddress), result.config.HostAddress)
}

func TestMain(m *testing.M) {

	testContext := StartTestContainers()

	code := m.Run()

	StopTestContainers(testContext)

	os.Exit(code)

}
