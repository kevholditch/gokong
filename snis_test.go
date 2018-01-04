package gokong

import (
	"testing"

	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func Test_SnisCreate(t *testing.T) {

	client := NewClient(NewDefaultConfig())

	certificateRequest := &CertificateRequest{
		Cert: "public key-" + uuid.NewV4().String(),
		Key:  "private key-" + uuid.NewV4().String(),
	}

	certificate, err := client.Certificates().Create(certificateRequest)
	assert.Nil(t, err)
	assert.NotNil(t, certificate)

	snisRequest := &SnisRequest{
		Name:             "example.com",
		SslCertificateID: certificate.ID,
	}

	result, err := client.Snis().Create(snisRequest)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, snisRequest.Name, result.Name)
	assert.Equal(t, snisRequest.SslCertificateID, result.SslCertificateID)
}

func Test_SnisCreateInvalid(t *testing.T) {

	client := NewClient(NewDefaultConfig())

	snisRequest := &SnisRequest{
		Name:             "example.com",
		SslCertificateID: "123",
	}

	result, err := client.Snis().Create(snisRequest)

	assert.NotNil(t, err)
	assert.Nil(t, result)
}

func Test_SnisGetByName(t *testing.T) {
	client := NewClient(NewDefaultConfig())

	certificateRequest := &CertificateRequest{
		Cert: "public key-" + uuid.NewV4().String(),
		Key:  "private key-" + uuid.NewV4().String(),
	}

	certificate, err := client.Certificates().Create(certificateRequest)
	assert.Nil(t, err)
	assert.NotNil(t, certificate)

	snisRequest := &SnisRequest{
		Name:             uuid.NewV4().String() + ".com",
		SslCertificateID: certificate.ID,
	}

	sni, err := client.Snis().Create(snisRequest)

	assert.Nil(t, err)
	assert.NotNil(t, sni)

	result, err := client.Snis().GetByName(sni.Name)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, sni.Name, result.Name)
	assert.Equal(t, sni.SslCertificateID, result.SslCertificateID)
}

func Test_SnisGetNonExistentByName(t *testing.T) {

	result, err := NewClient(NewDefaultConfig()).APIs().GetByID("dd1de132-ede6-4534-bd65-57bcf0beba4b")

	assert.Nil(t, err)
	assert.Nil(t, result)
}

func Test_SnisList(t *testing.T) {

	client := NewClient(NewDefaultConfig())

	certificateRequest := &CertificateRequest{
		Cert: "public key-" + uuid.NewV4().String(),
		Key:  "private key-" + uuid.NewV4().String(),
	}

	certificate, err := client.Certificates().Create(certificateRequest)
	assert.Nil(t, err)
	assert.NotNil(t, certificate)

	snisRequest := &SnisRequest{
		Name:             uuid.NewV4().String() + ".com",
		SslCertificateID: certificate.ID,
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
		Cert: "public key-" + uuid.NewV4().String(),
		Key:  "private key-" + uuid.NewV4().String(),
	}

	certificate, err := client.Certificates().Create(certificateRequest)
	assert.Nil(t, err)
	assert.NotNil(t, certificate)

	snisRequest := &SnisRequest{
		Name:             uuid.NewV4().String() + ".com",
		SslCertificateID: certificate.ID,
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
		Cert: "public key-" + uuid.NewV4().String(),
		Key:  "private key-" + uuid.NewV4().String(),
	}

	certificate, err := client.Certificates().Create(certificateRequest)
	assert.Nil(t, err)
	assert.NotNil(t, certificate)

	snisRequest := &SnisRequest{
		Name:             uuid.NewV4().String() + ".com",
		SslCertificateID: certificate.ID,
	}

	sni, err := client.Snis().Create(snisRequest)
	assert.Nil(t, err)
	assert.NotNil(t, sni)

	certificateRequest2 := &CertificateRequest{
		Cert: "public key-" + uuid.NewV4().String(),
		Key:  "private key-" + uuid.NewV4().String(),
	}

	certificate2, err := client.Certificates().Create(certificateRequest2)
	assert.Nil(t, err)
	assert.NotNil(t, certificate2)

	snisRequest.SslCertificateID = certificate2.ID

	result, err := client.Snis().UpdateByName(snisRequest.Name, snisRequest)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, snisRequest.Name, result.Name)
	assert.Equal(t, certificate2.ID, result.SslCertificateID)

}

func Test_SnisUpdateByNameInvalid(t *testing.T) {

	client := NewClient(NewDefaultConfig())

	certificateRequest := &CertificateRequest{
		Cert: "public key-" + uuid.NewV4().String(),
		Key:  "private key-" + uuid.NewV4().String(),
	}

	certificate, err := client.Certificates().Create(certificateRequest)
	assert.Nil(t, err)
	assert.NotNil(t, certificate)

	snisRequest := &SnisRequest{
		Name:             uuid.NewV4().String() + ".com",
		SslCertificateID: certificate.ID,
	}

	sni, err := client.Snis().Create(snisRequest)
	assert.Nil(t, err)
	assert.NotNil(t, sni)

	snisRequest.SslCertificateID = "234"

	result, err := client.Snis().UpdateByName(snisRequest.Name, snisRequest)

	assert.NotNil(t, err)
	assert.Nil(t, result)

}
