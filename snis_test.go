package gokong

import (
	"testing"

	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func Test_SnisCreate(t *testing.T) {

	client := NewClient(NewDefaultConfig())

	certificateRequest := &CertificateRequest{
		Cert: String("public key-" + uuid.NewV4().String()),
		Key:  String("private key-" + uuid.NewV4().String()),
	}

	certificate, err := client.Certificates().Create(certificateRequest)
	assert.Nil(t, err)
	assert.NotNil(t, certificate)

	snisRequest := &SnisRequest{
		Name:             "example.com",
		SslCertificateId: *certificate.Id,
	}

	result, err := client.Snis().Create(snisRequest)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, snisRequest.Name, result.Name)
	assert.Equal(t, snisRequest.SslCertificateId, result.SslCertificateId)
}

func Test_SnisCreateInvalid(t *testing.T) {

	client := NewClient(NewDefaultConfig())

	snisRequest := &SnisRequest{
		Name:             "example.com",
		SslCertificateId: "123",
	}

	result, err := client.Snis().Create(snisRequest)

	assert.NotNil(t, err)
	assert.Nil(t, result)
}

func Test_SnisGetByName(t *testing.T) {
	client := NewClient(NewDefaultConfig())

	certificateRequest := &CertificateRequest{
		Cert: String("public key-" + uuid.NewV4().String()),
		Key:  String("private key-" + uuid.NewV4().String()),
	}

	certificate, err := client.Certificates().Create(certificateRequest)
	assert.Nil(t, err)
	assert.NotNil(t, certificate)

	snisRequest := &SnisRequest{
		Name:             uuid.NewV4().String() + ".com",
		SslCertificateId: *certificate.Id,
	}

	sni, err := client.Snis().Create(snisRequest)

	assert.Nil(t, err)
	assert.NotNil(t, sni)

	result, err := client.Snis().GetByName(sni.Name)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, sni.Name, result.Name)
	assert.Equal(t, sni.SslCertificateId, result.SslCertificateId)
}

func Test_SnisList(t *testing.T) {

	client := NewClient(NewDefaultConfig())

	certificateRequest := &CertificateRequest{
		Cert: String("public key-" + uuid.NewV4().String()),
		Key:  String("private key-" + uuid.NewV4().String()),
	}

	certificate, err := client.Certificates().Create(certificateRequest)
	assert.Nil(t, err)
	assert.NotNil(t, certificate)

	snisRequest := &SnisRequest{
		Name:             uuid.NewV4().String() + ".com",
		SslCertificateId: *certificate.Id,
	}

	sni, err := client.Snis().Create(snisRequest)

	assert.Nil(t, err)
	assert.NotNil(t, sni)

	results, err := client.Snis().List()

	assert.Nil(t, err)
	assert.NotNil(t, results)
	assert.True(t, len(results.Results) > 0)

}

func Test_SnisDeleteByName(t *testing.T) {

	client := NewClient(NewDefaultConfig())

	certificateRequest := &CertificateRequest{
		Cert: String("public key-" + uuid.NewV4().String()),
		Key:  String("private key-" + uuid.NewV4().String()),
	}

	certificate, err := client.Certificates().Create(certificateRequest)
	assert.Nil(t, err)
	assert.NotNil(t, certificate)

	snisRequest := &SnisRequest{
		Name:             uuid.NewV4().String() + ".com",
		SslCertificateId: *certificate.Id,
	}

	sni, err := client.Snis().Create(snisRequest)

	assert.Nil(t, err)
	assert.NotNil(t, sni)

	err = client.Snis().DeleteByName(sni.Name)

	assert.Nil(t, err)

	result, err := client.Snis().GetByName(sni.Name)
	assert.Nil(t, err)
	assert.Nil(t, result)

}

func Test_SnisUpdateByName(t *testing.T) {

	client := NewClient(NewDefaultConfig())

	certificateRequest := &CertificateRequest{
		Cert: String("public key-" + uuid.NewV4().String()),
		Key:  String("private key-" + uuid.NewV4().String()),
	}

	certificate, err := client.Certificates().Create(certificateRequest)
	assert.Nil(t, err)
	assert.NotNil(t, certificate)

	snisRequest := &SnisRequest{
		Name:             uuid.NewV4().String() + ".com",
		SslCertificateId: *certificate.Id,
	}

	sni, err := client.Snis().Create(snisRequest)
	assert.Nil(t, err)
	assert.NotNil(t, sni)

	certificateRequest2 := &CertificateRequest{
		Cert: String("public key-" + uuid.NewV4().String()),
		Key:  String("private key-" + uuid.NewV4().String()),
	}

	certificate2, err := client.Certificates().Create(certificateRequest2)
	assert.Nil(t, err)
	assert.NotNil(t, certificate2)

	snisRequest.SslCertificateId = *certificate2.Id

	result, err := client.Snis().UpdateByName(snisRequest.Name, snisRequest)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, snisRequest.Name, result.Name)
	assert.Equal(t, *certificate2.Id, result.SslCertificateId)

}

