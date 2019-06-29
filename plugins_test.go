package gokong

import (
	"fmt"
	"testing"

	uuid "github.com/satori/go.uuid"
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

func Test_PluginsCreateForAll(t *testing.T) {
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

	assert.Equal(t, pluginRequest.Name, createdPlugin.Name)
	assert.True(t, createdPlugin.Enabled)
	assert.Nil(t, createdPlugin.ConsumerId)

	err = client.Plugins().DeleteById(createdPlugin.Id)

	assert.Nil(t, err)

}

func Test_PluginsCreateForAll_ExplicitEnabled(t *testing.T) {
	pluginRequest := &PluginRequest{
		Name: "request-size-limiting",
		Config: map[string]interface{}{
			"allowed_payload_size": 128,
		},
		Enabled: Bool(true),
	}

	client := NewClient(NewDefaultConfig())
	createdPlugin, err := client.Plugins().Create(pluginRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdPlugin)

	assert.Equal(t, pluginRequest.Name, createdPlugin.Name)
	assert.True(t, createdPlugin.Enabled)
	assert.Nil(t, createdPlugin.ConsumerId)

	err = client.Plugins().DeleteById(createdPlugin.Id)

	assert.Nil(t, err)

}

func Test_PluginsCreateForAll_Disabled(t *testing.T) {
	pluginRequest := &PluginRequest{
		Name: "request-size-limiting",
		Config: map[string]interface{}{
			"allowed_payload_size": 128,
		},
		Enabled: Bool(false),
	}

	client := NewClient(NewDefaultConfig())
	createdPlugin, err := client.Plugins().Create(pluginRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdPlugin)

	assert.Equal(t, pluginRequest.Name, createdPlugin.Name)
	assert.False(t, createdPlugin.Enabled)
	assert.Nil(t, createdPlugin.ConsumerId)

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
		Name:       "request-size-limiting",
		ConsumerId: ToId(createdConsumer.Id),
		Config: map[string]interface{}{
			"allowed_payload_size": 128,
		},
	}

	createdPlugin, err := client.Plugins().Create(pluginRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdPlugin)

	assert.Equal(t, pluginRequest.Name, createdPlugin.Name)
	assert.True(t, createdPlugin.Enabled)
	assert.Equal(t, createdConsumer.Id, IdToString(createdPlugin.ConsumerId))
	assert.Nil(t, createdPlugin.ServiceId)
	assert.Nil(t, createdPlugin.RouteId)

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
	createdService, err := client.Services().Create(serviceRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdService)

	pluginRequest := &PluginRequest{
		Name:      "request-size-limiting",
		ServiceId: ToId(*createdService.Id),
		Config: map[string]interface{}{
			"allowed_payload_size": 128,
		},
	}

	createdPlugin, err := client.Plugins().Create(pluginRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdPlugin)

	assert.Equal(t, pluginRequest.Name, createdPlugin.Name)
	assert.True(t, createdPlugin.Enabled)
	assert.Equal(t, *createdService.Id, IdToString(createdPlugin.ServiceId))
	assert.Nil(t, createdPlugin.RouteId)
	assert.Nil(t, createdPlugin.ConsumerId)

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
	createdService, err := client.Services().Create(serviceRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdService)

	routeRequest := &RouteRequest{
		Protocols:    StringSlice([]string{"http"}),
		Methods:      StringSlice([]string{"GET"}),
		Hosts:        StringSlice([]string{fmt.Sprintf("%s.example.com", uuid.NewV4().String())}),
		Paths:        StringSlice([]string{"/"}),
		StripPath:    Bool(true),
		PreserveHost: Bool(false),
		Service:      ToId(*createdService.Id),
	}

	createdRoute, err := client.Routes().Create(routeRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdRoute)

	pluginRequest := &PluginRequest{
		Name:    "request-size-limiting",
		RouteId: ToId(*createdRoute.Id),
		Config: map[string]interface{}{
			"allowed_payload_size": 128,
		},
	}

	createdPlugin, err := client.Plugins().Create(pluginRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdPlugin)

	assert.Equal(t, pluginRequest.Name, createdPlugin.Name)
	assert.True(t, createdPlugin.Enabled)
	assert.Equal(t, *createdRoute.Id, IdToString(createdPlugin.RouteId))
	assert.Nil(t, createdPlugin.ConsumerId)
	assert.Nil(t, createdPlugin.ServiceId)

	err = client.Plugins().DeleteById(createdPlugin.Id)

	assert.Nil(t, err)

	err = client.Routes().DeleteById(*createdRoute.Id)

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
		Name:    "rate-limiting",
		RouteId: ToId("123"),
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
	assert.Equal(t, pluginRequest.Config["minute"].(float64), result.Config["minute"].(float64))
	assert.Equal(t, pluginRequest.Config["hour"].(float64), result.Config["hour"].(float64))

	err = client.Plugins().DeleteById(createdPlugin.Id)

	assert.Nil(t, err)

}

func Test_PluginsUpdateInvalid(t *testing.T) {

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
		Name: "request-size-limiting",
		Config: map[string]interface{}{
			"allowed_payload_size": 128,
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
		Name: "request-size-limiting",
		Config: map[string]interface{}{
			"allowed_payload_size": 128,
		},
	}

	createdPlugin2, err := client.Plugins().Create(pluginRequest2)

	assert.Nil(t, err)
	assert.NotNil(t, createdPlugin2)

	plugins, err := client.Plugins().List(&PluginQueryString{})

	assert.Nil(t, err)
	assert.NotNil(t, plugins)
	assert.True(t, len(plugins) > 1)

	err = client.Plugins().DeleteById(createdPlugin.Id)
	assert.Nil(t, err)

	err = client.Plugins().DeleteById(createdPlugin2.Id)
	assert.Nil(t, err)

}
