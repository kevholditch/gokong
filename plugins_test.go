package gokong

import (
	"testing"

	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func Test_PluginsGetByID(t *testing.T) {
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

	result, err := client.Plugins().GetByID(createdPlugin.ID)

	assert.Equal(t, createdPlugin, result)

	err = client.Plugins().DeleteByID(createdPlugin.ID)

	assert.Nil(t, err)
}

func Test_PluginsGetNonExistentByID(t *testing.T) {

	result, err := NewClient(NewDefaultConfig()).Plugins().GetByID("cc8e128c-c38d-421c-93cd-b045f64d5d44")

	assert.Nil(t, err)
	assert.Nil(t, result)
}

func Test_PluginsCreateForAllAPIsAndConsumers(t *testing.T) {
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
	assert.Equal(t, "", createdPlugin.ConsumerID)
	assert.Equal(t, "", createdPlugin.APIID)

	err = client.Plugins().DeleteByID(createdPlugin.ID)

	assert.Nil(t, err)

}

func Test_PluginsCreateForASpecificAPI(t *testing.T) {

	apiRequest := &APIRequest{
		Name:                   "test-" + uuid.NewV4().String(),
		Hosts:                  []string{"example.com"},
		URIs:                   []string{"/example"},
		Methods:                []string{"GET", "POST"},
		UpstreamURL:            "http://localhost:4140/testservice",
		StripURI:               true,
		PreserveHost:           true,
		Retries:                3,
		UpstreamConnectTimeout: 1000,
		UpstreamSendTimeout:    2000,
		UpstreamReadTimeout:    3000,
		HTTPSOnly:              true,
		HTTPIfTerminated:       true,
	}

	client := NewClient(NewDefaultConfig())
	createdAPI, err := client.APIs().Create(apiRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdAPI)

	pluginRequest := &PluginRequest{
		Name:  "basic-auth",
		APIID: createdAPI.ID,
		Config: map[string]interface{}{
			"hide_credentials": true,
		},
	}

	createdPlugin, err := client.Plugins().Create(pluginRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdPlugin)

	assert.Equal(t, pluginRequest.Name, createdPlugin.Name)
	assert.True(t, createdPlugin.Enabled)
	assert.Equal(t, "", createdPlugin.ConsumerID)
	assert.Equal(t, createdAPI.ID, createdPlugin.APIID)

	err = client.Plugins().DeleteByID(createdPlugin.ID)

	assert.Nil(t, err)

}

func Test_PluginsCreateForASpecificConsumer(t *testing.T) {

	consumerRequest := &ConsumerRequest{
		Username: "username-" + uuid.NewV4().String(),
		CustomID: "test-" + uuid.NewV4().String(),
	}

	client := NewClient(NewDefaultConfig())
	createdConsumer, err := client.Consumers().Create(consumerRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdConsumer)

	pluginRequest := &PluginRequest{
		Name:       "response-ratelimiting",
		ConsumerID: createdConsumer.ID,
		Config: map[string]interface{}{
			"limits.sms.minute": 20,
		},
	}

	createdPlugin, err := client.Plugins().Create(pluginRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdPlugin)

	assert.Equal(t, pluginRequest.Name, createdPlugin.Name)
	assert.True(t, createdPlugin.Enabled)
	assert.Equal(t, createdConsumer.ID, createdPlugin.ConsumerID)
	assert.Equal(t, "", createdPlugin.APIID)

	err = client.Plugins().DeleteByID(createdPlugin.ID)

	assert.Nil(t, err)

}

func Test_PluginsCreateForASpecificAPIAndConsumer(t *testing.T) {

	consumerRequest := &ConsumerRequest{
		Username: "username-" + uuid.NewV4().String(),
		CustomID: "test-" + uuid.NewV4().String(),
	}

	client := NewClient(NewDefaultConfig())
	createdConsumer, err := client.Consumers().Create(consumerRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdConsumer)

	apiRequest := &APIRequest{
		Name:                   "test-" + uuid.NewV4().String(),
		Hosts:                  []string{"example.com"},
		URIs:                   []string{"/example"},
		Methods:                []string{"GET", "POST"},
		UpstreamURL:            "http://localhost:4140/testservice",
		StripURI:               true,
		PreserveHost:           true,
		Retries:                3,
		UpstreamConnectTimeout: 1000,
		UpstreamSendTimeout:    2000,
		UpstreamReadTimeout:    3000,
		HTTPSOnly:              true,
		HTTPIfTerminated:       true,
	}

	createdAPI, err := client.APIs().Create(apiRequest)

	pluginRequest := &PluginRequest{
		Name:       "response-ratelimiting",
		ConsumerID: createdConsumer.ID,
		APIID:      createdAPI.ID,
		Config: map[string]interface{}{
			"limits.sms.minute": 20,
		},
	}

	createdPlugin, err := client.Plugins().Create(pluginRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdPlugin)

	assert.Equal(t, pluginRequest.Name, createdPlugin.Name)
	assert.True(t, createdPlugin.Enabled)
	assert.Equal(t, createdConsumer.ID, createdPlugin.ConsumerID)
	assert.Equal(t, createdAPI.ID, createdPlugin.APIID)

	err = client.Plugins().DeleteByID(createdPlugin.ID)

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
		APIID: "123",
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

	result, err := client.Plugins().UpdateByID(createdPlugin.ID, pluginRequest)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, pluginRequest.Name, result.Name)
	assert.Equal(t, pluginRequest.ConsumerID, result.ConsumerID)
	assert.Equal(t, pluginRequest.APIID, result.APIID)
	assert.Equal(t, pluginRequest.Config["minute"].(float64), result.Config["minute"].(float64))
	assert.Equal(t, pluginRequest.Config["hour"].(float64), result.Config["hour"].(float64))

	err = client.Plugins().DeleteByID(createdPlugin.ID)

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

	result, err := client.Plugins().UpdateByID(createdPlugin.ID, pluginRequest)

	assert.NotNil(t, err)
	assert.Nil(t, result)

	err = client.Plugins().DeleteByID(createdPlugin.ID)
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

	err = client.Plugins().DeleteByID(createdPlugin.ID)
	assert.Nil(t, err)

	plugin, err := client.Plugins().GetByID(createdPlugin.ID)
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

	err = client.Plugins().DeleteByID(createdPlugin.ID)
	assert.Nil(t, err)

	err = client.Plugins().DeleteByID(createdPlugin2.ID)
	assert.Nil(t, err)

}

