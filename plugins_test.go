package gokong

import (
	"fmt"
	"os"
	"testing"

	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func Test_PluginsGetById(t *testing.T) {
	pluginRequest := &PluginRequest{
		Name: "request-size-limiting",
		Config: map[string]interface{}{
			"allowed_payload_size": 128,
		},
	}

	client := NewClient(NewDefaultConfig())
	createdPlugin, err := client.Plugins().Create(pluginRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdPlugin)

	result, err := client.Plugins().GetById(createdPlugin.Id)

	assert.Equal(t, createdPlugin, result)

	err = client.Plugins().DeleteById(createdPlugin.Id)

	assert.Nil(t, err)
}

func Test_PluginsGetNonExistentById(t *testing.T) {

	result, err := NewClient(NewDefaultConfig()).Plugins().GetById("cc8e128c-c38d-421c-93cd-b045f64d5d44")

	assert.Nil(t, err)
	assert.Nil(t, result)
}

func Test_PluginsCreateForAllApisAndConsumers(t *testing.T) {
	pluginRequest := &PluginRequest{
		Name: "response-ratelimiting",
		Config: map[string]interface{}{
			"limits.sms.minute": 20,
		},
	}

	client := NewClient(NewDefaultConfig())
	createdPlugin, err := client.Plugins().Create(pluginRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdPlugin)

	assert.Equal(t, pluginRequest.Name, createdPlugin.Name)
	assert.True(t, createdPlugin.Enabled)
	assert.Equal(t, "", createdPlugin.ConsumerId)
	assert.Equal(t, "", createdPlugin.ApiId)

	err = client.Plugins().DeleteById(createdPlugin.Id)

	assert.Nil(t, err)

}

func Test_PluginsCreateForASpecificApi(t *testing.T) {

	apiRequest := &ApiRequest{
		Name:                   String("test-" + uuid.NewV4().String()),
		Hosts:                  StringSlice([]string{"example.com"}),
		Uris:                   StringSlice([]string{"/example"}),
		Methods:                StringSlice([]string{"GET", "POST"}),
		UpstreamUrl:            String("http://localhost:4140/testservice"),
		StripUri:               Bool(true),
		PreserveHost:           Bool(true),
		Retries:                Int(3),
		UpstreamConnectTimeout: Int(1000),
		UpstreamSendTimeout:    Int(2000),
		UpstreamReadTimeout:    Int(3000),
		HttpsOnly:              Bool(true),
		HttpIfTerminated:       Bool(true),
	}

	client := NewClient(NewDefaultConfig())
	createdApi, err := client.Apis().Create(apiRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdApi)

	pluginRequest := &PluginRequest{
		Name:  "basic-auth",
		ApiId: *createdApi.Id,
		Config: map[string]interface{}{
			"hide_credentials": true,
		},
	}

	createdPlugin, err := client.Plugins().Create(pluginRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdPlugin)

	assert.Equal(t, pluginRequest.Name, createdPlugin.Name)
	assert.True(t, createdPlugin.Enabled)
	assert.Equal(t, "", createdPlugin.ConsumerId)
	assert.Equal(t, "", createdPlugin.ServiceId)
	assert.Equal(t, "", createdPlugin.RouteId)
	assert.Equal(t, *createdApi.Id, createdPlugin.ApiId)

	err = client.Plugins().DeleteById(createdPlugin.Id)

	assert.Nil(t, err)

}

func Test_PluginsCreateForASpecificConsumer(t *testing.T) {

	consumerRequest := &ConsumerRequest{
		Username: "username-" + uuid.NewV4().String(),
		CustomId: "test-" + uuid.NewV4().String(),
	}

	client := NewClient(NewDefaultConfig())
	createdConsumer, err := client.Consumers().Create(consumerRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdConsumer)

	pluginRequest := &PluginRequest{
		Name:       "response-ratelimiting",
		ConsumerId: createdConsumer.Id,
		Config: map[string]interface{}{
			"limits.sms.minute": 20,
		},
	}

	createdPlugin, err := client.Plugins().Create(pluginRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdPlugin)

	assert.Equal(t, pluginRequest.Name, createdPlugin.Name)
	assert.True(t, createdPlugin.Enabled)
	assert.Equal(t, createdConsumer.Id, createdPlugin.ConsumerId)
	assert.Equal(t, "", createdPlugin.ApiId)
	assert.Equal(t, "", createdPlugin.ServiceId)
	assert.Equal(t, "", createdPlugin.RouteId)

	err = client.Plugins().DeleteById(createdPlugin.Id)

	assert.Nil(t, err)

}

func Test_PluginsCreateForASpecificApiAndConsumer(t *testing.T) {

	consumerRequest := &ConsumerRequest{
		Username: "username-" + uuid.NewV4().String(),
		CustomId: "test-" + uuid.NewV4().String(),
	}

	client := NewClient(NewDefaultConfig())
	createdConsumer, err := client.Consumers().Create(consumerRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdConsumer)

	apiRequest := &ApiRequest{
		Name:                   String("test-" + uuid.NewV4().String()),
		Hosts:                  StringSlice([]string{"example.com"}),
		Uris:                   StringSlice([]string{"/example"}),
		Methods:                StringSlice([]string{"GET", "POST"}),
		UpstreamUrl:            String("http://localhost:4140/testservice"),
		StripUri:               Bool(true),
		PreserveHost:           Bool(true),
		Retries:                Int(3),
		UpstreamConnectTimeout: Int(1000),
		UpstreamSendTimeout:    Int(2000),
		UpstreamReadTimeout:    Int(3000),
		HttpsOnly:              Bool(true),
		HttpIfTerminated:       Bool(true),
	}

	createdApi, err := client.Apis().Create(apiRequest)

	pluginRequest := &PluginRequest{
		Name:       "response-ratelimiting",
		ConsumerId: createdConsumer.Id,
		ApiId:      *createdApi.Id,
		Config: map[string]interface{}{
			"limits.sms.minute": 20,
		},
	}

	createdPlugin, err := client.Plugins().Create(pluginRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdPlugin)

	assert.Equal(t, pluginRequest.Name, createdPlugin.Name)
	assert.True(t, createdPlugin.Enabled)
	assert.Equal(t, createdConsumer.Id, createdPlugin.ConsumerId)
	assert.Equal(t, *createdApi.Id, createdPlugin.ApiId)

	err = client.Plugins().DeleteById(createdPlugin.Id)

	assert.Nil(t, err)

}

func Test_PluginsCreateForASpecificService(t *testing.T) {

	serviceRequest := &ServiceRequest{
		Name:     String(fmt.Sprintf("service-%s", uuid.NewV4().String())),
		Protocol: String("http"),
		Host:     String(fmt.Sprintf("%s.example.com", uuid.NewV4().String())),
	}

	client := NewClient(NewDefaultConfig())
	createdService, err := client.Services().AddService(serviceRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdService)

	pluginRequest := &PluginRequest{
		Name:      "response-ratelimiting",
		ServiceId: *createdService.Id,
		Config: map[string]interface{}{
			"limits.sms.minute": 20,
		},
	}

	createdPlugin, err := client.Plugins().Create(pluginRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdPlugin)

	assert.Equal(t, pluginRequest.Name, createdPlugin.Name)
	assert.True(t, createdPlugin.Enabled)
	assert.Equal(t, *createdService.Id, createdPlugin.ServiceId)
	assert.Equal(t, "", createdPlugin.ApiId)
	assert.Equal(t, "", createdPlugin.RouteId)
	assert.Equal(t, "", createdPlugin.ConsumerId)

	err = client.Plugins().DeleteById(createdPlugin.Id)

	assert.Nil(t, err)

	err = client.Services().DeleteServiceById(*createdService.Id)

	assert.Nil(t, err)
}

func Test_PluginsCreateForASpecificRoute(t *testing.T) {

	serviceRequest := &ServiceRequest{
		Name:     String(fmt.Sprintf("service-%s", uuid.NewV4().String())),
		Protocol: String("http"),
		Host:     String(fmt.Sprintf("%s.example.com", uuid.NewV4().String())),
	}

	client := NewClient(NewDefaultConfig())
	createdService, err := client.Services().AddService(serviceRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdService)

	routeRequest := &RouteRequest{
		Protocols:    StringSlice([]string{"http"}),
		Methods:      StringSlice([]string{"GET"}),
		Hosts:        StringSlice([]string{fmt.Sprintf("%s.example.com", uuid.NewV4().String())}),
		Paths:        StringSlice([]string{"/"}),
		StripPath:    Bool(true),
		PreserveHost: Bool(false),
		Service:      &RouteServiceObject{Id: *createdService.Id},
	}

	createdRoute, err := client.Routes().AddRoute(routeRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdRoute)

	pluginRequest := &PluginRequest{
		Name:    "response-ratelimiting",
		RouteId: *createdRoute.Id,
		Config: map[string]interface{}{
			"limits.sms.minute": 20,
		},
	}

	createdPlugin, err := client.Plugins().Create(pluginRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdPlugin)

	assert.Equal(t, pluginRequest.Name, createdPlugin.Name)
	assert.True(t, createdPlugin.Enabled)
	assert.Equal(t, *createdRoute.Id, createdPlugin.RouteId)
	assert.Equal(t, "", createdPlugin.ConsumerId)
	assert.Equal(t, "", createdPlugin.ServiceId)
	assert.Equal(t, "", createdPlugin.ApiId)

	err = client.Plugins().DeleteById(createdPlugin.Id)

	assert.Nil(t, err)

	err = client.Routes().DeleteRoute(*createdRoute.Id)

	assert.Nil(t, err)

	err = client.Services().DeleteServiceById(*createdService.Id)

	assert.Nil(t, err)

}

func Test_PluginsCreatePluginNonExistant(t *testing.T) {

	pluginRequest := &PluginRequest{
		Name: "non-existant-plugin",
		Config: map[string]interface{}{
			"some-setting": 20,
		},
	}

	result, err := NewClient(NewDefaultConfig()).Plugins().Create(pluginRequest)

	assert.NotNil(t, err)
	assert.Nil(t, result)

}

func Test_PluginsCreatePluginInvalid(t *testing.T) {

	pluginRequest := &PluginRequest{
		Name:  "rate-limiting",
		ApiId: "123",
		Config: map[string]interface{}{
			"some-setting": 20,
		},
	}

	result, err := NewClient(NewDefaultConfig()).Plugins().Create(pluginRequest)

	assert.NotNil(t, err)
	assert.Nil(t, result)

}

func Test_PluginsUpdate(t *testing.T) {
	pluginRequest := &PluginRequest{
		Name: "rate-limiting",
		Config: map[string]interface{}{
			"minute": float64(20),
			"hour":   float64(500),
		},
	}

	client := NewClient(NewDefaultConfig())
	createdPlugin, err := client.Plugins().Create(pluginRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdPlugin)
	assert.Equal(t, pluginRequest.Config["minute"].(float64), createdPlugin.Config["minute"].(float64))
	assert.Equal(t, pluginRequest.Config["hour"].(float64), createdPlugin.Config["hour"].(float64))

	pluginRequest.Config = map[string]interface{}{
		"minute": float64(11),
		"hour":   float64(123),
	}

	result, err := client.Plugins().UpdateById(createdPlugin.Id, pluginRequest)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, pluginRequest.Name, result.Name)
	assert.Equal(t, pluginRequest.ConsumerId, result.ConsumerId)
	assert.Equal(t, pluginRequest.ApiId, result.ApiId)
	assert.Equal(t, pluginRequest.Config["minute"].(float64), result.Config["minute"].(float64))
	assert.Equal(t, pluginRequest.Config["hour"].(float64), result.Config["hour"].(float64))

	err = client.Plugins().DeleteById(createdPlugin.Id)

	assert.Nil(t, err)

}

