/*
Package platforms Examples

Appengine:

 import (
 	"github.com/jsonrouter/platforms/appengine"
 	ht "net/http"
 	"github.com/jsonrouter/core/openapi/v2"
 	"github.com/jsonrouter/core/validation"
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
 
 func (app *App) api_product_get(req http.Request) *http.Status {
 	
 	// Do something 
 	
 	return nil
 }

Fasthttp:

 import (
 	"github.com/jsonrouter/logging/testing"
 	"github.com/jsonrouter/platforms/fasthttp"
 	"github.com/jsonrouter/core/openapi/v2"
 	"github.com/jsonrouter/core/validation"
 
 )
 
 func TestServer() {
 	log := logs.NewClient().NewLogger()
 	spec := openapiv2.New("localhost", "TITLE")
 	
 	_, service = jsonrouter.New(log, spec)
 
 
 	api := service.Add("/api")
 
 	api.Add("/product").Param(validation.StringExact(30), "productID").GET(
 		app.api_product_get,
 	).Description(
 		"Gets the specified product",
 	).Response(
 		Product{},
 	)
 
 	panic(
 		service.Serve(PORT),
 	)
 }
 
 func (app *App) api_product_get(req http.Request) *http.Status {
 	
 	// Do something 
 	
 	return nil
 }

Standard:

 import (
 	"github.com/jsonrouter/logging/testing"
 	"github.com/jsonrouter/platforms/standard"
 	"github.com/jsonrouter/core/openapi/v2"
 	"github.com/jsonrouter/core/validation"
 	ht "net/http"
 )
 
 type Product struct {
 	name string
 }
 
 func TestServer() {
 	log := logs.NewClient().NewLogger()
 	spec := openapiv2.New("localhost", "TITLE")
 	
 	service, err = jsonrouter.New(log, spec)
 	if (err != nil){
 		// Handle error 
 	}
 
 
 	api := service.Add("/api")
 
 	api.Add("/product").Param(validation.StringExact(30), "productID").GET(
 		app.api_product_get,
 	).Description(
 		"Gets the specified product",
 	).Response(
 		Product{},
 	)
 
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
 
 func (app *App) api_product_get(req http.Request) *http.Status {
 	
 	// Do something 
 	
 	return nil
 }

*/
package platforms
