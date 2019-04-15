package platforms

import (
	"testing"
	//
	"github.com/jsonrouter/core/http"
)

func StandardTests_parameters(t *testing.T, req http.Request) {

	req.SetParam("test", true)
	if !req.Param("test").(bool) {
		t.Fail()
		return
	}

	req.SetBodyParam("test", true)
	if !req.BodyParam("test").(bool) {
		t.Fail()
		return
	}

	req.SetParams(nil)
	if req.Params() != nil {
		t.Fail()
		return
	}

	req.SetBodyParams(nil)
	if req.BodyParams() != nil {
		t.Fail()
		return
	}

	p := map[string]interface{}{
		"test": true,
	}

	req.SetParams(p)
	if !req.Param("test").(bool) {
		t.Fail()
		return
	}

	req.SetBodyParams(p)
	if !req.BodyParam("test").(bool) {
		t.Fail()
		return
	}

}
