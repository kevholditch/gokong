package gokong

import (
	"testing"

	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func Test_APIsGetByID(t *testing.T) {
	apiRequest := &APIRequest{
		Name:                   "test-" + uuid.NewV4().String(),
		Hosts:                  []string{"example.com"},
		URIs:                   []string{"/example"},
		Methods:                []string{"GET", "POST"},
		UpstreamURL:            "http://localhost:4140/testservice",
		StripURI:               true,
		PreserveHost:           true,
		Retries:                3,
		UpstreamConnectTimeout: 1000,
		UpstreamSendTimeout:    2000,
		UpstreamReadTimeout:    3000,
		HTTPSOnly:              true,
		HTTPIfTerminated:       true,
	}

	client := NewClient(NewDefaultConfig())
	createdAPI, err := client.APIs().Create(apiRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdAPI)

	result, err := client.APIs().GetByID(createdAPI.ID)

	assert.Equal(t, createdAPI, result)

}

func Test_APIsGetByName(t *testing.T) {
	apiRequest := &APIRequest{
		Name:                   "test-" + uuid.NewV4().String(),
		Hosts:                  []string{"example.com"},
		URIs:                   []string{"/example"},
		Methods:                []string{"GET", "POST"},
		UpstreamURL:            "http://localhost:4140/testservice",
		StripURI:               true,
		PreserveHost:           true,
		Retries:                3,
		UpstreamConnectTimeout: 1000,
		UpstreamSendTimeout:    2000,
		UpstreamReadTimeout:    3000,
		HTTPSOnly:              true,
		HTTPIfTerminated:       true,
	}

	client := NewClient(NewDefaultConfig())
	createdAPI, err := client.APIs().Create(apiRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdAPI)

	result, err := client.APIs().GetByName(createdAPI.Name)

	assert.Equal(t, createdAPI, result)

}

func Test_APIsGetNonExistentByID(t *testing.T) {

	result, err := NewClient(NewDefaultConfig()).APIs().GetByID("e5da4f1e-6b96-4b3b-a1aa-bdd71779e403")

	assert.Nil(t, err)
	assert.Nil(t, result)
}

func Test_APIsGetNonExistentByName(t *testing.T) {

	result, err := NewClient(NewDefaultConfig()).APIs().GetByName("9706f478-fd83-413c-b086-5608f7849db0")

	assert.Nil(t, err)
	assert.Nil(t, result)
}

func Test_APIsList(t *testing.T) {

	apiRequest := &APIRequest{
		Name:                   "test-" + uuid.NewV4().String(),
		Hosts:                  []string{"another.com"},
		URIs:                   []string{"/another"},
		Methods:                []string{"PUT", "POST"},
		UpstreamURL:            "http://linkerd:4140/myservice",
		StripURI:               false,
		PreserveHost:           false,
		Retries:                5,
		UpstreamConnectTimeout: 2222,
		UpstreamSendTimeout:    1233,
		UpstreamReadTimeout:    1234,
		HTTPSOnly:              false,
		HTTPIfTerminated:       false,
	}

	client := NewClient(NewDefaultConfig())
	createdAPI, err := client.APIs().Create(apiRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdAPI)

	results, err := client.APIs().List()

	assert.True(t, results.Total > 0)
	assert.True(t, len(results.Results) > 0)

}

func Test_APIsListFilteredByID(t *testing.T) {

	apiRequest := &APIRequest{
		Name:                   "test-" + uuid.NewV4().String(),
		Hosts:                  []string{"filter1.com"},
		URIs:                   []string{"/filter"},
		Methods:                []string{"PUT", "POST"},
		UpstreamURL:            "http://linkerd:4140/myservice",
		StripURI:               false,
		PreserveHost:           false,
		Retries:                5,
		UpstreamConnectTimeout: 2222,
		UpstreamSendTimeout:    1233,
		UpstreamReadTimeout:    1234,
		HTTPSOnly:              false,
		HTTPIfTerminated:       false,
	}
	apiRequest2 := &APIRequest{
		Name:                   "test-" + uuid.NewV4().String(),
		Hosts:                  []string{"filter2.com"},
		URIs:                   []string{"/filter2"},
		Methods:                []string{"PUT", "POST"},
		UpstreamURL:            "http://linkerd:4140/myservice2",
		StripURI:               false,
		PreserveHost:           false,
		Retries:                5,
		UpstreamConnectTimeout: 2222,
		UpstreamSendTimeout:    1233,
		UpstreamReadTimeout:    1234,
		HTTPSOnly:              false,
		HTTPIfTerminated:       false,
	}

	client := NewClient(NewDefaultConfig())
	createdAPI, err := client.APIs().Create(apiRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdAPI)

	createdAPI2, err := client.APIs().Create(apiRequest2)

	assert.Nil(t, err)
	assert.NotNil(t, createdAPI2)

	results, err := client.APIs().ListFiltered(&APIFilter{ID: createdAPI2.ID})

	assert.True(t, results.Total == 1)
	assert.True(t, len(results.Results) == 1)

	result := results.Results[0]

	assert.Equal(t, createdAPI2, result)

}

func Test_APIsListFilteredByName(t *testing.T) {

	apiRequest := &APIRequest{
		Name:                   "test-" + uuid.NewV4().String(),
		Hosts:                  []string{"filter1.com"},
		URIs:                   []string{"/filter"},
		Methods:                []string{"PUT", "POST"},
		UpstreamURL:            "http://linkerd:4140/myservice",
		StripURI:               false,
		PreserveHost:           false,
		Retries:                5,
		UpstreamConnectTimeout: 2222,
		UpstreamSendTimeout:    1233,
		UpstreamReadTimeout:    1234,
		HTTPSOnly:              false,
		HTTPIfTerminated:       false,
	}
	apiRequest2 := &APIRequest{
		Name:                   "test-" + uuid.NewV4().String(),
		Hosts:                  []string{"filter2.com"},
		URIs:                   []string{"/filter2"},
		Methods:                []string{"PUT", "POST"},
		UpstreamURL:            "http://linkerd:4140/myservice2",
		StripURI:               false,
		PreserveHost:           false,
		Retries:                5,
		UpstreamConnectTimeout: 2222,
		UpstreamSendTimeout:    1233,
		UpstreamReadTimeout:    1234,
		HTTPSOnly:              false,
		HTTPIfTerminated:       false,
	}

	client := NewClient(NewDefaultConfig())
	createdAPI, err := client.APIs().Create(apiRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdAPI)

	createdAPI2, err := client.APIs().Create(apiRequest2)

	assert.Nil(t, err)
	assert.NotNil(t, createdAPI2)

	results, err := client.APIs().ListFiltered(&APIFilter{Name: createdAPI2.Name})

	assert.Nil(t, err)
	assert.True(t, len(results.Results) == 1)

	result := results.Results[0]

	assert.Equal(t, createdAPI2, result)

}

func Test_APIsListFilteredByUpstreamURL(t *testing.T) {

	apiRequest := &APIRequest{
		Name:                   "test-" + uuid.NewV4().String(),
		Hosts:                  []string{"filter1.com"},
		URIs:                   []string{"/filter"},
		Methods:                []string{"PUT", "POST"},
		UpstreamURL:            "http://linkerd:4140/someurl",
		StripURI:               false,
		PreserveHost:           false,
		Retries:                5,
		UpstreamConnectTimeout: 2222,
		UpstreamSendTimeout:    1233,
		UpstreamReadTimeout:    1234,
		HTTPSOnly:              false,
		HTTPIfTerminated:       false,
	}
	apiRequest2 := &APIRequest{
		Name:                   "test-" + uuid.NewV4().String(),
		Hosts:                  []string{"filter2.com"},
		URIs:                   []string{"/filter2"},
		Methods:                []string{"PUT", "POST"},
		UpstreamURL:            "http://upstreamunique:4140/uniqueurl",
		StripURI:               false,
		PreserveHost:           false,
		Retries:                5,
		UpstreamConnectTimeout: 2222,
		UpstreamSendTimeout:    1233,
		UpstreamReadTimeout:    1234,
		HTTPSOnly:              false,
		HTTPIfTerminated:       false,
	}

	client := NewClient(NewDefaultConfig())
	createdAPI, err := client.APIs().Create(apiRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdAPI)

	createdAPI2, err := client.APIs().Create(apiRequest2)

	assert.Nil(t, err)
	assert.NotNil(t, createdAPI2)

	results, err := client.APIs().ListFiltered(&APIFilter{UpstreamURL: createdAPI2.UpstreamURL})

	assert.True(t, results.Total == 1)
	assert.True(t, len(results.Results) == 1)

	result := results.Results[0]

	assert.Equal(t, createdAPI2, result)
}

func Test_APIsListFilteredByRetries(t *testing.T) {

	apiRequest := &APIRequest{
		Name:                   "test-" + uuid.NewV4().String(),
		Hosts:                  []string{"filter1.com"},
		URIs:                   []string{"/filter"},
		Methods:                []string{"PUT", "POST"},
		UpstreamURL:            "http://linkerd:4140/someurl",
		StripURI:               false,
		PreserveHost:           false,
		Retries:                5,
		UpstreamConnectTimeout: 2222,
		UpstreamSendTimeout:    1233,
		UpstreamReadTimeout:    1234,
		HTTPSOnly:              false,
		HTTPIfTerminated:       false,
	}
	apiRequest2 := &APIRequest{
		Name:                   "test-" + uuid.NewV4().String(),
		Hosts:                  []string{"filter2.com"},
		URIs:                   []string{"/filter2"},
		Methods:                []string{"PUT", "POST"},
		UpstreamURL:            "http://linkerd:4140/uniqueurl",
		StripURI:               false,
		PreserveHost:           false,
		Retries:                1234,
		UpstreamConnectTimeout: 2222,
		UpstreamSendTimeout:    1233,
		UpstreamReadTimeout:    1234,
		HTTPSOnly:              false,
		HTTPIfTerminated:       false,
	}

	client := NewClient(NewDefaultConfig())
	createdAPI, err := client.APIs().Create(apiRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdAPI)

	createdAPI2, err := client.APIs().Create(apiRequest2)

	assert.Nil(t, err)
	assert.NotNil(t, createdAPI2)

	results, err := client.APIs().ListFiltered(&APIFilter{Retries: createdAPI2.Retries})

	assert.True(t, results.Total == 1)
	assert.True(t, len(results.Results) == 1)

	result := results.Results[0]

	assert.Equal(t, createdAPI2, result)
}

func Test_APIsListFilteredBySize(t *testing.T) {

	apiRequest := &APIRequest{
		Name:                   "test-" + uuid.NewV4().String(),
		Hosts:                  []string{"filter1.com"},
		URIs:                   []string{"/filter"},
		Methods:                []string{"PUT", "POST"},
		UpstreamURL:            "http://linkerd:4140/someurl",
		StripURI:               false,
		PreserveHost:           false,
		Retries:                5,
		UpstreamConnectTimeout: 2222,
		UpstreamSendTimeout:    1233,
		UpstreamReadTimeout:    1234,
		HTTPSOnly:              false,
		HTTPIfTerminated:       false,
	}
	apiRequest2 := &APIRequest{
		Name:                   "test-" + uuid.NewV4().String(),
		Hosts:                  []string{"filter2.com"},
		URIs:                   []string{"/filter2"},
		Methods:                []string{"PUT", "POST"},
		UpstreamURL:            "http://linkerd:4140/uniqueurl",
		StripURI:               false,
		PreserveHost:           false,
		Retries:                1234,
		UpstreamConnectTimeout: 2222,
		UpstreamSendTimeout:    1233,
		UpstreamReadTimeout:    1234,
		HTTPSOnly:              false,
		HTTPIfTerminated:       false,
	}

	client := NewClient(NewDefaultConfig())
	createdAPI, err := client.APIs().Create(apiRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdAPI)

	createdAPI2, err := client.APIs().Create(apiRequest2)

	assert.Nil(t, err)
	assert.NotNil(t, createdAPI2)

	results, err := client.APIs().ListFiltered(&APIFilter{Size: 1})

	assert.True(t, len(results.Results) == 1)

}

func Test_APIsCreate(t *testing.T) {

	apiRequest := &APIRequest{
		Name:                   "test-" + uuid.NewV4().String(),
		Hosts:                  []string{"example.com"},
		URIs:                   []string{"/example"},
		Methods:                []string{"GET", "POST"},
		UpstreamURL:            "http://localhost:4140/testservice",
		StripURI:               false,
		PreserveHost:           true,
		Retries:                3,
		UpstreamConnectTimeout: 1000,
		UpstreamSendTimeout:    2000,
		UpstreamReadTimeout:    3000,
		HTTPSOnly:              true,
		HTTPIfTerminated:       false,
	}

	result, err := NewClient(NewDefaultConfig()).APIs().Create(apiRequest)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, apiRequest.Name, result.Name)
	assert.Equal(t, apiRequest.Hosts, result.Hosts)
	assert.Equal(t, apiRequest.URIs, result.URIs)
	assert.Equal(t, apiRequest.Methods, result.Methods)
	assert.Equal(t, apiRequest.UpstreamURL, result.UpstreamURL)
	assert.Equal(t, apiRequest.StripURI, result.StripURI)
	assert.Equal(t, apiRequest.PreserveHost, result.PreserveHost)
	assert.Equal(t, apiRequest.UpstreamConnectTimeout, result.UpstreamConnectTimeout)
	assert.Equal(t, apiRequest.UpstreamSendTimeout, result.UpstreamSendTimeout)
	assert.Equal(t, apiRequest.UpstreamReadTimeout, result.UpstreamReadTimeout)
	assert.Equal(t, apiRequest.HTTPSOnly, result.HTTPSOnly)
	assert.Equal(t, apiRequest.HTTPIfTerminated, result.HTTPIfTerminated)
}

