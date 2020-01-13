package gokong

import (
	"fmt"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestRoutes_GetById(t *testing.T) {
	serviceRequest := &ServiceRequest{
		Name:     String("service-name" + uuid.NewV4().String()),
		Protocol: String("http"),
		Host:     String("foo.com"),
	}

	client := NewClient(NewDefaultConfig())

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
		Tags:         []*string{String("my-tag")},
	}

	createdRoute, err := client.Routes().Create(routeRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdRoute)

	result, err := client.Routes().GetById(*createdRoute.Id)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, createdRoute, result)

	client.Routes().DeleteById(*createdRoute.Id)
	client.Services().DeleteServiceById(*createdService.Id)

	route, err := client.Routes().GetById(*createdRoute.Id)
	assert.Nil(t, route)
	assert.Nil(t, err)
}

func TestRoutes_CreateWithSourcesAndDestinations(t *testing.T) {
	serviceRequest := &ServiceRequest{
		Name:     String("service-name" + uuid.NewV4().String()),
		Protocol: String("http"),
		Host:     String("foo.com"),
	}

	client := NewClient(NewDefaultConfig())

	createdService, err := client.Services().Create(serviceRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdService)

	routeRequest := &RouteRequest{
		Protocols:    StringSlice([]string{"tls"}),
		StripPath:    Bool(true),
		PreserveHost: Bool(true),
		Snis:         StringSlice([]string{"example.com"}),
		Sources:      IpPortSliceSlice([]IpPort{{Ip: String("192.168.1.1"), Port: Int(80)}, {Ip: String("192.168.1.2"), Port: Int(81)}}),
		Destinations: IpPortSliceSlice([]IpPort{{Ip: String("172.10.1.1"), Port: Int(83)}, {Ip: String("172.10.1.2"), Port: nil}}),
		Service:      ToId(*createdService.Id),
	}

	createdRoute, err := client.Routes().Create(routeRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdRoute)

	result, err := client.Routes().GetById(*createdRoute.Id)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, createdRoute, result)

	client.Routes().DeleteById(*createdRoute.Id)
	client.Services().DeleteServiceById(*createdService.Id)

	route, err := client.Routes().GetById(*createdRoute.Id)
	assert.Nil(t, route)
	assert.Nil(t, err)
}

func TestRoutes_List(t *testing.T) {
	serviceRequest := &ServiceRequest{
		Name:     String("service-name" + uuid.NewV4().String()),
		Protocol: String("http"),
		Host:     String("foo.com"),
	}

	client := NewClient(NewDefaultConfig())
	createdRoutes := &Routes{}
	createdService, err := client.Services().Create(serviceRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdService)

	routeRequest := &RouteRequest{
		Protocols:    StringSlice([]string{"http"}),
		Methods:      StringSlice([]string{"GET"}),
		Hosts:        StringSlice([]string{"foo.com"}),
		StripPath:    Bool(true),
		PreserveHost: Bool(true),
		Service:      ToId(*createdService.Id),
	}

	for i := 0; i < 5; i++ {
		routeRequest.Paths = StringSlice([]string{fmt.Sprintf("/bar-%s", uuid.NewV4().String())})
		createdRoute, err := client.Routes().Create(routeRequest)

		assert.Nil(t, err)
		assert.NotNil(t, createdService)

		createdRoutes.Data = append(createdRoutes.Data, createdRoute)
	}

	result, err := client.Routes().List(&RouteQueryString{})

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Subset(t, createdRoutes.Data, result)

	for _, route := range createdRoutes.Data {
		err := client.Routes().DeleteById(*route.Id)
		assert.Nil(t, err)

		route, err := client.Routes().GetById(*route.Id)
		assert.Nil(t, route)
		assert.Nil(t, err)
	}

	err = client.Services().DeleteServiceById(*createdService.Id)
	assert.Nil(t, err)

}

func TestRoutes_GetRoutesFromServiceId(t *testing.T) {
	serviceRequest := &ServiceRequest{
		Name:     String("service-name" + uuid.NewV4().String()),
		Protocol: String("http"),
		Host:     String("foo.com"),
	}

	client := NewClient(NewDefaultConfig())

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
	assert.NotNil(t, createdRoute)

	result, err := client.Routes().GetRoutesFromServiceId(*createdService.Id)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, result[0], createdRoute)

	err = client.Routes().DeleteById(*createdRoute.Id)
	assert.Nil(t, err)

	route, err := client.Routes().GetById(*createdRoute.Id)
	assert.Nil(t, route)
	assert.Nil(t, err)

	err = client.Services().DeleteServiceById(*createdService.Id)
	assert.Nil(t, err)

}

