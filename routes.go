package gokong

import (
	"encoding/json"
	"fmt"
)

type RouteClient interface {
	GetByName(name string) (*Route, error)
	GetById(id string) (*Route, error)
	Create(routeRequest *RouteRequest) (*Route, error)
	List(query *RouteQueryString) ([]*Route, error)
	GetRoutesFromServiceName(name string) ([]*Route, error)
	GetRoutesFromServiceId(id string) ([]*Route, error)
	UpdateByName(name string, routeRequest *RouteRequest) (*Route, error)
	UpdateById(id string, routeRequest *RouteRequest) (*Route, error)
	DeleteByName(name string) error
	DeleteById(id string) error
}

type routeClient struct {
	config *Config
}

type RouteRequest struct {
	Name          *string   `json:"name" yaml:"name"`
	Protocols     []*string `json:"protocols" yaml:"protocols"`
	Methods       []*string `json:"methods" yaml:"methods"`
	Hosts         []*string `json:"hosts" yaml:"hosts"`
	Paths         []*string `json:"paths" yaml:"paths"`
	RegexPriority *int      `json:"regex_priority" yaml:"regex_priority"`
	StripPath     *bool     `json:"strip_path" yaml:"strip_path"`
	PreserveHost  *bool     `json:"preserve_host" yaml:"preserve_host"`
	Snis          []*string `json:"snis" yaml:"snis"`
	Sources       []*IpPort `json:"sources" yaml:"sources"`
	Destinations  []*IpPort `json:"destinations" yaml:"destinations"`
	Service       *Id       `json:"service" yaml:"service"`
}

type Route struct {
	Id            *string   `json:"id" yaml:"id"`
	Name          *string   `json:"name" yaml:"name"`
	CreatedAt     *int      `json:"created_at" yaml:"created_at"`
	UpdatedAt     *int      `json:"updated_at" yaml:"updated_at"`
	Protocols     []*string `json:"protocols" yaml:"protocols"`
	Methods       []*string `json:"methods" yaml:"methods"`
	Hosts         []*string `json:"hosts" yaml:"hosts"`
	Paths         []*string `json:"paths" yaml:"paths"`
	RegexPriority *int      `json:"regex_priority" yaml:"regex_priority"`
	StripPath     *bool     `json:"strip_path" yaml:"strip_path"`
	PreserveHost  *bool     `json:"preserve_host" yaml:"preserve_host"`
	Snis          []*string `json:"snis" yaml:"snis"`
	Sources       []*IpPort `json:"sources" yaml:"sources"`
	Destinations  []*IpPort `json:"destinations" yaml:"destinations"`
	Service       *Id       `json:"service" yaml:"service"`
}

type IpPort struct {
	Ip   *string `json:"ip" yaml:"ip"`
	Port *int    `json:"port" yaml:"port"`
}

type Routes struct {
	Data   []*Route `json:"data" yaml:"data"`
	Next   *string  `json:"next" yaml:"next"`
	Offset string   `json:"offset,omitempty" yaml:"offset,omitempty"`
}

type RouteQueryString struct {
	Offset string `json:"offset,omitempty"`
	Size   int    `json:"size"`
}

const RoutesPath = "/routes/"

func (routeClient *routeClient) GetByName(name string) (*Route, error) {
	return routeClient.GetById(name)
}

func (routeClient *routeClient) GetById(id string) (*Route, error) {
	r, body, errs := newGet(routeClient.config, routeClient.config.HostAddress+RoutesPath+id).End()
	if errs != nil {
		return nil, fmt.Errorf("could not get the route, error: %v", errs)
	}

	if r.StatusCode == 401 || r.StatusCode == 403 {
		return nil, fmt.Errorf("not authorised, message from kong: %s", body)
	}

	route := &Route{}
	err := json.Unmarshal([]byte(body), route)
	if err != nil {
		return nil, fmt.Errorf("could not parse route get response, error: %v", err)
	}

	if route.Id == nil {
		return nil, nil
	}

	return route, nil
}

