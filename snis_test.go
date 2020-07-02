package gokong

import (
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func Test_SnisCreate(t *testing.T) {

	client := NewClient(NewDefaultConfig())

	certificateRequest := &CertificateRequest{
		Cert: String(testCert1),
		Key:  String(testKey1),
	}

	certificate, err := client.Certificates().Create(certificateRequest)
	assert.Nil(t, err)
	assert.NotNil(t, certificate)

	snisRequest := &SnisRequest{
		Name:          "example.com",
		CertificateId: ToId(*certificate.Id),
		Tags:          []*string{String("my-tag")},
	}

	result, err := client.Snis().Create(snisRequest)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, snisRequest.Name, result.Name)
	assert.Equal(t, IdToString(snisRequest.CertificateId), IdToString(result.CertificateId))
	assert.Equal(t, snisRequest.Tags, result.Tags)

	err = client.Snis().DeleteByName(result.Name)
	assert.Nil(t, err)

	err = client.Certificates().DeleteById(*certificate.Id)
	assert.Nil(t, err)

}

func Test_SnisCreateInvalid(t *testing.T) {

	client := NewClient(NewDefaultConfig())

	snisRequest := &SnisRequest{
		Name:          "example.com",
		CertificateId: ToId("123"),
	}

	result, err := client.Snis().Create(snisRequest)

	assert.NotNil(t, err)
	assert.Nil(t, result)
}

func Test_SnisGetByName(t *testing.T) {
	client := NewClient(NewDefaultConfig())

	certificateRequest := &CertificateRequest{
		Cert: String(testCert1),
		Key:  String(testKey1),
	}

	certificate, err := client.Certificates().Create(certificateRequest)
	assert.Nil(t, err)
	assert.NotNil(t, certificate)

	snisRequest := &SnisRequest{
		Name:          uuid.NewV4().String() + ".com",
		CertificateId: ToId(*certificate.Id),
	}

	sni, err := client.Snis().Create(snisRequest)

	assert.Nil(t, err)
	assert.NotNil(t, sni)

	result, err := client.Snis().GetByName(sni.Name)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, sni.Name, result.Name)
	assert.Equal(t, IdToString(sni.CertificateId), IdToString(result.CertificateId))

	err = client.Snis().DeleteByName(result.Name)
	assert.Nil(t, err)

	err = client.Certificates().DeleteById(*certificate.Id)
	assert.Nil(t, err)
}

func Test_SnisList(t *testing.T) {

	client := NewClient(NewDefaultConfig())

	certificateRequest := &CertificateRequest{
		Cert: String(testCert1),
		Key:  String(testKey1),
	}

	certificate, err := client.Certificates().Create(certificateRequest)
	assert.Nil(t, err)
	assert.NotNil(t, certificate)

	snisRequest := &SnisRequest{
		Name:          uuid.NewV4().String() + ".com",
		CertificateId: ToId(*certificate.Id),
	}

	sni, err := client.Snis().Create(snisRequest)

	assert.Nil(t, err)
	assert.NotNil(t, sni)

	results, err := client.Snis().List()

	assert.Nil(t, err)
	assert.NotNil(t, results)
	assert.True(t, len(results.Results) > 0)

	for _, r := range results.Results {
		err = client.Snis().DeleteByName(r.Name)
		assert.Nil(t, err)
	}

	err = client.Certificates().DeleteById(*certificate.Id)
	assert.Nil(t, err)

}

func Test_SnisDeleteByName(t *testing.T) {

	client := NewClient(NewDefaultConfig())

	certificateRequest := &CertificateRequest{
		Cert: String(testCert1),
		Key:  String(testKey1),
	}

	certificate, err := client.Certificates().Create(certificateRequest)
	assert.Nil(t, err)
	assert.NotNil(t, certificate)

	snisRequest := &SnisRequest{
		Name:          uuid.NewV4().String() + ".com",
		CertificateId: ToId(*certificate.Id),
	}

	sni, err := client.Snis().Create(snisRequest)

	assert.Nil(t, err)
	assert.NotNil(t, sni)

	err = client.Snis().DeleteByName(sni.Name)

	assert.Nil(t, err)

	result, err := client.Snis().GetByName(sni.Name)
	assert.Nil(t, err)
	assert.Nil(t, result)

	err = client.Certificates().DeleteById(*certificate.Id)
	assert.Nil(t, err)

}