func TestRoutes_UpdateRouteById(t *testing.T) {
	serviceRequest := &ServiceRequest{
		Name:     String("service-name" + uuid.NewV4().String()),
		Protocol: String("http"),
		Host:     String("foo.com"),
	}

	client := NewClient(NewDefaultConfig())

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
	assert.NotNil(t, createdRoute)

	routeRequest.Paths = StringSlice([]string{"/qux"})
	updatedRoute, err := client.Routes().UpdateById(*createdRoute.Id, routeRequest)
	result, err := client.Routes().GetById(*createdRoute.Id)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, updatedRoute, result)

	err = client.Routes().DeleteById(*createdRoute.Id)
	assert.Nil(t, err)

	err = client.Services().DeleteServiceById(*createdService.Id)
	assert.Nil(t, err)
}

func TestRoutes_UpdateRouteMethodsToEmptyArray(t *testing.T) {
	serviceRequest := &ServiceRequest{
		Name:     String("service-name" + uuid.NewV4().String()),
		Protocol: String("http"),
		Host:     String("foo.com"),
	}

	client := NewClient(NewDefaultConfig())
	createdService, err := client.Services().Create(serviceRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdService)

	routeRequest := &RouteRequest{
		Protocols:    StringSlice([]string{"http"}),
		Methods:      StringSlice([]string{"GET"}),
		Hosts:        StringSlice([]string{"foo.com"}),
		Paths:        StringSlice([]string{"/foo"}),
		StripPath:    Bool(true),
		PreserveHost: Bool(true),
		Service:      ToId(*createdService.Id),
	}

	createdRoute, err := client.Routes().Create(routeRequest)

	assert.Nil(t, err)

	routeRequest.Methods = StringSlice([]string{})

	updatedRoute, err := client.Routes().UpdateById(*createdRoute.Id, routeRequest)

	assert.Nil(t, err)
	assert.Equal(t, routeRequest.Protocols, updatedRoute.Protocols)
	assert.Equal(t, StringSlice([]string{}), updatedRoute.Methods)
	assert.Equal(t, routeRequest.Hosts, updatedRoute.Hosts)
	assert.Equal(t, routeRequest.Paths, updatedRoute.Paths)
	assert.Equal(t, routeRequest.StripPath, updatedRoute.StripPath)
	assert.Equal(t, routeRequest.PreserveHost, updatedRoute.PreserveHost)

	fetchedRoute, err := client.Routes().GetById(*createdRoute.Id)

	assert.Nil(t, err)
	assert.Equal(t, routeRequest.Protocols, fetchedRoute.Protocols)
	assert.Equal(t, StringSlice([]string{}), fetchedRoute.Methods)
	assert.Equal(t, routeRequest.Hosts, fetchedRoute.Hosts)
	assert.Equal(t, routeRequest.Paths, fetchedRoute.Paths)
	assert.Equal(t, routeRequest.StripPath, fetchedRoute.StripPath)
	assert.Equal(t, routeRequest.PreserveHost, fetchedRoute.PreserveHost)

	err = client.Routes().DeleteById(*createdRoute.Id)
	assert.Nil(t, err)

	err = client.Services().DeleteServiceById(*createdService.Id)
	assert.Nil(t, err)

}

func TestRoutes_UpdateRouteHostsToEmptyArray(t *testing.T) {
	serviceRequest := &ServiceRequest{
		Name:     String("service-name" + uuid.NewV4().String()),
		Protocol: String("http"),
		Host:     String("foo.com"),
	}

	client := NewClient(NewDefaultConfig())
	createdService, err := client.Services().Create(serviceRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdService)

	routeRequest := &RouteRequest{
		Protocols:    StringSlice([]string{"http"}),
		Methods:      StringSlice([]string{"GET"}),
		Hosts:        StringSlice([]string{"foo.com"}),
		Paths:        StringSlice([]string{"/foo"}),
		StripPath:    Bool(true),
		PreserveHost: Bool(true),
		Service:      ToId(*createdService.Id),
	}

	createdRoute, err := client.Routes().Create(routeRequest)

	assert.Nil(t, err)

	routeRequest.Hosts = StringSlice([]string{})

	updatedRoute, err := client.Routes().UpdateById(*createdRoute.Id, routeRequest)

	assert.Nil(t, err)
	assert.Equal(t, routeRequest.Protocols, updatedRoute.Protocols)
	assert.Equal(t, routeRequest.Methods, updatedRoute.Methods)
	assert.Equal(t, StringSlice([]string{}), updatedRoute.Hosts)
	assert.Equal(t, routeRequest.Paths, updatedRoute.Paths)
	assert.Equal(t, routeRequest.StripPath, updatedRoute.StripPath)
	assert.Equal(t, routeRequest.PreserveHost, updatedRoute.PreserveHost)

	fetchedRoute, err := client.Routes().GetById(*createdRoute.Id)

	assert.Nil(t, err)
	assert.Equal(t, routeRequest.Protocols, fetchedRoute.Protocols)
	assert.Equal(t, routeRequest.Methods, fetchedRoute.Methods)
	assert.Equal(t, StringSlice([]string{}), fetchedRoute.Hosts)
	assert.Equal(t, routeRequest.Paths, fetchedRoute.Paths)
	assert.Equal(t, routeRequest.StripPath, fetchedRoute.StripPath)
	assert.Equal(t, routeRequest.PreserveHost, fetchedRoute.PreserveHost)

	err = client.Routes().DeleteById(*createdRoute.Id)
	assert.Nil(t, err)

	err = client.Services().DeleteServiceById(*createdService.Id)
	assert.Nil(t, err)

}

