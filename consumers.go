package gokong

import (
	"encoding/json"
	"fmt"

	"github.com/parnurzeal/gorequest"
)

type ConsumerClient struct {
	config *Config
}

type ConsumerRequest struct {
	Username string `json:"username,omitempty"`
	CustomID string `json:"custom_id,omitempty"`
}

type Consumer struct {
	ID       string `json:"id,omitempty"`
	CustomID string `json:"custom_id,omitempty"`
	Username string `json:"username,omitempty"`
}

type Consumers struct {
	Results []*Consumer `json:"data,omitempty"`
	Total   int         `json:"total,omitempty"`
	Next    string      `json:"next,omitempty"`
}

type ConsumerFilter struct {
	ID       string `url:"id,omitempty"`
	CustomID string `url:"custom_id,omitempty"`
	Username string `url:"username,omitempty"`
	Size     int    `url:"size,omitempty"`
	Offset   int    `url:"offset,omitempty"`
}

const ConsumersPath = "/consumers/"

func (consumerClient *ConsumerClient) GetByUsername(username string) (*Consumer, error) {
	return consumerClient.GetByID(username)
}

func (consumerClient *ConsumerClient) GetByID(id string) (*Consumer, error) {

	_, body, errs := gorequest.New().Get(consumerClient.config.HostAddress + ConsumersPath + id).End()
	if errs != nil {
		return nil, fmt.Errorf("could not get consumer, error: %v", errs)
	}

	consumer := &Consumer{}
	err := json.Unmarshal([]byte(body), consumer)
	if err != nil {
		return nil, fmt.Errorf("could not parse consumer get response, error: %v", err)
	}

	if consumer.ID == "" {
		return nil, nil
	}

	return consumer, nil
}

func (consumerClient *ConsumerClient) Create(consumerRequest *ConsumerRequest) (*Consumer, error) {

	_, body, errs := gorequest.New().Post(consumerClient.config.HostAddress + ConsumersPath).Send(consumerRequest).End()
	if errs != nil {
		return nil, fmt.Errorf("could not create new consumer, error: %v", errs)
	}

	createdConsumer := &Consumer{}
	err := json.Unmarshal([]byte(body), createdConsumer)
	if err != nil {
		return nil, fmt.Errorf("could not parse consumer creation response, error: %v", err)
	}

	if createdConsumer.ID == "" {
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

	_, body, errs := gorequest.New().Get(address).End()
	if errs != nil {
		return nil, fmt.Errorf("could not get consumers, error: %v", errs)
	}

	consumers := &Consumers{}
	err = json.Unmarshal([]byte(body), consumers)
	if err != nil {
		return nil, fmt.Errorf("could not parse consumers list response, error: %v", err)
	}

	return consumers, nil
}

func (consumerClient *ConsumerClient) DeleteByUsername(username string) error {
	return consumerClient.DeleteByID(username)
}

func (consumerClient *ConsumerClient) DeleteByID(id string) error {

	res, _, errs := gorequest.New().Delete(consumerClient.config.HostAddress + ConsumersPath + id).End()
	if errs != nil {
		return fmt.Errorf("could not delete consumer, result: %v error: %v", res, errs)
	}

	return nil
}

func (consumerClient *ConsumerClient) UpdateByUsername(username string, consumerRequest *ConsumerRequest) (*Consumer, error) {
	return consumerClient.UpdateByID(username, consumerRequest)
}

func (consumerClient *ConsumerClient) UpdateByID(id string, consumerRequest *ConsumerRequest) (*Consumer, error) {

	_, body, errs := gorequest.New().Patch(consumerClient.config.HostAddress + ConsumersPath + id).Send(consumerRequest).End()
	if errs != nil {
		return nil, fmt.Errorf("could not update consumer, error: %v", errs)
	}

	updatedConsumer := &Consumer{}
	err := json.Unmarshal([]byte(body), updatedConsumer)
	if err != nil {
		return nil, fmt.Errorf("could not parse consumer update response, error: %v", err)
	}

	if updatedConsumer.ID == "" {
		return nil, fmt.Errorf("could not update consumer, error: %v", body)
	}

	return updatedConsumer, nil
}
