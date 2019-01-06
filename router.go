package platforms

import (
	"net/http"
	//
	"github.com/jsonrouter/core/tree"
)

const (
	CONST_SPEC_PATH_V2 = "/openapi.spec.v2.json"
	CONST_SPEC_PATH_V3 = "/openapi.spec.v3.json"
)

func NewRouter(handler http.Handler, node *tree.Node) *Router {

	node.Add(
		CONST_SPEC_PATH_V2,
	).GET(
		node.Config.ServeSpec,
	)

	node.Add(
		CONST_SPEC_PATH_V3,
	).GET(
		node.Config.ServeSpec,
	)

	return &Router{
		handler,
		node,
	}
}

type Router struct {
	http.Handler
	*tree.Node
}
