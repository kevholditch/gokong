package gokong

import (
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func Test_CertificatesGetById(t *testing.T) {

	certificateRequest := &CertificateRequest{
		Cert: String(testCert1),
		Key:  String(testKey1),
	}

	client := NewClient(NewDefaultConfig())
	createdCertificate, err := client.Certificates().Create(certificateRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdCertificate)

	result, err := client.Certificates().GetById(*createdCertificate.Id)

	assert.NotNil(t, result)
	assert.Equal(t, createdCertificate, result)

	err = client.Certificates().DeleteById(*createdCertificate.Id)

}

func Test_CertificatesGetNonExistentById(t *testing.T) {

	result, err := NewClient(NewDefaultConfig()).Certificates().GetById("8df4d1ed-c973-4b9a-868d-3e67d5c417da")

	assert.Nil(t, err)
	assert.Nil(t, result)
}

func Test_CertificatesCreate(t *testing.T) {
	certificateRequest := &CertificateRequest{
		Cert: String(testCert1),
		Key:  String(testKey1),
		Tags: []*string{String("my-tag")},
	}

	client := NewClient(NewDefaultConfig())
	result, err := client.Certificates().Create(certificateRequest)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.True(t, result.Id != nil)
	assert.Equal(t, certificateRequest.Cert, result.Cert)
	assert.Equal(t, certificateRequest.Key, result.Key)
	assert.Equal(t, certificateRequest.Tags, result.Tags)

	err = client.Certificates().DeleteById(*result.Id)
	assert.Nil(t, err)

}

func Test_CertificatesCreateInvalid(t *testing.T) {
	certificateRequest := &CertificateRequest{}

	result, err := NewClient(NewDefaultConfig()).Certificates().Create(certificateRequest)

	assert.NotNil(t, err)
	assert.Nil(t, result)

}

func Test_CertificatesUpdateById(t *testing.T) {
	certificateRequest := &CertificateRequest{
		Cert: String(testCert1),
		Key:  String(testKey1),
	}

	client := NewClient(NewDefaultConfig())
	createdCertificate, err := client.Certificates().Create(certificateRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdCertificate)

	certificateRequest.Cert = String(testCert2)
	certificateRequest.Key = String(testKey2)

	result, err := client.Certificates().UpdateById(*createdCertificate.Id, certificateRequest)

	assert.Nil(t, err)
	assert.Equal(t, certificateRequest.Cert, result.Cert)
	assert.Equal(t, certificateRequest.Key, result.Key)

	err = client.Certificates().DeleteById(*result.Id)
	assert.Nil(t, err)

}

func Test_CertificatesUpdateByIdInvalid(t *testing.T) {
	certificateRequest := &CertificateRequest{
		Cert: String(testCert1),
		Key:  String(testKey1),
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

	err = client.Certificates().DeleteById(*createdCertificate.Id)
	assert.Nil(t, err)

}

func Test_CertificatesDeleteById(t *testing.T) {
	certificateRequest := &CertificateRequest{
		Cert: String(testCert1),
		Key:  String(testKey1),
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
		Cert: String(testCert1),
		Key:  String(testKey1),
	}

	createdCertificate, err := client.Certificates().Create(certificateRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdCertificate)

	certificateRequest2 := &CertificateRequest{
		Cert: String(testCert2),
		Key:  String(testKey2),
	}

	createdCertificate2, err := client.Certificates().Create(certificateRequest2)

	assert.Nil(t, err)
	assert.NotNil(t, createdCertificate2)

	results, err := client.Certificates().List()

	assert.Nil(t, err)
	assert.True(t, len(results.Results) > 1)

	for _, result := range results.Results {
		err = client.Certificates().DeleteById(*result.Id)
		assert.Nil(t, err)
	}

}

func Test_AllCertificateEndpointsShouldReturnErrorWhenRequestUnauthorised(t *testing.T) {

	unauthorisedClient := NewClient(&Config{HostAddress: kong401Server})

	certificate, err := unauthorisedClient.Certificates().GetById(uuid.NewV4().String())
	assert.NotNil(t, err)
	assert.Nil(t, certificate)

	results, err := unauthorisedClient.Certificates().List()
	assert.NotNil(t, err)
	assert.Nil(t, results)

	err = unauthorisedClient.Certificates().DeleteById(uuid.NewV4().String())
	assert.NotNil(t, err)

	certificateResult, err := unauthorisedClient.Certificates().Create(&CertificateRequest{
		Cert: String(testCert1),
		Key:  String(testKey1),
	})
	assert.Nil(t, certificateResult)
	assert.NotNil(t, err)

	updatedCertificate, err := unauthorisedClient.Certificates().UpdateById(uuid.NewV4().String(), &CertificateRequest{})
	assert.Nil(t, updatedCertificate)
	assert.NotNil(t, err)

}
