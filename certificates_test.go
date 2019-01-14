package gokong

import (
	"os"
	"testing"

	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func Test_CertificatesGetById(t *testing.T) {
	certificateRequest := &CertificateRequest{
		Cert: String("public key-" + uuid.NewV4().String()),
		Key:  String("private key-" + uuid.NewV4().String()),
	}

	client := NewClient(NewDefaultConfig())
	createdCertificate, err := client.Certificates().Create(certificateRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdCertificate)

	result, err := client.Certificates().GetById(*createdCertificate.Id)

	assert.NotNil(t, result)
	assert.Equal(t, createdCertificate, result)

}

func Test_CertificatesGetNonExistentById(t *testing.T) {

	result, err := NewClient(NewDefaultConfig()).Certificates().GetById("8df4d1ed-c973-4b9a-868d-3e67d5c417da")

	assert.Nil(t, err)
	assert.Nil(t, result)
}

func Test_CertificatesCreate(t *testing.T) {
	certificateRequest := &CertificateRequest{
		Cert: String("public key-" + uuid.NewV4().String()),
		Key:  String("private key-" + uuid.NewV4().String()),
	}

	result, err := NewClient(NewDefaultConfig()).Certificates().Create(certificateRequest)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.True(t, result.Id != nil)
	assert.Equal(t, certificateRequest.Cert, result.Cert)
	assert.Equal(t, certificateRequest.Key, result.Key)

}

func Test_CertificatesCreateInvalid(t *testing.T) {
	certificateRequest := &CertificateRequest{}

	result, err := NewClient(NewDefaultConfig()).Certificates().Create(certificateRequest)

	assert.NotNil(t, err)
	assert.Nil(t, result)

}

func Test_CertificatesUpdateById(t *testing.T) {
	certificateRequest := &CertificateRequest{
		Cert: String("public key-" + uuid.NewV4().String()),
		Key:  String("private key-" + uuid.NewV4().String()),
	}

	client := NewClient(NewDefaultConfig())
	createdCertificate, err := client.Certificates().Create(certificateRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdCertificate)

	certificateRequest.Cert = String("public key-" + uuid.NewV4().String())

	result, err := client.Certificates().UpdateById(*createdCertificate.Id, certificateRequest)

	assert.Equal(t, certificateRequest.Cert, result.Cert)
	assert.Equal(t, certificateRequest.Key, result.Key)

}

func Test_CertificatesUpdateByIdInvalid(t *testing.T) {
	certificateRequest := &CertificateRequest{
		Cert: String("public key-" + uuid.NewV4().String()),
		Key:  String("private key-" + uuid.NewV4().String()),
	}

	client := NewClient(NewDefaultConfig())
	createdCertificate, err := client.Certificates().Create(certificateRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdCertificate)

	certificateRequest.Cert = String("")
	certificateRequest.Key = String("")

	result, err := client.Certificates().UpdateById(*createdCertificate.Id, certificateRequest)

	assert.NotNil(t, err)
	assert.Nil(t, result)

}

func Test_CertificatesDeleteById(t *testing.T) {
	certificateRequest := &CertificateRequest{
		Cert: String("public key-" + uuid.NewV4().String()),
		Key:  String("private key-" + uuid.NewV4().String()),
	}

	client := NewClient(NewDefaultConfig())
	createdCertificate, err := client.Certificates().Create(certificateRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdCertificate)

	err = client.Certificates().DeleteById(*createdCertificate.Id)

	assert.Nil(t, err)

	deletedCertificate, err := client.Certificates().GetById(*createdCertificate.Id)

	assert.Nil(t, err)
	assert.Nil(t, deletedCertificate)

}

func Test_CertificatesList(t *testing.T) {
	client := NewClient(NewDefaultConfig())

	certificateRequest := &CertificateRequest{
		Cert: String("public key-" + uuid.NewV4().String()),
		Key:  String("private key-" + uuid.NewV4().String()),
	}

	createdCertificate, err := client.Certificates().Create(certificateRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdCertificate)

	certificateRequest2 := &CertificateRequest{
		Cert: String("public key-" + uuid.NewV4().String()),
		Key:  String("private key-" + uuid.NewV4().String()),
	}

	createdCertificate2, err := client.Certificates().Create(certificateRequest2)

	assert.Nil(t, err)
	assert.NotNil(t, createdCertificate2)

	results, err := client.Certificates().List()

	assert.Nil(t, err)
	assert.True(t, len(results.Results) > 1)

}

func Test_AllCertificateEndpointsShouldReturnErrorWhenRequestUnauthorised(t *testing.T) {

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

	certificateRequest := &CertificateRequest{
		Cert: String("public key-" + uuid.NewV4().String()),
		Key:  String("private key-" + uuid.NewV4().String()),
	}
	createdCertificate, err := client.Certificates().Create(certificateRequest)

	assert.NotNil(t, createdCertificate)
	assert.Nil(t, err)

	kongApiAddress := os.Getenv(EnvKongApiHostAddress) + "/admin-api"
	unauthorisedClient := NewClient(&Config{HostAddress: kongApiAddress})

	certificate, err := unauthorisedClient.Certificates().GetById(*createdCertificate.Id)
	assert.NotNil(t, err)
	assert.Nil(t, certificate)

	results, err := unauthorisedClient.Certificates().List()
	assert.NotNil(t, err)
	assert.Nil(t, results)

	err = unauthorisedClient.Certificates().DeleteById(*createdCertificate.Id)
	assert.NotNil(t, err)

	certificateResult, err := unauthorisedClient.Certificates().Create(&CertificateRequest{
		Cert: String("public key-" + uuid.NewV4().String()),
		Key:  String("private key-" + uuid.NewV4().String()),
	})
	assert.Nil(t, certificateResult)
	assert.NotNil(t, err)

	updatedCertificate, err := unauthorisedClient.Certificates().UpdateById(*createdCertificate.Id, certificateRequest)
	assert.Nil(t, updatedCertificate)
	assert.NotNil(t, err)

	err = client.Plugins().DeleteById(createdPlugin.Id)
	assert.Nil(t, err)

	err = client.Apis().DeleteById(*createdApi.Id)
	assert.Nil(t, err)

}
