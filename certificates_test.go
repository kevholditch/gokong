package gokong

import (
	"testing"

	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func Test_CertificatesGetByID(t *testing.T) {
	certificateRequest := &CertificateRequest{
		Cert: "public key-" + uuid.NewV4().String(),
		Key:  "private key-" + uuid.NewV4().String(),
	}

	client := NewClient(NewDefaultConfig())
	createdCertificate, err := client.Certificates().Create(certificateRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdCertificate)

	result, err := client.Certificates().GetByID(createdCertificate.ID)

	assert.NotNil(t, result)
	assert.Equal(t, createdCertificate, result)

}

func Test_CertificatesGetNonExistentByID(t *testing.T) {

	result, err := NewClient(NewDefaultConfig()).Certificates().GetByID("8df4d1ed-c973-4b9a-868d-3e67d5c417da")

	assert.Nil(t, err)
	assert.Nil(t, result)
}

func Test_CertificatesCreate(t *testing.T) {
	certificateRequest := &CertificateRequest{
		Cert: "public key-" + uuid.NewV4().String(),
		Key:  "private key-" + uuid.NewV4().String(),
	}

	result, err := NewClient(NewDefaultConfig()).Certificates().Create(certificateRequest)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.True(t, result.ID != "")
	assert.Equal(t, certificateRequest.Cert, result.Cert)
	assert.Equal(t, certificateRequest.Key, result.Key)

}

func Test_CertificatesCreateInvalid(t *testing.T) {
	certificateRequest := &CertificateRequest{}

	result, err := NewClient(NewDefaultConfig()).Certificates().Create(certificateRequest)

	assert.NotNil(t, err)
	assert.Nil(t, result)

}

func Test_CertificatesUpdateByID(t *testing.T) {
	certificateRequest := &CertificateRequest{
		Cert: "public key-" + uuid.NewV4().String(),
		Key:  "private key-" + uuid.NewV4().String(),
	}

	client := NewClient(NewDefaultConfig())
	createdCertificate, err := client.Certificates().Create(certificateRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdCertificate)

	certificateRequest.Cert = "public key-" + uuid.NewV4().String()

	result, err := client.Certificates().UpdateByID(createdCertificate.ID, certificateRequest)

	assert.Equal(t, certificateRequest.Cert, result.Cert)
	assert.Equal(t, certificateRequest.Key, result.Key)

}

func Test_CertificatesUpdateByIDInvalid(t *testing.T) {
	certificateRequest := &CertificateRequest{
		Cert: "public key-" + uuid.NewV4().String(),
		Key:  "private key-" + uuid.NewV4().String(),
	}

	client := NewClient(NewDefaultConfig())
	createdCertificate, err := client.Certificates().Create(certificateRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdCertificate)

	certificateRequest.Cert = ""
	certificateRequest.Key = ""

	result, err := client.Certificates().UpdateByID(createdCertificate.ID, certificateRequest)

	assert.NotNil(t, err)
	assert.Nil(t, result)

}

func Test_CertificatesDeleteByID(t *testing.T) {
	certificateRequest := &CertificateRequest{
		Cert: "public key-" + uuid.NewV4().String(),
		Key:  "private key-" + uuid.NewV4().String(),
	}

	client := NewClient(NewDefaultConfig())
	createdCertificate, err := client.Certificates().Create(certificateRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdCertificate)

	err = client.Certificates().DeleteByID(createdCertificate.ID)

	assert.Nil(t, err)

	deletedCertificate, err := client.Certificates().GetByID(createdCertificate.ID)

	assert.Nil(t, err)
	assert.Nil(t, deletedCertificate)

}

func Test_CertificatesList(t *testing.T) {
	client := NewClient(NewDefaultConfig())

	certificateRequest := &CertificateRequest{
		Cert: "public key-" + uuid.NewV4().String(),
		Key:  "private key-" + uuid.NewV4().String(),
	}

	createdCertificate, err := client.Certificates().Create(certificateRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdCertificate)

	certificateRequest2 := &CertificateRequest{
		Cert: "public key-" + uuid.NewV4().String(),
		Key:  "private key-" + uuid.NewV4().String(),
	}

	createdCertificate2, err := client.Certificates().Create(certificateRequest2)

	assert.Nil(t, err)
	assert.NotNil(t, createdCertificate2)

	results, err := client.Certificates().List()

	assert.Nil(t, err)
	assert.True(t, len(results.Results) > 1)

}
