package router

import	(
	"strconv"
	//
	"github.com/valyala/fasthttp"
	//
	"github.com/jsonrouter/logging"
	"github.com/jsonrouter/core"
	"github.com/jsonrouter/core/tree"
)

type FastHttpRouter func (ctx *fasthttp.RequestCtx)

func (router FastHttpRouter) Serve(port int) error {

	return fasthttp.ListenAndServe(":"+strconv.Itoa(port), fasthttp.RequestHandler(router))
}

func (router FastHttpRouter) ServeTLS(port int, crt, key string) error {

	return fasthttp.ListenAndServeTLS(":"+strconv.Itoa(port), crt, key, fasthttp.RequestHandler(router))
}

func NewRouter(logger logging.Logger, spec interface{}) (*tree.Node, FastHttpRouter) {

	root := tree.NewNode()

	root.Config.Spec = spec
	root.Config.Log = logger

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
