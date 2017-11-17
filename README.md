[![Build Status](https://travis-ci.org/kevholditch/gokong.svg?branch=master)](https://travis-ci.org/kevholditch/gokong)

GoKong
======
A kong go client fully tested with no mocks!!

## GoKong
GoKong is a easy to use api client for [kong](https://getkong.org/).  The difference with the gokong library is all of its tests are written against a real running kong running inside a docker container, yep that's right you won't see a horrible mock anywhere!!

## Run build
Ensure docker is installed then run:
`make`


## Usage

To add gokong via `go get`:
```
go get github.com/kevholditch/gokong
```

To add gokong via `govendor`:

```
govendor fetch github.com/kevholditch/gokong
```

Import gokong
```
import (
  gokong "github.com/kevholditch/gokong"
)
```

Gokong uses the env variable `KONG_ADMIN_ADDR` for the host address for the kong admin api.
If the env variable is not set then the address is defaulted to `http://localhost:8001`


Getting the status of the kong server:
```
kongClient := gokong.NewClient()
status, err := kongClient.Status().Get()
```
Gokong is fluent so we can combine the above two lines into one:

```
status, err := gokong.NewClient().Status().Get()
```

## APIs
Create a new API:
```
newApi := &gokong.NewApi{
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

api, err := gokong.NewClient().Apis().Create(newApi)
```

Get an API by id:
```
api, err := gokong.NewClient().Apis().GetById("ExampleApi")
```


