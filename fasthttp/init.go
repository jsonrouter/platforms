package jsonrouter

import	(
	"strconv"
	//
	"github.com/valyala/fasthttp"
	//
	"github.com/jsonrouter/logging"
	"github.com/jsonrouter/core"
	"github.com/jsonrouter/core/tree"
	"github.com/jsonrouter/platforms"
	"github.com/jsonrouter/core/metrics"
)

type FastHttpRouter func (ctx *fasthttp.RequestCtx)

// Serve is a function which calls the ListenAndServe func in the fasthttp package.
func (router FastHttpRouter) Serve(port int) error {

	return fasthttp.ListenAndServe(":"+strconv.Itoa(port), fasthttp.RequestHandler(router))
}

// Serve is a function which calls the ListenAndServeTLS func in the fasthttp package.
func (router FastHttpRouter) ServeTLS(port int, crt, key string) error {

	return fasthttp.ListenAndServeTLS(":"+strconv.Itoa(port), crt, key, fasthttp.RequestHandler(router))
}

// New creates a JSONrouter for the fasthttp platform.
func New(logger logging.Logger, spec interface{}) (*tree.Node, FastHttpRouter) {

	config := &tree.Config{
		Spec: spec,
		Log: logger,
		Metrics: metrics.NewMetrics(),
		MetResults: map[string]interface{}{},
	}
	root := tree.NewNode(config)

	platforms.AddMetricsEndpoints(root)
	platforms.AddSpecEndpoints(root)

	return root, FastHttpRouter(
		func (ctx *fasthttp.RequestCtx) {

			core.MainHandler(
				NewRequestObject(root, ctx),
				root,
				string(ctx.Path()),
			)

		},
	)
}
