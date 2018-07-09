package gokong

import (
	"fmt"
	"testing"

	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestRouteClient_GetRoute(t *testing.T) {
	serviceRequest := &ServiceRequest{
		Name:     String("service-name" + uuid.NewV4().String()),
		Protocol: String("http"),
		Host:     String("foo.com"),
	}

	client := NewClient(NewDefaultConfig())

	createdService, err := client.Services().AddService(serviceRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdService)

	routeRequest := &RouteRequest{
		Protocols:    StringSlice([]string{"http"}),
		Methods:      StringSlice([]string{"GET"}),
		Hosts:        StringSlice([]string{"foo.com"}),
		Paths:        StringSlice([]string{"/bar"}),
		StripPath:    Bool(true),
		PreserveHost: Bool(true),
		Service:      &RouteServiceObject{Id: *createdService.Id},
	}

	createdRoute, err := client.Routes().AddRoute(routeRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdRoute)

	result, err := client.Routes().GetRoute(*createdRoute.Id)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, createdRoute, result)

	client.Routes().DeleteRoute(*createdRoute.Id)
	client.Services().DeleteServiceById(*createdService.Id)

	route, err := client.Routes().GetRoute(*createdRoute.Id)
	assert.Nil(t, route)
	assert.Nil(t, err)
}

func TestRouteClient_GetRoutes(t *testing.T) {
	serviceRequest := &ServiceRequest{
		Name:     String("service-name" + uuid.NewV4().String()),
		Protocol: String("http"),
		Host:     String("foo.com"),
	}

	client := NewClient(NewDefaultConfig())
	createdRoutes := &Routes{}
	createdService, err := client.Services().AddService(serviceRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdService)

	routeRequest := &RouteRequest{
		Protocols:    StringSlice([]string{"http"}),
		Methods:      StringSlice([]string{"GET"}),
		Hosts:        StringSlice([]string{"foo.com"}),
		StripPath:    Bool(true),
		PreserveHost: Bool(true),
		Service:      &RouteServiceObject{Id: *createdService.Id},
	}

	for i := 0; i < 5; i++ {
		routeRequest.Paths = StringSlice([]string{fmt.Sprintf("/bar-%s", uuid.NewV4().String())})
		createdRoute, err := client.Routes().AddRoute(routeRequest)

		assert.Nil(t, err)
		assert.NotNil(t, createdService)

		createdRoutes.Data = append(createdRoutes.Data, createdRoute)
	}

	result, err := client.Routes().GetRoutes(&RouteQueryString{})

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Subset(t, createdRoutes.Data, result)

	for _, route := range createdRoutes.Data {
		err := client.Routes().DeleteRoute(*route.Id)
		assert.Nil(t, err)

		route, err := client.Routes().GetRoute(*route.Id)
		assert.Nil(t, route)
		assert.Nil(t, err)
	}

	err = client.Services().DeleteServiceById(*createdService.Id)
	assert.Nil(t, err)

}

func TestRouteClient_GetRoutesFromServiceId(t *testing.T) {
	serviceRequest := &ServiceRequest{
		Name:     String("service-name" + uuid.NewV4().String()),
		Protocol: String("http"),
		Host:     String("foo.com"),
	}

	client := NewClient(NewDefaultConfig())

	createdService, err := client.Services().AddService(serviceRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdService)

	routeRequest := &RouteRequest{
		Protocols:    StringSlice([]string{"http"}),
		Methods:      StringSlice([]string{"GET"}),
		Hosts:        StringSlice([]string{"foo.com"}),
		Paths:        StringSlice([]string{"/bar"}),
		StripPath:    Bool(true),
		PreserveHost: Bool(true),
		Service:      &RouteServiceObject{Id: *createdService.Id},
	}

	createdRoute, err := client.Routes().AddRoute(routeRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdRoute)

	result, err := client.Routes().GetRoutesFromServiceId(*createdService.Id)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, result[0], createdRoute)

	err = client.Routes().DeleteRoute(*createdRoute.Id)
	assert.Nil(t, err)

	route, err := client.Routes().GetRoute(*createdRoute.Id)
	assert.Nil(t, route)
	assert.Nil(t, err)

	err = client.Services().DeleteServiceById(*createdService.Id)
	assert.Nil(t, err)

}

func TestRouteClient_UpdateRoute(t *testing.T) {
	serviceRequest := &ServiceRequest{
		Name:     String("service-name" + uuid.NewV4().String()),
		Protocol: String("http"),
		Host:     String("foo.com"),
	}

	client := NewClient(NewDefaultConfig())

	createdService, err := client.Services().AddService(serviceRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdService)

	routeRequest := &RouteRequest{
		Protocols:    StringSlice([]string{"http"}),
		Methods:      StringSlice([]string{"GET"}),
		Hosts:        StringSlice([]string{"foo.com"}),
		Paths:        StringSlice([]string{"/bar"}),
		StripPath:    Bool(true),
		PreserveHost: Bool(true),
		Service:      &RouteServiceObject{Id: *createdService.Id},
	}

	createdRoute, err := client.Routes().AddRoute(routeRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdRoute)

	routeRequest.Paths = StringSlice([]string{"/qux"})
	updatedRoute, err := client.Routes().UpdateRoute(*createdRoute.Id, routeRequest)
	result, err := client.Routes().GetRoute(*createdRoute.Id)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, updatedRoute, result)

	err = client.Routes().DeleteRoute(*createdRoute.Id)
	assert.Nil(t, err)

	err = client.Services().DeleteServiceById(*createdService.Id)
	assert.Nil(t, err)
}