func Test_APIsCreateInvalid(t *testing.T) {

	apiRequest := &APIRequest{
		Name: "test-" + uuid.NewV4().String(),
	}

	result, err := NewClient(NewDefaultConfig()).APIs().Create(apiRequest)

	assert.NotNil(t, err)
	assert.Nil(t, result)
}

func Test_APIsCreateWithOnlyRequiredFields(t *testing.T) {
	apiRequest := &APIRequest{
		Name:        "test-" + uuid.NewV4().String(),
		Hosts:       []string{"example.com"},
		UpstreamURL: "http://localhost:4140/testservice",
	}

	result, err := NewClient(NewDefaultConfig()).APIs().Create(apiRequest)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, apiRequest.Name, result.Name)
	assert.Equal(t, apiRequest.Hosts, result.Hosts)
	assert.Nil(t, result.URIs)
	assert.Nil(t, result.Methods)
	assert.Equal(t, apiRequest.UpstreamURL, result.UpstreamURL)
	assert.Equal(t, false, result.StripURI)
	assert.Equal(t, false, result.PreserveHost)
	assert.Equal(t, 60000, result.UpstreamConnectTimeout)
	assert.Equal(t, 60000, result.UpstreamSendTimeout)
	assert.Equal(t, 60000, result.UpstreamReadTimeout)
	assert.Equal(t, false, result.HTTPSOnly)
	assert.Equal(t, false, result.HTTPIfTerminated)
}

