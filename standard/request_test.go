package jsonrouter

import (
	"bytes"
	"testing"
	"net/http/httptest"
	//
	"github.com/jsonrouter/core/tree"
	"github.com/jsonrouter/platforms"
)

func TestStandardRequest(t *testing.T) {

	platforms.RunStandardTests(
		t,
		NewRequestObject(
			&tree.Node{},
			httptest.NewRecorder(),
			httptest.NewRequest("POST", "https://google.com", bytes.NewBuffer(nil)),
		),
	)
}
