package platforms

import (
	"regexp"
	"strconv"
	"net/http"
)

type WildcardRouter struct {
	handler http.Handler
}

func (router *WildcardRouter) Handler(pattern *regexp.Regexp, handler http.Handler) {}

func (router *WildcardRouter) HandleFunc(pattern *regexp.Regexp, handler func(http.ResponseWriter, *http.Request)) {}

func (router *WildcardRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) { router.handler.ServeHTTP(w, r) }

func (router *WildcardRouter) Serve(port int) error {
	return http.ListenAndServe(":"+strconv.Itoa(port), router)
}
