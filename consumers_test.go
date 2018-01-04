package gokong

import (
	"testing"

	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func Test_ConsumersGetByID(t *testing.T) {
	consumerRequest := &ConsumerRequest{
		Username: "username-" + uuid.NewV4().String(),
		CustomID: "test-" + uuid.NewV4().String(),
	}

	client := NewClient(NewDefaultConfig())
	createdConsumer, err := client.Consumers().Create(consumerRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdConsumer)

	result, err := client.Consumers().GetByID(createdConsumer.ID)

	assert.Equal(t, createdConsumer, result)

}

func Test_ConsumersGetNonExistentByID(t *testing.T) {

	result, err := NewClient(NewDefaultConfig()).Consumers().GetByID("7c924010-fca4-4314-8a3f-725cf749eac6")

	assert.Nil(t, err)
	assert.Nil(t, result)
}

func Test_ConsumersGetByUsername(t *testing.T) {
	consumerRequest := &ConsumerRequest{
		Username: "username-" + uuid.NewV4().String(),
		CustomID: "test-" + uuid.NewV4().String(),
	}

	client := NewClient(NewDefaultConfig())
	createdConsumer, err := client.Consumers().Create(consumerRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdConsumer)

	result, err := client.Consumers().GetByUsername(createdConsumer.Username)

	assert.Equal(t, createdConsumer, result)

}

func Test_ConsumersGetNonExistentByUsername(t *testing.T) {

	result, err := NewClient(NewDefaultConfig()).Consumers().GetByID("408b5b13-b7c0-4ffd-afa1-aea957f00252")

	assert.Nil(t, err)
	assert.Nil(t, result)
}

func Test_ConsumersCreate(t *testing.T) {
	consumerRequest := &ConsumerRequest{
		Username: "username-" + uuid.NewV4().String(),
		CustomID: "test-" + uuid.NewV4().String(),
	}

	result, err := NewClient(NewDefaultConfig()).Consumers().Create(consumerRequest)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, consumerRequest.Username, result.Username)
	assert.Equal(t, consumerRequest.CustomID, result.CustomID)

}

func Test_ConsumersCreateInvalid(t *testing.T) {
	consumerRequest := &ConsumerRequest{}

	result, err := NewClient(NewDefaultConfig()).Consumers().Create(consumerRequest)

	assert.NotNil(t, err)
	assert.Nil(t, result)

}

func Test_ConsumersCreateOnlyUsername(t *testing.T) {
	consumerRequest := &ConsumerRequest{
		Username: "username-" + uuid.NewV4().String(),
	}

	result, err := NewClient(NewDefaultConfig()).Consumers().Create(consumerRequest)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, consumerRequest.Username, result.Username)
	assert.Equal(t, "", result.CustomID)

}

func Test_ConsumersCreateOnlyCustomID(t *testing.T) {
	consumerRequest := &ConsumerRequest{
		CustomID: "test-" + uuid.NewV4().String(),
	}

	result, err := NewClient(NewDefaultConfig()).Consumers().Create(consumerRequest)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "", result.Username)
	assert.Equal(t, consumerRequest.CustomID, result.CustomID)

}

func Test_ConsumersList(t *testing.T) {

	consumerRequest := &ConsumerRequest{
		Username: "username-" + uuid.NewV4().String(),
		CustomID: "test-" + uuid.NewV4().String(),
	}

	client := NewClient(NewDefaultConfig())
	createdConsumer, err := client.Consumers().Create(consumerRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdConsumer)

	results, err := client.Consumers().List()

	assert.True(t, results.Total > 0)
	assert.True(t, len(results.Results) > 0)

}

