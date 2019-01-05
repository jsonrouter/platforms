// Package router implements the tarantula http router for net/http projects
package router

import	(
		"regexp"
		"strconv"
		"net/http"
		//
		"github.com/jsonrouter/logging"
		"github.com/jsonrouter/core"
		"github.com/jsonrouter/core/tree"
		)

type WildcardRouter struct {
	handler http.Handler
}

func (router *WildcardRouter) Handler(pattern *regexp.Regexp, handler http.Handler) {}

func (router *WildcardRouter) HandleFunc(pattern *regexp.Regexp, handler func(http.ResponseWriter, *http.Request)) {}

func (router *WildcardRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) { router.handler.ServeHTTP(w, r) }

func (router *WildcardRouter) Serve(port int) error {
	return http.ListenAndServe(":"+strconv.Itoa(port), router)
}

// Creates a new router.
func NewRouter(log logging.Logger, spec interface{}) (*tree.Node, *WildcardRouter) {

	root := tree.NewNode()

	root.Config.Spec = spec
	root.Config.Log = log

	f := func (res http.ResponseWriter, r *http.Request) {

		req := NewRequestObject(root, res, r)

		core.MainHandler(req, root, r.URL.Path)

	}

	return root, &WildcardRouter{http.HandlerFunc(f)}
}
