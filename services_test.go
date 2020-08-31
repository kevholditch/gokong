package gokong

import (
	"fmt"
	"testing"

	uuid "github.com/satori/go.uuid"
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

	createdService, err := client.Services().Create(serviceRequest)

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

func TestServiceClient_GetServices(t *testing.T) {
	serviceRequest := &ServiceRequest{
		Protocol: String("http"),
		Host:     String("foo.com"),
	}
	createdServices := &Services{}
	client := NewClient(NewDefaultConfig())

	for i := 0; i < 5; i++ {
		serviceRequest.Name = String(fmt.Sprintf("service-name-%s", uuid.NewV4().String()))
		createdService, err := client.Services().Create(serviceRequest)

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

func TestServiceClient_DeleteServiceByIdWithRouteError(t *testing.T) {
	serviceRequest := &ServiceRequest{
		Protocol: String("http"),
		Host:     String("foo.com"),
	}
	client := NewClient(NewDefaultConfig())

	serviceRequest.Name = String(fmt.Sprintf("service-name-%s", uuid.NewV4().String()))
	createdService, err := client.Services().Create(serviceRequest)
	assert.Nil(t, err)
	assert.NotNil(t, createdService)

	routeRequest := &RouteRequest{
		Protocols:    StringSlice([]string{"http"}),
		Methods:      StringSlice([]string{"GET"}),
		Hosts:        StringSlice([]string{"foo.com"}),
		Paths:        StringSlice([]string{"/bar"}),
		StripPath:    Bool(true),
		PreserveHost: Bool(true),
		Service:      ToId(*createdService.Id),
	}
	createdRoute, err := client.Routes().Create(routeRequest)
	assert.Nil(t, err)

	err = client.Services().DeleteServiceById(*createdService.Id)
	expectedError := "bad request, message from kong: {\"message\":\"an existing 'routes' entity references this 'services' entity\",\"name\":\"foreign key violation\",\"fields\":{\"@referenced_by\":\"routes\"},\"code\":4}"
	assert.Equal(t, expectedError, err.Error())
	assert.NotNil(t, createdRoute)

	err = client.Routes().DeleteById(*createdRoute.Id)
	assert.Nil(t, err)

	err = client.Services().DeleteServiceById(*createdService.Id)
	assert.Nil(t, err)
}

func TestServiceClient_UpdateServiceById(t *testing.T) {
	serviceRequest := &ServiceRequest{
		Name:     String(fmt.Sprintf("service-name-%s", uuid.NewV4().String())),
		Protocol: String("http"),
		Host:     String("foo.com"),
	}

	client := NewClient(NewDefaultConfig())

	createdService, err := client.Services().Create(serviceRequest)

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

	unauthorisedClient := NewClient(&Config{HostAddress: kong401Server})

	s, err := unauthorisedClient.Services().GetServiceByName("foo")
	assert.NotNil(t, err)
	assert.Nil(t, s)

	s, err = unauthorisedClient.Services().GetServiceById(uuid.NewV4().String())
	assert.NotNil(t, err)
	assert.Nil(t, s)

	results, err := unauthorisedClient.Services().GetServices(&ServiceQueryString{})
	assert.NotNil(t, err)
	assert.Nil(t, results)

	err = unauthorisedClient.Services().DeleteServiceById(uuid.NewV4().String())
	assert.NotNil(t, err)

	err = unauthorisedClient.Services().DeleteServiceByName("foo")
	assert.NotNil(t, err)

	createServiceRequest := &ServiceRequest{
		Name:     String("service-name" + uuid.NewV4().String()),
		Protocol: String("http"),
		Host:     String("foo.com"),
	}

	newService, err := unauthorisedClient.Services().Create(createServiceRequest)
	assert.Nil(t, newService)
	assert.NotNil(t, err)

	updatedService, err := unauthorisedClient.Services().UpdateServiceById(uuid.NewV4().String(), createServiceRequest)
	assert.Nil(t, updatedService)
	assert.NotNil(t, err)

	updatedService, err = unauthorisedClient.Services().UpdateServiceByName("foo", createServiceRequest)
	assert.Nil(t, updatedService)
	assert.NotNil(t, err)

}

func Test_CreateShouldRerturnErrorWhenBadRequest(t *testing.T) {
	serviceRequest := &ServiceRequest{
		Name:     String("service-name" + uuid.NewV4().String()),
		Protocol: String("http"),
	}

	client := NewClient(NewDefaultConfig())
	createdService, err := client.Services().Create(serviceRequest)

	errorMessage := `bad request, message from kong: {"message":"schema violation (host: required field missing)","name":"schema violation","fields":{"host":"required field missing"},"code":2}`

	assert.Nil(t, createdService)
	assert.Equal(t, err.Error(), errorMessage)
}

func Test_UpdateShouldRerturnErrorWhenBadRequest(t *testing.T) {
	serviceRequest := &ServiceRequest{
		Name:     String("service-name" + uuid.NewV4().String()),
		Host:     String("foo.com"),
		Protocol: String("http"),
	}

	client := NewClient(NewDefaultConfig())
	createdService, err := client.Services().Create(serviceRequest)

	assert.Nil(t, err)

	serviceRequestUpdate := &ServiceRequest{
		Name:     createdService.Name,
		Host:     createdService.Host,
		Protocol: String("foo"),
	}

	errorMessage := `bad request, message from kong: {"message":"schema violation (protocol: expected one of: http, https, tcp, tls)","name":"schema violation","fields":{"protocol":"expected one of: http, https, tcp, tls"},"code":2}`

	updatedService, err := client.Services().UpdateServiceById(*createdService.Id, serviceRequestUpdate)

	assert.Nil(t, updatedService)
	assert.Equal(t, err.Error(), errorMessage)

	err = client.Services().DeleteServiceById(*createdService.Id)
	assert.Nil(t, err)
}