func Test_ConsumersListFilteredByID(t *testing.T) {

	consumerRequest := &ConsumerRequest{
		Username: "username-" + uuid.NewV4().String(),
		CustomID: "test-" + uuid.NewV4().String(),
	}

	client := NewClient(NewDefaultConfig())
	createdConsumer, err := client.Consumers().Create(consumerRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdConsumer)

	consumerRequest2 := &ConsumerRequest{
		Username: "username-" + uuid.NewV4().String(),
		CustomID: "test-" + uuid.NewV4().String(),
	}
	createdConsumer2, err := client.Consumers().Create(consumerRequest2)

	assert.Nil(t, err)
	assert.NotNil(t, createdConsumer2)

	results, err := client.Consumers().ListFiltered(&ConsumerFilter{ID: createdConsumer.ID})

	assert.True(t, len(results.Results) == 1)
	result := results.Results[0]

	assert.Equal(t, createdConsumer.ID, result.ID)
	assert.Equal(t, createdConsumer.CustomID, result.CustomID)
	assert.Equal(t, createdConsumer.Username, result.Username)

}

func Test_ConsumersListFilteredByCustomID(t *testing.T) {

	consumerRequest := &ConsumerRequest{
		Username: "username-" + uuid.NewV4().String(),
		CustomID: "test-" + uuid.NewV4().String(),
	}

	client := NewClient(NewDefaultConfig())
	createdConsumer, err := client.Consumers().Create(consumerRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdConsumer)

	consumerRequest2 := &ConsumerRequest{
		Username: "username-" + uuid.NewV4().String(),
		CustomID: "test-" + uuid.NewV4().String(),
	}
	createdConsumer2, err := client.Consumers().Create(consumerRequest2)

	assert.Nil(t, err)
	assert.NotNil(t, createdConsumer2)

	results, err := client.Consumers().ListFiltered(&ConsumerFilter{CustomID: createdConsumer.CustomID})

	assert.True(t, len(results.Results) == 1)
	result := results.Results[0]

	assert.Equal(t, createdConsumer.ID, result.ID)
	assert.Equal(t, createdConsumer.CustomID, result.CustomID)
	assert.Equal(t, createdConsumer.Username, result.Username)

}

func Test_ConsumersListFilteredByUsername(t *testing.T) {

	consumerRequest := &ConsumerRequest{
		Username: "username-" + uuid.NewV4().String(),
		CustomID: "test-" + uuid.NewV4().String(),
	}

	client := NewClient(NewDefaultConfig())
	createdConsumer, err := client.Consumers().Create(consumerRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdConsumer)

	consumerRequest2 := &ConsumerRequest{
		Username: "username-" + uuid.NewV4().String(),
		CustomID: "test-" + uuid.NewV4().String(),
	}
	createdConsumer2, err := client.Consumers().Create(consumerRequest2)

	assert.Nil(t, err)
	assert.NotNil(t, createdConsumer2)

	results, err := client.Consumers().ListFiltered(&ConsumerFilter{Username: createdConsumer.Username})

	assert.True(t, len(results.Results) == 1)
	result := results.Results[0]

	assert.Equal(t, createdConsumer.ID, result.ID)
	assert.Equal(t, createdConsumer.CustomID, result.CustomID)
	assert.Equal(t, createdConsumer.Username, result.Username)

}

func Test_ConsumersListFilteredBySize(t *testing.T) {

	consumerRequest := &ConsumerRequest{
		Username: "username-" + uuid.NewV4().String(),
		CustomID: "test-" + uuid.NewV4().String(),
	}

	client := NewClient(NewDefaultConfig())
	createdConsumer, err := client.Consumers().Create(consumerRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdConsumer)

	consumerRequest2 := &ConsumerRequest{
		Username: "username-" + uuid.NewV4().String(),
		CustomID: "test-" + uuid.NewV4().String(),
	}
	createdConsumer2, err := client.Consumers().Create(consumerRequest2)

	assert.Nil(t, err)
	assert.NotNil(t, createdConsumer2)

	consumerRequest3 := &ConsumerRequest{
		Username: "username-" + uuid.NewV4().String(),
		CustomID: "test-" + uuid.NewV4().String(),
	}
	createdConsumer3, err := client.Consumers().Create(consumerRequest3)

	assert.Nil(t, err)
	assert.NotNil(t, createdConsumer3)

	results, err := client.Consumers().ListFiltered(&ConsumerFilter{Size: 2})

	assert.True(t, len(results.Results) == 2)

}

