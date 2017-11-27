package gokong

import (
	"encoding/json"
	"fmt"
	"github.com/parnurzeal/gorequest"
)

type PluginClient struct {
	config *Config
	client *gorequest.SuperAgent
}

type PluginRequest struct {
	Name       string                 `json:"name"`
	ApiId      string                 `json:"api_id,omitempty"`
	ConsumerId string                 `json:"consumer_id,omitempty"`
	Config     map[string]interface{} `json:"config,omitempty"`
}

type Plugin struct {
	Id         string                 `json:"id"`
	Name       string                 `json:"name"`
	ApiId      string                 `json:"api_id,omitempty"`
	ConsumerId string                 `json:"consumer_id,omitempty"`
	Config     map[string]interface{} `json:"config,omitempty"`
	Enabled    bool                   `json:"enabled,omitempty"`
}

const PluginsPath = "/plugins/"

func (pluginClient *PluginClient) GetById(id string) (*Plugin, error) {

	_, body, errs := pluginClient.client.Get(pluginClient.config.HostAddress + PluginsPath + id).End()
	if errs != nil {
		return nil, fmt.Errorf("could not get plugin, error: %v", errs)
	}

	plugin := &Plugin{}
	err := json.Unmarshal([]byte(body), plugin)
	if err != nil {
		return nil, fmt.Errorf("could not parse plugin plugin response, error: %v", err)
	}

	if plugin.Id == "" {
		return nil, nil
	}

	return plugin, nil
}

func (pluginClient *PluginClient) Create(pluginRequest *PluginRequest) (*Plugin, error) {

	_, body, errs := pluginClient.client.Post(pluginClient.config.HostAddress + PluginsPath).Send(pluginRequest).End()
	if errs != nil {
		return nil, fmt.Errorf("could not create new plugin, error: %v", errs)
	}

	createdPlugin := &Plugin{}
	err := json.Unmarshal([]byte(body), createdPlugin)
	if err != nil {
		return nil, fmt.Errorf("could not parse plugin creation response, error: %v", err)
	}

	return createdPlugin, nil
}

func (pluginClient *PluginClient) UpdateById(id string, pluginRequest *PluginRequest) (*Plugin, error) {

	_, body, errs := pluginClient.client.Patch(pluginClient.config.HostAddress + PluginsPath + id).Send(pluginRequest).End()
	if errs != nil {
		return nil, fmt.Errorf("could not update plugin, error: %v", errs)
	}

	updatedPlugin := &Plugin{}
	err := json.Unmarshal([]byte(body), updatedPlugin)
	if err != nil {
		return nil, fmt.Errorf("could not parse plugin update response, error: %v", err)
	}

	return updatedPlugin, nil
}

func (pluginClient *PluginClient) DeleteById(id string) error {

	res, _, errs := pluginClient.client.Delete(pluginClient.config.HostAddress + ApisPath + id).End()
	if errs != nil {
		return fmt.Errorf("could not delete plugin, result: %v error: %v", res, errs)
	}

	return nil
}
