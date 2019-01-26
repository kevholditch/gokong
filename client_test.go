package gokong

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/phayes/freeport"

	"github.com/kevholditch/gokong/containers"
	"github.com/stretchr/testify/assert"
)

const defaultKongVersion = "1.0.0"
const kong401Server = "KONG_401_SERVER"

func Test_Newclient(t *testing.T) {
	result := NewClient(NewDefaultConfig())

	assert.NotNil(t, result)
	assert.Equal(t, os.Getenv(EnvKongAdminHostAddress), result.config.HostAddress)
	assert.Equal(t, os.Getenv(EnvKongAdminUsername), result.config.Username)
	assert.Equal(t, os.Getenv(EnvKongAdminPassword), result.config.Password)
}

func TestMain(m *testing.M) {

	testContext := containers.StartKong(GetEnvVarOrDefault("KONG_VERSION", defaultKongVersion))

	err := os.Setenv(EnvKongAdminHostAddress, testContext.KongHostAddress)
	if err != nil {
		log.Fatalf("Could not set kong host address env variable: %v", err)
	}

	stopSignal := make(chan bool)
	serverPort, _ := freeport.GetFreePort()

	err = os.Setenv(kong401Server, fmt.Sprintf("http://localhost:%d", serverPort))
	if err != nil {
		log.Fatalf("Could not set kong api host address env variable: %v", err)
	}

	go func() { StartServer(serverPort, stopSignal) }()

	code := m.Run()

	stopSignal <- true

	containers.StopKong(testContext)

	os.Exit(code)

}

type serveHttp struct {
	code int
}

func (s *serveHttp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(s.code)
}

func StartServer(port int, ch <-chan bool) {
	address := fmt.Sprintf(":%d", port)
	server := &http.Server{Addr: address, Handler: &serveHttp{code: 401}}
	go func() {
		fmt.Printf("Server started on %s \n", address)
		<-ch
		fmt.Printf("Shutting down \n")
		server.Shutdown(context.Background())
	}()
	fmt.Printf("listening on localhost:%d \n", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
