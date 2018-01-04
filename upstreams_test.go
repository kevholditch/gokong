package gokong

import (
	"testing"

	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func Test_UpstreamsGetByID(t *testing.T) {

	upstreamRequest := &UpstreamRequest{
		Name:  "upstream-" + uuid.NewV4().String(),
		Slots: 10,
	}

	client := NewClient(NewDefaultConfig())
	createdUpstream, err := client.Upstreams().Create(upstreamRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdUpstream)

	result, err := client.Upstreams().GetByID(createdUpstream.ID)

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

func Test_UpstreamsGetByIDForNonExistentUpstreamByID(t *testing.T) {

	result, err := NewClient(NewDefaultConfig()).Upstreams().GetByID(uuid.NewV4().String())

	assert.Nil(t, err)
	assert.Nil(t, result)

}

func Test_UpstreamsGetByIDForNonExistentUpstreamByName(t *testing.T) {

	result, err := NewClient(NewDefaultConfig()).Upstreams().GetByName(uuid.NewV4().String())

	assert.Nil(t, err)
	assert.Nil(t, result)

}

func Test_UpstreamsCreate(t *testing.T) {
	upstreamRequest := &UpstreamRequest{
		Name:      "upstream-" + uuid.NewV4().String(),
		Slots:     10,
		OrderList: []int{2, 1, 3, 4, 5, 6, 7, 8, 9, 10},
	}

	result, err := NewClient(NewDefaultConfig()).Upstreams().Create(upstreamRequest)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.True(t, result.ID != "")
	assert.Equal(t, upstreamRequest.Name, result.Name)
	assert.Equal(t, upstreamRequest.Slots, result.Slots)
	assert.Equal(t, upstreamRequest.OrderList, result.OrderList)

}

func Test_UpstreamsCreateInvalid(t *testing.T) {
	upstreamRequest := &UpstreamRequest{
		Name:      "upstream-" + uuid.NewV4().String(),
		Slots:     2,
		OrderList: []int{2, 1, 3, 4, 5, 6, 7, 8, 9, 10},
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

func Test_UpstreamsListFilteredByID(t *testing.T) {
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

	results, err := client.Upstreams().ListFiltered(&UpstreamFilter{ID: createdUpstream.ID})

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

func Test_UpstreamsDeleteByID(t *testing.T) {

	client := NewClient(NewDefaultConfig())

	upstreamRequest := &UpstreamRequest{
		Name:      "upstream-" + uuid.NewV4().String(),
		Slots:     10,
		OrderList: []int{2, 1, 3, 4, 5, 6, 7, 8, 9, 10},
	}

	createdUpstream, err := client.Upstreams().Create(upstreamRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdUpstream)

	err = client.Upstreams().DeleteByID(createdUpstream.ID)
	assert.Nil(t, err)

	result, err := client.Upstreams().GetByID(createdUpstream.ID)
	assert.Nil(t, err)
	assert.Nil(t, result)

}

func Test_UpstreamsDeleteByName(t *testing.T) {

	client := NewClient(NewDefaultConfig())

	upstreamRequest := &UpstreamRequest{
		Name:      "upstream-" + uuid.NewV4().String(),
		Slots:     10,
		OrderList: []int{2, 1, 3, 4, 5, 6, 7, 8, 9, 10},
	}

	createdUpstream, err := client.Upstreams().Create(upstreamRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdUpstream)

	err = client.Upstreams().DeleteByName(createdUpstream.Name)
	assert.Nil(t, err)

	result, err := client.Upstreams().GetByID(createdUpstream.ID)
	assert.Nil(t, err)
	assert.Nil(t, result)

}

func Test_UpstreamsUpdateByID(t *testing.T) {

	client := NewClient(NewDefaultConfig())

	upstreamRequest := &UpstreamRequest{
		Name:  "upstream-" + uuid.NewV4().String(),
		Slots: 10,
	}

	createdUpstream, err := client.Upstreams().Create(upstreamRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdUpstream)
	assert.Equal(t, 10, createdUpstream.Slots)

	upstreamRequest.Slots = 11
	upstreamRequest.OrderList = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	result, err := client.Upstreams().UpdateByID(createdUpstream.ID, upstreamRequest)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, upstreamRequest.Name, result.Name)
	assert.Equal(t, upstreamRequest.Slots, result.Slots)
	assert.Equal(t, upstreamRequest.OrderList, result.OrderList)

}

func Test_UpstreamsUpdateByName(t *testing.T) {

	client := NewClient(NewDefaultConfig())

	upstreamRequest := &UpstreamRequest{
		Name:  "upstream-" + uuid.NewV4().String(),
		Slots: 10,
	}

	createdUpstream, err := client.Upstreams().Create(upstreamRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdUpstream)
	assert.Equal(t, 10, createdUpstream.Slots)

	upstreamRequest.Slots = 12
	upstreamRequest.OrderList = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

	result, err := client.Upstreams().UpdateByName(createdUpstream.Name, upstreamRequest)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, upstreamRequest.Name, result.Name)
	assert.Equal(t, upstreamRequest.Slots, result.Slots)
	assert.Equal(t, upstreamRequest.OrderList, result.OrderList)

}

func Test_UpstreamsUpdateByIDInvalid(t *testing.T) {

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

	result, err := client.Upstreams().UpdateByID(createdUpstream.ID, upstreamRequest)

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
