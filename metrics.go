package platforms

import (
	"github.com/jsonrouter/core/metrics"
)

func InitMetrics() metrics.Metrics {
	return metrics.Metrics{
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
			"requestMethods" : &metrics.MultiCounter{
				Name : "requestMethods",
				Counters : map[string]*metrics.Counter{},
			},
		},
		BenchMarks: map[string]*metrics.BenchMark{
				"requestMethodsBench": &metrics.BenchMark {
					Name: "methodsBenchMark",
					Timers: map[string]*metrics.Timer{
						"GET": &metrics.Timer{
							Name : "GET",
							BufferSize : 1000,
						},
						"POST": &metrics.Timer{
							Name : "POST",
							BufferSize : 1000,
						},
					},
				},
			},
		Results: map[string]interface{}{},
	}
}
