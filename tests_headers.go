package platforms

import (
	"testing"
	//
	"github.com/jsonrouter/core/http"
)

func StandardTests_headers(t *testing.T, req http.Request) {

	req.SetRequestHeader("hello", "world")
	if req.GetRequestHeader("hello") != "world" {
		t.Error("REQUEST HEADER FAILED")
		t.Fail()
		return
	}

	req.SetResponseHeader("hello", "world")
	if req.GetResponseHeader("hello") != "world" {
		t.Error("RESPONSE HEADER FAILED: '"+req.GetResponseHeader("hello")+"'")
		t.Fail()
		return
	}

}
