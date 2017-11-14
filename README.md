[![Build Status](https://travis-ci.org/kevholditch/konggo.svg?branch=master)](https://travis-ci.org/kevholditch/konggo)

Konggo
======
A kong go client fully tested with no mocks!!

## Kong go
Kong go is a easy to use api client for [kong](https://getkong.org/).  The difference with the konggo library is all of its tests are written against a real running kong running inside a docker container, yep that's right you won't see a horrible mock anywhere!!

## Run build
Ensure docker is installed then run:
`make`


## Usage

To add konggo via `go get`:
```
go get github.com/kevholditch/konggo
```

To add konggo via `govendor`:

```
govendor fetch github.com/kevholditch/konggo
```

Import konggo
```
import (
  konggo "github.com/kevholditch/konggo"
)
```

Konggo uses the env variable `KONG_ADMIN_ADDR` for the host address for the kong admin api.
If the env variable is not set then the address is defaulted to `http://localhost:8001`


Getting the status of the kong server:
```
kongClient := konggo.NewClient()
status, err := kongClient.Status().Get()
```
Konggo is fluent so we can combine the above two lines into one:

```
status, err := konggo.NewClient().Status().Get()
```

## APIs
Create a new API:
```
newApi := &konggo.NewApi{
	Name:                   "Example",
	Hosts:                  []string{"example.com"},
	Uris:                   []string{"/example"},
	Methods:                []string{"GET", "POST"},
	UpstreamUrl:            "http://localhost:4140/testservice",
	StripUri:               true,
	PreserveHost:           true,
	Retries:                3,
	UpstreamConnectTimeout: 1000,
	UpstreamSendTimeout:    2000,
	UpstreamReadTimeout:    3000,
	HttpsOnly:              true,
	HttpIfTerminated:       true,
}

api, err := konggo.NewClient().Apis().Create(newApi)
```

Get an API by id:
```
api, err := konggo.NewClient().Apis().GetById("ExampleApi")
```


