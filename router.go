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

func NewRouter(node *tree.Node, handlerFunc func (res http.ResponseWriter, r *http.Request)) *Router {

	return &Router{
		&WildcardRouter{
			http.HandlerFunc(handlerFunc),
		},
		node,
	}
}

type Router struct {
	http.Handler
	*tree.Node
}

func AddSpecEndpoints(root *tree.Node) {
	root.Add("/openapi.spec.json").GET(
		root.Config.SpecHandler,
	)
}
