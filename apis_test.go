package gokong

import (
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ApisGetById(t *testing.T) {
	apiRequest := &ApiRequest{
		Name:                   "test-" + uuid.NewV4().String(),
		Hosts:                  []string{"example.com"},
		Uris:                   []string{"/example"},
		Methods:                []string{"GET", "POST"},
		UpstreamUrl:            "http://localhost:4140/testservice",
		StripUri:               true,
		PreserveHost:           true,
		Retries:                3,
		UpstreamConnectTimeout: 1000,
		UpstreamSendTimeout:    2000,
		UpstreamReadTimeout:    3000,
		HttpsOnly:              true,
		HttpIfTerminated:       true,
	}

	apiClient := NewClient(NewDefaultConfig()).Apis()
	createdApi, err := apiClient.Create(apiRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdApi)

	result, err := apiClient.GetById(createdApi.Id)

	assert.Equal(t, createdApi, result)

}

func Test_ApisGetByName(t *testing.T) {
	apiRequest := &ApiRequest{
		Name:                   "test-" + uuid.NewV4().String(),
		Hosts:                  []string{"example.com"},
		Uris:                   []string{"/example"},
		Methods:                []string{"GET", "POST"},
		UpstreamUrl:            "http://localhost:4140/testservice",
		StripUri:               true,
		PreserveHost:           true,
		Retries:                3,
		UpstreamConnectTimeout: 1000,
		UpstreamSendTimeout:    2000,
		UpstreamReadTimeout:    3000,
		HttpsOnly:              true,
		HttpIfTerminated:       true,
	}

	apiClient := NewClient(NewDefaultConfig()).Apis()
	createdApi, err := apiClient.Create(apiRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdApi)

	result, err := apiClient.GetByName(createdApi.Name)

	assert.Equal(t, createdApi, result)

}

func Test_ApisGetAll(t *testing.T) {

	apiRequest := &ApiRequest{
		Name:                   "test-" + uuid.NewV4().String(),
		Hosts:                  []string{"another.com"},
		Uris:                   []string{"/another"},
		Methods:                []string{"PUT", "POST"},
		UpstreamUrl:            "http://linkerd:4140/myservice",
		StripUri:               false,
		PreserveHost:           false,
		Retries:                5,
		UpstreamConnectTimeout: 2222,
		UpstreamSendTimeout:    1233,
		UpstreamReadTimeout:    1234,
		HttpsOnly:              false,
		HttpIfTerminated:       false,
	}
	apiClient := NewClient(NewDefaultConfig()).Apis()
	createdApi, err := apiClient.Create(apiRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdApi)

	results, err := apiClient.GetAll()

	assert.True(t, results.Total > 0)
	assert.True(t, len(results.Results) > 0)

}

func Test_ApisGetAllFilteredById(t *testing.T) {

	apiRequest := &ApiRequest{
		Name:                   "test-" + uuid.NewV4().String(),
		Hosts:                  []string{"filter1.com"},
		Uris:                   []string{"/filter"},
		Methods:                []string{"PUT", "POST"},
		UpstreamUrl:            "http://linkerd:4140/myservice",
		StripUri:               false,
		PreserveHost:           false,
		Retries:                5,
		UpstreamConnectTimeout: 2222,
		UpstreamSendTimeout:    1233,
		UpstreamReadTimeout:    1234,
		HttpsOnly:              false,
		HttpIfTerminated:       false,
	}
	apiRequest2 := &ApiRequest{
		Name:                   "test-" + uuid.NewV4().String(),
		Hosts:                  []string{"filter2.com"},
		Uris:                   []string{"/filter2"},
		Methods:                []string{"PUT", "POST"},
		UpstreamUrl:            "http://linkerd:4140/myservice2",
		StripUri:               false,
		PreserveHost:           false,
		Retries:                5,
		UpstreamConnectTimeout: 2222,
		UpstreamSendTimeout:    1233,
		UpstreamReadTimeout:    1234,
		HttpsOnly:              false,
		HttpIfTerminated:       false,
	}
	apiClient := NewClient(NewDefaultConfig()).Apis()

	createdApi, err := apiClient.Create(apiRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdApi)

	createdApi2, err := apiClient.Create(apiRequest2)

	assert.Nil(t, err)
	assert.NotNil(t, createdApi2)

	results, err := apiClient.GetAllFiltered(&GetAllFilter{Id: createdApi2.Id})

	assert.True(t, results.Total == 1)
	assert.True(t, len(results.Results) == 1)

	result := results.Results[0]

	assert.Equal(t, createdApi2, result)

}

func Test_ApisGetAllFilteredByName(t *testing.T) {

	apiRequest := &ApiRequest{
		Name:                   "test-" + uuid.NewV4().String(),
		Hosts:                  []string{"filter1.com"},
		Uris:                   []string{"/filter"},
		Methods:                []string{"PUT", "POST"},
		UpstreamUrl:            "http://linkerd:4140/myservice",
		StripUri:               false,
		PreserveHost:           false,
		Retries:                5,
		UpstreamConnectTimeout: 2222,
		UpstreamSendTimeout:    1233,
		UpstreamReadTimeout:    1234,
		HttpsOnly:              false,
		HttpIfTerminated:       false,
	}
	apiRequest2 := &ApiRequest{
		Name:                   "test-" + uuid.NewV4().String(),
		Hosts:                  []string{"filter2.com"},
		Uris:                   []string{"/filter2"},
		Methods:                []string{"PUT", "POST"},
		UpstreamUrl:            "http://linkerd:4140/myservice2",
		StripUri:               false,
		PreserveHost:           false,
		Retries:                5,
		UpstreamConnectTimeout: 2222,
		UpstreamSendTimeout:    1233,
		UpstreamReadTimeout:    1234,
		HttpsOnly:              false,
		HttpIfTerminated:       false,
	}
	apiClient := NewClient(NewDefaultConfig()).Apis()

	createdApi, err := apiClient.Create(apiRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdApi)

	createdApi2, err := apiClient.Create(apiRequest2)

	assert.Nil(t, err)
	assert.NotNil(t, createdApi2)

	results, err := apiClient.GetAllFiltered(&GetAllFilter{Name: createdApi2.Name})

	assert.True(t, results.Total == 1)
	assert.True(t, len(results.Results) == 1)

	result := results.Results[0]

	assert.Equal(t, createdApi2, result)

}

func Test_ApisGetAllFilteredByUpstreamUrl(t *testing.T) {

	apiRequest := &ApiRequest{
		Name:                   "test-" + uuid.NewV4().String(),
		Hosts:                  []string{"filter1.com"},
		Uris:                   []string{"/filter"},
		Methods:                []string{"PUT", "POST"},
		UpstreamUrl:            "http://linkerd:4140/someurl",
		StripUri:               false,
		PreserveHost:           false,
		Retries:                5,
		UpstreamConnectTimeout: 2222,
		UpstreamSendTimeout:    1233,
		UpstreamReadTimeout:    1234,
		HttpsOnly:              false,
		HttpIfTerminated:       false,
	}
	apiRequest2 := &ApiRequest{
		Name:                   "test-" + uuid.NewV4().String(),
		Hosts:                  []string{"filter2.com"},
		Uris:                   []string{"/filter2"},
		Methods:                []string{"PUT", "POST"},
		UpstreamUrl:            "http://upstreamunique:4140/uniqueurl",
		StripUri:               false,
		PreserveHost:           false,
		Retries:                5,
		UpstreamConnectTimeout: 2222,
		UpstreamSendTimeout:    1233,
		UpstreamReadTimeout:    1234,
		HttpsOnly:              false,
		HttpIfTerminated:       false,
	}
	apiClient := NewClient(NewDefaultConfig()).Apis()

	createdApi, err := apiClient.Create(apiRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdApi)

	createdApi2, err := apiClient.Create(apiRequest2)

	assert.Nil(t, err)
	assert.NotNil(t, createdApi2)

	results, err := apiClient.GetAllFiltered(&GetAllFilter{UpstreamUrl: createdApi2.UpstreamUrl})

	assert.True(t, results.Total == 1)
	assert.True(t, len(results.Results) == 1)

	result := results.Results[0]

	assert.Equal(t, createdApi2, result)
}

func Test_ApisGetAllFilteredByRetries(t *testing.T) {

	apiRequest := &ApiRequest{
		Name:                   "test-" + uuid.NewV4().String(),
		Hosts:                  []string{"filter1.com"},
		Uris:                   []string{"/filter"},
		Methods:                []string{"PUT", "POST"},
		UpstreamUrl:            "http://linkerd:4140/someurl",
		StripUri:               false,
		PreserveHost:           false,
		Retries:                5,
		UpstreamConnectTimeout: 2222,
		UpstreamSendTimeout:    1233,
		UpstreamReadTimeout:    1234,
		HttpsOnly:              false,
		HttpIfTerminated:       false,
	}
	apiRequest2 := &ApiRequest{
		Name:                   "test-" + uuid.NewV4().String(),
		Hosts:                  []string{"filter2.com"},
		Uris:                   []string{"/filter2"},
		Methods:                []string{"PUT", "POST"},
		UpstreamUrl:            "http://linkerd:4140/uniqueurl",
		StripUri:               false,
		PreserveHost:           false,
		Retries:                1234,
		UpstreamConnectTimeout: 2222,
		UpstreamSendTimeout:    1233,
		UpstreamReadTimeout:    1234,
		HttpsOnly:              false,
		HttpIfTerminated:       false,
	}
	apiClient := NewClient(NewDefaultConfig()).Apis()

	createdApi, err := apiClient.Create(apiRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdApi)

	createdApi2, err := apiClient.Create(apiRequest2)

	assert.Nil(t, err)
	assert.NotNil(t, createdApi2)

	results, err := apiClient.GetAllFiltered(&GetAllFilter{Retries: createdApi2.Retries})

	assert.True(t, results.Total == 1)
	assert.True(t, len(results.Results) == 1)

	result := results.Results[0]

	assert.Equal(t, createdApi2, result)
}

func Test_ApisGetAllFilteredBySize(t *testing.T) {

	apiRequest := &ApiRequest{
		Name:                   "test-" + uuid.NewV4().String(),
		Hosts:                  []string{"filter1.com"},
		Uris:                   []string{"/filter"},
		Methods:                []string{"PUT", "POST"},
		UpstreamUrl:            "http://linkerd:4140/someurl",
		StripUri:               false,
		PreserveHost:           false,
		Retries:                5,
		UpstreamConnectTimeout: 2222,
		UpstreamSendTimeout:    1233,
		UpstreamReadTimeout:    1234,
		HttpsOnly:              false,
		HttpIfTerminated:       false,
	}
	apiRequest2 := &ApiRequest{
		Name:                   "test-" + uuid.NewV4().String(),
		Hosts:                  []string{"filter2.com"},
		Uris:                   []string{"/filter2"},
		Methods:                []string{"PUT", "POST"},
		UpstreamUrl:            "http://linkerd:4140/uniqueurl",
		StripUri:               false,
		PreserveHost:           false,
		Retries:                1234,
		UpstreamConnectTimeout: 2222,
		UpstreamSendTimeout:    1233,
		UpstreamReadTimeout:    1234,
		HttpsOnly:              false,
		HttpIfTerminated:       false,
	}
	apiClient := NewClient(NewDefaultConfig()).Apis()

	createdApi, err := apiClient.Create(apiRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdApi)

	createdApi2, err := apiClient.Create(apiRequest2)

	assert.Nil(t, err)
	assert.NotNil(t, createdApi2)

	results, err := apiClient.GetAllFiltered(&GetAllFilter{Size: 1})

	assert.True(t, len(results.Results) == 1)

}

func Test_ApisCreate(t *testing.T) {

	apiRequest := &ApiRequest{
		Name:                   "test-" + uuid.NewV4().String(),
		Hosts:                  []string{"example.com"},
		Uris:                   []string{"/example"},
		Methods:                []string{"GET", "POST"},
		UpstreamUrl:            "http://localhost:4140/testservice",
		StripUri:               false,
		PreserveHost:           true,
		Retries:                3,
		UpstreamConnectTimeout: 1000,
		UpstreamSendTimeout:    2000,
		UpstreamReadTimeout:    3000,
		HttpsOnly:              true,
		HttpIfTerminated:       false,
	}

	result, err := NewClient(NewDefaultConfig()).Apis().Create(apiRequest)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, apiRequest.Name, result.Name)
	assert.Equal(t, apiRequest.Hosts, result.Hosts)
	assert.Equal(t, apiRequest.Uris, result.Uris)
	assert.Equal(t, apiRequest.Methods, result.Methods)
	assert.Equal(t, apiRequest.UpstreamUrl, result.UpstreamUrl)
	assert.Equal(t, apiRequest.StripUri, result.StripUri)
	assert.Equal(t, apiRequest.PreserveHost, result.PreserveHost)
	assert.Equal(t, apiRequest.UpstreamConnectTimeout, result.UpstreamConnectTimeout)
	assert.Equal(t, apiRequest.UpstreamSendTimeout, result.UpstreamSendTimeout)
	assert.Equal(t, apiRequest.UpstreamReadTimeout, result.UpstreamReadTimeout)
	assert.Equal(t, apiRequest.HttpsOnly, result.HttpsOnly)
	assert.Equal(t, apiRequest.HttpIfTerminated, result.HttpIfTerminated)
}