func Test_ConsumersDeleteByID(t *testing.T) {
	consumerRequest := &ConsumerRequest{
		Username: "username-" + uuid.NewV4().String(),
		CustomID: "test-" + uuid.NewV4().String(),
	}

	client := NewClient(NewDefaultConfig())
	createdConsumer, err := client.Consumers().Create(consumerRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdConsumer)

	err = client.Consumers().DeleteByID(createdConsumer.ID)
	assert.Nil(t, err)

	deletedConsumer, err := client.Consumers().GetByID(createdConsumer.ID)
	assert.Nil(t, deletedConsumer)
}

func Test_ConsumersDeleteByUsername(t *testing.T) {
	consumerRequest := &ConsumerRequest{
		Username: "username-" + uuid.NewV4().String(),
		CustomID: "test-" + uuid.NewV4().String(),
	}

	client := NewClient(NewDefaultConfig())
	createdConsumer, err := client.Consumers().Create(consumerRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdConsumer)

	err = client.Consumers().DeleteByUsername(createdConsumer.Username)
	assert.Nil(t, err)

	deletedConsumer, err := client.Consumers().GetByID(createdConsumer.ID)
	assert.Nil(t, deletedConsumer)
}

func Test_ConsumersUpdateByID(t *testing.T) {

	consumerRequest := &ConsumerRequest{
		Username: "username-" + uuid.NewV4().String(),
		CustomID: "test-" + uuid.NewV4().String(),
	}

	client := NewClient(NewDefaultConfig())
	createdConsumer, err := client.Consumers().Create(consumerRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdConsumer)
	assert.Equal(t, consumerRequest.CustomID, createdConsumer.CustomID)

	consumerRequest.CustomID = "test-" + uuid.NewV4().String()

	result, err := client.Consumers().UpdateByID(createdConsumer.ID, consumerRequest)

	assert.Nil(t, err)
	assert.NotNil(t, result)

	assert.Equal(t, consumerRequest.CustomID, result.CustomID)
	assert.Equal(t, consumerRequest.Username, result.Username)
}

func Test_ConsumersUpdateByIDInvalid(t *testing.T) {

	consumerRequest := &ConsumerRequest{
		Username: "username-" + uuid.NewV4().String(),
		CustomID: "test-" + uuid.NewV4().String(),
	}

	client := NewClient(NewDefaultConfig())
	createdConsumer, err := client.Consumers().Create(consumerRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdConsumer)
	assert.Equal(t, consumerRequest.CustomID, createdConsumer.CustomID)

	consumerRequest.Username = ""
	consumerRequest.CustomID = ""

	result, err := client.Consumers().UpdateByID(createdConsumer.ID, consumerRequest)

	assert.NotNil(t, err)
	assert.Nil(t, result)
}

func Test_ConsumersUpdateByUsername(t *testing.T) {

	consumerRequest := &ConsumerRequest{
		Username: "username-" + uuid.NewV4().String(),
		CustomID: "test-" + uuid.NewV4().String(),
	}

	client := NewClient(NewDefaultConfig())
	createdConsumer, err := client.Consumers().Create(consumerRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdConsumer)
	assert.Equal(t, consumerRequest.CustomID, createdConsumer.CustomID)

	consumerRequest.Username = "username-" + uuid.NewV4().String()

	result, err := client.Consumers().UpdateByUsername(createdConsumer.Username, consumerRequest)

	assert.Nil(t, err)
	assert.NotNil(t, result)

	assert.Equal(t, consumerRequest.CustomID, result.CustomID)
	assert.Equal(t, consumerRequest.Username, result.Username)
}

func Test_ConsumersUpdateByUsernameInvalid(t *testing.T) {

	consumerRequest := &ConsumerRequest{
		Username: "username-" + uuid.NewV4().String(),
		CustomID: "test-" + uuid.NewV4().String(),
	}

	client := NewClient(NewDefaultConfig())
	createdConsumer, err := client.Consumers().Create(consumerRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdConsumer)
	assert.Equal(t, consumerRequest.CustomID, createdConsumer.CustomID)

	consumerRequest.Username = ""
	consumerRequest.CustomID = ""

	result, err := client.Consumers().UpdateByUsername(createdConsumer.Username, consumerRequest)

	assert.NotNil(t, err)
	assert.Nil(t, result)
}
