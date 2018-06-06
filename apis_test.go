package gokong

import (
	"testing"

	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func Test_ApisGetById(t *testing.T) {
	apiRequest := &ApiRequest{
		Name:                   String("test-" + uuid.NewV4().String()),
		Hosts:                  StringSlice([]string{"example.com"}),
		Uris:                   StringSlice([]string{"/example"}),
		Methods:                StringSlice([]string{"GET", "POST"}),
		UpstreamUrl:            String("http://localhost:4140/testservice"),
		StripUri:               Bool(true),
		PreserveHost:           Bool(true),
		Retries:                Int(3),
		UpstreamConnectTimeout: Int(1000),
		UpstreamSendTimeout:    Int(2000),
		UpstreamReadTimeout:    Int(3000),
		HttpsOnly:              Bool(true),
		HttpIfTerminated:       Bool(true),
	}

	client := NewClient(NewDefaultConfig())
	createdApi, err := client.Apis().Create(apiRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdApi)

	result, err := client.Apis().GetById(*createdApi.Id)

	assert.Equal(t, createdApi, result)

}

func Test_ApisGetByName(t *testing.T) {
	apiRequest := &ApiRequest{
		Name:                   String("test-" + uuid.NewV4().String()),
		Hosts:                  StringSlice([]string{"example.com"}),
		Uris:                   StringSlice([]string{"/example"}),
		Methods:                StringSlice([]string{"GET", "POST"}),
		UpstreamUrl:            String("http://localhost:4140/testservice"),
		StripUri:               Bool(true),
		PreserveHost:           Bool(true),
		Retries:                Int(3),
		UpstreamConnectTimeout: Int(1000),
		UpstreamSendTimeout:    Int(2000),
		UpstreamReadTimeout:    Int(3000),
		HttpsOnly:              Bool(true),
		HttpIfTerminated:       Bool(true),
	}

	client := NewClient(NewDefaultConfig())
	createdApi, err := client.Apis().Create(apiRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdApi)

	result, err := client.Apis().GetByName(*createdApi.Name)

	assert.Equal(t, createdApi, result)

}

func Test_ApisGetNonExistentById(t *testing.T) {

	result, err := NewClient(NewDefaultConfig()).Apis().GetById("e5da4f1e-6b96-4b3b-a1aa-bdd71779e403")

	assert.Nil(t, err)
	assert.Nil(t, result)
}

func Test_ApisGetNonExistentByName(t *testing.T) {

	result, err := NewClient(NewDefaultConfig()).Apis().GetByName("9706f478-fd83-413c-b086-5608f7849db0")

	assert.Nil(t, err)
	assert.Nil(t, result)
}

func Test_ApisList(t *testing.T) {

	apiRequest := &ApiRequest{
		Name:                   String("test-" + uuid.NewV4().String()),
		Hosts:                  StringSlice([]string{"another.com"}),
		Uris:                   StringSlice([]string{"/another"}),
		Methods:                StringSlice([]string{"PUT", "POST"}),
		UpstreamUrl:            String("http://linkerd:4140/myservice"),
		StripUri:               Bool(false),
		PreserveHost:           Bool(false),
		Retries:                Int(5),
		UpstreamConnectTimeout: Int(222),
		UpstreamSendTimeout:    Int(233),
		UpstreamReadTimeout:    Int(234),
		HttpsOnly:              Bool(false),
		HttpIfTerminated:       Bool(false),
	}

	client := NewClient(NewDefaultConfig())
	createdApi, err := client.Apis().Create(apiRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdApi)

	results, err := client.Apis().List()

	assert.True(t, results.Total > 0)
	assert.True(t, len(results.Results) > 0)

}

func Test_ApisListFilteredById(t *testing.T) {

	apiRequest := &ApiRequest{
		Name:                   String("test-" + uuid.NewV4().String()),
		Hosts:                  StringSlice([]string{"filter1.com"}),
		Uris:                   StringSlice([]string{"/filter"}),
		Methods:                StringSlice([]string{"PUT", "POST"}),
		UpstreamUrl:            String("http://linkerd:4140/myservice"),
		StripUri:               Bool(false),
		PreserveHost:           Bool(false),
		Retries:                Int(5),
		UpstreamConnectTimeout: Int(2222),
		UpstreamSendTimeout:    Int(1233),
		UpstreamReadTimeout:    Int(1234),
		HttpsOnly:              Bool(false),
		HttpIfTerminated:       Bool(false),
	}
	apiRequest2 := &ApiRequest{
		Name:                   String("test-" + uuid.NewV4().String()),
		Hosts:                  StringSlice([]string{"filter2.com"}),
		Uris:                   StringSlice([]string{"/filter2"}),
		Methods:                StringSlice([]string{"PUT", "POST"}),
		UpstreamUrl:            String("http://linkerd:4140/myservice2"),
		StripUri:               Bool(false),
		PreserveHost:           Bool(false),
		Retries:                Int(5),
		UpstreamConnectTimeout: Int(2222),
		UpstreamSendTimeout:    Int(1233),
		UpstreamReadTimeout:    Int(1234),
		HttpsOnly:              Bool(false),
		HttpIfTerminated:       Bool(false),
	}

	client := NewClient(NewDefaultConfig())
	createdApi, err := client.Apis().Create(apiRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdApi)

	createdApi2, err := client.Apis().Create(apiRequest2)

	assert.Nil(t, err)
	assert.NotNil(t, createdApi2)

	results, err := client.Apis().ListFiltered(&ApiFilter{Id: *createdApi2.Id})

	assert.True(t, results.Total == 1)
	assert.True(t, len(results.Results) == 1)

	result := results.Results[0]

	assert.Equal(t, createdApi2, result)

}

func Test_ApisListFilteredByName(t *testing.T) {

	apiRequest := &ApiRequest{
		Name:                   String("test-" + uuid.NewV4().String()),
		Hosts:                  StringSlice([]string{"filter1.com"}),
		Uris:                   StringSlice([]string{"/filter"}),
		Methods:                StringSlice([]string{"PUT", "POST"}),
		UpstreamUrl:            String("http://linkerd:4140/myservice"),
		StripUri:               Bool(false),
		PreserveHost:           Bool(false),
		Retries:                Int(5),
		UpstreamConnectTimeout: Int(2222),
		UpstreamSendTimeout:    Int(1233),
		UpstreamReadTimeout:    Int(1234),
		HttpsOnly:              Bool(false),
		HttpIfTerminated:       Bool(false),
	}
	apiRequest2 := &ApiRequest{
		Name:                   String("test-" + uuid.NewV4().String()),
		Hosts:                  StringSlice([]string{"filter2.com"}),
		Uris:                   StringSlice([]string{"/filter2"}),
		Methods:                StringSlice([]string{"PUT", "POST"}),
		UpstreamUrl:            String("http://linkerd:4140/myservice2"),
		StripUri:               Bool(false),
		PreserveHost:           Bool(false),
		Retries:                Int(5),
		UpstreamConnectTimeout: Int(2222),
		UpstreamSendTimeout:    Int(1233),
		UpstreamReadTimeout:    Int(1234),
		HttpsOnly:              Bool(false),
		HttpIfTerminated:       Bool(false),
	}

	client := NewClient(NewDefaultConfig())
	createdApi, err := client.Apis().Create(apiRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdApi)

	createdApi2, err := client.Apis().Create(apiRequest2)

	assert.Nil(t, err)
	assert.NotNil(t, createdApi2)

	results, err := client.Apis().ListFiltered(&ApiFilter{Name: *createdApi2.Name})

	assert.Nil(t, err)
	assert.True(t, len(results.Results) == 1)

	result := results.Results[0]

	assert.Equal(t, createdApi2, result)

}

func Test_ApisListFilteredByUpstreamUrl(t *testing.T) {

	apiRequest := &ApiRequest{
		Name:                   String("test-" + uuid.NewV4().String()),
		Hosts:                  StringSlice([]string{"filter1.com"}),
		Uris:                   StringSlice([]string{"/filter"}),
		Methods:                StringSlice([]string{"PUT", "POST"}),
		UpstreamUrl:            String("http://linkerd:4140/someurl"),
		StripUri:               Bool(false),
		PreserveHost:           Bool(false),
		Retries:                Int(5),
		UpstreamConnectTimeout: Int(2222),
		UpstreamSendTimeout:    Int(1233),
		UpstreamReadTimeout:    Int(1234),
		HttpsOnly:              Bool(false),
		HttpIfTerminated:       Bool(false),
	}

	apiRequest2 := &ApiRequest{
		Name:                   String("test-" + uuid.NewV4().String()),
		Hosts:                  StringSlice([]string{"filter2.com"}),
		Uris:                   StringSlice([]string{"/filter2"}),
		Methods:                StringSlice([]string{"PUT", "POST"}),
		UpstreamUrl:            String("http://upstreamunique:4140/uniqueurl"),
		StripUri:               Bool(false),
		PreserveHost:           Bool(false),
		Retries:                Int(5),
		UpstreamConnectTimeout: Int(2222),
		UpstreamSendTimeout:    Int(1233),
		UpstreamReadTimeout:    Int(1234),
		HttpsOnly:              Bool(false),
		HttpIfTerminated:       Bool(false),
	}

	client := NewClient(NewDefaultConfig())
	createdApi, err := client.Apis().Create(apiRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdApi)

	createdApi2, err := client.Apis().Create(apiRequest2)

	assert.Nil(t, err)
	assert.NotNil(t, createdApi2)

	results, err := client.Apis().ListFiltered(&ApiFilter{UpstreamUrl: *createdApi2.UpstreamUrl})

	assert.True(t, results.Total == 1)
	assert.True(t, len(results.Results) == 1)

	result := results.Results[0]

	assert.Equal(t, createdApi2, result)
}

func Test_ApisListFilteredByRetries(t *testing.T) {

	apiRequest := &ApiRequest{
		Name:                   String("test-" + uuid.NewV4().String()),
		Hosts:                  StringSlice([]string{"filter1.com"}),
		Uris:                   StringSlice([]string{"/filter"}),
		Methods:                StringSlice([]string{"PUT", "POST"}),
		UpstreamUrl:            String("http://linkerd:4140/someurl"),
		StripUri:               Bool(false),
		PreserveHost:           Bool(false),
		Retries:                Int(5),
		UpstreamConnectTimeout: Int(2222),
		UpstreamSendTimeout:    Int(1233),
		UpstreamReadTimeout:    Int(1234),
		HttpsOnly:              Bool(false),
		HttpIfTerminated:       Bool(false),
	}
	apiRequest2 := &ApiRequest{
		Name:                   String("test-" + uuid.NewV4().String()),
		Hosts:                  StringSlice([]string{"filter2.com"}),
		Uris:                   StringSlice([]string{"/filter2"}),
		Methods:                StringSlice([]string{"PUT", "POST"}),
		UpstreamUrl:            String("http://linkerd:4140/uniqueurl"),
		StripUri:               Bool(false),
		PreserveHost:           Bool(false),
		Retries:                Int(1234),
		UpstreamConnectTimeout: Int(2222),
		UpstreamSendTimeout:    Int(1233),
		UpstreamReadTimeout:    Int(1234),
		HttpsOnly:              Bool(false),
		HttpIfTerminated:       Bool(false),
	}

	client := NewClient(NewDefaultConfig())
	createdApi, err := client.Apis().Create(apiRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdApi)

	createdApi2, err := client.Apis().Create(apiRequest2)

	assert.Nil(t, err)
	assert.NotNil(t, createdApi2)

	results, err := client.Apis().ListFiltered(&ApiFilter{Retries: *createdApi2.Retries})

	assert.True(t, results.Total == 1)
	assert.True(t, len(results.Results) == 1)

	result := results.Results[0]

	assert.Equal(t, createdApi2, result)
}

func Test_ApisListFilteredBySize(t *testing.T) {

	apiRequest := &ApiRequest{
		Name:                   String("test-" + uuid.NewV4().String()),
		Hosts:                  StringSlice([]string{"filter1.com"}),
		Uris:                   StringSlice([]string{"/filter"}),
		Methods:                StringSlice([]string{"PUT", "POST"}),
		UpstreamUrl:            String("http://linkerd:4140/someurl"),
		StripUri:               Bool(false),
		PreserveHost:           Bool(false),
		Retries:                Int(5),
		UpstreamConnectTimeout: Int(2222),
		UpstreamSendTimeout:    Int(1233),
		UpstreamReadTimeout:    Int(1234),
		HttpsOnly:              Bool(false),
		HttpIfTerminated:       Bool(false),
	}

	apiRequest2 := &ApiRequest{
		Name:                   String("test-" + uuid.NewV4().String()),
		Hosts:                  StringSlice([]string{"filter2.com"}),
		Uris:                   StringSlice([]string{"/filter2"}),
		Methods:                StringSlice([]string{"PUT", "POST"}),
		UpstreamUrl:            String("http://upstreamunique:4140/uniqueurl"),
		StripUri:               Bool(false),
		PreserveHost:           Bool(false),
		Retries:                Int(5),
		UpstreamConnectTimeout: Int(2222),
		UpstreamSendTimeout:    Int(1233),
		UpstreamReadTimeout:    Int(1234),
		HttpsOnly:              Bool(false),
		HttpIfTerminated:       Bool(false),
	}

	client := NewClient(NewDefaultConfig())
	createdApi, err := client.Apis().Create(apiRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdApi)

	createdApi2, err := client.Apis().Create(apiRequest2)

	assert.Nil(t, err)
	assert.NotNil(t, createdApi2)

	results, err := client.Apis().ListFiltered(&ApiFilter{Size: 1})

	assert.True(t, len(results.Results) == 1)

}

func Test_ApisCreate(t *testing.T) {

	apiRequest := &ApiRequest{
		Name:                   String("test-" + uuid.NewV4().String()),
		Hosts:                  StringSlice([]string{"example.com"}),
		Uris:                   StringSlice([]string{"/example"}),
		Methods:                StringSlice([]string{"GET", "POST"}),
		UpstreamUrl:            String("http://localhost:4140/testservice"),
		StripUri:               Bool(false),
		PreserveHost:           Bool(true),
		Retries:                Int(3),
		UpstreamConnectTimeout: Int(1000),
		UpstreamSendTimeout:    Int(2000),
		UpstreamReadTimeout:    Int(3000),
		HttpsOnly:              Bool(true),
		HttpIfTerminated:       Bool(false),
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

func Test_ApisCreateInvalid(t *testing.T) {

	apiRequest := &ApiRequest{
		Name: String("test-" + uuid.NewV4().String()),
	}

	result, err := NewClient(NewDefaultConfig()).Apis().Create(apiRequest)

	assert.NotNil(t, err)
	assert.Nil(t, result)
}

func Test_ApisCreateWithOnlyRequiredFields(t *testing.T) {
	apiRequest := &ApiRequest{
		Name:        String("test-" + uuid.NewV4().String()),
		Hosts:       StringSlice([]string{"example.com"}),
		UpstreamUrl: String("http://localhost:4140/testservice"),
	}

	result, err := NewClient(NewDefaultConfig()).Apis().Create(apiRequest)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, apiRequest.Name, result.Name)
	assert.Equal(t, apiRequest.Hosts, result.Hosts)
	assert.Nil(t, result.Uris)
	assert.Nil(t, result.Methods)
	assert.Equal(t, apiRequest.UpstreamUrl, result.UpstreamUrl)
	assert.Equal(t, 5, *result.Retries)
	assert.Equal(t, true, *result.StripUri)
	assert.Equal(t, false, *result.PreserveHost)
	assert.Equal(t, 60000, *result.UpstreamConnectTimeout)
	assert.Equal(t, 60000, *result.UpstreamSendTimeout)
	assert.Equal(t, 60000, *result.UpstreamReadTimeout)
	assert.Equal(t, false, *result.HttpsOnly)
	assert.Equal(t, false, *result.HttpIfTerminated)
}

func Test_ApisDeleteById(t *testing.T) {

	apiRequest := &ApiRequest{
		Name:                   String("test-" + uuid.NewV4().String()),
		Hosts:                  StringSlice([]string{"delete.com"}),
		Uris:                   StringSlice([]string{"/delete"}),
		Methods:                StringSlice([]string{"GET", "POST"}),
		UpstreamUrl:            String("http://localhost:4140/testservice"),
		StripUri:               Bool(true),
		PreserveHost:           Bool(true),
		Retries:                Int(3),
		UpstreamConnectTimeout: Int(1000),
		UpstreamSendTimeout:    Int(2000),
		UpstreamReadTimeout:    Int(3000),
		HttpsOnly:              Bool(true),
		HttpIfTerminated:       Bool(true),
	}

	client := NewClient(NewDefaultConfig())
	createdApi, err := client.Apis().Create(apiRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdApi)

	err = client.Apis().DeleteById(*createdApi.Id)

	assert.Nil(t, err)

	deletedApi, err := client.Apis().GetById(*createdApi.Id)
	assert.Nil(t, err)
	assert.Nil(t, deletedApi)

}

func Test_ApisDeleteByName(t *testing.T) {

	apiRequest := &ApiRequest{
		Name:                   String("test-" + uuid.NewV4().String()),
		Hosts:                  StringSlice([]string{"delete.com"}),
		Uris:                   StringSlice([]string{"/delete"}),
		Methods:                StringSlice([]string{"GET", "POST"}),
		UpstreamUrl:            String("http://localhost:4140/testservice"),
		StripUri:               Bool(true),
		PreserveHost:           Bool(true),
		Retries:                Int(3),
		UpstreamConnectTimeout: Int(1000),
		UpstreamSendTimeout:    Int(2000),
		UpstreamReadTimeout:    Int(3000),
		HttpsOnly:              Bool(true),
		HttpIfTerminated:       Bool(true),
	}

	client := NewClient(NewDefaultConfig())
	createdApi, err := client.Apis().Create(apiRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdApi)

	err = client.Apis().DeleteByName(*createdApi.Id)

	assert.Nil(t, err)

	deletedApi, err := client.Apis().GetById(*createdApi.Id)
	assert.Nil(t, err)
	assert.Nil(t, deletedApi)

}

func Test_ApisUpdateApiById(t *testing.T) {

	apiRequest := &ApiRequest{
		Name:                   String("test-" + uuid.NewV4().String()),
		Hosts:                  StringSlice([]string{"filter1.com"}),
		Uris:                   StringSlice([]string{"/filter"}),
		Methods:                StringSlice([]string{"PUT", "POST"}),
		UpstreamUrl:            String("http://linkerd:4140/someurl"),
		StripUri:               Bool(false),
		PreserveHost:           Bool(false),
		Retries:                Int(5),
		UpstreamConnectTimeout: Int(2222),
		UpstreamSendTimeout:    Int(1233),
		UpstreamReadTimeout:    Int(1234),
		HttpsOnly:              Bool(false),
		HttpIfTerminated:       Bool(false),
	}

	client := NewClient(NewDefaultConfig())

	createdApi, err := client.Apis().Create(apiRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdApi)

	apiRequest.Methods = StringSlice([]string{"GET"})
	apiRequest.Name = String("kevin")
	apiRequest.StripUri = Bool(true)
	apiRequest.PreserveHost = Bool(true)
	apiRequest.Retries = Int(10)
	apiRequest.UpstreamConnectTimeout = Int(1000)
	apiRequest.UpstreamSendTimeout = Int(4000)
	apiRequest.UpstreamReadTimeout = Int(7000)
	apiRequest.HttpsOnly = Bool(true)
	apiRequest.HttpIfTerminated = Bool(true)

	result, err := client.Apis().UpdateById(*createdApi.Id, apiRequest)

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
		Name:                   String("test-" + uuid.NewV4().String()),
		Hosts:                  StringSlice([]string{"filter1.com"}),
		Uris:                   StringSlice([]string{"/filter"}),
		Methods:                StringSlice([]string{"PUT", "POST"}),
		UpstreamUrl:            String("http://linkerd:4140/someurl"),
		StripUri:               Bool(false),
		PreserveHost:           Bool(false),
		Retries:                Int(5),
		UpstreamConnectTimeout: Int(2222),
		UpstreamSendTimeout:    Int(1233),
		UpstreamReadTimeout:    Int(1234),
		HttpsOnly:              Bool(false),
		HttpIfTerminated:       Bool(false),
	}

	client := NewClient(NewDefaultConfig())
	createdApi, err := client.Apis().Create(apiRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdApi)

	apiRequest.Methods = StringSlice([]string{"POST"})
	apiRequest.StripUri = Bool(true)
	apiRequest.PreserveHost = Bool(true)
	apiRequest.Retries = Int(3)
	apiRequest.UpstreamConnectTimeout = Int(1000)
	apiRequest.UpstreamSendTimeout = Int(888)
	apiRequest.UpstreamReadTimeout = Int(234)
	apiRequest.HttpsOnly = Bool(true)
	apiRequest.HttpIfTerminated = Bool(true)

	result, err := client.Apis().UpdateByName(*createdApi.Id, apiRequest)

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
