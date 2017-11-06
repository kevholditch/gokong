package konggo

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"gopkg.in/ory-am/dockertest.v3"
	"log"
	"os"
	"testing"
)

var db *sql.DB

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

func TestMain(m *testing.M) {

	var err error
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}


	postgresContainer := createPostgres(pool)

	code := m.Run()

	if err := pool.Purge(postgresContainer); err != nil {
		log.Fatalf("Could not purge containers: %s", err)
	}

	os.Exit(code)
}

func TestSomething(t *testing.T) {
	db.Query("SELECT 1")

}
