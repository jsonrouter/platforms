/*
Package platforms Examples

Appengine:

 import (
 	"github.com/jsonrouter/platforms/appengine"
 	ht "net/http"
 )

 func TestServer() {
 	spec := openapiv2.New("localhost", "TITLE")
 	
 	service, err = jsonrouter.New(spec)
 	if (err != nil){
		// Handle Error //
 	}

 	panic(
 		ht.ListenAndServe(
 			fmt.Sprintf(
 				":%d",
 				PORT,
 			),
 			service,
 		),
 	)
 }

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
 		service.Serve(PORT),
 	)
 }

Standard:

 import (
 	"github.com/jsonrouter/logging/testing"
 	"github.com/jsonrouter/platforms/standard"
 	ht "net/http"
 )

 func TestServer() {
 	log := logs.NewClient().NewLogger()
 	spec := openapiv2.New("localhost", "TITLE")
 	
 	service, err = jsonrouter.New(log, spec)
 	if (err != nil){
		// Handle Error //
 	}

 	panic(
 		ht.ListenAndServe(
 			fmt.Sprintf(
 				":%d",
 				PORT,
 			),
 			service,
 		),
 	)
 }

*/
package platforms
