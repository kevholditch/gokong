package gokong

import (
	"encoding/json"
	"fmt"
)

type ConsumerClient struct {
	config *Config
}

type ConsumerRequest struct {
	Username string `json:"username,omitempty"`
	CustomId string `json:"custom_id,omitempty"`
}

type Consumer struct {
	Id       string `json:"id,omitempty"`
	CustomId string `json:"custom_id,omitempty"`
	Username string `json:"username,omitempty"`
}

type Consumers struct {
	Results []*Consumer `json:"data,omitempty"`
	Total   int         `json:"total,omitempty"`
	Next    string      `json:"next,omitempty"`
}

type ConsumerFilter struct {
	Id       string `url:"id,omitempty"`
	CustomId string `url:"custom_id,omitempty"`
	Username string `url:"username,omitempty"`
	Size     int    `url:"size,omitempty"`
	Offset   int    `url:"offset,omitempty"`
}

type ConsumerPluginConfigs struct {
	Results []*ConsumerPluginConfig `json:"data,omitempty"`
	Total   int                     `json:"total,omitempty"`
	Next    string                  `json:"next,omitempty"`
}

type ConsumerPluginConfig struct {
	Id   string `json:"id,omitempty"`
	Body string
}

/* Type used specifically to unmarshal the ConsumerPluginConfig into an Id and the
 * body of the config
 */
type consumerPluginConfigMarshal struct {
	Id string `json:"id,omitempty"`
}

func (c *ConsumerPluginConfig) UnmarshalJSON(data []byte) error {
	consumerPluginConfigM := consumerPluginConfigMarshal{}
	err := json.Unmarshal([]byte(data), &consumerPluginConfigM)

	if err != nil {
		return err
	}

	c.Id = consumerPluginConfigM.Id
	c.Body = string(data[:])
	return nil
}

const ConsumersPath = "/consumers/"

func (consumerClient *ConsumerClient) GetByUsername(username string) (*Consumer, error) {
	return consumerClient.GetById(username)
}

func (consumerClient *ConsumerClient) GetById(id string) (*Consumer, error) {

	r, body, errs := newGet(consumerClient.config, consumerClient.config.HostAddress+ConsumersPath+id).End()
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

	r, body, errs := newPost(consumerClient.config, consumerClient.config.HostAddress+ConsumersPath).Send(consumerRequest).End()
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

func (consumerClient *ConsumerClient) List() (*Consumers, error) {
	return consumerClient.ListFiltered(nil)
}

func (consumerClient *ConsumerClient) ListFiltered(filter *ConsumerFilter) (*Consumers, error) {

	address, err := addQueryString(consumerClient.config.HostAddress+ConsumersPath, filter)

	if err != nil {
		return nil, fmt.Errorf("could not build query string for consumer filter, error: %v", err)
	}

	r, body, errs := newGet(consumerClient.config, address).End()
	if errs != nil {
		return nil, fmt.Errorf("could not get consumers, error: %v", errs)
	}

	if r.StatusCode == 401 || r.StatusCode == 403 {
		return nil, fmt.Errorf("not authorised, message from kong: %s", body)
	}

	consumers := &Consumers{}
	err = json.Unmarshal([]byte(body), consumers)
	if err != nil {
		return nil, fmt.Errorf("could not parse consumers list response, error: %v", err)
	}

	return consumers, nil
}

func (consumerClient *ConsumerClient) DeleteByUsername(username string) error {
	return consumerClient.DeleteById(username)
}

func (consumerClient *ConsumerClient) DeleteById(id string) error {

	r, body, errs := newDelete(consumerClient.config, consumerClient.config.HostAddress+ConsumersPath+id).End()
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

	r, body, errs := newPatch(consumerClient.config, consumerClient.config.HostAddress+ConsumersPath+id).Send(consumerRequest).End()
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

	r, body, errs := newPost(consumerClient.config, consumerClient.config.HostAddress+ConsumersPath+consumerId+"/"+pluginName).Send(pluginConfig).End()
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

	return createdConsumerPluginConfig, nil
}

func (consumerClient *ConsumerClient) ListPluginConfig(consumerId string, pluginName string) (*ConsumerPluginConfigs, error) {
	r, body, errs := newGet(consumerClient.config, consumerClient.config.HostAddress+ConsumersPath+consumerId+"/"+pluginName).End()
	if errs != nil {
		return nil, fmt.Errorf("could not get plugin config for consumer, error: %v", errs)
	}

	if r.StatusCode == 401 || r.StatusCode == 403 {
		return nil, fmt.Errorf("not authorised, message from kong: %s", body)
	}

	consumerPluginConfigs := ConsumerPluginConfigs{}
	err := json.Unmarshal([]byte(body), &consumerPluginConfigs)
	if err != nil {
		return nil, fmt.Errorf("could not parse consumer plugin configs response, error: %v", err)
	}

	return &consumerPluginConfigs, nil
}

func (consumerClient *ConsumerClient) GetPluginConfig(consumerId string, pluginName string, id string) (*ConsumerPluginConfig, error) {

	r, body, errs := newGet(consumerClient.config, consumerClient.config.HostAddress+ConsumersPath+consumerId+"/"+pluginName+"/"+id).End()
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

	return consumerPluginConfig, nil
}

func (consumerClient *ConsumerClient) DeletePluginConfig(consumerId string, pluginName string, id string) error {

	r, body, errs := newDelete(consumerClient.config, consumerClient.config.HostAddress+ConsumersPath+consumerId+"/"+pluginName+"/"+id).End()
	if errs != nil {
		return fmt.Errorf("could not delete plugin config for consumer, error: %v", errs)
	}

	if r.StatusCode == 401 || r.StatusCode == 403 {
		return fmt.Errorf("not authorised, message from kong: %s", body)
	}

	return nil
}
