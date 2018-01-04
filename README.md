[![Build Status](https://travis-ci.org/kevholditch/gokong.svg?branch=master)](https://travis-ci.org/kevholditch/gokong)

GoKong
======
A kong go client fully tested with no mocks!!

## GoKong
GoKong is a easy to use api client for [kong](https://getkong.org/).  The difference with the gokong library is all of its tests are written against a real running kong running inside a docker container, yep that's right you won't see a horrible mock anywhere!!

## Importing

To add gokong via `go get`:
```
go get github.com/kevholditch/gokong
```

To add gokong via `govendor`:
```
govendor fetch github.com/kevholditch/gokong
```

## Usage

Import gokong
```go
import (
  "github.com/kevholditch/gokong"
)
```

To create a default config for use with the client:
```go
config := gokong.NewDefaultConfig()
```

`NewDefaultConfig` creates a config with the host address set to the value of the env variable `KONG_ADMIN_ADDR`.
If the env variable is not set then the address is defaulted to `http://localhost:8001`.

You can of course create your own config with the address set to whatever you want:
```go
config := gokong.Config{HostAddress:"http://localhost:1234"}
```


Getting the status of the kong server:
```go
kongClient := gokong.NewClient(gokong.NewDefaultConfig())
status, err := kongClient.Status().Get()
```

Gokong is fluent so we can combine the above two lines into one:
```go
status, err := gokong.NewClient(gokong.NewDefaultConfig()).Status().Get()
```

## APIs
Create a new API ([for more information on the API fields see the Kong documentation](https://getkong.org/docs/0.11.x/admin-api/#api-object):
```go
apiRequest := &gokong.APIRequest{
	Name:                   "Example",
	Hosts:                  []string{"example.com"},
	URIs:                   []string{"/example"},
	Methods:                []string{"GET", "POST"},
	UpstreamURL:            "http://localhost:4140/testservice",
	StripURI:               true,
	PreserveHost:           true,
	Retries:                3,
	UpstreamConnectTimeout: 1000,
	UpstreamSendTimeout:    2000,
	UpstreamReadTimeout:    3000,
	HTTPSOnly:              true,
	HTTPIfTerminated:       true,
}

api, err := gokong.NewClient(gokong.NewDefaultConfig()).APIs().Create(apiRequest)
```

Get an API by id:
```go
api, err := gokong.NewClient(gokong.NewDefaultConfig()).APIs().GetById("cdf5372e-1c10-4ea5-a3dd-1e4c31bb99f5")
```

Get an API by name:
```go
api, err :=  gokong.NewClient(gokong.NewDefaultConfig()).APIs().GetByName("Example")
```

List all APIs:
```go
apis, err := gokong.NewClient(gokong.NewDefaultConfig()).APIs().List()
```

List all APIs with a filter:
```go
apis, err := gokong.NewClient(gokong.NewDefaultConfig()).APIs().ListFiltered(&gokong.APIFilter{Id:"936ad391-c30d-43db-b624-2f820d6fd38d", Name:"MyAPI"})
```

Delete an API by id:
```go
err :=  gokong.NewClient(gokong.NewDefaultConfig()).APIs().DeleteById("f138641a-a15b-43c3-bd76-7157a68eae24")
```

Delete an API by name:
```go
err :=  gokong.NewClient(gokong.NewDefaultConfig()).APIs().DeleteByName("Example")
```

Update an API by id:
```go
apiRequest := &gokong.APIRequest{
  Name:                   "Example",
  Hosts:                  []string{"example.com"},
  URIs:                   []string{"/example"},
  Methods:                []string{"GET", "POST"},
  UpstreamURL:            "http://localhost:4140/testservice",
  StripURI:               true,
  PreserveHost:           true,
  Retries:                3,
  UpstreamConnectTimeout: 1000,
  UpstreamSendTimeout:    2000,
  UpstreamReadTimeout:    3000,
  HTTPSOnly:              true,
  HTTPIfTerminated:       true,
}

updatedAPI, err :=  gokong.NewClient(gokong.NewDefaultConfig()).APIs().UpdateById("1213a00d-2b12-4d65-92ad-5a02d6c710c2", apiRequest)
```

Update an API by name:
```go
apiRequest := &gokong.APIRequest{
  Name:                   "Example",
  Hosts:                  []string{"example.com"},
  URIs:                   []string{"/example"},
  Methods:                []string{"GET", "POST"},
  UpstreamURL:            "http://localhost:4140/testservice",
  StripURI:               true,
  PreserveHost:           true,
  Retries:                3,
  UpstreamConnectTimeout: 1000,
  UpstreamSendTimeout:    2000,
  UpstreamReadTimeout:    3000,
  HTTPSOnly:              true,
  HTTPIfTerminated:       true,
}

updatedAPI, err :=  gokong.NewClient(gokong.NewDefaultConfig()).APIs().UpdateByName("Example", apiRequest)
```


## Consumers
Create a new Consumer ([for more information on the Consumer Fields see the Kong documentation](https://getkong.org/docs/0.11.x/admin-api/#consumer-object)):
```go
consumerRequest := &gokong.ConsumerRequest{
  Username: "User1",
  CustomId: "SomeId",
}

consumer, err := gokong.NewClient(gokong.NewDefaultConfig()).Consumers().Create(consumerRequest)
```

Get a Consumer by id:
```go
consumer, err := gokong.NewClient(gokong.NewDefaultConfig()).Consumers().GetById("e8ccbf13-a662-45be-9b6a-b549cc739c18")
```

Get a Consumer by username:
```go
consumer, err := gokong.NewClient(gokong.NewDefaultConfig()).Consumers().GetByUsername("User1")
```

List all Consumers:
```go
consumers, err := gokong.NewClient(gokong.NewDefaultConfig()).Consumers().List()
```

List all Consumers with a filter:
```go
consumers, err := gokong.NewClient(gokong.NewDefaultConfig()).Consumers().ListFiltered(&gokong.ConsumerFilter{CustomId:"1234", Username: "User1"})
```

Delete a Consumer by id:
```go
err :=  gokong.NewClient(gokong.NewDefaultConfig()).Consumers().DeleteById("7c8741b7-3cf5-4d90-8674-b34153efbcd6")
```

Delete a Consumer by username:
```go
err :=  gokong.NewClient(gokong.NewDefaultConfig()).Consumers().DeleteByUsername("User1")
```

Update a Consumer by id:
```go
consumerRequest := &gokong.ConsumerRequest{
  Username: "User1",
  CustomId: "SomeId",
}

updatedConsumer, err :=  gokong.NewClient(gokong.NewDefaultConfig()).Consumers().UpdateById("44a37c3d-a252-4968-ab55-58c41b0289c2", consumerRequest)
```

Update a Consumer by username:
```go
consumerRequest := &gokong.ConsumerRequest{
  Username: "User2",
  CustomId: "SomeId",
}

updatedConsumer, err :=  gokong.NewClient(gokong.NewDefaultConfig()).Consumers().UpdateByUsername("User2", consumerRequest)
```

## Plugins
Create a new Plugin to be applied to all APIs and consumers do not set `APIId` or `ConsumerId`.  Not all plugins can be configured in this way
 ([for more information on the Plugin Fields see the Kong documentation](https://getkong.org/docs/0.11.x/admin-api/#add-plugin)):

```go
pluginRequest := &gokong.PluginRequest{
  Name: "response-ratelimiting",
  Config: map[string]interface{}{
    "limits.sms.minute": 20,
  },
}

createdPlugin, err := gokong.NewClient(gokong.NewDefaultConfig()).Plugins().Create(pluginRequest)
```

Create a new Plugin for a single API (only set `APIId`), not all plugins can be configured in this way ([for more information on the Plugin Fields see the Kong documentation](https://getkong.org/docs/0.11.x/admin-api/#plugin-object)):
```go
client := gokong.NewClient(gokong.NewDefaultConfig())

apiRequest := &gokong.APIRequest{
  Name:                   "test-api",
  Hosts:                  []string{"example.com"},
  URIs:                   []string{"/example"},
  Methods:                []string{"GET", "POST"},
  UpstreamURL:            "http://localhost:4140/testservice",
  StripURI:               true,
  PreserveHost:           true,
  Retries:                3,
  UpstreamConnectTimeout: 1000,
  UpstreamSendTimeout:    2000,
  UpstreamReadTimeout:    3000,
  HTTPSOnly:              true,
  HTTPIfTerminated:       true,
}

createdAPI, err := client.APIs().Create(apiRequest)

pluginRequest := &gokong.PluginRequest{
  Name: "response-ratelimiting",
  APIId: createdAPI.Id,
  Config: map[string]interface{}{
    "limits.sms.minute": 20,
  },
}

createdPlugin, err :=  client.Plugins().Create(pluginRequest)
```

Create a new Plugin for a single Consumer (only set `ConsumerId`), Not all plugins can be configured in this way ([for more information on the Plugin Fields see the Kong documentation](https://getkong.org/docs/0.11.x/admin-api/#plugin-object)):
```go
client := gokong.NewClient(gokong.NewDefaultConfig())

consumerRequest := &gokong.ConsumerRequest{
  Username: "User1",
  CustomId: "test",
}

createdConsumer, err := client.Consumers().Create(consumerRequest)

pluginRequest := &gokong.PluginRequest{
  Name: "response-ratelimiting",
  ConsumerId: createdConsumer.Id,
  Config: map[string]interface{}{
    "limits.sms.minute": 20,
  },
}

createdPlugin, err :=  client.Plugins().Create(pluginRequest)
```

Create a new Plugin for a single Consumer and API (set `ConsumerId` and `APIId`), Not all plugins can be configured in this way ([for more information on the Plugin Fields see the Kong documentation](https://getkong.org/docs/0.11.x/admin-api/#plugin-object)):
```go
client := gokong.NewClient(gokong.NewDefaultConfig())

consumerRequest := &gokong.ConsumerRequest{
  Username: "User1",
  CustomId: "test",
}

createdConsumer, err := client.Consumers().Create(consumerRequest)

apiRequest := &gokong.APIRequest{
  Name:                   "test-api",
  Hosts:                  []string{"example.com"},
  URIs:                   []string{"/example"},
  Methods:                []string{"GET", "POST"},
  UpstreamURL:            "http://localhost:4140/testservice",
  StripURI:               true,
  PreserveHost:           true,
  Retries:                3,
  UpstreamConnectTimeout: 1000,
  UpstreamSendTimeout:    2000,
  UpstreamReadTimeout:    3000,
  HTTPSOnly:              true,
  HTTPIfTerminated:       true,
}

createdAPI, err := client.APIs().Create(apiRequest)

pluginRequest := &gokong.PluginRequest{
  Name:       "response-ratelimiting",
  ConsumerId: createdConsumer.Id,
  APIId:      createdAPI.Id,
  Config: map[string]interface{}{
    "limits.sms.minute": 20,
  },
}

createdPlugin, err :=  client.Plugins().Create(pluginRequest)
```

Get a plugin by id:
```go
plugin, err := gokong.NewClient(gokong.NewDefaultConfig()).Plugins().GetById("04bda233-d035-4b8a-8cf2-a53f3dd990f3")
```

List all plugins:
```go
plugins, err := gokong.NewClient(gokong.NewDefaultConfig()).Plugins().List()
```

List all plugins with a filter:
```go
plugins, err := gokong.NewClient(gokong.NewDefaultConfig()).Plugins().ListFiltered(&gokong.PluginFilter{Name: "response-ratelimiting", ConsumerId: "7009a608-b40c-4a21-9a90-9219d5fd1ac7"})
```

Delete a plugin by id:
```go
err := gokong.NewClient(gokong.NewDefaultConfig()).Plugins().DeleteById("f2bbbab8-3e6f-4d9d-bada-d486600b3b4c")
```

Update a plugin by id:
```go
updatePluginRequest := &gokong.PluginRequest{
  Name:       "response-ratelimiting",
  ConsumerId: createdConsumer.Id,
  APIId:      createdAPI.Id,
  Config: map[string]interface{}{
    "limits.sms.minute": 20,
  },
}

updatedPlugin, err := gokong.NewClient(gokong.NewDefaultConfig()).Plugins().UpdateById("70692eed-2293-486d-b992-db44a6459360", updatePluginRequest)
```

## Certificates
Create a Certificate ([for more information on the Certificate Fields see the Kong documentation](https://getkong.org/docs/0.11.x/admin-api/#certificate-object)):

```go
certificateRequest := &gokong.CertificateRequest{
  Cert: "public key --- 123",
  Key:  "private key --- 456",
}

createdCertificate, err := gokong.NewClient(gokong.NewDefaultConfig()).Certificates().Create(certificateRequest)
```

Get a Certificate by id:
```go
certificate, err := gokong.NewClient(gokong.NewDefaultConfig()).Certificates().GetById("0408cbd4-e856-4565-bc11-066326de9231")
```

List all certificates:
```go
certificates, err := gokong.NewClient(gokong.NewDefaultConfig()).Certificates().List()
```

Delete a Certificate:
```go
err := gokong.NewClient(gokong.NewDefaultConfig()).Certificates().DeleteById("db884cf2-9dd7-4e33-9ef5-628165076a42")
```

Update a Certificate:
```go
updateCertificateRequest := &gokong.CertificateRequest{
  Cert: "public key --- 789",
  Key:  "private key --- 111",
}

updatedCertificate, err := gokong.NewClient(gokong.NewDefaultConfig()).Certificates().UpdateById("1dc11281-30a6-4fb9-aec2-c6ff33445375", updateCertificateRequest)
```


## SNIs
Create an SNI ([for more information on the Sni Fields see the Kong documentation](https://getkong.org/docs/0.11.x/admin-api/#sni-objects)):
```go
client := gokong.NewClient(gokong.NewDefaultConfig())

certificateRequest := &gokong.CertificateRequest{
  Cert: "public key --- 123",
  Key:  "private key --- 111",
}

certificate, err := client.Certificates().Create(certificateRequest)

snisRequest := &gokong.SnisRequest{
  Name:             "example.com",
  SslCertificateId: certificate.Id,
}

sni, err := client.Snis().Create(snisRequest)
```

Get an SNI by name:
```go
sni, err := client.Snis().GetByName("example.com")
```

List all SNIs:
```
snis, err := client.Snis().List()
```

Delete an SNI by name:
```go
err := client.Snis().DeleteByName("example.com")
```

Update an SNI by name:
```go
updateSniRequest := &gokong.SnisRequest{
  Name:             "example.com",
  SslCertificateId: "a9797703-3ae6-44a9-9f0a-4ebb5d7f301f",
}

updatedSni, err := client.Snis().UpdateByName("example.com", updateSniRequest)
```

## Upstreams
Create an Upstream ([for more information on the Upstream Fields see the Kong documentation](https://getkong.org/docs/0.11.x/admin-api/#upstream-objects)):
```go
upstreamRequest := &gokong.UpstreamRequest{
  Name: "test-upstream",
  Slots: 10,
}

createdUpstream, err := gokong.NewClient(gokong.NewDefaultConfig()).Upstreams().Create(upstreamRequest)
```

Get an Upstream by id:
```go
upstream, err := gokong.NewClient(gokong.NewDefaultConfig()).Upstreams().GetById("3705d962-caa8-4d0b-b291-4f0e85fe227a")
```

Get an Upstream by name:
```go
upstream, err := gokong.NewClient(gokong.NewDefaultConfig()).Upstreams().GetByName("test-upstream")
```

List all Upstreams:
```go
upstreams, err := gokong.NewClient(gokong.NewDefaultConfig()).Upstreams().List()
```

List all Upstreams with a filter:
```go
upstreams, err := gokong.NewClient(gokong.NewDefaultConfig()).Upstreams().ListFiltered(&gokong.UpstreamFilter{Name:"test-upstream", Slots:10})
```

Delete an Upstream by id:
```go
err := gokong.NewClient(gokong.NewDefaultConfig()).Upstreams().DeleteById("3a46b122-47ee-4c5d-b2de-49be84a672e6")
```

Delete an Upstream by name:
```go
err := gokong.NewClient(gokong.NewDefaultConfig()).Upstreams().DeleteById("3a46b122-47ee-4c5d-b2de-49be84a672e6")
```

Delete an Upstream by id:
```go
err := gokong.NewClient(gokong.NewDefaultConfig()).Upstreams().DeleteByName("test-upstream")
```

Update an Upstream by id:
```
updateUpstreamRequest := &gokong.UpstreamRequest{
  Name: "test-upstream",
  Slots: 10,
}

updatedUpstream, err := gokong.NewClient(gokong.NewDefaultConfig()).Upstreams().UpdateById("3a46b122-47ee-4c5d-b2de-49be84a672e6", updateUpstreamRequest)
```

Update an Upstream by name:
```go
updateUpstreamRequest := &gokong.UpstreamRequest{
  Name: "test-upstream",
  Slots: 10,
}

updatedUpstream, err := gokong.NewClient(gokong.NewDefaultConfig()).Upstreams().UpdateByName("test-upstream", updateUpstreamRequest)
```

# Contributing
I would love to get contributions to the project so please feel free to submit a PR.  To setup your dev station you need go and docker installed.

Once you have cloned the repository the `make` command will build the code and run all of the tests.  If they all pass then you are good to go!

If when you run the make command you get the following error:
```
gofmt needs running on the following files:
```
Then all you need to do is run `make fmt` this will reformat all of the code (I know awesome)!!

