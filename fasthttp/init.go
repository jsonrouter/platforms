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
	"github.com/chrysmore/metrics"
)

type FastHttpRouter func (ctx *fasthttp.RequestCtx)

func (router FastHttpRouter) Serve(port int) error {

	return fasthttp.ListenAndServe(":"+strconv.Itoa(port), fasthttp.RequestHandler(router))
}

func (router FastHttpRouter) ServeTLS(port int, crt, key string) error {

	return fasthttp.ListenAndServeTLS(":"+strconv.Itoa(port), crt, key, fasthttp.RequestHandler(router))
}

func New(logger logging.Logger, spec interface{}) (*tree.Node, FastHttpRouter) {

	config := &tree.Config{
		Spec: spec,
		Log: logger,
		Metrics: metrics.Metrics{
			Timers: map[string]*metrics.Timer{
				"requestTime": &metrics.Timer{
					Name : "requestTime",
				},
			},
			Counters: map[string]*metrics.Counter{
				"requestCount" : &metrics.Counter{
					Name : "requestCount",
				},
			},
			MultiCounters: map[string]*metrics.MultiCounter{
				"responseCodes" : &metrics.MultiCounter{
					Name : "responseCodes",
					Counters : map[string]*metrics.Counter{},
				},
			},
			Results: map[string]interface{}{},
		},
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