func Test_PluginsUpdateInvalid(t *testing.T) {

	pluginRequest := &PluginRequest{
		Name: "rate-limiting",
		Config: map[string]interface{}{
			"minute": float64(20),
			"hour":   float64(500),
		},
	}

	client := NewClient(NewDefaultConfig())
	createdPlugin, err := client.Plugins().Create(pluginRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdPlugin)

	pluginRequest.Config = map[string]interface{}{
		"asd":   float64(11),
		"asdfs": float64(123),
	}

	result, err := client.Plugins().UpdateById(createdPlugin.Id, pluginRequest)

	assert.NotNil(t, err)
	assert.Nil(t, result)

	err = client.Plugins().DeleteById(createdPlugin.Id)
}

func Test_PluginsDelete(t *testing.T) {
	pluginRequest := &PluginRequest{
		Name: "rate-limiting",
		Config: map[string]interface{}{
			"minute": float64(20),
			"hour":   float64(500),
		},
	}

	client := NewClient(NewDefaultConfig())
	createdPlugin, err := client.Plugins().Create(pluginRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdPlugin)

	err = client.Plugins().DeleteById(createdPlugin.Id)
	assert.Nil(t, err)

	plugin, err := client.Plugins().GetById(createdPlugin.Id)
	assert.Nil(t, plugin)

}

