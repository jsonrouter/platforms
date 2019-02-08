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
	"github.com/chrysmore/metrics"
)

// Creates a new Router object that is ready to serve.
func New(log logging.Logger, spec interface{}) (*platforms.Router, error) {

	if err := openapi.ValidSpec(spec); err != nil {
		return nil, err
	}

	config := &tree.Config{
		Spec: spec,
		Log: log,
		Metrics: metrics.Metrics{
			Timers: map[string]*metrics.Timer{
				"requestTime": &metrics.Timer{
					Name : "requestTime",
					BufferSize: 1000,
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
