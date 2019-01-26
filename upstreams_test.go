package gokong

import (
	"testing"

	uuid "github.com/satori/go.uuid"
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
	assert.True(t, len(results.Results) > 0)

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

	unauthorisedClient := NewClient(&Config{HostAddress: kong401Server})

	upstream, err := unauthorisedClient.Upstreams().GetByName("foo")
	assert.NotNil(t, err)
	assert.Nil(t, upstream)

	upstream, err = unauthorisedClient.Upstreams().GetById(uuid.NewV4().String())
	assert.NotNil(t, err)
	assert.Nil(t, upstream)

	results, err := unauthorisedClient.Upstreams().List()
	assert.NotNil(t, err)
	assert.Nil(t, results)

	err = unauthorisedClient.Upstreams().DeleteByName("bar")
	assert.NotNil(t, err)

	err = unauthorisedClient.Upstreams().DeleteById(uuid.NewV4().String())
	assert.NotNil(t, err)

	upstreamResult, err := unauthorisedClient.Upstreams().Create(&UpstreamRequest{
		Name:  "upstream-" + uuid.NewV4().String(),
		Slots: 10,
	})
	assert.Nil(t, upstreamResult)
	assert.NotNil(t, err)

	updatedUpstream, err := unauthorisedClient.Upstreams().UpdateByName("foo", &UpstreamRequest{
		Name:  "upstream-" + uuid.NewV4().String(),
		Slots: 10,
	})
	assert.Nil(t, updatedUpstream)
	assert.NotNil(t, err)

	updatedUpstream, err = unauthorisedClient.Upstreams().UpdateById(uuid.NewV4().String(), &UpstreamRequest{
		Name:  "upstream-" + uuid.NewV4().String(),
		Slots: 10,
	})
	assert.Nil(t, updatedUpstream)
	assert.NotNil(t, err)

}
