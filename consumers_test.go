package gokong

import (
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func Test_ConsumersGetById(t *testing.T) {
	consumerRequest := &ConsumerRequest{
		Username: "username-" + uuid.NewV4().String(),
		CustomId: "test-" + uuid.NewV4().String(),
	}

	client := NewClient(NewDefaultConfig())
	createdConsumer, err := client.Consumers().Create(consumerRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdConsumer)

	result, err := client.Consumers().GetById(createdConsumer.Id)

	assert.Equal(t, createdConsumer, result)

}

func Test_ConsumersGetNonExistentById(t *testing.T) {

	result, err := NewClient(NewDefaultConfig()).Consumers().GetById("7c924010-fca4-4314-8a3f-725cf749eac6")

	assert.Nil(t, err)
	assert.Nil(t, result)
}

func Test_ConsumersGetByUsername(t *testing.T) {
	consumerRequest := &ConsumerRequest{
		Username: "username-" + uuid.NewV4().String(),
		CustomId: "test-" + uuid.NewV4().String(),
	}

	client := NewClient(NewDefaultConfig())
	createdConsumer, err := client.Consumers().Create(consumerRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdConsumer)

	result, err := client.Consumers().GetByUsername(createdConsumer.Username)

	assert.Equal(t, createdConsumer, result)

}

func Test_ConsumersGetNonExistentByUsername(t *testing.T) {

	result, err := NewClient(NewDefaultConfig()).Consumers().GetById("408b5b13-b7c0-4ffd-afa1-aea957f00252")

	assert.Nil(t, err)
	assert.Nil(t, result)
}

func Test_ConsumersCreate(t *testing.T) {
	consumerRequest := &ConsumerRequest{
		Username: "username-" + uuid.NewV4().String(),
		CustomId: "test-" + uuid.NewV4().String(),
	}

	result, err := NewClient(NewDefaultConfig()).Consumers().Create(consumerRequest)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, consumerRequest.Username, result.Username)
	assert.Equal(t, consumerRequest.CustomId, result.CustomId)

}

func Test_ConsumersCreateInvalid(t *testing.T) {
	consumerRequest := &ConsumerRequest{}

	result, err := NewClient(NewDefaultConfig()).Consumers().Create(consumerRequest)

	assert.NotNil(t, err)
	assert.Nil(t, result)

}

func Test_ConsumersCreateOnlyUsername(t *testing.T) {
	consumerRequest := &ConsumerRequest{
		Username: "username-" + uuid.NewV4().String(),
	}

	result, err := NewClient(NewDefaultConfig()).Consumers().Create(consumerRequest)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, consumerRequest.Username, result.Username)
	assert.Equal(t, "", result.CustomId)

}

func Test_ConsumersCreateOnlyCustomId(t *testing.T) {
	consumerRequest := &ConsumerRequest{
		CustomId: "test-" + uuid.NewV4().String(),
	}

	result, err := NewClient(NewDefaultConfig()).Consumers().Create(consumerRequest)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "", result.Username)
	assert.Equal(t, consumerRequest.CustomId, result.CustomId)

}

func Test_ConsumersList(t *testing.T) {

	consumerRequest := &ConsumerRequest{
		Username: "username-" + uuid.NewV4().String(),
		CustomId: "test-" + uuid.NewV4().String(),
	}

	client := NewClient(NewDefaultConfig())
	createdConsumer, err := client.Consumers().Create(consumerRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdConsumer)

	results, err := client.Consumers().List()

	assert.True(t, len(results.Results) > 0)

}

func Test_ConsumersDeleteById(t *testing.T) {
	consumerRequest := &ConsumerRequest{
		Username: "username-" + uuid.NewV4().String(),
		CustomId: "test-" + uuid.NewV4().String(),
	}

	client := NewClient(NewDefaultConfig())
	createdConsumer, err := client.Consumers().Create(consumerRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdConsumer)

	err = client.Consumers().DeleteById(createdConsumer.Id)
	assert.Nil(t, err)

	deletedConsumer, err := client.Consumers().GetById(createdConsumer.Id)
	assert.Nil(t, deletedConsumer)
}

