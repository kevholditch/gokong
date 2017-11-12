package konggo

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/parnurzeal/gorequest"
)

const EnvKongAdminHostAddress = "KONG_ADMIN_ADDR"

type KongAdminClient struct {
	hostAddress string
	client      *gorequest.SuperAgent
}

type Status struct {
	Server   serverStatus   `json:"server"`
	Database databaseStatus `json:"database"`
}

type serverStatus struct {
	TotalRequests       int `json:"total_requests"`
	ConnectionsActive   int `json:"connections_active"`
	ConnectionsAccepted int `json:"connections_accepted"`
	ConnectionsHandled  int `json:"connections_handled"`
	ConnectionsReading  int `json:"connections_reading"`
	ConnectionsWriting  int `json:"connections_writing"`
	ConnectionsWaiting  int `json:"connections_waiting"`
}

type databaseStatus struct {
	Reachable bool `json:"reachable"`
}

func NewKongAdminClient() *KongAdminClient {
	return &KongAdminClient{
		hostAddress: GetEnvOrDefault("KONG_ADMIN_ADDR", "http://localhost:8001"),
		client:      gorequest.New(),
	}
}

func (kongAdminClient *KongAdminClient) GetStatus() (*Status, error) {

	_, body, errs := kongAdminClient.client.Get(kongAdminClient.hostAddress + "/status").End()
	if errs != nil {
		return nil, errors.New(fmt.Sprintf("Could not call kong api client, error: %v", errs))
	}

	status := &Status{}
	err := json.Unmarshal([]byte(body), status)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Could not parse status response, error: %v", err))
	}

	return status, nil

}