func (routeClient *routeClient) Create(routeRequest *RouteRequest) (*Route, error) {
	r, body, errs := newPost(routeClient.config, routeClient.config.HostAddress+RoutesPath).Send(routeRequest).End()
	if errs != nil {
		return nil, fmt.Errorf("could not register the route, error: %v", errs)
	}

	if r.StatusCode == 401 || r.StatusCode == 403 {
		return nil, fmt.Errorf("not authorised, message from kong: %s", body)
	}

	createdRoute := &Route{}
	err := json.Unmarshal([]byte(body), createdRoute)
	if err != nil {
		return nil, fmt.Errorf("could not parse route get response, error: %v", err)
	}

	if createdRoute.Id == nil {
		return nil, fmt.Errorf("could not register the route, error: %v", body)
	}

	return createdRoute, nil
}

func (routeClient *routeClient) List(query *RouteQueryString) ([]*Route, error) {
	routes := make([]*Route, 0)

	if query.Size < 100 {
		query.Size = 100
	}

	if query.Size > 1000 {
		query.Size = 1000
	}

	for {
		data := &Routes{}

		r, body, errs := newGet(routeClient.config, routeClient.config.HostAddress+RoutesPath).Query(*query).End()
		if errs != nil {
			return nil, fmt.Errorf("could not get the route, error: %v", errs)
		}

		if r.StatusCode == 401 || r.StatusCode == 403 {
			return nil, fmt.Errorf("not authorised, message from kong: %s", body)
		}

		err := json.Unmarshal([]byte(body), data)
		if err != nil {
			return nil, fmt.Errorf("could not parse route get response, error: %v", err)
		}

		routes = append(routes, data.Data...)

		if data.Next == nil || *data.Next == "" {
			break
		}

		query.Offset = data.Offset
	}

	return routes, nil
}

func (routeClient *routeClient) GetRoutesFromServiceName(name string) ([]*Route, error) {
	return routeClient.GetRoutesFromServiceId(name)
}

func (routeClient *routeClient) GetRoutesFromServiceId(id string) ([]*Route, error) {
	routes := make([]*Route, 0)
	data := &Routes{}

	for {
		r, body, errs := newGet(routeClient.config, routeClient.config.HostAddress+fmt.Sprintf("/services/%s/routes", id)).End()
		if errs != nil {
			return nil, fmt.Errorf("could not get the route, error: %v", errs)
		}

		if r.StatusCode == 401 || r.StatusCode == 403 {
			return nil, fmt.Errorf("not authorised, message from kong: %s", body)
		}

		err := json.Unmarshal([]byte(body), data)
		if err != nil {
			return nil, fmt.Errorf("could not parse route get response, error: %v", err)
		}

		routes = append(routes, data.Data...)

		if data.Next == nil || *data.Next == "" {
			break
		}

	}
	return routes, nil
}

func (routeClient *routeClient) UpdateByName(name string, routeRequest *RouteRequest) (*Route, error) {
	return routeClient.UpdateById(name, routeRequest)
}

func (routeClient *routeClient) UpdateById(id string, routeRequest *RouteRequest) (*Route, error) {
	r, body, errs := newPatch(routeClient.config, routeClient.config.HostAddress+RoutesPath+id).Send(routeRequest).End()
	if errs != nil {
		return nil, fmt.Errorf("could not update route, error: %v", errs)
	}

	if r.StatusCode == 401 || r.StatusCode == 403 {
		return nil, fmt.Errorf("not authorised, message from kong: %s", body)
	}

	updatedRoute := &Route{}
	err := json.Unmarshal([]byte(body), updatedRoute)
	if err != nil {
		return nil, fmt.Errorf("could not parse route update response, error: %v", err)
	}

	if updatedRoute.Id == nil {
		return nil, fmt.Errorf("could not update route, error: %v", body)
	}

	return updatedRoute, nil
}

func (routeClient *routeClient) DeleteByName(name string) error {
	return routeClient.DeleteById(name)
}

func (routeClient *routeClient) DeleteById(id string) error {
	r, body, errs := newDelete(routeClient.config, routeClient.config.HostAddress+RoutesPath+id).End()
	if errs != nil {
		return fmt.Errorf("could not delete the route, result: %v error: %v", r, errs)
	}

	if r.StatusCode == 401 || r.StatusCode == 403 {
		return fmt.Errorf("not authorised, message from kong: %s", body)
	}

	return nil
}
