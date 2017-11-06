package go_sync

import (
	"os"
	"testing"
)

func testPreCheck(t *testing.T) {
	skip := len(os.Getenv("FORM3_ACC")) == 0
	if skip {
		t.Log("form3 tests require setting FORM3_ACC=1 environment variable")
		t.Skip()
	}
}
