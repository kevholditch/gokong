package kongfeatures

import "github.com/blang/semver"

type KongFeature int

const (
	Apis KongFeature = iota
	Certificates
	Consumers
	Plugins
	Routes
	Services
	Snis
	Upstreams
)

func IsSupported(feature KongFeature, version string) bool  {

	v, err := semver.Make(version)

	if err != nil {
		v, _ = semver.Make(version + ".0")
	}

	switch feature {
		case Apis:
			apiVersion, _ := semver.New("0.14.0")
			return v.LT(*apiVersion)
		case Routes:
			routeVersion, _ := semver.New("0.13.0")
			return v.GTE(*routeVersion)
		case Services:
			serviceVersion, _ := semver.New("0.13.0")
			return v.GTE(*serviceVersion)
		case Upstreams:
			upstreamVersion, _ := semver.New("0.12.0")
			return v.GTE(*upstreamVersion)

	}

	return true
}