func Test_APIsDeleteByID(t *testing.T) {

	apiRequest := &APIRequest{
		Name:                   "test-" + uuid.NewV4().String(),
		Hosts:                  []string{"delete.com"},
		URIs:                   []string{"/delete"},
		Methods:                []string{"GET", "POST"},
		UpstreamURL:            "http://localhost:4140/testservice",
		StripURI:               true,
		PreserveHost:           true,
		Retries:                3,
		UpstreamConnectTimeout: 1000,
		UpstreamSendTimeout:    2000,
		UpstreamReadTimeout:    3000,
		HTTPSOnly:              true,
		HTTPIfTerminated:       true,
	}

	client := NewClient(NewDefaultConfig())
	createdAPI, err := client.APIs().Create(apiRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdAPI)

	err = client.APIs().DeleteByID(createdAPI.ID)

	assert.Nil(t, err)

	deletedAPI, err := client.APIs().GetByID(createdAPI.ID)
	assert.Nil(t, err)
	assert.Nil(t, deletedAPI)

}

func Test_APIsDeleteByName(t *testing.T) {

	apiRequest := &APIRequest{
		Name:                   "test-" + uuid.NewV4().String(),
		Hosts:                  []string{"delete.com"},
		URIs:                   []string{"/delete"},
		Methods:                []string{"GET", "POST"},
		UpstreamURL:            "http://localhost:4140/testservice",
		StripURI:               true,
		PreserveHost:           true,
		Retries:                3,
		UpstreamConnectTimeout: 1000,
		UpstreamSendTimeout:    2000,
		UpstreamReadTimeout:    3000,
		HTTPSOnly:              true,
		HTTPIfTerminated:       true,
	}

	client := NewClient(NewDefaultConfig())
	createdAPI, err := client.APIs().Create(apiRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdAPI)

	err = client.APIs().DeleteByName(createdAPI.ID)

	assert.Nil(t, err)

	deletedAPI, err := client.APIs().GetByID(createdAPI.ID)
	assert.Nil(t, err)
	assert.Nil(t, deletedAPI)

}