func Test_ConsumersDeleteByUsername(t *testing.T) {
	consumerRequest := &ConsumerRequest{
		Username: "username-" + uuid.NewV4().String(),
		CustomId: "test-" + uuid.NewV4().String(),
	}

	client := NewClient(NewDefaultConfig())
	createdConsumer, err := client.Consumers().Create(consumerRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdConsumer)

	err = client.Consumers().DeleteByUsername(createdConsumer.Username)
	assert.Nil(t, err)

	deletedConsumer, err := client.Consumers().GetById(createdConsumer.Id)
	assert.Nil(t, deletedConsumer)
}

func Test_ConsumersUpdateById(t *testing.T) {

	consumerRequest := &ConsumerRequest{
		Username: "username-" + uuid.NewV4().String(),
		CustomId: "test-" + uuid.NewV4().String(),
	}

	client := NewClient(NewDefaultConfig())
	createdConsumer, err := client.Consumers().Create(consumerRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdConsumer)
	assert.Equal(t, consumerRequest.CustomId, createdConsumer.CustomId)

	consumerRequest.CustomId = "test-" + uuid.NewV4().String()

	result, err := client.Consumers().UpdateById(createdConsumer.Id, consumerRequest)

	assert.Nil(t, err)
	assert.NotNil(t, result)

	assert.Equal(t, consumerRequest.CustomId, result.CustomId)
	assert.Equal(t, consumerRequest.Username, result.Username)
}

func Test_ConsumersUpdateByIdInvalid(t *testing.T) {

	consumerRequest := &ConsumerRequest{
		Username: "username-" + uuid.NewV4().String(),
		CustomId: "test-" + uuid.NewV4().String(),
	}

	client := NewClient(NewDefaultConfig())
	createdConsumer, err := client.Consumers().Create(consumerRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdConsumer)
	assert.Equal(t, consumerRequest.CustomId, createdConsumer.CustomId)

	consumerRequest.Username = ""
	consumerRequest.CustomId = ""

	result, err := client.Consumers().UpdateById(createdConsumer.Id, consumerRequest)

	assert.NotNil(t, err)
	assert.Nil(t, result)
}

func Test_ConsumersUpdateByUsername(t *testing.T) {

	consumerRequest := &ConsumerRequest{
		Username: "username-" + uuid.NewV4().String(),
		CustomId: "test-" + uuid.NewV4().String(),
	}

	client := NewClient(NewDefaultConfig())
	createdConsumer, err := client.Consumers().Create(consumerRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdConsumer)
	assert.Equal(t, consumerRequest.CustomId, createdConsumer.CustomId)

	consumerRequest.Username = "username-" + uuid.NewV4().String()

	result, err := client.Consumers().UpdateByUsername(createdConsumer.Username, consumerRequest)

	assert.Nil(t, err)
	assert.NotNil(t, result)

	assert.Equal(t, consumerRequest.CustomId, result.CustomId)
	assert.Equal(t, consumerRequest.Username, result.Username)
}

func Test_ConsumersUpdateByUsernameInvalid(t *testing.T) {

	consumerRequest := &ConsumerRequest{
		Username: "username-" + uuid.NewV4().String(),
		CustomId: "test-" + uuid.NewV4().String(),
	}

	client := NewClient(NewDefaultConfig())
	createdConsumer, err := client.Consumers().Create(consumerRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdConsumer)
	assert.Equal(t, consumerRequest.CustomId, createdConsumer.CustomId)

	consumerRequest.Username = ""
	consumerRequest.CustomId = ""

	result, err := client.Consumers().UpdateByUsername(createdConsumer.Username, consumerRequest)

	assert.NotNil(t, err)
	assert.Nil(t, result)
}

