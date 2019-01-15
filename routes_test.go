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
		Service:      ToId(*createdService.Id),
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
		Service:      ToId(*createdService.Id),
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
		Service:      ToId(*createdService.Id),
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
		Service:      ToId(*createdService.Id),
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

func TestRouteClient_UpdateRouteMethodsToEmptyArray(t *testing.T) {
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
		Paths:        StringSlice([]string{"/foo"}),
		StripPath:    Bool(true),
		PreserveHost: Bool(true),
		Service:      ToId(*createdService.Id),
	}

	createdRoute, err := client.Routes().AddRoute(routeRequest)

	assert.Nil(t, err)

	routeRequest.Methods = StringSlice([]string{})

	updatedRoute, err := client.Routes().UpdateRoute(*createdRoute.Id, routeRequest)

	assert.Nil(t, err)
	assert.Equal(t, routeRequest.Protocols, updatedRoute.Protocols)
	assert.Equal(t, StringSlice([]string{}), updatedRoute.Methods)
	assert.Equal(t, routeRequest.Hosts, updatedRoute.Hosts)
	assert.Equal(t, routeRequest.Paths, updatedRoute.Paths)
	assert.Equal(t, routeRequest.StripPath, updatedRoute.StripPath)
	assert.Equal(t, routeRequest.PreserveHost, updatedRoute.PreserveHost)

	fetchedRoute, err := client.Routes().GetRoute(*createdRoute.Id)

	assert.Nil(t, err)
	assert.Equal(t, routeRequest.Protocols, fetchedRoute.Protocols)
	assert.Equal(t, StringSlice([]string{}), fetchedRoute.Methods)
	assert.Equal(t, routeRequest.Hosts, fetchedRoute.Hosts)
	assert.Equal(t, routeRequest.Paths, fetchedRoute.Paths)
	assert.Equal(t, routeRequest.StripPath, fetchedRoute.StripPath)
	assert.Equal(t, routeRequest.PreserveHost, fetchedRoute.PreserveHost)

	err = client.Routes().DeleteRoute(*createdRoute.Id)
	assert.Nil(t, err)

	err = client.Services().DeleteServiceById(*createdService.Id)
	assert.Nil(t, err)

}

func TestRouteClient_UpdateRouteHostsToEmptyArray(t *testing.T) {
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
		Paths:        StringSlice([]string{"/foo"}),
		StripPath:    Bool(true),
		PreserveHost: Bool(true),
		Service:      ToId(*createdService.Id),
	}

	createdRoute, err := client.Routes().AddRoute(routeRequest)

	assert.Nil(t, err)

	routeRequest.Hosts = StringSlice([]string{})

	updatedRoute, err := client.Routes().UpdateRoute(*createdRoute.Id, routeRequest)

	assert.Nil(t, err)
	assert.Equal(t, routeRequest.Protocols, updatedRoute.Protocols)
	assert.Equal(t, routeRequest.Methods, updatedRoute.Methods)
	assert.Equal(t, StringSlice([]string{}), updatedRoute.Hosts)
	assert.Equal(t, routeRequest.Paths, updatedRoute.Paths)
	assert.Equal(t, routeRequest.StripPath, updatedRoute.StripPath)
	assert.Equal(t, routeRequest.PreserveHost, updatedRoute.PreserveHost)

	fetchedRoute, err := client.Routes().GetRoute(*createdRoute.Id)

	assert.Nil(t, err)
	assert.Equal(t, routeRequest.Protocols, fetchedRoute.Protocols)
	assert.Equal(t, routeRequest.Methods, fetchedRoute.Methods)
	assert.Equal(t, StringSlice([]string{}), fetchedRoute.Hosts)
	assert.Equal(t, routeRequest.Paths, fetchedRoute.Paths)
	assert.Equal(t, routeRequest.StripPath, fetchedRoute.StripPath)
	assert.Equal(t, routeRequest.PreserveHost, fetchedRoute.PreserveHost)

	err = client.Routes().DeleteRoute(*createdRoute.Id)
	assert.Nil(t, err)

	err = client.Services().DeleteServiceById(*createdService.Id)
	assert.Nil(t, err)

}

func TestRouteClient_UpdateRoutePathsToEmptyArray(t *testing.T) {
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
		Paths:        StringSlice([]string{"/foo"}),
		StripPath:    Bool(true),
		PreserveHost: Bool(true),
		Service:      ToId(*createdService.Id),
	}

	createdRoute, err := client.Routes().AddRoute(routeRequest)

	assert.Nil(t, err)

	routeRequest.Paths = StringSlice([]string{})

	updatedRoute, err := client.Routes().UpdateRoute(*createdRoute.Id, routeRequest)

	assert.Nil(t, err)
	assert.Equal(t, routeRequest.Protocols, updatedRoute.Protocols)
	assert.Equal(t, routeRequest.Methods, updatedRoute.Methods)
	assert.Equal(t, routeRequest.Hosts, updatedRoute.Hosts)
	assert.Equal(t, StringSlice([]string{}), updatedRoute.Paths)
	assert.Equal(t, routeRequest.StripPath, updatedRoute.StripPath)
	assert.Equal(t, routeRequest.PreserveHost, updatedRoute.PreserveHost)

	fetchedRoute, err := client.Routes().GetRoute(*createdRoute.Id)

	assert.Nil(t, err)
	assert.Equal(t, routeRequest.Protocols, fetchedRoute.Protocols)
	assert.Equal(t, routeRequest.Methods, fetchedRoute.Methods)
	assert.Equal(t, routeRequest.Hosts, fetchedRoute.Hosts)
	assert.Equal(t, StringSlice([]string{}), fetchedRoute.Paths)
	assert.Equal(t, routeRequest.StripPath, fetchedRoute.StripPath)
	assert.Equal(t, routeRequest.PreserveHost, fetchedRoute.PreserveHost)

	err = client.Routes().DeleteRoute(*createdRoute.Id)
	assert.Nil(t, err)

	err = client.Services().DeleteServiceById(*createdService.Id)
	assert.Nil(t, err)

}