func Test_ApisCreateWithOnlyRequiredFields(t *testing.T) {
	apiRequest := &ApiRequest{
		Name:        "test-" + uuid.NewV4().String(),
		Hosts:       []string{"example.com"},
		UpstreamUrl: "http://localhost:4140/testservice",
	}

	result, err := NewClient(NewDefaultConfig()).Apis().Create(apiRequest)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, apiRequest.Name, result.Name)
	assert.Equal(t, apiRequest.Hosts, result.Hosts)
	assert.Nil(t, result.Uris)
	assert.Nil(t, result.Methods)
	assert.Equal(t, apiRequest.UpstreamUrl, result.UpstreamUrl)
	assert.Equal(t, false, result.StripUri)
	assert.Equal(t, false, result.PreserveHost)
	assert.Equal(t, 60000, result.UpstreamConnectTimeout)
	assert.Equal(t, 60000, result.UpstreamSendTimeout)
	assert.Equal(t, 60000, result.UpstreamReadTimeout)
	assert.Equal(t, false, result.HttpsOnly)
	assert.Equal(t, false, result.HttpIfTerminated)
}

func Test_ApisDeleteById(t *testing.T) {

	apiRequest := &ApiRequest{
		Name:                   "test-" + uuid.NewV4().String(),
		Hosts:                  []string{"delete.com"},
		Uris:                   []string{"/delete"},
		Methods:                []string{"GET", "POST"},
		UpstreamUrl:            "http://localhost:4140/testservice",
		StripUri:               true,
		PreserveHost:           true,
		Retries:                3,
		UpstreamConnectTimeout: 1000,
		UpstreamSendTimeout:    2000,
		UpstreamReadTimeout:    3000,
		HttpsOnly:              true,
		HttpIfTerminated:       true,
	}

	apiClient := NewClient(NewDefaultConfig()).Apis()
	createdApi, err := apiClient.Create(apiRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdApi)

	err = apiClient.DeleteById(createdApi.Id)

	assert.Nil(t, err)

	deletedApi, err := apiClient.GetById(createdApi.Id)
	assert.Nil(t, err)
	assert.Nil(t, deletedApi)

}