func Test_PluginsList(t *testing.T) {
	pluginRequest := &PluginRequest{
		Name: "rate-limiting",
		Config: map[string]interface{}{
			"minute": float64(20),
			"hour":   float64(500),
		},
	}

	client := NewClient(NewDefaultConfig())
	createdPlugin, err := client.Plugins().Create(pluginRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdPlugin)

	pluginRequest2 := &PluginRequest{
		Name: "response-ratelimiting",
		Config: map[string]interface{}{
			"limits.sms.minute": 20,
		},
	}

	createdPlugin2, err := client.Plugins().Create(pluginRequest2)

	assert.Nil(t, err)
	assert.NotNil(t, createdPlugin2)

	results, err := client.Plugins().List()

	assert.Nil(t, err)
	assert.NotNil(t, results)
	assert.True(t, len(results.Results) > 1)

	err = client.Plugins().DeleteById(createdPlugin.Id)
	assert.Nil(t, err)

	err = client.Plugins().DeleteById(createdPlugin2.Id)
	assert.Nil(t, err)

}

func Test_PluginsListFilteredById(t *testing.T) {
	pluginRequest := &PluginRequest{
		Name: "rate-limiting",
		Config: map[string]interface{}{
			"minute": float64(20),
			"hour":   float64(500),
		},
	}

	client := NewClient(NewDefaultConfig())
	createdPlugin, err := client.Plugins().Create(pluginRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdPlugin)

	pluginRequest2 := &PluginRequest{
		Name: "response-ratelimiting",
		Config: map[string]interface{}{
			"limits.sms.minute": 20,
		},
	}

	createdPlugin2, err := client.Plugins().Create(pluginRequest2)

	assert.Nil(t, err)
	assert.NotNil(t, createdPlugin2)

	results, err := client.Plugins().ListFiltered(&PluginFilter{Id: createdPlugin.Id})

	assert.Nil(t, err)
	assert.NotNil(t, results)
	assert.True(t, len(results.Results) == 1)
	assert.Equal(t, createdPlugin, results.Results[0])

	err = client.Plugins().DeleteById(createdPlugin.Id)
	assert.Nil(t, err)

	err = client.Plugins().DeleteById(createdPlugin2.Id)
	assert.Nil(t, err)

}

