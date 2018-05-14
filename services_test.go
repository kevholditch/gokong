package gokong

import (
	"fmt"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestServiceClient_GetServiceById(t *testing.T) {
	serviceRequest := &ServiceRequest{
		Name:     fmt.Sprintf("service-name-%s", uuid.NewV4().String()),
		Protocol: "http",
		Host:     "foo.com",
	}

	client := NewClient(NewDefaultConfig())

	createdService, err := client.Services().AddService(serviceRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdService)

	result, err := client.Services().GetServiceById(createdService.Id)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, createdService, result)

	err = client.Services().DeleteServiceById(createdService.Id)
	assert.Nil(t, err)
}

func TestServiceClient_GetServices(t *testing.T) {
	serviceRequest := &ServiceRequest{
		Protocol: "http",
		Host:     "foo.com",
	}
	createdServices := &Services{}
	client := NewClient(NewDefaultConfig())

	for i := 0; i < 5; i++ {
		serviceRequest.Name = fmt.Sprintf("service-name-%s", uuid.NewV4().String())
		createdService, err := client.Services().AddService(serviceRequest)

		assert.Nil(t, err)
		assert.NotNil(t, createdService)

		createdServices.Data = append(createdServices.Data, createdService)
	}

	result, err := client.Services().GetServices(&ServiceQueryString{})

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Subset(t, createdServices.Data, result)

	for _, service := range createdServices.Data {
		err = client.Services().DeleteServiceById(service.Id)
		assert.Nil(t, err)
	}
}

func TestServiceClient_UpdateServiceById(t *testing.T) {
	serviceRequest := &ServiceRequest{
		Name:     fmt.Sprintf("service-name-%s", uuid.NewV4().String()),
		Protocol: "http",
		Host:     "foo.com",
	}

	client := NewClient(NewDefaultConfig())

	createdService, err := client.Services().AddService(serviceRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdService)

	serviceRequest.Host = "bar.io"
	updatedService, err := client.Services().UpdateServiceById(createdService.Id, serviceRequest)
	result, err := client.Services().GetServiceById(createdService.Id)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, updatedService, result)

	err = client.Services().DeleteServiceById(createdService.Id)
	assert.Nil(t, err)
}