func Test_PluginsListFilteredByID(t *testing.T) {
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

	results, err := client.Plugins().ListFiltered(&PluginFilter{ID: createdPlugin.ID})

	assert.Nil(t, err)
	assert.NotNil(t, results)
	assert.True(t, len(results.Results) == 1)
	assert.Equal(t, createdPlugin, results.Results[0])

	err = client.Plugins().DeleteByID(createdPlugin.ID)
	assert.Nil(t, err)

	err = client.Plugins().DeleteByID(createdPlugin2.ID)
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

	err = client.Plugins().DeleteByID(createdPlugin.ID)
	assert.Nil(t, err)

	err = client.Plugins().DeleteByID(createdPlugin2.ID)
	assert.Nil(t, err)

}

func Test_PluginsListFilteredByAPIID(t *testing.T) {

	apiRequest := &APIRequest{
		Name:                   "test-" + uuid.NewV4().String(),
		Hosts:                  []string{"example.com"},
		URIs:                   []string{"/example"},
		Methods:                []string{"GET", "POST"},
		UpstreamURL:            "http://localhost:4140/testservice",
		StripURI:               true,
		PreserveHost:           true,
		Retries:                3,
		UpstreamConnectTimeout: 1000,
		UpstreamSendTimeout:    2000,
		UpstreamReadTimeout:    3000,
		HTTPSOnly:              true,
		HTTPIfTerminated:       true,
	}

	client := NewClient(NewDefaultConfig())
	createdAPI, err := client.APIs().Create(apiRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdAPI)

	pluginRequest := &PluginRequest{
		Name:  "rate-limiting",
		APIID: createdAPI.ID,
		Config: map[string]interface{}{
			"minute": float64(22),
			"hour":   float64(111),
		},
	}
	createdPlugin, err := client.Plugins().Create(pluginRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdPlugin)

	apiRequest2 := &APIRequest{
		Name:                   "test-" + uuid.NewV4().String(),
		Hosts:                  []string{"example.com"},
		URIs:                   []string{"/example"},
		Methods:                []string{"GET", "POST"},
		UpstreamURL:            "http://localhost:4140/testservice",
		StripURI:               true,
		PreserveHost:           true,
		Retries:                3,
		UpstreamConnectTimeout: 1000,
		UpstreamSendTimeout:    2000,
		UpstreamReadTimeout:    3000,
		HTTPSOnly:              true,
		HTTPIfTerminated:       true,
	}

	createdAPI2, err := client.APIs().Create(apiRequest2)

	assert.Nil(t, err)
	assert.NotNil(t, createdAPI2)

	pluginRequest2 := &PluginRequest{
		Name:  "response-ratelimiting",
		APIID: createdAPI2.ID,
		Config: map[string]interface{}{
			"limits.sms.minute": 20,
		},
	}

	createdPlugin2, err := client.Plugins().Create(pluginRequest2)

	assert.Nil(t, err)
	assert.NotNil(t, createdPlugin2)

	results, err := client.Plugins().ListFiltered(&PluginFilter{APIID: createdAPI.ID})

	assert.Nil(t, err)
	assert.NotNil(t, results)
	assert.True(t, len(results.Results) == 1)
	assert.Equal(t, createdPlugin, results.Results[0])

	err = client.Plugins().DeleteByID(createdPlugin.ID)
	assert.Nil(t, err)

	err = client.Plugins().DeleteByID(createdPlugin2.ID)
	assert.Nil(t, err)

}

