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
	}
	createdTarget, err := client.Targets().CreateFromUpstreamId(createdUpstream.Id, targetRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdTarget)
	assert.True(t, *createdTarget.Id != "")
	assert.Equal(t, *createdTarget.Target, targetRequest.Target)
	assert.Equal(t, *createdTarget.Weight, targetRequest.Weight)
	assert.Equal(t, *createdTarget.UpstreamId, createdUpstream.Id)

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
