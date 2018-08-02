package gokong

import (
	"testing"

	"os"

	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func Test_UpstreamsGetById(t *testing.T) {

	upstreamRequest := &UpstreamRequest{
		Name:  "upstream-" + uuid.NewV4().String(),
		Slots: 10,
	}

	client := NewClient(NewDefaultConfig())
	createdUpstream, err := client.Upstreams().Create(upstreamRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdUpstream)

	result, err := client.Upstreams().GetById(createdUpstream.Id)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, createdUpstream, result)

}

func Test_UpstreamsGetByName(t *testing.T) {

	upstreamRequest := &UpstreamRequest{
		Name:  "upstream-" + uuid.NewV4().String(),
		Slots: 10,
	}

	client := NewClient(NewDefaultConfig())
	createdUpstream, err := client.Upstreams().Create(upstreamRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdUpstream)

	result, err := client.Upstreams().GetByName(createdUpstream.Name)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, createdUpstream, result)

}

func Test_UpstreamsGetByIdForNonExistentUpstreamById(t *testing.T) {

	result, err := NewClient(NewDefaultConfig()).Upstreams().GetById(uuid.NewV4().String())

	assert.Nil(t, err)
	assert.Nil(t, result)

}

func Test_UpstreamsGetByIdForNonExistentUpstreamByName(t *testing.T) {

	result, err := NewClient(NewDefaultConfig()).Upstreams().GetByName(uuid.NewV4().String())

	assert.Nil(t, err)
	assert.Nil(t, result)

}

func Test_UpstreamsCreate(t *testing.T) {
	upstreamRequest := &UpstreamRequest{
		Name:  "upstream-" + uuid.NewV4().String(),
		Slots: 10,
		HealthChecks: &UpstreamHealthCheck{
			Active: &UpstreamHealthCheckActive{
				Concurrency: 10,
				HttpPath:    "/",
				Timeout:     1,
				Healthy: &ActiveHealthy{
					HttpStatuses: []int{200, 302},
					Interval:     0,
					Successes:    0,
				},
				Unhealthy: &ActiveUnhealthy{
					HttpFailures: 0,
					HttpStatuses: []int{429, 404, 500, 501, 502, 503, 504, 505},
					Interval:     0,
					TcpFailures:  0,
					Timeouts:     0,
				},
			},
			Passive: &UpstreamHealthCheckPassive{
				Healthy: &PassiveHealthy{
					HttpStatuses: []int{200, 201, 202, 203, 204, 205, 206, 207, 208, 226, 300, 301, 302, 303, 304, 305, 306, 307, 308},
					Successes:    0,
				},
				Unhealthy: &PassiveUnhealthy{
					HttpFailures: 0,
					HttpStatuses: []int{429, 500, 503},
					TcpFailures:  0,
					Timeouts:     0,
				},
			},
		},
	}

	result, err := NewClient(NewDefaultConfig()).Upstreams().Create(upstreamRequest)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.True(t, result.Id != "")
	assert.Equal(t, upstreamRequest.Name, result.Name)
	assert.Equal(t, upstreamRequest.Slots, result.Slots)
	assert.Equal(t, upstreamRequest.HealthChecks, result.HealthChecks)

}

func Test_UpstreamsCreateInvalid(t *testing.T) {
	upstreamRequest := &UpstreamRequest{
		Name:         "upstream-" + uuid.NewV4().String(),
		Slots:        2,
		HealthChecks: &UpstreamHealthCheck{},
	}

	result, err := NewClient(NewDefaultConfig()).Upstreams().Create(upstreamRequest)

	assert.NotNil(t, err)
	assert.Nil(t, result)

}

func Test_UpstreamsList(t *testing.T) {
	client := NewClient(NewDefaultConfig())

	upstreamRequest := &UpstreamRequest{
		Name:  "upstream-" + uuid.NewV4().String(),
		Slots: 10,
	}

	createdUpstream, err := client.Upstreams().Create(upstreamRequest)
	assert.Nil(t, err)
	assert.NotNil(t, createdUpstream)

	results, err := client.Upstreams().List()

	assert.Nil(t, err)
	assert.True(t, results.Total > 0)
	assert.True(t, len(results.Results) > 0)

}

