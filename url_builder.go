package konggo

import "strings"

type UrlBuilder struct {
	url string
}

func NewUrlBuilder(hostAddress string) *UrlBuilder {
	return &UrlBuilder{url: strings.TrimRight(hostAddress, "/")}
}

func (urlBuilder *UrlBuilder) Status() *UrlBuilder {
	return &UrlBuilder{url: urlBuilder.url + "/status"}
}

func (urlBuilder *UrlBuilder) Apis() *UrlBuilder {
	return &UrlBuilder{url: urlBuilder.url + "/apis/"}
}

func (urlBuilder *UrlBuilder) Build() string {
	return urlBuilder.url
}