func Test_ApisDeleteByName(t *testing.T) {

	apiRequest := &ApiRequest{
		Name:                   "test-" + uuid.NewV4().String(),
		Hosts:                  []string{"delete.com"},
		Uris:                   []string{"/delete"},
		Methods:                []string{"GET", "POST"},
		UpstreamUrl:            "http://localhost:4140/testservice",
		StripUri:               true,
		PreserveHost:           true,
		Retries:                3,
		UpstreamConnectTimeout: 1000,
		UpstreamSendTimeout:    2000,
		UpstreamReadTimeout:    3000,
		HttpsOnly:              true,
		HttpIfTerminated:       true,
	}

	apiClient := NewClient(NewDefaultConfig()).Apis()
	createdApi, err := apiClient.Create(apiRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdApi)

	err = apiClient.DeleteByName(createdApi.Id)

	assert.Nil(t, err)

	deletedApi, err := apiClient.GetById(createdApi.Id)
	assert.Nil(t, err)
	assert.Nil(t, deletedApi)

}

func Test_ApisUpdateApiById(t *testing.T) {

	apiRequest := &ApiRequest{
		Name:                   "test-" + uuid.NewV4().String(),
		Hosts:                  []string{"filter1.com"},
		Uris:                   []string{"/filter"},
		Methods:                []string{"PUT", "POST"},
		UpstreamUrl:            "http://linkerd:4140/someurl",
		StripUri:               false,
		PreserveHost:           false,
		Retries:                5,
		UpstreamConnectTimeout: 2222,
		UpstreamSendTimeout:    1233,
		UpstreamReadTimeout:    1234,
		HttpsOnly:              false,
		HttpIfTerminated:       false,
	}

	apiClient := NewClient(NewDefaultConfig()).Apis()

	createdApi, err := apiClient.Create(apiRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdApi)

	apiRequest.Methods = []string{"GET"}
	apiRequest.Name = "kevin"
	apiRequest.StripUri = true
	apiRequest.PreserveHost = true
	apiRequest.Retries = 10
	apiRequest.UpstreamConnectTimeout = 1000
	apiRequest.UpstreamSendTimeout = 4000
	apiRequest.UpstreamReadTimeout = 7000
	apiRequest.HttpsOnly = true
	apiRequest.HttpIfTerminated = true

	result, err := apiClient.UpdateById(createdApi.Id, apiRequest)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, apiRequest.Name, result.Name)
	assert.Equal(t, apiRequest.Hosts, result.Hosts)
	assert.Equal(t, apiRequest.Methods, result.Methods)
	assert.Equal(t, apiRequest.UpstreamUrl, result.UpstreamUrl)
	assert.Equal(t, apiRequest.StripUri, result.StripUri)
	assert.Equal(t, apiRequest.PreserveHost, result.PreserveHost)
	assert.Equal(t, apiRequest.Retries, result.Retries)
	assert.Equal(t, apiRequest.UpstreamConnectTimeout, result.UpstreamConnectTimeout)
	assert.Equal(t, apiRequest.UpstreamSendTimeout, result.UpstreamSendTimeout)
	assert.Equal(t, apiRequest.UpstreamReadTimeout, result.UpstreamReadTimeout)
	assert.Equal(t, apiRequest.HttpsOnly, result.HttpsOnly)
	assert.Equal(t, apiRequest.HttpIfTerminated, result.HttpIfTerminated)
}