func Test_APIsUpdateAPIByID(t *testing.T) {

	apiRequest := &APIRequest{
		Name:                   "test-" + uuid.NewV4().String(),
		Hosts:                  []string{"filter1.com"},
		URIs:                   []string{"/filter"},
		Methods:                []string{"PUT", "POST"},
		UpstreamURL:            "http://linkerd:4140/someurl",
		StripURI:               false,
		PreserveHost:           false,
		Retries:                5,
		UpstreamConnectTimeout: 2222,
		UpstreamSendTimeout:    1233,
		UpstreamReadTimeout:    1234,
		HTTPSOnly:              false,
		HTTPIfTerminated:       false,
	}

	client := NewClient(NewDefaultConfig())

	createdAPI, err := client.APIs().Create(apiRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdAPI)

	apiRequest.Methods = []string{"GET"}
	apiRequest.Name = "kevin"
	apiRequest.StripURI = true
	apiRequest.PreserveHost = true
	apiRequest.Retries = 10
	apiRequest.UpstreamConnectTimeout = 1000
	apiRequest.UpstreamSendTimeout = 4000
	apiRequest.UpstreamReadTimeout = 7000
	apiRequest.HTTPSOnly = true
	apiRequest.HTTPIfTerminated = true

	result, err := client.APIs().UpdateByID(createdAPI.ID, apiRequest)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, apiRequest.Name, result.Name)
	assert.Equal(t, apiRequest.Hosts, result.Hosts)
	assert.Equal(t, apiRequest.Methods, result.Methods)
	assert.Equal(t, apiRequest.UpstreamURL, result.UpstreamURL)
	assert.Equal(t, apiRequest.StripURI, result.StripURI)
	assert.Equal(t, apiRequest.PreserveHost, result.PreserveHost)
	assert.Equal(t, apiRequest.Retries, result.Retries)
	assert.Equal(t, apiRequest.UpstreamConnectTimeout, result.UpstreamConnectTimeout)
	assert.Equal(t, apiRequest.UpstreamSendTimeout, result.UpstreamSendTimeout)
	assert.Equal(t, apiRequest.UpstreamReadTimeout, result.UpstreamReadTimeout)
	assert.Equal(t, apiRequest.HTTPSOnly, result.HTTPSOnly)
	assert.Equal(t, apiRequest.HTTPIfTerminated, result.HTTPIfTerminated)
}

