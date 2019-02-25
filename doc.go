/*
Package jsonrouter Examples

Appengine:

Fasthttp:

 import (
 	"github.com/jsonrouter/logging/testing"
 	"github.com/jsonrouter/platforms/fasthttp"
 )

 func TestServer() {
 	log := logs.NewClient().NewLogger()
 	spec := openapiv2.New("localhost", "TITLE")
 	
 	_, service = jsonrouter.New(log, spec)

 	panic(
 		service.Serve("8080"),
 	)
 }

Standard:
*/
package jsonrouter