func Test_AllRouteEndpointsShouldReturnErrorWhenRequestUnauthorised(t *testing.T) {

	//apiRequest := &ApiRequest{
	//	Name:        String("admin-api"),
	//	Uris:        StringSlice([]string{"/admin-api"}),
	//	UpstreamUrl: String("http://localhost:8001"),
	//}
	//
	//client := NewClient(NewDefaultConfig())
	//createdApi, err := client.Apis().Create(apiRequest)
	//
	//assert.Nil(t, err)
	//assert.NotNil(t, createdApi)
	//
	//consumerRequest := &ConsumerRequest{
	//	Username: "username-" + uuid.NewV4().String(),
	//	CustomId: "test-" + uuid.NewV4().String(),
	//}
	//
	//createdConsumer, err := client.Consumers().Create(consumerRequest)
	//
	//assert.Nil(t, err)
	//assert.NotNil(t, createdConsumer)
	//
	//pluginRequest := &PluginRequest{
	//	Name:  "key-auth",
	//	ApiId: *createdApi.Id,
	//	Config: map[string]interface{}{
	//		"hide_credentials": true,
	//	},
	//}
	//
	//createdPlugin, err := client.Plugins().Create(pluginRequest)
	//
	//assert.Nil(t, err)
	//assert.NotNil(t, createdPlugin)
	//
	//_, err = client.Consumers().CreatePluginConfig(createdConsumer.Id, "key-auth", "")
	//assert.Nil(t, err)
	//
	//serviceRequest := &ServiceRequest{
	//	Name:     String("service-name" + uuid.NewV4().String()),
	//	Protocol: String("http"),
	//	Host:     String("foo.com"),
	//}
	//
	//createdService, err := client.Services().AddService(serviceRequest)
	//
	//assert.Nil(t, err)
	//assert.NotNil(t, createdService)
	//
	//routeRequest := &RouteRequest{
	//	Protocols:    StringSlice([]string{"http"}),
	//	Methods:      StringSlice([]string{"GET"}),
	//	Hosts:        StringSlice([]string{"foo.com"}),
	//	Paths:        StringSlice([]string{"/bar"}),
	//	StripPath:    Bool(true),
	//	PreserveHost: Bool(true),
	//	Service:      &RouteServiceObject{Id: *createdService.Id},
	//}
	//
	//createdRoute, err := client.Routes().AddRoute(routeRequest)
	//
	//assert.Nil(t, err)
	//assert.NotNil(t, createdRoute)
	//
	//kongApiAddress := os.Getenv(EnvKongApiHostAddress) + "/admin-api"
	//unauthorisedClient := NewClient(&Config{HostAddress: kongApiAddress})
	//
	//r, err := unauthorisedClient.Routes().GetRoute(*createdRoute.Id)
	//assert.NotNil(t, err)
	//assert.Nil(t, r)
	//
	//results, err := unauthorisedClient.Routes().GetRoutes(&RouteQueryString{})
	//assert.NotNil(t, err)
	//assert.Nil(t, results)
	//
	//results, err = unauthorisedClient.Routes().GetRoutesFromServiceId(*createdService.Id)
	//assert.NotNil(t, err)
	//assert.Nil(t, results)
	//
	//results, err = unauthorisedClient.Routes().GetRoutesFromServiceName(*createdService.Name)
	//assert.NotNil(t, err)
	//assert.Nil(t, results)
	//
	//err = unauthorisedClient.Routes().DeleteRoute(*createdRoute.Id)
	//assert.NotNil(t, err)
	//
	//createNewRouteRequest := &RouteRequest{
	//	Protocols:    StringSlice([]string{"http"}),
	//	Methods:      StringSlice([]string{"POST"}),
	//	Hosts:        StringSlice([]string{"foo.com"}),
	//	Paths:        StringSlice([]string{"/bar"}),
	//	StripPath:    Bool(true),
	//	PreserveHost: Bool(true),
	//	Service:      &RouteServiceObject{Id: *createdService.Id},
	//}
	//
	//newRoute, err := unauthorisedClient.Routes().AddRoute(createNewRouteRequest)
	//assert.Nil(t, newRoute)
	//assert.NotNil(t, err)
	//
	//updatedRoute, err := unauthorisedClient.Routes().UpdateRoute(*createdRoute.Id, createNewRouteRequest)
	//assert.Nil(t, updatedRoute)
	//assert.NotNil(t, err)
	//
	//err = client.Plugins().DeleteById(createdPlugin.Id)
	//assert.Nil(t, err)
	//
	//err = client.Apis().DeleteById(*createdApi.Id)
	//assert.Nil(t, err)
	//
	//err = client.Routes().DeleteRoute(*createdRoute.Id)
	//assert.Nil(t, err)
	//
	//err = client.Services().DeleteServiceById(*createdService.Id)
	//assert.Nil(t, err)

}