func Test_PluginsListFilteredByName(t *testing.T) {
	pluginRequest := &PluginRequest{
		Name: "rate-limiting",
		Config: map[string]interface{}{
			"minute": float64(20),
			"hour":   float64(500),
		},
	}

	client := NewClient(NewDefaultConfig())
	createdPlugin, err := client.Plugins().Create(pluginRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdPlugin)

	pluginRequest2 := &PluginRequest{
		Name: "response-ratelimiting",
		Config: map[string]interface{}{
			"limits.sms.minute": 20,
		},
	}

	createdPlugin2, err := client.Plugins().Create(pluginRequest2)

	assert.Nil(t, err)
	assert.NotNil(t, createdPlugin2)

	results, err := client.Plugins().ListFiltered(&PluginFilter{Name: createdPlugin.Name})

	assert.Nil(t, err)
	assert.NotNil(t, results)
	assert.True(t, len(results.Results) == 1)
	assert.Equal(t, createdPlugin, results.Results[0])

	err = client.Plugins().DeleteById(createdPlugin.Id)
	assert.Nil(t, err)

	err = client.Plugins().DeleteById(createdPlugin2.Id)
	assert.Nil(t, err)

}

func Test_PluginsListFilteredByApiId(t *testing.T) {

	apiRequest := &ApiRequest{
		Name:                   String("test-" + uuid.NewV4().String()),
		Hosts:                  StringSlice([]string{"example.com"}),
		Uris:                   StringSlice([]string{"/example"}),
		Methods:                StringSlice([]string{"GET", "POST"}),
		UpstreamUrl:            String("http://localhost:4140/testservice"),
		StripUri:               Bool(true),
		PreserveHost:           Bool(true),
		Retries:                Int(3),
		UpstreamConnectTimeout: Int(1000),
		UpstreamSendTimeout:    Int(2000),
		UpstreamReadTimeout:    Int(3000),
		HttpsOnly:              Bool(true),
		HttpIfTerminated:       Bool(true),
	}

	client := NewClient(NewDefaultConfig())
	createdApi, err := client.Apis().Create(apiRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdApi)

	pluginRequest := &PluginRequest{
		Name:  "rate-limiting",
		ApiId: *createdApi.Id,
		Config: map[string]interface{}{
			"minute": float64(22),
			"hour":   float64(111),
		},
	}
	createdPlugin, err := client.Plugins().Create(pluginRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdPlugin)

	apiRequest2 := &ApiRequest{
		Name:                   String("test-" + uuid.NewV4().String()),
		Hosts:                  StringSlice([]string{"example.com"}),
		Uris:                   StringSlice([]string{"/example"}),
		Methods:                StringSlice([]string{"GET", "POST"}),
		UpstreamUrl:            String("http://localhost:4140/testservice"),
		StripUri:               Bool(true),
		PreserveHost:           Bool(true),
		Retries:                Int(3),
		UpstreamConnectTimeout: Int(1000),
		UpstreamSendTimeout:    Int(2000),
		UpstreamReadTimeout:    Int(3000),
		HttpsOnly:              Bool(true),
		HttpIfTerminated:       Bool(true),
	}

	createdApi2, err := client.Apis().Create(apiRequest2)

	assert.Nil(t, err)
	assert.NotNil(t, createdApi2)

	pluginRequest2 := &PluginRequest{
		Name:  "response-ratelimiting",
		ApiId: *createdApi2.Id,
		Config: map[string]interface{}{
			"limits.sms.minute": 20,
		},
	}

	createdPlugin2, err := client.Plugins().Create(pluginRequest2)

	assert.Nil(t, err)
	assert.NotNil(t, createdPlugin2)

	results, err := client.Plugins().ListFiltered(&PluginFilter{ApiId: *createdApi.Id})

	assert.Nil(t, err)
	assert.NotNil(t, results)
	assert.True(t, len(results.Results) == 1)
	assert.Equal(t, createdPlugin, results.Results[0])

	err = client.Plugins().DeleteById(createdPlugin.Id)
	assert.Nil(t, err)

	err = client.Plugins().DeleteById(createdPlugin2.Id)
	assert.Nil(t, err)

}

