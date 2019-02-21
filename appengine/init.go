package jsonrouter

import	(
		www "net/http"
		//
		"github.com/jsonrouter/core"
		"github.com/jsonrouter/core/tree"
		"github.com/jsonrouter/platforms"
		)

// Creates a JSONrouter for App Engine platforms
func New(spec interface{}) (*platforms.Router, error) {

	config := &tree.Config{
		Spec: spec,
		Metrics: platforms.InitMetrics(),
		MetResults: map[string]interface{}{},
	}
	root := tree.NewNode(config)


	platforms.AddMetricsEndpoints(root)
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
