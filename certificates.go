package gokong

import (
	"encoding/json"
	"fmt"

	"github.com/parnurzeal/gorequest"
)

type CertificateClient struct {
	config *Config
}

type CertificateRequest struct {
	Cert string `json:"cert,omitempty"`
	Key  string `json:"key,omitempty"`
}

type Certificate struct {
	ID   string `json:"id,omitempty"`
	Cert string `json:"cert,omitempty"`
	Key  string `json:"key,omitempty"`
}

type Certificates struct {
	Results []*Certificate `json:"data,omitempty"`
	Total   int            `json:"total,omitempty"`
}

const CertificatesPath = "/certificates/"

func (certificateClient *CertificateClient) GetByID(id string) (*Certificate, error) {

	_, body, errs := gorequest.New().Get(certificateClient.config.HostAddress + CertificatesPath + id).End()
	if errs != nil {
		return nil, fmt.Errorf("could not get certificate, error: %v", errs)
	}

	certificate := &Certificate{}
	err := json.Unmarshal([]byte(body), certificate)
	if err != nil {
		return nil, fmt.Errorf("could not parse certificate get response, error: %v", err)
	}

	if certificate.ID == "" {
		return nil, nil
	}

	return certificate, nil
}

func (certificateClient *CertificateClient) Create(certificateRequest *CertificateRequest) (*Certificate, error) {

	_, body, errs := gorequest.New().Post(certificateClient.config.HostAddress + CertificatesPath).Send(certificateRequest).End()
	if errs != nil {
		return nil, fmt.Errorf("could not create new certificate, error: %v", errs)
	}

	createdCertificate := &Certificate{}
	err := json.Unmarshal([]byte(body), createdCertificate)
	if err != nil {
		return nil, fmt.Errorf("could not parse certificate creation response, error: %v", err)
	}

	if createdCertificate.ID == "" {
		return nil, fmt.Errorf("could not create certificate, error: %v", body)
	}

	return createdCertificate, nil
}

func (certificateClient *CertificateClient) DeleteByID(id string) error {

	res, _, errs := gorequest.New().Delete(certificateClient.config.HostAddress + CertificatesPath + id).End()
	if errs != nil {
		return fmt.Errorf("could not delete certificate, result: %v error: %v", res, errs)
	}

	return nil
}

func (certificateClient *CertificateClient) List() (*Certificates, error) {

	_, body, errs := gorequest.New().Get(certificateClient.config.HostAddress + CertificatesPath).End()
	if errs != nil {
		return nil, fmt.Errorf("could not get certificates, error: %v", errs)
	}

	certificates := &Certificates{}
	err := json.Unmarshal([]byte(body), certificates)
	if err != nil {
		return nil, fmt.Errorf("could not parse certificates list response, error: %v", err)
	}

	return certificates, nil
}

func (certificateClient *CertificateClient) UpdateByID(id string, certificateRequest *CertificateRequest) (*Certificate, error) {

	_, body, errs := gorequest.New().Patch(certificateClient.config.HostAddress + CertificatesPath + id).Send(certificateRequest).End()
	if errs != nil {
		return nil, fmt.Errorf("could not update certificate, error: %v", errs)
	}

	updatedCertificate := &Certificate{}
	err := json.Unmarshal([]byte(body), updatedCertificate)
	if err != nil {
		return nil, fmt.Errorf("could not parse certificate update response, error: %v", err)
	}

	if updatedCertificate.ID == "" {
		return nil, fmt.Errorf("could not update certificate, error: %v", body)
	}

	return updatedCertificate, nil
}
