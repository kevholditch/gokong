// +build all community

package gokong

import (
	"testing"
)

func Test_ApiKeyPassedViaHeader(t *testing.T) {

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
	//c, err := client.Consumers().CreatePluginConfig(createdConsumer.Id, "key-auth", "")
	//
	//m := make(map[string]interface{})
	//json.Unmarshal([]byte(c.Body), &m)
	//
	//kongApiAddress := os.Getenv(EnvKongApiHostAddress) + "/admin-api"
	//unauthorisedClient := NewClient(&Config{HostAddress: kongApiAddress})
	//
	//api, err := unauthorisedClient.Apis().GetByName("admin-api")
	//
	//assert.NotNil(t, err)
	//assert.Nil(t, api)
	//
	//authorisedClient := NewClient(&Config{HostAddress: kongApiAddress, ApiKey: m["key"].(string)})
	//
	//api, err = authorisedClient.Apis().GetByName("admin-api")
	//assert.Nil(t, err)
	//assert.NotNil(t, api)
	//
	//err = client.Plugins().DeleteById(createdPlugin.Id)
	//
	//assert.Nil(t, err)
	//
	//err = client.Apis().DeleteById(*createdApi.Id)
	//
	//assert.Nil(t, err)

}