func Test_UpstreamsListFilteredById(t *testing.T) {
	client := NewClient(NewDefaultConfig())

	upstreamRequest := &UpstreamRequest{
		Name:  "upstream-" + uuid.NewV4().String(),
		Slots: 10,
	}

	createdUpstream, err := client.Upstreams().Create(upstreamRequest)
	assert.Nil(t, err)
	assert.NotNil(t, createdUpstream)

	upstreamRequest2 := &UpstreamRequest{
		Name:  "upstream-" + uuid.NewV4().String(),
		Slots: 10,
	}
	createdUpstream2, err := client.Upstreams().Create(upstreamRequest2)
	assert.Nil(t, err)
	assert.NotNil(t, createdUpstream2)

	results, err := client.Upstreams().ListFiltered(&UpstreamFilter{Id: createdUpstream.Id})

	assert.Nil(t, err)
	assert.NotNil(t, results)
	assert.Equal(t, len(results.Results), 1)
	assert.Equal(t, createdUpstream, results.Results[0])

}

func Test_UpstreamsListFilteredByName(t *testing.T) {
	client := NewClient(NewDefaultConfig())

	upstreamRequest := &UpstreamRequest{
		Name:  "upstream-" + uuid.NewV4().String(),
		Slots: 10,
	}

	createdUpstream, err := client.Upstreams().Create(upstreamRequest)
	assert.Nil(t, err)
	assert.NotNil(t, createdUpstream)

	upstreamRequest2 := &UpstreamRequest{
		Name:  "upstream-" + uuid.NewV4().String(),
		Slots: 10,
	}
	createdUpstream2, err := client.Upstreams().Create(upstreamRequest2)
	assert.Nil(t, err)
	assert.NotNil(t, createdUpstream2)

	results, err := client.Upstreams().ListFiltered(&UpstreamFilter{Name: createdUpstream.Name})

	assert.Nil(t, err)
	assert.NotNil(t, results)
	assert.Equal(t, len(results.Results), 1)
	assert.Equal(t, createdUpstream, results.Results[0])

}

func Test_UpstreamsListFilteredBySlots(t *testing.T) {
	client := NewClient(NewDefaultConfig())

	upstreamRequest := &UpstreamRequest{
		Name:  "upstream-" + uuid.NewV4().String(),
		Slots: 111,
	}

	createdUpstream, err := client.Upstreams().Create(upstreamRequest)
	assert.Nil(t, err)
	assert.NotNil(t, createdUpstream)

	upstreamRequest2 := &UpstreamRequest{
		Name:  "upstream-" + uuid.NewV4().String(),
		Slots: 10,
	}
	createdUpstream2, err := client.Upstreams().Create(upstreamRequest2)
	assert.Nil(t, err)
	assert.NotNil(t, createdUpstream2)

	results, err := client.Upstreams().ListFiltered(&UpstreamFilter{Slots: 111})

	assert.Nil(t, err)
	assert.NotNil(t, results)
	assert.Equal(t, len(results.Results), 1)
	assert.Equal(t, createdUpstream, results.Results[0])

}

func Test_UpstreamsListFilteredBySize(t *testing.T) {
	client := NewClient(NewDefaultConfig())

	for i := 0; i < 5; i++ {
		upstreamRequest := &UpstreamRequest{
			Name:  "upstream-" + uuid.NewV4().String(),
			Slots: 10,
		}

		createdUpstream, err := client.Upstreams().Create(upstreamRequest)
		assert.Nil(t, err)
		assert.NotNil(t, createdUpstream)
	}

	results, err := client.Upstreams().ListFiltered(&UpstreamFilter{Size: 3})

	assert.Nil(t, err)
	assert.NotNil(t, results)
	assert.Equal(t, 3, len(results.Results))

}

func Test_UpstreamsDeleteById(t *testing.T) {

	client := NewClient(NewDefaultConfig())

	upstreamRequest := &UpstreamRequest{
		Name:         "upstream-" + uuid.NewV4().String(),
		Slots:        10,
		HealthChecks: &UpstreamHealthCheck{},
	}

	createdUpstream, err := client.Upstreams().Create(upstreamRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdUpstream)

	err = client.Upstreams().DeleteById(createdUpstream.Id)
	assert.Nil(t, err)

	result, err := client.Upstreams().GetById(createdUpstream.Id)
	assert.Nil(t, err)
	assert.Nil(t, result)

}

