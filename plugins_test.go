package gokong

import (
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"testing"
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
		Name:                   "test-" + uuid.NewV4().String(),
		Hosts:                  []string{"example.com"},
		Uris:                   []string{"/example"},
		Methods:                []string{"GET", "POST"},
		UpstreamUrl:            "http://localhost:4140/testservice",
		StripUri:               true,
		PreserveHost:           true,
		Retries:                "3",
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
		Name:  "basic-auth",
		ApiId: createdApi.Id,
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
		Retries:                "3",
		UpstreamConnectTimeout: 1000,
		UpstreamSendTimeout:    2000,
		UpstreamReadTimeout:    3000,
		HttpsOnly:              true,
		HttpIfTerminated:       true,
	}

	createdApi, err := client.Apis().Create(apiRequest)

	pluginRequest := &PluginRequest{
		Name:       "response-ratelimiting",
		ConsumerId: createdConsumer.Id,
		ApiId:      createdApi.Id,
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
	assert.Equal(t, createdApi.Id, createdPlugin.ApiId)

	err = client.Plugins().DeleteById(createdPlugin.Id)

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
		Name:                   "test-" + uuid.NewV4().String(),
		Hosts:                  []string{"example.com"},
		Uris:                   []string{"/example"},
		Methods:                []string{"GET", "POST"},
		UpstreamUrl:            "http://localhost:4140/testservice",
		StripUri:               true,
		PreserveHost:           true,
		Retries:                "3",
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
			"minute": float64(22),
			"hour":   float64(111),
		},
	}
	createdPlugin, err := client.Plugins().Create(pluginRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdPlugin)

	apiRequest2 := &ApiRequest{
		Name:                   "test-" + uuid.NewV4().String(),
		Hosts:                  []string{"example.com"},
		Uris:                   []string{"/example"},
		Methods:                []string{"GET", "POST"},
		UpstreamUrl:            "http://localhost:4140/testservice",
		StripUri:               true,
		PreserveHost:           true,
		Retries:                "3",
		UpstreamConnectTimeout: 1000,
		UpstreamSendTimeout:    2000,
		UpstreamReadTimeout:    3000,
		HttpsOnly:              true,
		HttpIfTerminated:       true,
	}

	createdApi2, err := client.Apis().Create(apiRequest2)

	assert.Nil(t, err)
	assert.NotNil(t, createdApi2)

	pluginRequest2 := &PluginRequest{
		Name:  "response-ratelimiting",
		ApiId: createdApi2.Id,
		Config: map[string]interface{}{
			"limits.sms.minute": 20,
		},
	}

	createdPlugin2, err := client.Plugins().Create(pluginRequest2)

	assert.Nil(t, err)
	assert.NotNil(t, createdPlugin2)

	results, err := client.Plugins().ListFiltered(&PluginFilter{ApiId: createdApi.Id})

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