func Test_PluginsListFilteredByConsumerId(t *testing.T) {

	consumerRequest := &ConsumerRequest{
		Username: "username-" + uuid.NewV4().String(),
		CustomId: "test-" + uuid.NewV4().String(),
	}

	client := NewClient(NewDefaultConfig())
	createdConsumer, err := client.Consumers().Create(consumerRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdConsumer)

	pluginRequest := &PluginRequest{
		Name:       "rate-limiting",
		ConsumerId: createdConsumer.Id,
		Config: map[string]interface{}{
			"minute": float64(22),
			"hour":   float64(111),
		},
	}

	createdPlugin, err := client.Plugins().Create(pluginRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdPlugin)

	consumerRequest2 := &ConsumerRequest{
		Username: "username-" + uuid.NewV4().String(),
		CustomId: "test-" + uuid.NewV4().String(),
	}

	createdConsumer2, err := client.Consumers().Create(consumerRequest2)

	assert.Nil(t, err)
	assert.NotNil(t, createdConsumer2)

	pluginRequest2 := &PluginRequest{
		Name:       "response-ratelimiting",
		ConsumerId: createdConsumer2.Id,
		Config: map[string]interface{}{
			"limits.sms.minute": 20,
		},
	}

	createdPlugin2, err := client.Plugins().Create(pluginRequest2)

	assert.Nil(t, err)
	assert.NotNil(t, createdPlugin2)

	results, err := client.Plugins().ListFiltered(&PluginFilter{ConsumerId: createdConsumer.Id})

	assert.Nil(t, err)
	assert.NotNil(t, results)
	assert.True(t, len(results.Results) == 1)
	assert.Equal(t, createdPlugin, results.Results[0])

	err = client.Plugins().DeleteById(createdPlugin.Id)
	assert.Nil(t, err)

	err = client.Plugins().DeleteById(createdPlugin2.Id)
	assert.Nil(t, err)

}

