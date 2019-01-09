// Package jsonrouter implements the http-router for net/http projects
package jsonrouter

import	(
	"net/http"
	//
	"github.com/jsonrouter/platforms"
	"github.com/jsonrouter/logging"
	"github.com/jsonrouter/core"
	"github.com/jsonrouter/core/tree"
	"github.com/jsonrouter/core/openapi"
)

// Creates a new Router object that is ready to serve.
func New(log logging.Logger, spec interface{}) (*platforms.Router, error) {

	if err := openapi.ValidSpec(spec); err != nil {
		return nil, err
	}

	config := &tree.Config{
		Spec: spec,
		Log: log,
	}
	root := tree.NewNode(config)

	platforms.AddSpecEndpoints(root)

	return platforms.NewRouter(
		root,
		func (res http.ResponseWriter, r *http.Request) {

					core.MainHandler(
						NewRequestObject(root, res, r),
						root,
						r.URL.Path,
					)

		},
	), nil
}