func Test_SnisUpdateByNameInvalid(t *testing.T) {

	client := NewClient(NewDefaultConfig())

	certificateRequest := &CertificateRequest{
		Cert: String("public key-" + uuid.NewV4().String()),
		Key:  String("private key-" + uuid.NewV4().String()),
	}

	certificate, err := client.Certificates().Create(certificateRequest)
	assert.Nil(t, err)
	assert.NotNil(t, certificate)

	snisRequest := &SnisRequest{
		Name:             uuid.NewV4().String() + ".com",
		SslCertificateId: *certificate.Id,
	}

	sni, err := client.Snis().Create(snisRequest)
	assert.Nil(t, err)
	assert.NotNil(t, sni)

	snisRequest.SslCertificateId = "234"

	result, err := client.Snis().UpdateByName(snisRequest.Name, snisRequest)

	assert.NotNil(t, err)
	assert.Nil(t, result)

}

func Test_AllSniEndpointsShouldReturnErrorWhenRequestUnauthorised(t *testing.T) {

	//apiRequest := &ApiRequest{
	//	Name:        String("admin-api"),
	//	Uris:        StringSlice([]string{"/admin-api"}),
	//	UpstreamUrl: String("http://localhost:8001"),
	//}
	//
	//client := NewClient(NewDefaultConfig())
	//createdApi, err := client.Apis().Create(apiRequest)
	//
	//assert.Nil(t, err)
	//assert.NotNil(t, createdApi)
	//
	//consumerRequest := &ConsumerRequest{
	//	Username: "username-" + uuid.NewV4().String(),
	//	CustomId: "test-" + uuid.NewV4().String(),
	//}
	//
	//createdConsumer, err := client.Consumers().Create(consumerRequest)
	//
	//assert.Nil(t, err)
	//assert.NotNil(t, createdConsumer)
	//
	//pluginRequest := &PluginRequest{
	//	Name:  "key-auth",
	//	ApiId: *createdApi.Id,
	//	Config: map[string]interface{}{
	//		"hide_credentials": true,
	//	},
	//}
	//
	//createdPlugin, err := client.Plugins().Create(pluginRequest)
	//
	//assert.Nil(t, err)
	//assert.NotNil(t, createdPlugin)
	//
	//_, err = client.Consumers().CreatePluginConfig(createdConsumer.Id, "key-auth", "")
	//assert.Nil(t, err)
	//
	//certificate, err := client.Certificates().Create(&CertificateRequest{
	//	Cert: String("public key-" + uuid.NewV4().String()),
	//	Key:  String("private key-" + uuid.NewV4().String()),
	//})
	//assert.Nil(t, err)
	//assert.NotNil(t, certificate)
	//
	//snisRequest := &SnisRequest{
	//	Name:             uuid.NewV4().String() + ".example.com",
	//	SslCertificateId: *certificate.Id,
	//}
	//
	//createdSni, err := client.Snis().Create(snisRequest)
	//assert.Nil(t, err)
	//assert.NotNil(t, createdSni)
	//
	//kongApiAddress := os.Getenv(EnvKongApiHostAddress) + "/admin-api"
	//unauthorisedClient := NewClient(&Config{HostAddress: kongApiAddress})
	//
	//sni, err := unauthorisedClient.Snis().GetByName(createdSni.Name)
	//assert.NotNil(t, err)
	//assert.Nil(t, sni)
	//
	//results, err := unauthorisedClient.Snis().List()
	//assert.NotNil(t, err)
	//assert.Nil(t, results)
	//
	//err = unauthorisedClient.Snis().DeleteByName(createdSni.Name)
	//assert.NotNil(t, err)
	//
	//sniResult, err := unauthorisedClient.Snis().Create(&SnisRequest{
	//	Name:             uuid.NewV4().String() + ".example.com",
	//	SslCertificateId: *certificate.Id,
	//})
	//assert.Nil(t, sniResult)
	//assert.NotNil(t, err)
	//
	//updatedSni, err := unauthorisedClient.Snis().UpdateByName(createdSni.Name, &SnisRequest{
	//	Name:             uuid.NewV4().String() + ".example.com",
	//	SslCertificateId: *certificate.Id,
	//})
	//assert.Nil(t, updatedSni)
	//assert.NotNil(t, err)
	//
	//err = client.Plugins().DeleteById(createdPlugin.Id)
	//assert.Nil(t, err)
	//
	//err = client.Apis().DeleteById(*createdApi.Id)
	//assert.Nil(t, err)

}
