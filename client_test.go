package gokong

import (
	"github.com/kevholditch/gokong/containers"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

const defaultKongVersion = "0.11"

func Test_Newclient(t *testing.T) {
	result := NewClient(NewDefaultConfig())

	assert.NotNil(t, result)
	assert.Equal(t, os.Getenv(EnvKongAdminHostAddress), result.config.HostAddress)
}

func TestMain(m *testing.M) {

	testContext := containers.StartKong(GetEnvVarOrDefault("KONG_VERSION", defaultKongVersion))

	err := os.Setenv(EnvKongAdminHostAddress, testContext.KongHostAddress)
	if err != nil {
		log.Fatalf("Could not set kong host address env variable: %v", err)
	}

	code := m.Run()

	containers.StopKong(testContext)

	os.Exit(code)

}
