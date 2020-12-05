package gokong

import (
	"encoding/json"
	"fmt"
)

type ConsumerClient struct {
	config *Config
}

type ConsumerRequest struct {
	Username string `json:"username,omitempty" yaml:"username,omitempty"`
	CustomId string `json:"custom_id,omitempty" yaml:"custom_id,omitempty"`
}

type Consumer struct {
	Id       string `json:"id,omitempty" yaml:"id,omitempty"`
	CustomId string `json:"custom_id,omitempty" yaml:"custom_id,omitempty"`
	Username string `json:"username,omitempty" yaml:"custom_id,omitempty"`
}

type Consumers struct {
	Data   []*Consumer `json:"data,omitempty" yaml:"data,omitempty"`
	Next   string      `json:"next,omitempty" yaml:"next,omitempty"`
	Offset string      `json:"offset,omitempty" yaml:"offset,omitempty"`
}

type ConsumerQueryString struct {
	Offset string `json:"offset,omitempty"`
	Size   int    `json:"size"`
}

type ConsumerPluginConfig struct {
	Id   string `json:"id,omitempty" yaml:"id,omitempty"`
	Body string
}

type ConsumerPluginConfigs struct {
	Data   []map[string]interface{} `json:"data,omitempty" yaml:"data,omitempty"`
	Next   string                   `json:"next,omitempty" yaml:"next,omitempty"`
	Offset string                   `json:"offset,omitempty" yaml:"offset,omitempty"`
}

const ConsumersPath = "/consumers/"

func (consumerClient *ConsumerClient) GetByUsername(username string) (*Consumer, error) {
	return consumerClient.GetById(username)
}

func (consumerClient *ConsumerClient) GetById(id string) (*Consumer, error) {
	r, body, errs := newGet(consumerClient.config, ConsumersPath+id).End()
	if errs != nil {
		return nil, fmt.Errorf("could not get consumer, error: %v", errs)
	}

	if r.StatusCode == 401 || r.StatusCode == 403 {
		return nil, fmt.Errorf("not authorised, message from kong: %s", body)
	}

	consumer := &Consumer{}
	err := json.Unmarshal([]byte(body), consumer)
	if err != nil {
		return nil, fmt.Errorf("could not parse consumer get response, error: %v", err)
	}

	if consumer.Id == "" {
		return nil, nil
	}

	return consumer, nil
}

func (consumerClient *ConsumerClient) Create(consumerRequest *ConsumerRequest) (*Consumer, error) {
	r, body, errs := newPost(consumerClient.config, ConsumersPath).Send(consumerRequest).End()
	if errs != nil {
		return nil, fmt.Errorf("could not create new consumer, error: %v", errs)
	}

	if r.StatusCode == 401 || r.StatusCode == 403 {
		return nil, fmt.Errorf("not authorised, message from kong: %s", body)
	}

	createdConsumer := &Consumer{}
	err := json.Unmarshal([]byte(body), createdConsumer)
	if err != nil {
		return nil, fmt.Errorf("could not parse consumer creation response, error: %v", err)
	}

	if createdConsumer.Id == "" {
		return nil, fmt.Errorf("could not create consumer, error: %v", body)
	}

	return createdConsumer, nil
}

func (consumerClient *ConsumerClient) List(query *ConsumerQueryString) ([]*Consumer, error) {
	consumers := make([]*Consumer, 0)

	if query.Size < 100 {
		query.Size = 100
	}

	if query.Size > 1000 {
		query.Size = 1000
	}

	for {
		data := &Consumers{}

		r, body, errs := newGet(consumerClient.config, ConsumersPath).Query(*query).End()
		if errs != nil {
			return nil, fmt.Errorf("could not get the consumer, error: %v", errs)
		}

		if r.StatusCode == 401 || r.StatusCode == 403 {
			return nil, fmt.Errorf("not authorised, message from kong: %s", body)
		}

		err := json.Unmarshal([]byte(body), data)
		if err != nil {
			return nil, fmt.Errorf("could not parse consumer get response, error: %v", err)
		}

		consumers = append(consumers, data.Data...)
		if data.Next == "" {
			break
		}

		query.Offset = data.Offset
	}

	return consumers, nil
}

func (consumerClient *ConsumerClient) DeleteByUsername(username string) error {
	return consumerClient.DeleteById(username)
}

func (consumerClient *ConsumerClient) DeleteById(id string) error {
	r, body, errs := newDelete(consumerClient.config, ConsumersPath+id).End()
	if errs != nil {
		return fmt.Errorf("could not delete consumer, result: %v error: %v", r, errs)
	}

	if r.StatusCode == 401 || r.StatusCode == 403 {
		return fmt.Errorf("not authorised, message from kong: %s", body)
	}

	return nil
}