func Test_ConsumersPluginConfig(t *testing.T) {
	consumerRequest := &ConsumerRequest{
		Username: "username-" + uuid.NewV4().String(),
		CustomId: "test-" + uuid.NewV4().String(),
	}

	client := NewClient(NewDefaultConfig())
	createdConsumer, err := client.Consumers().Create(consumerRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdConsumer)

	pluginRequest := &PluginRequest{
		Name: "jwt",
		Config: map[string]interface{}{
			"claims_to_verify": []string{"exp"},
		},
	}

	plugin, err := client.Plugins().Create(pluginRequest)
	assert.Nil(t, err)
	assert.NotNil(t, plugin)

	createdPluginConfig, err := client.Consumers().CreatePluginConfig(createdConsumer.Id, "jwt", "{\"key\": \"a36c3049b36249a3c9f8891cb127243c\"}")

	assert.Nil(t, err)
	assert.NotNil(t, createdPluginConfig)
	assert.NotEqual(t, "", createdPluginConfig.Id)
	assert.Contains(t, createdPluginConfig.Body, "a36c3049b36249a3c9f8891cb127243c")

	retrievedPluginConfig, err := client.Consumers().GetPluginConfig(createdConsumer.Id, "jwt", createdPluginConfig.Id)

	assert.Nil(t, err)

	err = client.Consumers().DeletePluginConfig(createdConsumer.Id, "jwt", createdPluginConfig.Id)
	assert.Nil(t, err)

	retrievedPluginConfig, err = client.Consumers().GetPluginConfig(createdConsumer.Id, "jwt", createdPluginConfig.Id)

	assert.Nil(t, retrievedPluginConfig)
	assert.Nil(t, err)

}

func Test_AllConsumerEndpointsShouldReturnErrorWhenRequestUnauthorised(t *testing.T) {

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
	//kongApiAddress := os.Getenv(EnvKongApiHostAddress) + "/admin-api"
	//unauthorisedClient := NewClient(&Config{HostAddress: kongApiAddress})
	//
	//consumer, err := unauthorisedClient.Consumers().GetById(createdConsumer.Id)
	//assert.NotNil(t, err)
	//assert.Nil(t, consumer)
	//
	//consumer, err = unauthorisedClient.Consumers().GetByUsername(createdConsumer.Username)
	//assert.NotNil(t, err)
	//assert.Nil(t, consumer)
	//
	//results, err := unauthorisedClient.Consumers().List()
	//assert.NotNil(t, err)
	//assert.Nil(t, results)
	//
	//err = unauthorisedClient.Consumers().DeleteById(createdConsumer.Id)
	//assert.NotNil(t, err)
	//
	//err = unauthorisedClient.Consumers().DeleteByUsername(createdConsumer.Username)
	//assert.NotNil(t, err)
	//
	//createNewConsumer := &ConsumerRequest{
	//	Username: "username-" + uuid.NewV4().String(),
	//	CustomId: "test-" + uuid.NewV4().String(),
	//}
	//newConsumer, err := unauthorisedClient.Consumers().Create(createNewConsumer)
	//assert.Nil(t, newConsumer)
	//assert.NotNil(t, err)
	//
	//updatedConsumer, err := unauthorisedClient.Consumers().UpdateById(createdConsumer.Id, createNewConsumer)
	//assert.Nil(t, updatedConsumer)
	//assert.NotNil(t, err)
	//
	//updatedConsumer, err = unauthorisedClient.Consumers().UpdateByUsername(createdConsumer.Username, createNewConsumer)
	//assert.Nil(t, updatedConsumer)
	//assert.NotNil(t, err)
	//
	//createdPluginConfig, err := unauthorisedClient.Consumers().CreatePluginConfig(createdConsumer.Id, "jwt", "{\"key\": \"a36c3049b36249a3c9f8891cb127243c\"}")
	//assert.Nil(t, createdPluginConfig)
	//assert.NotNil(t, err)
	//
	//pluginConfig, err := unauthorisedClient.Consumers().GetPluginConfig(createdConsumer.Id, "jwt", "id")
	//assert.Nil(t, pluginConfig)
	//assert.NotNil(t, err)
	//
	//err = unauthorisedClient.Consumers().DeletePluginConfig(createdConsumer.Id, "jwt", "id")
	//assert.NotNil(t, err)
	//
	//err = client.Plugins().DeleteById(createdPlugin.Id)
	//assert.Nil(t, err)
	//
	//err = client.Apis().DeleteById(*createdApi.Id)
	//assert.Nil(t, err)

}
