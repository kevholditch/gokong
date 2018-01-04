package gokong

import (
	"encoding/json"
	"fmt"

	"github.com/parnurzeal/gorequest"
)

type PluginClient struct {
	config *Config
}

type PluginRequest struct {
	ID         string                 `json:"id,omitempty"`
	Name       string                 `json:"name"`
	CreatedAt  int                    `json:"created_at,omitempty"`
	APIID      string                 `json:"api_id,omitempty"`
	ConsumerID string                 `json:"consumer_id,omitempty"`
	Config     map[string]interface{} `json:"config,omitempty"`
}

type Plugin struct {
	ID         string                 `json:"id"`
	Name       string                 `json:"name"`
	CreatedAt  int                    `json:"created_at"`
	APIID      string                 `json:"api_id,omitempty"`
	ConsumerID string                 `json:"consumer_id,omitempty"`
	Config     map[string]interface{} `json:"config,omitempty"`
	Enabled    bool                   `json:"enabled,omitempty"`
}

type Plugins struct {
	Results []*Plugin `json:"data,omitempty"`
	Total   int       `json:"total,omitempty"`
	Next    string    `json:"next,omitempty"`
}

type PluginFilter struct {
	ID         string `url:"id,omitempty"`
	Name       string `url:"name,omitempty"`
	APIID      string `url:"api_id,omitempty"`
	ConsumerID string `url:"consumer_id,omitempty"`
	Size       int    `url:"size,omitempty"`
	Offset     int    `url:"offset,omitempty"`
}

const PluginsPath = "/plugins/"

func (pluginClient *PluginClient) GetByID(id string) (*Plugin, error) {

	_, body, errs := gorequest.New().Get(pluginClient.config.HostAddress + PluginsPath + id).End()
	if errs != nil {
		return nil, fmt.Errorf("could not get plugin, error: %v", errs)
	}

	plugin := &Plugin{}
	err := json.Unmarshal([]byte(body), plugin)
	if err != nil {
		return nil, fmt.Errorf("could not parse plugin plugin response, error: %v", err)
	}

	if plugin.ID == "" {
		return nil, nil
	}

	return plugin, nil
}

func (pluginClient *PluginClient) List() (*Plugins, error) {
	return pluginClient.ListFiltered(nil)
}

func (pluginClient *PluginClient) ListFiltered(filter *PluginFilter) (*Plugins, error) {

	ret := &Plugins{}
	address, err := addQueryString(pluginClient.config.HostAddress+PluginsPath, filter)

	if err != nil {
		return nil, fmt.Errorf("could not build query string for plugins filter, error: %v", err)
	}

	for {
		_, body, errs := gorequest.New().Get(address).End()
		if errs != nil {
			return nil, fmt.Errorf("could not get plugins, error: %v", errs)
		}

		plugins := &Plugins{}
		err = json.Unmarshal([]byte(body), plugins)
		if err != nil {
			return nil, fmt.Errorf("could not parse plugins list response, error: %v", err)
		}

		ret.Results = append(ret.Results, plugins.Results...)
		ret.Total += plugins.Total
		ret.Next = plugins.Next

		if plugins.Next != "" {
			address = plugins.Next
		} else {
			break
		}

	}

	return ret, nil
}

func (pluginClient *PluginClient) Create(pluginRequest *PluginRequest) (*Plugin, error) {

	_, body, errs := gorequest.New().Post(pluginClient.config.HostAddress + PluginsPath).Send(pluginRequest).End()
	if errs != nil {
		return nil, fmt.Errorf("could not create new plugin, error: %v", errs)
	}

	createdPlugin := &Plugin{}
	err := json.Unmarshal([]byte(body), createdPlugin)
	if err != nil {
		return nil, fmt.Errorf("could not parse plugin creation response, error: %v kong response: %s", err, body)
	}

	if createdPlugin.ID == "" {
		return nil, fmt.Errorf("could not create plugin, err: %v", body)
	}

	return createdPlugin, nil
}

func (pluginClient *PluginClient) UpdateOrAdd(pluginRequest *PluginRequest) (*Plugin, error) {

	var address string
	req := gorequest.New()
	if pluginRequest.APIID != "" {
		search := &PluginFilter{
			Name:  pluginRequest.Name,
			APIID: pluginRequest.APIID,
		}
		plugins, _ := pluginClient.ListFiltered(search)
		if plugins.Total == 1 {
			pluginRequest.ID = plugins.Results[0].ID
			pluginRequest.CreatedAt = plugins.Results[0].CreatedAt

		}
		address = pluginClient.config.HostAddress + APIsPath + pluginRequest.APIID + PluginsPath

		req = req.Put(address)
	} else {
		// global
		plugins, _ := pluginClient.ListFiltered(&PluginFilter{
			Name: pluginRequest.Name,
		})
		for _, v := range plugins.Results {
			if v.APIID == "" {
				pluginRequest.ID = plugins.Results[0].ID
				break
			}
		}

		if pluginRequest.ID != "" {
			address = pluginClient.config.HostAddress + PluginsPath + pluginRequest.ID
			req = req.Patch(address)
		} else {
			address = pluginClient.config.HostAddress + PluginsPath
			req = req.Post(address)
		}
	}

	_, body, errs := req.Send(pluginRequest).End()
	if errs != nil {
		return nil, fmt.Errorf("could not update or add plugin, error: %v %+v", errs, pluginRequest)
	}

	updatedPlugin := &Plugin{}
	err := json.Unmarshal([]byte(body), updatedPlugin)
	if err != nil {
		return nil, fmt.Errorf("could not parse plugin update or add response, error: %v kong response: %s %+v", err, body, pluginRequest)
	}

	if updatedPlugin.ID == "" {
		return nil, fmt.Errorf("could not update or add plugin, error: %v %s %+v", body, address, pluginRequest)
	}

	return updatedPlugin, nil
}

func (pluginClient *PluginClient) UpdateByID(id string, pluginRequest *PluginRequest) (*Plugin, error) {

	_, body, errs := gorequest.New().Patch(pluginClient.config.HostAddress + PluginsPath + id).Send(pluginRequest).End()
	if errs != nil {
		return nil, fmt.Errorf("could not update plugin, error: %v", errs)
	}

	updatedPlugin := &Plugin{}
	err := json.Unmarshal([]byte(body), updatedPlugin)
	if err != nil {
		return nil, fmt.Errorf("could not parse plugin update response, error: %v kong response: %s", err, body)
	}

	if updatedPlugin.ID == "" {
		return nil, fmt.Errorf("could not update plugin, error: %v", body)
	}

	return updatedPlugin, nil
}

func (pluginClient *PluginClient) DeleteByID(id string) error {

	res, _, errs := gorequest.New().Delete(pluginClient.config.HostAddress + PluginsPath + id).End()
	if errs != nil {
		return fmt.Errorf("could not delete plugin, result: %v error: %v", res, errs)
	}

	return nil
}
