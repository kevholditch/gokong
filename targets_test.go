package gokong

import (
	"testing"
	// "time"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestTargets_GetTargetsFromUpstreamId(t *testing.T) {
	upstreamRequest := &UpstreamRequest{
		Name:  "upstream-" + uuid.NewV4().String(),
		Slots: 10,
	}

	client := NewClient(NewDefaultConfig())
	createdUpstream, err := client.Upstreams().Create(upstreamRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdUpstream)

	targetRequest := &TargetRequest{
		Target: "www.example.com:80",
		Weight: 200,
	}
	createdTarget, err := client.Targets().CreateFromUpstreamId(createdUpstream.Id, targetRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdTarget)

	result, err := client.Targets().GetTargetsFromUpstreamId(createdUpstream.Id)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result, 1)
	assert.Equal(t, createdTarget, result[0])

	client.Targets().DeleteFromUpstreamById(createdUpstream.Id, *createdTarget.Id)
	client.Upstreams().DeleteById(createdUpstream.Id)
}

func TestTargets_GetTargetsFromUpstreamName(t *testing.T) {
	upstreamRequest := &UpstreamRequest{
		Name:  "upstream-" + uuid.NewV4().String(),
		Slots: 10,
	}

	client := NewClient(NewDefaultConfig())
	createdUpstream, err := client.Upstreams().Create(upstreamRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdUpstream)

	targetRequest := &TargetRequest{
		Target: "www.example.com:80",
		Weight: 200,
	}
	createdTarget, err := client.Targets().CreateFromUpstreamName(createdUpstream.Name, targetRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdTarget)

	result, err := client.Targets().GetTargetsFromUpstreamName(createdUpstream.Name)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result, 1)
	assert.Equal(t, createdTarget, result[0])

	client.Targets().DeleteFromUpstreamById(createdUpstream.Name, *createdTarget.Target)
	client.Upstreams().DeleteById(createdUpstream.Id)
}

func TestTargets_GetForNonExistentUpstream(t *testing.T) {
	result, err := NewClient(NewDefaultConfig()).Targets().GetTargetsFromUpstreamName(uuid.NewV4().String())

	assert.Nil(t, result)
	assert.NotNil(t, err)
}

func TestTargets_Create(t *testing.T) {
	upstreamRequest := &UpstreamRequest{
		Name:  "upstream-" + uuid.NewV4().String(),
		Slots: 10,
	}

	client := NewClient(NewDefaultConfig())
	createdUpstream, err := client.Upstreams().Create(upstreamRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdUpstream)

	targetRequest := &TargetRequest{
		Target: "www.example.com:80",
		Weight: 200,
		Tags:   []*string{String("my-tag")},
	}
	createdTarget, err := client.Targets().CreateFromUpstreamId(createdUpstream.Id, targetRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdTarget)
	assert.True(t, *createdTarget.Id != "")
	assert.Equal(t, *createdTarget.Target, targetRequest.Target)
	assert.Equal(t, *createdTarget.Weight, targetRequest.Weight)
	assert.Equal(t, createdTarget.Tags, targetRequest.Tags)

	client.Targets().DeleteFromUpstreamById(createdUpstream.Id, *createdTarget.Id)
	client.Upstreams().DeleteById(createdUpstream.Id)
}

func TestTargets_Delete(t *testing.T) {
	upstreamRequest := &UpstreamRequest{
		Name:  "upstream-" + uuid.NewV4().String(),
		Slots: 10,
	}

	client := NewClient(NewDefaultConfig())
	createdUpstream, err := client.Upstreams().Create(upstreamRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdUpstream)

	targetRequest := &TargetRequest{
		Target: "www.example.com:80",
		Weight: 200,
	}
	createdTarget, err := client.Targets().CreateFromUpstreamId(createdUpstream.Id, targetRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdTarget)

	client.Targets().DeleteFromUpstreamByHostPort(createdUpstream.Name, *createdTarget.Target)

	targets, err := client.Targets().GetTargetsFromUpstreamName(createdUpstream.Name)
	assert.Nil(t, err)
	assert.Len(t, targets, 0)

	client.Upstreams().DeleteById(createdUpstream.Id)
}

