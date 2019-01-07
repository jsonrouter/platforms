package router

import	(
		"regexp"
		www "net/http"
		//
		"github.com/jsonrouter/core"
		"github.com/jsonrouter/core/tree"
		)

type WildcardRouter struct {
	handler www.Handler
}

func (router *WildcardRouter) Handler(pattern *regexp.Regexp, handler www.Handler) {}

func (router *WildcardRouter) HandleFunc(pattern *regexp.Regexp, handler func(www.ResponseWriter, *www.Request)) {}

func (router *WildcardRouter) ServeHTTP(w www.ResponseWriter, r *www.Request) { router.handler.ServeHTTP(w, r) }

// create a new router for an app
func NewRouter(spec interface{}) (*tree.Node, *WildcardRouter) {

	config := &tree.Config{
		Spec: spec,
	}
	root := tree.NewNode(config)

	f := func (res www.ResponseWriter, r *www.Request) {

		core.MainHandler(
			NewRequestObject(root, res, r),
			root,
			r.URL.Path,
		)

	}

	return root, &WildcardRouter{www.HandlerFunc(f)}
}
