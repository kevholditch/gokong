package gokong

import (
	"fmt"
	"testing"

	"os"

	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestServiceClient_GetServiceById(t *testing.T) {
	serviceRequest := &ServiceRequest{
		Name:     String(fmt.Sprintf("service-name-%s", uuid.NewV4().String())),
		Protocol: String("http"),
		Host:     String("foo.com"),
		Port:     Int(8080),
	}

	client := NewClient(NewDefaultConfig())

	createdService, err := client.Services().AddService(serviceRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdService)
	assert.EqualValues(t, createdService.Name, serviceRequest.Name)
	assert.EqualValues(t, createdService.Protocol, serviceRequest.Protocol)
	assert.EqualValues(t, createdService.Host, serviceRequest.Host)
	assert.EqualValues(t, createdService.Port, serviceRequest.Port)

	result, err := client.Services().GetServiceById(*createdService.Id)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, createdService, result)

	err = client.Services().DeleteServiceById(*createdService.Id)
	assert.Nil(t, err)
}

func TestServiceClient_GetServiceByIdWithUrl(t *testing.T) {
	serviceRequest := &ServiceRequest{
		Name: String(fmt.Sprintf("service-name-%s", uuid.NewV4().String())),
		Url:  String("http://foo.com:8080"),
	}

	client := NewClient(NewDefaultConfig())

	createdService, err := client.Services().AddService(serviceRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdService)
	assert.EqualValues(t, createdService.Name, serviceRequest.Name)
	assert.EqualValues(t, createdService.Protocol, String("http"))
	assert.EqualValues(t, createdService.Host, String("foo.com"))
	assert.EqualValues(t, createdService.Port, Int(8080))

	result, err := client.Services().GetServiceById(*createdService.Id)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, createdService, result)

	err = client.Services().DeleteServiceById(*createdService.Id)
	assert.Nil(t, err)
}

func TestServiceClient_GetServices(t *testing.T) {
	serviceRequest := &ServiceRequest{
		Protocol: String("http"),
		Host:     String("foo.com"),
	}
	createdServices := &Services{}
	client := NewClient(NewDefaultConfig())

	for i := 0; i < 5; i++ {
		serviceRequest.Name = String(fmt.Sprintf("service-name-%s", uuid.NewV4().String()))
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
		err = client.Services().DeleteServiceById(*service.Id)
		assert.Nil(t, err)
	}
}

func TestServiceClient_UpdateServiceById(t *testing.T) {
	serviceRequest := &ServiceRequest{
		Name:     String(fmt.Sprintf("service-name-%s", uuid.NewV4().String())),
		Protocol: String("http"),
		Host:     String("foo.com"),
	}

	client := NewClient(NewDefaultConfig())

	createdService, err := client.Services().AddService(serviceRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdService)

	serviceRequest.Host = String("bar.io")
	updatedService, err := client.Services().UpdateServiceById(*createdService.Id, serviceRequest)
	result, err := client.Services().GetServiceById(*createdService.Id)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, updatedService, result)

	err = client.Services().DeleteServiceById(*createdService.Id)
	assert.Nil(t, err)
}

func Test_ServicesGetNonExistentById(t *testing.T) {

	service, err := NewClient(NewDefaultConfig()).Services().GetServiceById(uuid.NewV4().String())

	assert.Nil(t, service)
	assert.Nil(t, err)
}

func Test_ServicesGetNonExistentByName(t *testing.T) {

	service, err := NewClient(NewDefaultConfig()).Services().GetServiceByName(uuid.NewV4().String())

	assert.Nil(t, service)
	assert.Nil(t, err)
}

func Test_AllServiceEndpointsShouldReturnErrorWhenRequestUnauthorised(t *testing.T) {

	apiRequest := &ApiRequest{
		Name:        String("admin-api"),
		Uris:        StringSlice([]string{"/admin-api"}),
		UpstreamUrl: String("http://localhost:8001"),
	}

	client := NewClient(NewDefaultConfig())
	createdApi, err := client.Apis().Create(apiRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdApi)

	consumerRequest := &ConsumerRequest{
		Username: "username-" + uuid.NewV4().String(),
		CustomId: "test-" + uuid.NewV4().String(),
	}

	createdConsumer, err := client.Consumers().Create(consumerRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdConsumer)

	pluginRequest := &PluginRequest{
		Name:  "key-auth",
		ApiId: *createdApi.Id,
		Config: map[string]interface{}{
			"hide_credentials": true,
		},
	}

	createdPlugin, err := client.Plugins().Create(pluginRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdPlugin)

	_, err = client.Consumers().CreatePluginConfig(createdConsumer.Id, "key-auth", "")
	assert.Nil(t, err)

	serviceRequest := &ServiceRequest{
		Name:     String("service-name" + uuid.NewV4().String()),
		Protocol: String("http"),
		Host:     String("foo.com"),
	}

	createdService, err := client.Services().AddService(serviceRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdService)

	kongApiAddress := os.Getenv(EnvKongApiHostAddress) + "/admin-api"
	unauthorisedClient := NewClient(&Config{HostAddress: kongApiAddress})

	s, err := unauthorisedClient.Services().GetServiceByName(*createdService.Name)
	assert.NotNil(t, err)
	assert.Nil(t, s)

	s, err = unauthorisedClient.Services().GetServiceById(*createdService.Id)
	assert.NotNil(t, err)
	assert.Nil(t, s)

	results, err := unauthorisedClient.Services().GetServices(&ServiceQueryString{})
	assert.NotNil(t, err)
	assert.Nil(t, results)

	err = unauthorisedClient.Services().DeleteServiceById(*createdService.Id)
	assert.NotNil(t, err)

	err = unauthorisedClient.Services().DeleteServiceByName(*createdService.Name)
	assert.NotNil(t, err)

	createServiceRequest := &ServiceRequest{
		Name:     String("service-name" + uuid.NewV4().String()),
		Protocol: String("http"),
		Host:     String("foo.com"),
	}

	newService, err := unauthorisedClient.Services().AddService(createServiceRequest)
	assert.Nil(t, newService)
	assert.NotNil(t, err)

	updatedService, err := unauthorisedClient.Services().UpdateServiceById(*createdService.Id, createServiceRequest)
	assert.Nil(t, updatedService)
	assert.NotNil(t, err)

	updatedService, err = unauthorisedClient.Services().UpdateServiceByName(*createdService.Name, createServiceRequest)
	assert.Nil(t, updatedService)
	assert.NotNil(t, err)

	err = client.Plugins().DeleteById(createdPlugin.Id)
	assert.Nil(t, err)

	err = client.Apis().DeleteById(*createdApi.Id)
	assert.Nil(t, err)

	err = client.Services().DeleteServiceById(*createdService.Id)
	assert.Nil(t, err)

}