// WWOM: The following test runs locally without issue and without need for hack contained therein
// However, on the build server there seems to be a timing issue of sorts whereby trhe Kong container
// hasn't completed the registration of a target and/or related health checks when we attempt to
// manually set the health status (to healthy/unhealthy) and subsequently check that it was set
// func TestTargets_SetTargetHealthFromUpstreamById(t *testing.T) {
// 	upstreamRequest := &UpstreamRequest{
// 		Name:  "upstream-" + uuid.NewV4().String(),
// 		Slots: 10,
// 		HealthChecks: &UpstreamHealthCheck{
// 			Active: &UpstreamHealthCheckActive{
// 				Concurrency: 10,
// 				HttpPath:    "/",
// 				Timeout:     1,
// 				Healthy: &ActiveHealthy{
// 					HttpStatuses: []int{200, 302},
// 					Interval:     1000,
// 					Successes:    10,
// 				},
// 				Unhealthy: &ActiveUnhealthy{
// 					HttpFailures: 10,
// 					HttpStatuses: []int{429, 404, 500, 501, 502, 503, 504, 505},
// 					Interval:     1000,
// 					TcpFailures:  10,
// 					Timeouts:     10,
// 				},
// 			},
// 		},
// 	}
//
// 	client := NewClient(NewDefaultConfig())
// 	createdUpstream, err := client.Upstreams().Create(upstreamRequest)
//
// 	assert.Nil(t, err)
// 	assert.NotNil(t, createdUpstream)
//
// 	targetRequest := &TargetRequest{
// 		Target: "www.example.com:80",
// 		Weight: 200,
// 	}
// 	createdTarget, err := client.Targets().CreateFromUpstreamName(createdUpstream.Name, targetRequest)
//
// 	assert.Nil(t, err)
// 	assert.NotNil(t, createdTarget)
//
// 	// HACK: This is a bit hack-y - but tests fail on the build occassionly as Kong hasn't setup the load balancer
// 	// and health checks for the upstream/targets by the time we start trying to set their health status below
// 	targets, err := client.Targets().GetTargetsWithHealthFromUpstreamName(createdUpstream.Name)
// 	retry := 1
// 	for (*targets[0].Health == "HEALTHCHECKS_OFF" || *targets[0].Health == "DNS_ERROR") && retry < 10 {
// 		t.Log("Health-checks still off on target. Sleep and try again until we have another status.")
// 		assert.Len(t, targets, 1)
//
// 		time.Sleep(2 * time.Second)
// 		targets, err = client.Targets().GetTargetsWithHealthFromUpstreamName(createdUpstream.Name)
// 		retry++
// 	}
//
// 	result := client.Targets().SetTargetFromUpstreamByIdAsUnhealthy(createdUpstream.Id, *createdTarget.Id)
//
// 	assert.Nil(t, result)
//
// 	targets, err = client.Targets().GetTargetsWithHealthFromUpstreamName(createdUpstream.Name)
//
// 	assert.Nil(t, err)
// 	assert.NotNil(t, targets)
// 	assert.Len(t, targets, 1)
// 	target := targets[0]
// 	assert.Equal(t, "UNHEALTHY", *target.Health)
//
// 	result = client.Targets().SetTargetFromUpstreamByIdAsHealthy(createdUpstream.Id, *createdTarget.Id)
//
// 	assert.Nil(t, result)
//
// 	targets, err = client.Targets().GetTargetsWithHealthFromUpstreamName(createdUpstream.Name)
//
// 	assert.Nil(t, err)
// 	assert.NotNil(t, targets)
// 	assert.Len(t, targets, 1)
// 	target = targets[0]
// 	assert.Equal(t, "HEALTHY", *target.Health)
//
// 	client.Targets().DeleteFromUpstreamById(createdUpstream.Name, *createdTarget.Target)
// 	client.Upstreams().DeleteById(createdUpstream.Id)
// }
