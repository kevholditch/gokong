package konggo

type KongUrlBuilder struct {
	url string
}

func NewUrlBuilder(hostAddress string) *KongUrlBuilder {
	return &KongUrlBuilder{url: hostAddress}
}

func (kongUrlBuilder *KongUrlBuilder) Status() *KongUrlBuilder {
	return &KongUrlBuilder{url: kongUrlBuilder.url + "/status"}
}

func (kongUrlBuilder *KongUrlBuilder) Build() string {
	return kongUrlBuilder.url
}