func Test_PluginsListFilteredByServiceId(t *testing.T) {

	serviceRequest := &ServiceRequest{
		Name:     String(fmt.Sprintf("service-%s", uuid.NewV4().String())),
		Protocol: String("http"),
		Host:     String(fmt.Sprintf("%s.example.com", uuid.NewV4().String())),
	}

	client := NewClient(NewDefaultConfig())
	createdService, err := client.Services().AddService(serviceRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdService)

	pluginRequest := &PluginRequest{
		Name:      "rate-limiting",
		ServiceId: *createdService.Id,
		Config: map[string]interface{}{
			"minute": float64(22),
			"hour":   float64(111),
		},
	}

	createdPlugin, err := client.Plugins().Create(pluginRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdPlugin)

	serviceRequest2 := &ServiceRequest{
		Name:     String(fmt.Sprintf("service-%s", uuid.NewV4().String())),
		Protocol: String("http"),
		Host:     String(fmt.Sprintf("%s.example.com", uuid.NewV4().String())),
	}

	createdService2, err := client.Services().AddService(serviceRequest2)

	assert.Nil(t, err)
	assert.NotNil(t, createdService2)

	pluginRequest2 := &PluginRequest{
		Name:      "response-ratelimiting",
		ServiceId: *createdService2.Id,
		Config: map[string]interface{}{
			"limits.sms.minute": 20,
		},
	}

	createdPlugin2, err := client.Plugins().Create(pluginRequest2)

	assert.Nil(t, err)
	assert.NotNil(t, createdPlugin2)

	results, err := client.Plugins().ListFiltered(&PluginFilter{ServiceId: *createdService.Id})

	assert.Nil(t, err)
	assert.NotNil(t, results)
	assert.True(t, len(results.Results) == 1)
	assert.Equal(t, createdPlugin, results.Results[0])

	err = client.Plugins().DeleteById(createdPlugin.Id)
	assert.Nil(t, err)

	err = client.Plugins().DeleteById(createdPlugin2.Id)
	assert.Nil(t, err)

	err = client.Services().DeleteServiceById(*createdService.Id)
	assert.Nil(t, err)

	err = client.Services().DeleteServiceById(*createdService2.Id)
	assert.Nil(t, err)
}

func Test_PluginsListFilteredByRouteId(t *testing.T) {

	serviceRequest := &ServiceRequest{
		Name:     String(fmt.Sprintf("service-%s", uuid.NewV4().String())),
		Protocol: String("http"),
		Host:     String(fmt.Sprintf("%s.example.com", uuid.NewV4().String())),
	}

	client := NewClient(NewDefaultConfig())
	createdService, err := client.Services().AddService(serviceRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdService)

	routeRequest := &RouteRequest{
		Protocols:    StringSlice([]string{"http"}),
		Methods:      StringSlice([]string{"GET"}),
		Hosts:        StringSlice([]string{fmt.Sprintf("%s.example.com", uuid.NewV4().String())}),
		Paths:        StringSlice([]string{"/"}),
		StripPath:    Bool(true),
		PreserveHost: Bool(false),
		Service:      &RouteServiceObject{Id: *createdService.Id},
	}

	createdRoute, err := client.Routes().AddRoute(routeRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdRoute)

	pluginRequest := &PluginRequest{
		Name:    "rate-limiting",
		RouteId: *createdRoute.Id,
		Config: map[string]interface{}{
			"minute": float64(22),
			"hour":   float64(111),
		},
	}

	createdPlugin, err := client.Plugins().Create(pluginRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdPlugin)

	serviceRequest2 := &ServiceRequest{
		Name:     String(fmt.Sprintf("service-%s", uuid.NewV4().String())),
		Protocol: String("http"),
		Host:     String(fmt.Sprintf("%s.example.com", uuid.NewV4().String())),
	}

	createdService2, err := client.Services().AddService(serviceRequest2)

	assert.Nil(t, err)
	assert.NotNil(t, createdService2)

	routeRequest2 := &RouteRequest{
		Protocols:    StringSlice([]string{"http"}),
		Methods:      StringSlice([]string{"GET"}),
		Hosts:        StringSlice([]string{fmt.Sprintf("%s.example.com", uuid.NewV4().String())}),
		Paths:        StringSlice([]string{"/"}),
		StripPath:    Bool(true),
		PreserveHost: Bool(false),
		Service:      &RouteServiceObject{Id: *createdService2.Id},
	}

	createdRoute2, err := client.Routes().AddRoute(routeRequest2)

	assert.Nil(t, err)
	assert.NotNil(t, createdRoute2)

	pluginRequest2 := &PluginRequest{
		Name:    "response-ratelimiting",
		RouteId: *createdRoute2.Id,
		Config: map[string]interface{}{
			"limits.sms.minute": 20,
		},
	}

	createdPlugin2, err := client.Plugins().Create(pluginRequest2)

	assert.Nil(t, err)
	assert.NotNil(t, createdPlugin2)

	results, err := client.Plugins().ListFiltered(&PluginFilter{RouteId: *createdRoute.Id})

	assert.Nil(t, err)
	assert.NotNil(t, results)
	assert.True(t, len(results.Results) == 1)
	assert.Equal(t, createdPlugin, results.Results[0])

	err = client.Plugins().DeleteById(createdPlugin.Id)
	assert.Nil(t, err)

	err = client.Plugins().DeleteById(createdPlugin2.Id)
	assert.Nil(t, err)

	err = client.Routes().DeleteRoute(*createdRoute.Id)
	assert.Nil(t, err)

	err = client.Routes().DeleteRoute(*createdRoute2.Id)
	assert.Nil(t, err)

	err = client.Services().DeleteServiceById(*createdService.Id)
	assert.Nil(t, err)

	err = client.Services().DeleteServiceById(*createdService2.Id)
	assert.Nil(t, err)
}