func Test_UpstreamsDeleteByName(t *testing.T) {

	client := NewClient(NewDefaultConfig())

	upstreamRequest := &UpstreamRequest{
		Name:         "upstream-" + uuid.NewV4().String(),
		Slots:        10,
		HealthChecks: &UpstreamHealthCheck{},
	}

	createdUpstream, err := client.Upstreams().Create(upstreamRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdUpstream)

	err = client.Upstreams().DeleteByName(createdUpstream.Name)
	assert.Nil(t, err)

	result, err := client.Upstreams().GetById(createdUpstream.Id)
	assert.Nil(t, err)
	assert.Nil(t, result)

}

func Test_UpstreamsUpdateById(t *testing.T) {

	client := NewClient(NewDefaultConfig())

	upstreamRequest := &UpstreamRequest{
		Name:  "upstream-" + uuid.NewV4().String(),
		Slots: 10,
		HealthChecks: &UpstreamHealthCheck{
			Active: &UpstreamHealthCheckActive{
				Concurrency: 10,
				HttpPath:    "/",
				Timeout:     1,
				Healthy: &ActiveHealthy{
					HttpStatuses: []int{200, 302},
					Interval:     0,
					Successes:    0,
				},
				Unhealthy: &ActiveUnhealthy{
					HttpFailures: 0,
					HttpStatuses: []int{429, 404, 500, 501, 502, 503, 504, 505},
					Interval:     0,
					TcpFailures:  0,
					Timeouts:     0,
				},
			},
			Passive: &UpstreamHealthCheckPassive{
				Healthy: &PassiveHealthy{
					HttpStatuses: []int{200, 201, 202, 203, 204, 205, 206, 207, 208, 226, 300, 301, 302, 303, 304, 305, 306, 307, 308},
					Successes:    0,
				},
				Unhealthy: &PassiveUnhealthy{
					HttpFailures: 0,
					HttpStatuses: []int{429, 500, 503},
					TcpFailures:  0,
					Timeouts:     0,
				},
			},
		},
	}

	createdUpstream, err := client.Upstreams().Create(upstreamRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdUpstream)
	assert.Equal(t, 10, createdUpstream.Slots)

	upstreamRequest.Slots = 11

	result, err := client.Upstreams().UpdateById(createdUpstream.Id, upstreamRequest)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, upstreamRequest.Name, result.Name)
	assert.Equal(t, upstreamRequest.Slots, result.Slots)
	assert.Equal(t, upstreamRequest.HealthChecks, result.HealthChecks)

}

func Test_UpstreamsUpdateByName(t *testing.T) {

	client := NewClient(NewDefaultConfig())

	upstreamRequest := &UpstreamRequest{
		Name:  "upstream-" + uuid.NewV4().String(),
		Slots: 10,
		HealthChecks: &UpstreamHealthCheck{
			Active: &UpstreamHealthCheckActive{
				Concurrency: 10,
				HttpPath:    "/",
				Timeout:     1,
				Healthy: &ActiveHealthy{
					HttpStatuses: []int{200, 302},
					Interval:     0,
					Successes:    0,
				},
				Unhealthy: &ActiveUnhealthy{
					HttpFailures: 0,
					HttpStatuses: []int{429, 404, 500, 501, 502, 503, 504, 505},
					Interval:     0,
					TcpFailures:  0,
					Timeouts:     0,
				},
			},
			Passive: &UpstreamHealthCheckPassive{
				Healthy: &PassiveHealthy{
					HttpStatuses: []int{200, 201, 202, 203, 204, 205, 206, 207, 208, 226, 300, 301, 302, 303, 304, 305, 306, 307, 308},
					Successes:    0,
				},
				Unhealthy: &PassiveUnhealthy{
					HttpFailures: 0,
					HttpStatuses: []int{429, 500, 503},
					TcpFailures:  0,
					Timeouts:     0,
				},
			},
		},
	}

	createdUpstream, err := client.Upstreams().Create(upstreamRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdUpstream)
	assert.Equal(t, 10, createdUpstream.Slots)

	upstreamRequest.Slots = 12

	result, err := client.Upstreams().UpdateByName(createdUpstream.Name, upstreamRequest)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, upstreamRequest.Name, result.Name)
	assert.Equal(t, upstreamRequest.Slots, result.Slots)
	assert.Equal(t, upstreamRequest.HealthChecks, result.HealthChecks)

}