func Test_ApisUpdateApiByName(t *testing.T) {

	apiRequest := &ApiRequest{
		Name:                   "test-" + uuid.NewV4().String(),
		Hosts:                  []string{"filter1.com"},
		Uris:                   []string{"/filter"},
		Methods:                []string{"PUT", "POST"},
		UpstreamUrl:            "http://linkerd:4140/someurl",
		StripUri:               false,
		PreserveHost:           false,
		Retries:                5,
		UpstreamConnectTimeout: 2222,
		UpstreamSendTimeout:    1233,
		UpstreamReadTimeout:    1234,
		HttpsOnly:              false,
		HttpIfTerminated:       false,
	}

	apiClient := NewClient(NewDefaultConfig()).Apis()

	createdApi, err := apiClient.Create(apiRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdApi)

	apiRequest.Methods = []string{"POST"}
	apiRequest.StripUri = true
	apiRequest.PreserveHost = true
	apiRequest.Retries = 3
	apiRequest.UpstreamConnectTimeout = 1000
	apiRequest.UpstreamSendTimeout = 888
	apiRequest.UpstreamReadTimeout = 234
	apiRequest.HttpsOnly = true
	apiRequest.HttpIfTerminated = true

	result, err := apiClient.UpdateByName(createdApi.Id, apiRequest)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, apiRequest.Name, result.Name)
	assert.Equal(t, apiRequest.Hosts, result.Hosts)
	assert.Equal(t, apiRequest.Methods, result.Methods)
	assert.Equal(t, apiRequest.UpstreamUrl, result.UpstreamUrl)
	assert.Equal(t, apiRequest.StripUri, result.StripUri)
	assert.Equal(t, apiRequest.PreserveHost, result.PreserveHost)
	assert.Equal(t, apiRequest.Retries, result.Retries)
	assert.Equal(t, apiRequest.UpstreamConnectTimeout, result.UpstreamConnectTimeout)
	assert.Equal(t, apiRequest.UpstreamSendTimeout, result.UpstreamSendTimeout)
	assert.Equal(t, apiRequest.UpstreamReadTimeout, result.UpstreamReadTimeout)
	assert.Equal(t, apiRequest.HttpsOnly, result.HttpsOnly)
	assert.Equal(t, apiRequest.HttpIfTerminated, result.HttpIfTerminated)
}
