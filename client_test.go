package konggo

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_Newclient(t *testing.T) {
	result := NewClient()

	assert.NotNil(t, result)
	assert.Equal(t, os.Getenv(EnvKongAdminHostAddress), result.hostAddress)
}

func TestMain(m *testing.M) {

	testContext := StartTestContainers()

	code := m.Run()

	StopTestContainers(testContext)

	os.Exit(code)

}
