package jsonrouter

import	(
		www "net/http"
		//
		"github.com/jsonrouter/core"
		"github.com/jsonrouter/core/tree"
		"github.com/jsonrouter/platforms"
		"github.com/chrysmore/metrics"
		)

// create a new router for an app
func New(spec interface{}) (*platforms.Router, error) {

	config := &tree.Config{
		Spec: spec,
		Metrics: metrics.Metrics{
			Timers: map[string]*metrics.Timer{
				"requestTime": &metrics.Timer{
					Name : "requestTime",
					BufferSize : 1000,
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
		func (res www.ResponseWriter, r *www.Request) {

			core.MainHandler(
				NewRequestObject(root, res, r),
				root,
				r.URL.Path,
			)

		},
	), nil
}
