package konggo

import (
	"github.com/zimmski/tavor/test/assert"
	"testing"
)

func Test_BuildUrl(t *testing.T) {
	assert.Equal(t, "http://localhost:8001", NewUrlBuilder("http://localhost:8001").Build())
}

func Test_BuildUrlRemoveTrailingSlash(t *testing.T) {
	assert.Equal(t, "http://localhost:8001", NewUrlBuilder("http://localhost:8001/").Build())
}

func Test_BuildUrlToStatus(t *testing.T) {
	assert.Equal(t, "http://localhost:8001/status", NewUrlBuilder("http://localhost:8001").Status().Build())
}

func Test_BuildUrlToApis(t *testing.T) {
	assert.Equal(t, "http://localhost:8001/apis/", NewUrlBuilder("http://localhost:8001").Apis().Build())
}
