package konggo

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/ory-am/dockertest.v3"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {

	log.SetOutput(os.Stdout)
	var err error
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	postgres := NewPostgres(pool)
	kong := NewKong(pool, postgres)

	code := m.Run()

	for _, container := range []container{postgres, kong} {
		container.Stop()
	}

	os.Exit(code)
}

func Test_GetStatus(t *testing.T) {
	result, err := NewKongAdminClient().GetStatus()

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.True(t, result.Database.Reachable)
	assert.True(t, result.Server.ConnectionsAccepted >= 1)
}
