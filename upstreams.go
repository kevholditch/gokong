package gokong

import (
	"encoding/json"
	"fmt"

	"github.com/parnurzeal/gorequest"
)

type UpstreamClient struct {
	config *Config
}

type UpstreamRequest struct {
	Name      string `json:"name,omitempty"`
	Slots     int    `json:"slots,omitempty"`
	OrderList []int  `json:"orderlist,omitempty"`
}

type Upstream struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Slots     int    `json:"slots,omitempty"`
	OrderList []int  `json:"orderlist,omitempty"`
}

type Upstreams struct {
	Results []*Upstream `json:"data,omitempty"`
	Total   int         `json:"total,omitempty"`
	Next    string      `json:"next,omitempty"`
	Offset  string      `json:"offset,omitempty"`
}

type UpstreamFilter struct {
	ID     string `url:"id,omitempty"`
	Name   string `url:"name,omitempty"`
	Slots  int    `url:"slots,omitempty"`
	Size   int    `url:"size,omitempty"`
	Offset int    `url:"offset,omitempty"`
}

const UpstreamsPath = "/upstreams/"

func (upstreamClient *UpstreamClient) GetByName(name string) (*Upstream, error) {
	return upstreamClient.GetByID(name)
}

func (upstreamClient *UpstreamClient) GetByID(id string) (*Upstream, error) {

	_, body, errs := gorequest.New().Get(upstreamClient.config.HostAddress + UpstreamsPath + id).End()
	if errs != nil {
		return nil, fmt.Errorf("could not get upstream, error: %v", errs)
	}

	upstream := &Upstream{}
	err := json.Unmarshal([]byte(body), upstream)
	if err != nil {
		return nil, fmt.Errorf("could not parse upstream get response, error: %v", err)
	}

	if upstream.ID == "" {
		return nil, nil
	}

	return upstream, nil
}

func (upstreamClient *UpstreamClient) Create(upstreamRequest *UpstreamRequest) (*Upstream, error) {

	_, body, errs := gorequest.New().Post(upstreamClient.config.HostAddress + UpstreamsPath).Send(upstreamRequest).End()
	if errs != nil {
		return nil, fmt.Errorf("could not create new upstream, error: %v", errs)
	}

	createdUpstream := &Upstream{}
	err := json.Unmarshal([]byte(body), createdUpstream)
	if err != nil {
		return nil, fmt.Errorf("could not parse upstream creation response, error: %v", err)
	}

	if createdUpstream.ID == "" {
		return nil, fmt.Errorf("could not create update, error: %v", body)
	}

	return createdUpstream, nil
}

func (upstreamClient *UpstreamClient) DeleteByName(name string) error {
	return upstreamClient.DeleteByID(name)
}

func (upstreamClient *UpstreamClient) DeleteByID(id string) error {

	res, _, errs := gorequest.New().Delete(upstreamClient.config.HostAddress + UpstreamsPath + id).End()
	if errs != nil {
		return fmt.Errorf("could not delete upstream, result: %v error: %v", res, errs)
	}

	return nil
}

func (upstreamClient *UpstreamClient) List() (*Upstreams, error) {
	return upstreamClient.ListFiltered(nil)
}

func (upstreamClient *UpstreamClient) ListFiltered(filter *UpstreamFilter) (*Upstreams, error) {

	address, err := addQueryString(upstreamClient.config.HostAddress+UpstreamsPath, filter)

	if err != nil {
		return nil, fmt.Errorf("could not build query string for upstreams filter, error: %v", err)
	}

	_, body, errs := gorequest.New().Get(address).End()
	if errs != nil {
		return nil, fmt.Errorf("could not get upstreams, error: %v", errs)
	}

	upstreams := &Upstreams{}
	err = json.Unmarshal([]byte(body), upstreams)
	if err != nil {
		return nil, fmt.Errorf("could not parse upstreams list response, error: %v", err)
	}

	return upstreams, nil
}

func (upstreamClient *UpstreamClient) UpdateByName(name string, upstreamRequest *UpstreamRequest) (*Upstream, error) {
	return upstreamClient.UpdateByID(name, upstreamRequest)
}

func (upstreamClient *UpstreamClient) UpdateByID(id string, upstreamRequest *UpstreamRequest) (*Upstream, error) {

	_, body, errs := gorequest.New().Patch(upstreamClient.config.HostAddress + UpstreamsPath + id).Send(upstreamRequest).End()
	if errs != nil {
		return nil, fmt.Errorf("could not update upstream, error: %v", errs)
	}

	updatedUpstream := &Upstream{}
	err := json.Unmarshal([]byte(body), updatedUpstream)
	if err != nil {
		return nil, fmt.Errorf("could not parse upstream update response, error: %v", err)
	}

	if updatedUpstream.ID == "" {
		return nil, fmt.Errorf("could not update upstream, error: %v", body)
	}

	return updatedUpstream, nil
}