func Test_SnisUpdateByName(t *testing.T) {

	client := NewClient(NewDefaultConfig())

	certificateRequest := &CertificateRequest{
		Cert: String(testCert1),
		Key:  String(testKey1),
	}

	certificate, err := client.Certificates().Create(certificateRequest)
	assert.Nil(t, err)
	assert.NotNil(t, certificate)

	snisRequest := &SnisRequest{
		Name:          uuid.NewV4().String() + ".com",
		CertificateId: ToId(*certificate.Id),
	}

	sni, err := client.Snis().Create(snisRequest)
	assert.Nil(t, err)
	assert.NotNil(t, sni)

	certificateRequest2 := &CertificateRequest{
		Cert: String(testCert2),
		Key:  String(testKey2),
	}

	certificate2, err := client.Certificates().Create(certificateRequest2)
	assert.Nil(t, err)
	assert.NotNil(t, certificate2)

	snisRequest.CertificateId = ToId(*certificate2.Id)

	result, err := client.Snis().UpdateByName(snisRequest.Name, snisRequest)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, snisRequest.Name, result.Name)
	assert.Equal(t, *certificate2.Id, IdToString(result.CertificateId))

	err = client.Snis().DeleteByName(result.Name)
	assert.Nil(t, err)

	err = client.Certificates().DeleteById(*certificate.Id)
	assert.Nil(t, err)

}

func Test_SnisUpdateByNameInvalid(t *testing.T) {

	client := NewClient(NewDefaultConfig())

	certificateRequest := &CertificateRequest{
		Cert: String(testCert1),
		Key:  String(testKey1),
	}

	certificate, err := client.Certificates().Create(certificateRequest)
	assert.Nil(t, err)
	assert.NotNil(t, certificate)

	snisRequest := &SnisRequest{
		Name:          uuid.NewV4().String() + ".com",
		CertificateId: ToId(*certificate.Id),
	}

	sni, err := client.Snis().Create(snisRequest)
	assert.Nil(t, err)
	assert.NotNil(t, sni)

	snisRequest.CertificateId = ToId("234")

	result, err := client.Snis().UpdateByName(snisRequest.Name, snisRequest)

	assert.NotNil(t, err)
	assert.Nil(t, result)

	err = client.Snis().DeleteByName(sni.Name)
	assert.Nil(t, err)

	err = client.Certificates().DeleteById(*certificate.Id)
	assert.Nil(t, err)

}

func Test_AllSniEndpointsShouldReturnErrorWhenRequestUnauthorised(t *testing.T) {

	unauthorisedClient := NewClient(&Config{HostAddress: kong401Server})

	sni, err := unauthorisedClient.Snis().GetByName("foo")
	assert.NotNil(t, err)
	assert.Nil(t, sni)

	results, err := unauthorisedClient.Snis().List()
	assert.NotNil(t, err)
	assert.Nil(t, results)

	err = unauthorisedClient.Snis().DeleteByName("bar")
	assert.NotNil(t, err)

	sniResult, err := unauthorisedClient.Snis().Create(&SnisRequest{
		Name:          uuid.NewV4().String() + ".example.com",
		CertificateId: ToId(uuid.NewV4().String()),
	})
	assert.Nil(t, sniResult)
	assert.NotNil(t, err)

	updatedSni, err := unauthorisedClient.Snis().UpdateByName("foo", &SnisRequest{
		Name:          uuid.NewV4().String() + ".example.com",
		CertificateId: ToId(uuid.NewV4().String()),
	})
	assert.Nil(t, updatedSni)
	assert.NotNil(t, err)

}
