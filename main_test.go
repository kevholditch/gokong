package go_sync

import (
	"fmt"
	"gopkg.in/ory-am/dockertest.v3"
	"log"
	"net/http"
	"os"
	"testing"
)

func TestAcc_sync_certificate_is_successful(t *testing.T) {
	testPreCheck(t)

}

func TestMain(m *testing.M) {
	log.SetOutput(os.Stdout)
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	options := &dockertest.RunOptions{
		Repository: "consul",
		Tag:        "0.8.5",
	}

	consulResource, err := pool.RunWithOptions(options)
	if err != nil {
		log.Fatalf("Could not start consul: %s", err)
	}

	if err := pool.Retry(func() error {
		var err error
		resp, err := http.Get(fmt.Sprintf("http://localhost:%v/v1/catalog/services", consulResource.GetPort("8500/tcp")))
		if err != nil || resp.StatusCode >= 400 {
			return err
		}

		log.Printf("Consul up: %v", resp.StatusCode)
		return nil
	}); err != nil {
		log.Fatalf("Could not connect to consul: %s", err)
	}


	vaultResource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "vault",
		Tag:        "0.8.3",
		Links: []string{consulResource.Container.Name},
		Env: [] string {"VAULT_DEV_ROOT_TOKEN_ID=C4B4BDE9-D318-421C-BF65-9AE4C3DA169B"},
		Cmd: [] string {"-cap-add=IPC_LOCK"},
	})

	if err != nil {
		log.Fatal(err)
	}

	if err := pool.Retry(func() error {
		var err error
		resp, err := http.Get(fmt.Sprintf("http://localhost:%v/v1/sys/health", vaultResource.GetPort("8200/tcp")))
		if err != nil || resp.StatusCode >= 400 {
			return err
		}

		log.Printf("Vault up: %v", resp.StatusCode)
		return nil
	}); err != nil {
		log.Fatalf("Could not connect to vault: %s", err)
	}


	code := m.Run()

	// You can't defer this because os.Exit doesn't care for defer
	if err := pool.Purge(consulResource); err != nil {
		log.Fatalf("Could not purge consulResource: %s", err)
	}

	os.Exit(code)
}
