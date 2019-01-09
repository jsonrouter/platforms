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
	}
	root := tree.NewNode(config)

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
