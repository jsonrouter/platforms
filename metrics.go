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
		Results: map[string]interface{}{},
	}
}
