package jsonrouter

import (
	"testing"
	//
	"github.com/valyala/fasthttp"
	//
	"github.com/jsonrouter/core/tree"
	"github.com/jsonrouter/platforms"
)

func TestFasthttpRequest(t *testing.T) {

	ctx := &fasthttp.RequestCtx{}
	req := NewRequestObject(&tree.Node{}, ctx)

	platforms.RunStandardTests(t, req)

}
