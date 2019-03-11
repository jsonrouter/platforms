// Package jsonrouter implements the http-router for net/http projects
package jsonrouter

import	(
	"net/http"
	//"fmt"
	"github.com/jsonrouter/platforms"
	"github.com/jsonrouter/logging"
	"github.com/jsonrouter/core"
	"github.com/jsonrouter/core/metrics"
	"github.com/jsonrouter/core/tree"
	"github.com/jsonrouter/core/openapi"
)

// New creates a JSONrouter for the vanilla platform.
func New(log logging.Logger, spec interface{}) (*platforms.Router, error) {

	if err := openapi.ValidSpec(spec); err != nil {
		return nil, err
	}

	config := &tree.Config{
		Spec: spec,
		Log: log,
		Metrics: metrics.NewMetrics(),
		MetResults: map[string]interface{}{},
	}

	root := tree.NewNode(config)

	platforms.AddMetricsEndpoints(root)
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