func TestRoutes_UpdateRoutePathsToEmptyArray(t *testing.T) {
	serviceRequest := &ServiceRequest{
		Name:     String("service-name" + uuid.NewV4().String()),
		Protocol: String("http"),
		Host:     String("foo.com"),
	}

	client := NewClient(NewDefaultConfig())
	createdService, err := client.Services().Create(serviceRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdService)

	routeRequest := &RouteRequest{
		Protocols:    StringSlice([]string{"http"}),
		Methods:      StringSlice([]string{"GET"}),
		Hosts:        StringSlice([]string{"foo.com"}),
		Paths:        StringSlice([]string{"/foo"}),
		StripPath:    Bool(true),
		PreserveHost: Bool(true),
		Service:      ToId(*createdService.Id),
	}

	createdRoute, err := client.Routes().Create(routeRequest)

	assert.Nil(t, err)

	routeRequest.Paths = StringSlice([]string{})

	updatedRoute, err := client.Routes().UpdateById(*createdRoute.Id, routeRequest)

	assert.Nil(t, err)
	assert.Equal(t, routeRequest.Protocols, updatedRoute.Protocols)
	assert.Equal(t, routeRequest.Methods, updatedRoute.Methods)
	assert.Equal(t, routeRequest.Hosts, updatedRoute.Hosts)
	assert.Equal(t, StringSlice([]string{}), updatedRoute.Paths)
	assert.Equal(t, routeRequest.StripPath, updatedRoute.StripPath)
	assert.Equal(t, routeRequest.PreserveHost, updatedRoute.PreserveHost)

	fetchedRoute, err := client.Routes().GetById(*createdRoute.Id)

	assert.Nil(t, err)
	assert.Equal(t, routeRequest.Protocols, fetchedRoute.Protocols)
	assert.Equal(t, routeRequest.Methods, fetchedRoute.Methods)
	assert.Equal(t, routeRequest.Hosts, fetchedRoute.Hosts)
	assert.Equal(t, StringSlice([]string{}), fetchedRoute.Paths)
	assert.Equal(t, routeRequest.StripPath, fetchedRoute.StripPath)
	assert.Equal(t, routeRequest.PreserveHost, fetchedRoute.PreserveHost)

	err = client.Routes().DeleteById(*createdRoute.Id)
	assert.Nil(t, err)

	err = client.Services().DeleteServiceById(*createdService.Id)
	assert.Nil(t, err)

}

func Test_AllRouteEndpointsShouldReturnErrorWhenRequestUnauthorised(t *testing.T) {

	unauthorisedClient := NewClient(&Config{HostAddress: kong401Server})

	r, err := unauthorisedClient.Routes().GetById(uuid.NewV4().String())
	assert.NotNil(t, err)
	assert.Nil(t, r)

	results, err := unauthorisedClient.Routes().List(&RouteQueryString{})
	assert.NotNil(t, err)
	assert.Nil(t, results)

	results, err = unauthorisedClient.Routes().GetRoutesFromServiceId(uuid.NewV4().String())
	assert.NotNil(t, err)
	assert.Nil(t, results)

	results, err = unauthorisedClient.Routes().GetRoutesFromServiceName("foo")
	assert.NotNil(t, err)
	assert.Nil(t, results)

	err = unauthorisedClient.Routes().DeleteById(uuid.NewV4().String())
	assert.NotNil(t, err)

	createNewRouteRequest := &RouteRequest{
		Protocols:    StringSlice([]string{"http"}),
		Methods:      StringSlice([]string{"POST"}),
		Hosts:        StringSlice([]string{"foo.com"}),
		Paths:        StringSlice([]string{"/bar"}),
		StripPath:    Bool(true),
		PreserveHost: Bool(true),
		Service:      ToId(uuid.NewV4().String()),
	}

	newRoute, err := unauthorisedClient.Routes().Create(createNewRouteRequest)
	assert.Nil(t, newRoute)
	assert.NotNil(t, err)

	updatedRoute, err := unauthorisedClient.Routes().UpdateById(uuid.NewV4().String(), createNewRouteRequest)
	assert.Nil(t, updatedRoute)
	assert.NotNil(t, err)

}