func Test_PluginsListFilteredBySize(t *testing.T) {
	pluginRequest := &PluginRequest{
		Name: "rate-limiting",
		Config: map[string]interface{}{
			"minute": float64(20),
			"hour":   float64(500),
		},
	}

	client := NewClient(NewDefaultConfig())
	createdPlugin, err := client.Plugins().Create(pluginRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdPlugin)

	pluginRequest2 := &PluginRequest{
		Name: "response-ratelimiting",
		Config: map[string]interface{}{
			"limits.sms.minute": 20,
		},
	}

	createdPlugin2, err := client.Plugins().Create(pluginRequest2)

	assert.Nil(t, err)
	assert.NotNil(t, createdPlugin2)

	results, err := client.Plugins().ListFiltered(&PluginFilter{Size: 1})

	assert.Nil(t, err)
	assert.NotNil(t, results)
	assert.True(t, len(results.Results) == 1)

	err = client.Plugins().DeleteById(createdPlugin.Id)
	assert.Nil(t, err)

	err = client.Plugins().DeleteById(createdPlugin2.Id)
	assert.Nil(t, err)
}

func Test_AllPluginEndpointsShouldReturnErrorWhenRequestUnauthorised(t *testing.T) {

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

	testPluginRequest := &PluginRequest{
		Name: "request-size-limiting",
		Config: map[string]interface{}{
			"allowed_payload_size": 128,
		},
	}

	testPlugin, err := client.Plugins().Create(testPluginRequest)
	assert.NotNil(t, testPlugin)
	assert.Nil(t, err)

	_, err = client.Consumers().CreatePluginConfig(createdConsumer.Id, "key-auth", "")
	assert.Nil(t, err)

	kongApiAddress := os.Getenv(EnvKongApiHostAddress) + "/admin-api"
	unauthorisedClient := NewClient(&Config{HostAddress: kongApiAddress})

	p, err := unauthorisedClient.Plugins().GetById(testPlugin.Id)
	assert.NotNil(t, err)
	assert.Nil(t, p)

	results, err := unauthorisedClient.Plugins().List()
	assert.NotNil(t, err)
	assert.Nil(t, results)

	err = unauthorisedClient.Plugins().DeleteById(testPlugin.Id)
	assert.NotNil(t, err)

	createNewPluginRequest := &PluginRequest{
		Name: "response-ratelimiting",
		Config: map[string]interface{}{
			"limits.sms.minute": 20,
		},
	}

	newPlugin, err := unauthorisedClient.Plugins().Create(createNewPluginRequest)
	assert.Nil(t, newPlugin)
	assert.NotNil(t, err)

	updatedPlugin, err := unauthorisedClient.Plugins().UpdateById(testPlugin.Id, createNewPluginRequest)
	assert.Nil(t, updatedPlugin)
	assert.NotNil(t, err)

	err = client.Plugins().DeleteById(createdPlugin.Id)
	assert.Nil(t, err)

	err = client.Plugins().DeleteById(testPlugin.Id)
	assert.Nil(t, err)

	err = client.Apis().DeleteById(*createdApi.Id)
	assert.Nil(t, err)

}
