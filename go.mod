module github.comcast.com/dh-api-gw/gokong

go 1.12

require (
	github.com/google/go-querystring v1.0.0
	github.com/kevholditch/gokong v6.0.0+incompatible
	github.com/lib/pq v1.2.0
	github.com/ory/dockertest v3.3.5+incompatible
	github.com/parnurzeal/gorequest v0.2.16
	github.com/phayes/freeport v0.0.0-20180830031419-95f893ade6f2
	github.com/satori/go.uuid v1.2.0
	github.com/stretchr/testify v1.4.0
)

replace github.com/kevholditch/gokong => ./
