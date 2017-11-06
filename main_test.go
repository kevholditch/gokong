package konggo

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"gopkg.in/ory-am/dockertest.v3"
	"log"
	"os"
	"testing"
	"strings"
)

func getContainerName(container * dockertest.Resource) string  {
	return strings.TrimPrefix(container.Container.Name, "/")
}

func createPostgres(pool * dockertest.Pool) *dockertest.Resource {
	var db *sql.DB
	resource, err := pool.Run("postgres", "9.6", []string{"POSTGRES_PASSWORD=kong", "POSTGRES_DB=kong"})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	if err = pool.Retry(func() error {
		var err error
		db, err = sql.Open("postgres", fmt.Sprintf("postgres://postgres:kong@localhost:%s/kong?sslmode=disable", resource.GetPort("5432/tcp")))
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}
	return resource
}

func createKong(pool * dockertest.Pool, postgresContainer * dockertest.Resource) *dockertest.Resource {

	postgresContainerName := getContainerName(postgresContainer)


	options := &dockertest.RunOptions{
		Repository: "kong",
		Tag:        "0.11",
		Env:     []string{
			"KONG_DATABASE=postgres",
			fmt.Sprintf("KONG_PG_HOST=%v", postgresContainerName),
			"KONG_PG_USER=postgres",
			"KONG_PG_PASSWORD=kong",
		},
		Links:   []string{fmt.Sprintf("%s:%s", postgresContainerName, postgresContainerName)},
	}

	log.Printf("postgres container: %v", postgresContainerName)

	resource, err := pool.RunWithOptions(options)

	log.Printf("kong container: %v", getContainerName(resource))

	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	return resource

}

func stopAllContainers(pool * dockertest.Pool, containers []*dockertest.Resource) {
	for _, container := range containers {
		if err := pool.Purge(container); err != nil {
			log.Fatalf("Could not purge container %s, error: %s", container.Container.Name, err)
		}
	}
}

func TestMain(m *testing.M) {

	log.SetOutput(os.Stdout)
	var err error
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}


	postgresContainer := createPostgres(pool)
	kongContainer := createKong(pool, postgresContainer)

	code := m.Run()

	stopAllContainers(pool, []*dockertest.Resource{postgresContainer, kongContainer})

	os.Exit(code)
}

func TestSomething(t *testing.T) {

}