func Test_UpstreamsUpdateByIdInvalid(t *testing.T) {

	client := NewClient(NewDefaultConfig())

	upstreamRequest := &UpstreamRequest{
		Name:  "upstream-" + uuid.NewV4().String(),
		Slots: 10,
	}

	createdUpstream, err := client.Upstreams().Create(upstreamRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdUpstream)
	assert.Equal(t, 10, createdUpstream.Slots)

	upstreamRequest.Slots = 2

	result, err := client.Upstreams().UpdateById(createdUpstream.Id, upstreamRequest)

	assert.NotNil(t, err)
	assert.Nil(t, result)

}

func Test_UpstreamsUpdateByNameInvalid(t *testing.T) {

	client := NewClient(NewDefaultConfig())

	upstreamRequest := &UpstreamRequest{
		Name:  "upstream-" + uuid.NewV4().String(),
		Slots: 10,
	}

	createdUpstream, err := client.Upstreams().Create(upstreamRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdUpstream)
	assert.Equal(t, 10, createdUpstream.Slots)

	upstreamRequest.Slots = 2

	result, err := client.Upstreams().UpdateByName(createdUpstream.Name, upstreamRequest)

	assert.NotNil(t, err)
	assert.Nil(t, result)

}

func Test_AllUpstreamEndpointsShouldReturnErrorWhenRequestUnauthorised(t *testing.T) {

	apiRequest := &ApiRequest{
		Name:        String("admin-api"),
		Uris:        StringSlice([]string{"/admin-api"}),
		UpstreamUrl: String("http://localhost:8001"),
	}

	client := NewClient(NewDefaultConfig())
	createdApi, err := client.Apis().Create(apiRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdApi)

	consumerRequest := &ConsumerRequest{
		Username: "username-" + uuid.NewV4().String(),
		CustomId: "test-" + uuid.NewV4().String(),
	}

	createdConsumer, err := client.Consumers().Create(consumerRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdConsumer)

	pluginRequest := &PluginRequest{
		Name:  "key-auth",
		ApiId: *createdApi.Id,
		Config: map[string]interface{}{
			"hide_credentials": true,
		},
	}

	createdPlugin, err := client.Plugins().Create(pluginRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdPlugin)

	_, err = client.Consumers().CreatePluginConfig(createdConsumer.Id, "key-auth", "")
	assert.Nil(t, err)

	upstreamRequest := &UpstreamRequest{
		Name:  "upstream-" + uuid.NewV4().String(),
		Slots: 10,
	}

	createdUpstream, err := client.Upstreams().Create(upstreamRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdUpstream)

	kongApiAddress := os.Getenv(EnvKongApiHostAddress) + "/admin-api"
	unauthorisedClient := NewClient(&Config{HostAddress: kongApiAddress})

	upstream, err := unauthorisedClient.Upstreams().GetByName(createdUpstream.Name)
	assert.NotNil(t, err)
	assert.Nil(t, upstream)

	upstream, err = unauthorisedClient.Upstreams().GetById(createdUpstream.Id)
	assert.NotNil(t, err)
	assert.Nil(t, upstream)

	results, err := unauthorisedClient.Upstreams().List()
	assert.NotNil(t, err)
	assert.Nil(t, results)

	err = unauthorisedClient.Upstreams().DeleteByName(createdUpstream.Name)
	assert.NotNil(t, err)

	err = unauthorisedClient.Upstreams().DeleteById(createdUpstream.Id)
	assert.NotNil(t, err)

	upstreamResult, err := unauthorisedClient.Upstreams().Create(&UpstreamRequest{
		Name:  "upstream-" + uuid.NewV4().String(),
		Slots: 10,
	})
	assert.Nil(t, upstreamResult)
	assert.NotNil(t, err)

	updatedUpstream, err := unauthorisedClient.Upstreams().UpdateByName(createdUpstream.Name, &UpstreamRequest{
		Name:  "upstream-" + uuid.NewV4().String(),
		Slots: 10,
	})
	assert.Nil(t, updatedUpstream)
	assert.NotNil(t, err)

	updatedUpstream, err = unauthorisedClient.Upstreams().UpdateById(createdUpstream.Id, &UpstreamRequest{
		Name:  "upstream-" + uuid.NewV4().String(),
		Slots: 10,
	})
	assert.Nil(t, updatedUpstream)
	assert.NotNil(t, err)

	err = client.Plugins().DeleteById(createdPlugin.Id)
	assert.Nil(t, err)

	err = client.Apis().DeleteById(*createdApi.Id)
	assert.Nil(t, err)

}