func Test_APIsUpdateAPIByName(t *testing.T) {

	apiRequest := &APIRequest{
		Name:                   "test-" + uuid.NewV4().String(),
		Hosts:                  []string{"filter1.com"},
		URIs:                   []string{"/filter"},
		Methods:                []string{"PUT", "POST"},
		UpstreamURL:            "http://linkerd:4140/someurl",
		StripURI:               false,
		PreserveHost:           false,
		Retries:                5,
		UpstreamConnectTimeout: 2222,
		UpstreamSendTimeout:    1233,
		UpstreamReadTimeout:    1234,
		HTTPSOnly:              false,
		HTTPIfTerminated:       false,
	}

	client := NewClient(NewDefaultConfig())
	createdAPI, err := client.APIs().Create(apiRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdAPI)

	apiRequest.Methods = []string{"POST"}
	apiRequest.StripURI = true
	apiRequest.PreserveHost = true
	apiRequest.Retries = 3
	apiRequest.UpstreamConnectTimeout = 1000
	apiRequest.UpstreamSendTimeout = 888
	apiRequest.UpstreamReadTimeout = 234
	apiRequest.HTTPSOnly = true
	apiRequest.HTTPIfTerminated = true

	result, err := client.APIs().UpdateByName(createdAPI.ID, apiRequest)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, apiRequest.Name, result.Name)
	assert.Equal(t, apiRequest.Hosts, result.Hosts)
	assert.Equal(t, apiRequest.Methods, result.Methods)
	assert.Equal(t, apiRequest.UpstreamURL, result.UpstreamURL)
	assert.Equal(t, apiRequest.StripURI, result.StripURI)
	assert.Equal(t, apiRequest.PreserveHost, result.PreserveHost)
	assert.Equal(t, apiRequest.Retries, result.Retries)
	assert.Equal(t, apiRequest.UpstreamConnectTimeout, result.UpstreamConnectTimeout)
	assert.Equal(t, apiRequest.UpstreamSendTimeout, result.UpstreamSendTimeout)
	assert.Equal(t, apiRequest.UpstreamReadTimeout, result.UpstreamReadTimeout)
	assert.Equal(t, apiRequest.HTTPSOnly, result.HTTPSOnly)
	assert.Equal(t, apiRequest.HTTPIfTerminated, result.HTTPIfTerminated)
}