func Test_PluginsListFilteredByConsumerID(t *testing.T) {

	consumerRequest := &ConsumerRequest{
		Username: "username-" + uuid.NewV4().String(),
		CustomID: "test-" + uuid.NewV4().String(),
	}

	client := NewClient(NewDefaultConfig())
	createdConsumer, err := client.Consumers().Create(consumerRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdConsumer)

	pluginRequest := &PluginRequest{
		Name:       "rate-limiting",
		ConsumerID: createdConsumer.ID,
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
		CustomID: "test-" + uuid.NewV4().String(),
	}

	createdConsumer2, err := client.Consumers().Create(consumerRequest2)

	assert.Nil(t, err)
	assert.NotNil(t, createdConsumer2)

	pluginRequest2 := &PluginRequest{
		Name:       "response-ratelimiting",
		ConsumerID: createdConsumer2.ID,
		Config: map[string]interface{}{
			"limits.sms.minute": 20,
		},
	}

	createdPlugin2, err := client.Plugins().Create(pluginRequest2)

	assert.Nil(t, err)
	assert.NotNil(t, createdPlugin2)

	results, err := client.Plugins().ListFiltered(&PluginFilter{ConsumerID: createdConsumer.ID})

	assert.Nil(t, err)
	assert.NotNil(t, results)
	assert.True(t, len(results.Results) == 1)
	assert.Equal(t, createdPlugin, results.Results[0])

	err = client.Plugins().DeleteByID(createdPlugin.ID)
	assert.Nil(t, err)

	err = client.Plugins().DeleteByID(createdPlugin2.ID)
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

	err = client.Plugins().DeleteByID(createdPlugin.ID)
	assert.Nil(t, err)

	err = client.Plugins().DeleteByID(createdPlugin2.ID)
	assert.Nil(t, err)
}
