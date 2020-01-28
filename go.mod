module github.comcast.com/dh-api-gw/gokong

go 1.12

require (
	github.com/containerd/continuity v0.0.0-20200107194136-26c1120b8d41 // indirect
	github.com/google/go-querystring v1.0.0
	github.com/kevholditch/gokong v6.0.0+incompatible
	github.com/lib/pq v1.2.0
	github.com/ory/dockertest v3.3.5+incompatible
	github.com/parnurzeal/gorequest v0.2.16
	github.com/phayes/freeport v0.0.0-20180830031419-95f893ade6f2
	github.com/pkg/errors v0.9.1 // indirect
	github.com/satori/go.uuid v1.2.0
	github.com/stretchr/testify v1.4.0
	golang.org/x/net v0.0.0-20200114155413-6afb5195e5aa // indirect
	golang.org/x/sys v0.0.0-20200124204421-9fbb57f87de9 // indirect
)

replace github.com/kevholditch/gokong => github.comcast.com/dh-api-gw/gokong v0.0.0-20200128215130-cea50b2e8093
