package kongfeatures

import "testing"

var allVersions = []string { "0.11", "0.11.1", "0.11.2", "0.12", "0.13", "0.13.1", "0.14"}

func Test_FeatureApiIsSupported(t *testing.T)  {

	testData := map[string]bool {
		"0.11":true,
		"0.11.1":true,
		"0.11.2":true,
		"0.12":true,
		"0.12.3":true,
		"0.13":true,
		"0.13.1":true,
		"0.14":false,
	}

	for version, expected := range testData {
		result := IsSupported(Apis, version)

		if result != expected {
			t.Errorf("expected %v for version %s but found %v", expected, version, result)
		}
	}

}

func Test_FeatureCertificatesIsSupported(t *testing.T) {
	for _, version := range allVersions {
		result := IsSupported(Certificates, version)
		if !result {
			t.Errorf("version %s should be supported for certificates", version)
		}
	}
}

func Test_FeatureConsumersIsSupported(t *testing.T) {
	for _, version := range allVersions {
		result := IsSupported(Consumers, version)
		if !result {
			t.Errorf("version %s should be supported for consumers", version)
		}
	}
}

func Test_FeaturePluginsIsSupported(t *testing.T) {
	for _, version := range allVersions {
		result := IsSupported(Plugins, version)
		if !result {
			t.Errorf("version %s should be supported for plugins", version)
		}
	}
}

func Test_FeatureRoutesIsSupported(t *testing.T)  {

	testData := map[string]bool {
		"0.11":false,
		"0.11.1":false,
		"0.11.2":false,
		"0.12":false,
		"0.12.3":false,
		"0.13":true,
		"0.13.1":true,
		"0.14":true,
	}

	for version, expected := range testData {
		result := IsSupported(Routes, version)

		if result != expected {
			t.Errorf("expected %v for version %s but found %v", expected, version, result)
		}
	}

}

func Test_FeatureServicesIsSupported(t *testing.T)  {

	testData := map[string]bool {
		"0.11":false,
		"0.11.1":false,
		"0.11.2":false,
		"0.12":false,
		"0.12.3":false,
		"0.13":true,
		"0.13.1":true,
		"0.14":true,
	}

	for version, expected := range testData {
		result := IsSupported(Services, version)

		if result != expected {
			t.Errorf("expected %v for version %s but found %v", expected, version, result)
		}
	}

}

func Test_FeatureSnisIsSupported(t *testing.T) {
	for _, version := range allVersions {
		result := IsSupported(Snis, version)
		if !result {
			t.Errorf("version %s should be supported for snis", version)
		}
	}
}


func Test_FeatureUpstreamsIsSupported(t *testing.T)  {

	testData := map[string]bool {
		"0.11":false,
		"0.11.1":false,
		"0.11.2":false,
		"0.12":true,
		"0.12.3":true,
		"0.13":true,
		"0.13.1":true,
		"0.14":true,
	}

	for version, expected := range testData {
		result := IsSupported(Upstreams, version)

		if result != expected {
			t.Errorf("expected %v for version %s but found %v", expected, version, result)
		}
	}

}
