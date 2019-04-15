package platforms

import (
	"testing"
	//
	"github.com/jsonrouter/core/http"
)

func RunStandardTests(t *testing.T, req http.Request) {

	StandardTests_headers(t, req)
	StandardTests_parameters(t, req)

}
