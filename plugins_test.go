package gokong

import (
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_PluginsGetById(t *testing.T) {
	pluginRequest := &PluginRequest{
		Name: "rate-limiting",
		Config: map[string]interface{}{
			"minute": "20",
			"hour":   "500",
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

func Test_PluginsCreateForAllApisAndConsumers(t *testing.T) {
	pluginRequest := &PluginRequest{
		Name: "rate-limiting",
		Config: map[string]interface{}{
			"minute": "20",
			"hour":   "500",
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
		Name:                   "test-" + uuid.NewV4().String(),
		Hosts:                  []string{"example.com"},
		Uris:                   []string{"/example"},
		Methods:                []string{"GET", "POST"},
		UpstreamUrl:            "http://localhost:4140/testservice",
		StripUri:               true,
		PreserveHost:           true,
		Retries:                3,
		UpstreamConnectTimeout: 1000,
		UpstreamSendTimeout:    2000,
		UpstreamReadTimeout:    3000,
		HttpsOnly:              true,
		HttpIfTerminated:       true,
	}

	client := NewClient(NewDefaultConfig())
	createdApi, err := client.Apis().Create(apiRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdApi)

	pluginRequest := &PluginRequest{
		Name:  "rate-limiting",
		ApiId: createdApi.Id,
		Config: map[string]interface{}{
			"minute": "20",
			"hour":   "500",
		},
	}

	createdPlugin, err := client.Plugins().Create(pluginRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdPlugin)

	assert.Equal(t, pluginRequest.Name, createdPlugin.Name)
	assert.True(t, createdPlugin.Enabled)
	assert.Equal(t, "", createdPlugin.ConsumerId)
	assert.Equal(t, createdApi.Id, createdPlugin.ApiId)

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
		Name:       "rate-limiting",
		ConsumerId: createdConsumer.Id,
		Config: map[string]interface{}{
			"minute": "20",
			"hour":   "500",
		},
	}

	createdPlugin, err := client.Plugins().Create(pluginRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdPlugin)

	assert.Equal(t, pluginRequest.Name, createdPlugin.Name)
	assert.True(t, createdPlugin.Enabled)
	assert.Equal(t, createdConsumer.Id, createdPlugin.ConsumerId)
	assert.Equal(t, "", createdPlugin.ApiId)

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
		Name:                   "test-" + uuid.NewV4().String(),
		Hosts:                  []string{"example.com"},
		Uris:                   []string{"/example"},
		Methods:                []string{"GET", "POST"},
		UpstreamUrl:            "http://localhost:4140/testservice",
		StripUri:               true,
		PreserveHost:           true,
		Retries:                3,
		UpstreamConnectTimeout: 1000,
		UpstreamSendTimeout:    2000,
		UpstreamReadTimeout:    3000,
		HttpsOnly:              true,
		HttpIfTerminated:       true,
	}

	createdApi, err := client.Apis().Create(apiRequest)

	pluginRequest := &PluginRequest{
		Name:       "rate-limiting",
		ConsumerId: createdConsumer.Id,
		ApiId:      createdApi.Id,
		Config: map[string]interface{}{
			"minute": "20",
			"hour":   "500",
		},
	}

	createdPlugin, err := client.Plugins().Create(pluginRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdPlugin)

	assert.Equal(t, pluginRequest.Name, createdPlugin.Name)
	assert.True(t, createdPlugin.Enabled)
	assert.Equal(t, createdConsumer.Id, createdPlugin.ConsumerId)
	assert.Equal(t, createdApi.Id, createdPlugin.ApiId)

	err = client.Plugins().DeleteById(createdPlugin.Id)

	assert.Nil(t, err)

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