func (consumerClient *ConsumerClient) UpdateByUsername(username string, consumerRequest *ConsumerRequest) (*Consumer, error) {
	return consumerClient.UpdateById(username, consumerRequest)
}

func (consumerClient *ConsumerClient) UpdateById(id string, consumerRequest *ConsumerRequest) (*Consumer, error) {
	r, body, errs := newPatch(consumerClient.config, ConsumersPath+id).Send(consumerRequest).End()
	if errs != nil {
		return nil, fmt.Errorf("could not update consumer, error: %v", errs)
	}

	if r.StatusCode == 401 || r.StatusCode == 403 {
		return nil, fmt.Errorf("not authorised, message from kong: %s", body)
	}

	updatedConsumer := &Consumer{}
	err := json.Unmarshal([]byte(body), updatedConsumer)
	if err != nil {
		return nil, fmt.Errorf("could not parse consumer update response, error: %v", err)
	}

	if updatedConsumer.Id == "" {
		return nil, fmt.Errorf("could not update consumer, error: %v", body)
	}

	return updatedConsumer, nil
}

func (consumerClient *ConsumerClient) CreatePluginConfig(consumerId string, pluginName string, pluginConfig string) (*ConsumerPluginConfig, error) {
	r, body, errs := newPost(consumerClient.config, ConsumersPath+consumerId+"/"+pluginName).Send(pluginConfig).End()
	if errs != nil {
		return nil, fmt.Errorf("could not configure plugin for consumer, error: %v", errs)
	}

	if r.StatusCode == 401 || r.StatusCode == 403 {
		return nil, fmt.Errorf("not authorised, message from kong: %s", body)
	}

	createdConsumerPluginConfig := &ConsumerPluginConfig{}
	err := json.Unmarshal([]byte(body), createdConsumerPluginConfig)
	if err != nil {
		return nil, fmt.Errorf("could not parse consumer plugin config created response, error: %v", err)
	}

	if createdConsumerPluginConfig.Id == "" {
		return nil, fmt.Errorf("could not create consumer plugin config, error: %v", body)
	}

	createdConsumerPluginConfig.Body = body

	return createdConsumerPluginConfig, nil
}

func (consumerClient *ConsumerClient) GetPluginConfig(consumerId string, pluginName string, id string) (*ConsumerPluginConfig, error) {
	r, body, errs := newGet(consumerClient.config, ConsumersPath+consumerId+"/"+pluginName+"/"+id).End()
	if errs != nil {
		return nil, fmt.Errorf("could not get plugin config for consumer, error: %v", errs)
	}

	if r.StatusCode == 401 || r.StatusCode == 403 {
		return nil, fmt.Errorf("not authorised, message from kong: %s", body)
	}

	consumerPluginConfig := &ConsumerPluginConfig{}
	err := json.Unmarshal([]byte(body), consumerPluginConfig)
	if err != nil {
		return nil, fmt.Errorf("could not parse consumer plugin config response, error: %v", err)
	}

	if consumerPluginConfig.Id == "" {
		return nil, nil
	}

	consumerPluginConfig.Body = body

	return consumerPluginConfig, nil
}

func (consumerClient *ConsumerClient) GetPluginConfigs(consumerId string, pluginName string) ([]map[string]interface{}, error) {
	r, body, errs := newGet(consumerClient.config, ConsumersPath+consumerId+"/"+pluginName).End()
	if errs != nil {
		return nil, fmt.Errorf("could not get plugin config for consumer, error: %v", errs)
	}

	if r.StatusCode == 401 || r.StatusCode == 403 {
		return nil, fmt.Errorf("not authorised, message from kong: %s", body)
	}

	consumerPluginConfigs := &ConsumerPluginConfigs{}
	err := json.Unmarshal([]byte(body), consumerPluginConfigs)
	if err != nil {
		return nil, fmt.Errorf("could not parse consumer plugin config response, error: %v", err)
	}
	if len(consumerPluginConfigs.Data) == 0 {
		return nil, nil
	}

	return consumerPluginConfigs.Data, nil
}

func (consumerClient *ConsumerClient) DeletePluginConfig(consumerId string, pluginName string, id string) error {
	r, body, errs := newDelete(consumerClient.config, ConsumersPath+consumerId+"/"+pluginName+"/"+id).End()
	if errs != nil {
		return fmt.Errorf("could not delete plugin config for consumer, error: %v", errs)
	}

	if r.StatusCode == 401 || r.StatusCode == 403 {
		return fmt.Errorf("not authorised, message from kong: %s", body)
	}

	return nil
}
