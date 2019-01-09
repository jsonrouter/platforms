package router

import	(
		www "net/http"
		//
		"github.com/jsonrouter/core"
		"github.com/jsonrouter/core/tree"
		"github.com/jsonrouter/platforms"
		)

// create a new router for an app
func NewRouter(spec interface{}) (*platforms.Router, error) {

	config := &tree.Config{
		Spec: spec,
	}
	root := tree.NewNode(config)

	platforms.AddSpecEndpoints(root)

	return platforms.NewRouter(
		root,
		func (res www.ResponseWriter, r *www.Request) {

			core.MainHandler(
				NewRequestObject(root, res, r),
				root,
				r.URL.Path,
			)

		},
	), nil
}
