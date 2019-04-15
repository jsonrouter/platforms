package jsonrouter

import 	(
		"io"
		"sync"
		www "net/http"
		"encoding/json"
		//
		"github.com/valyala/fasthttp"
		"github.com/golangdaddy/go.uuid"
		//"github.com/hjmodha/goDevice"
		//
		"github.com/jsonrouter/logging"
		"github.com/jsonrouter/core/http"
		"github.com/jsonrouter/core/tree"
		"github.com/jsonrouter/platforms/parameters"
		)

type Request struct {
	*parameters.Parameters
	ctx *fasthttp.RequestCtx
	config *tree.Config
	path string
	Node *tree.Node
	method string
	Object map[string]interface{}
	Array []interface{}
	sync.RWMutex
}

// NewRequestObject constructs a new Request implementation for the fasthttp latform.
func NewRequestObject(node *tree.Node, ctx *fasthttp.RequestCtx) *Request {

	return &Request{
		Parameters: 	parameters.New(),
		ctx:			ctx,
		config:			node.Config,
		Node:			node,
		method:			string(ctx.Method()),
	}
}

// Testing returns whether or not this is a test implementation.
func (req *Request) Testing() bool {
	return false
}

// UID returns the UUIDv4 which was randomply generated for this request.
func (req *Request) UID() (string, error) {

	uid, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	return uid.String(), nil
}

// Log returns the logging client.
func (req *Request) Log() logging.Logger {

	return req.config.Log
}

// Does nothing useful in httprouter
func (req *Request) Res() www.ResponseWriter {

	x := new(www.ResponseWriter)

	return *x
}

// Does nothing useful in httprouter
func (req *Request) R() interface{} {

	return req.ctx
}

// IsTLS returns the state of whether this request is a secure request.
func (req *Request) IsTLS() bool {

	return req.ctx.IsTLS()
}

// BodyArray returns the HTTP body which was previously unmarshaled into a slice.
func (req *Request) BodyArray() []interface{} {

	return req.Array
}

// BodyObject returns the HTTP body which was previously unmarshaled into a map.
func (req *Request) BodyObject() map[string]interface{} {

	return req.Object
}

// FullPath returns the path for the http request.
func (req *Request) FullPath() string {

	if len(req.path) == 0 {

		req.path = req.Node.FullPath()

	}

	return req.path
}

// Method returns the HTTP method of the request, e.g. POST, GET, PUT etc
func (req *Request) Method() string {

	return req.method
}

// Device returns the device object.
func (req *Request) Device() string {

	//r := req.R().(*fasthttp.RequestCtx)

	return "?"
}

// Writer returns the responseWriter as an io.Writer.
func (req *Request) Writer() io.Writer {
	return req.ctx.Response.BodyWriter()
}

// Write calls the write method on the 'core/http' responsewriter
func (req *Request) Write(b []byte) (int, error) {
	return req.ctx.Write(b)
}

// WriteString calls the write method on the 'core/http' responsewriter after transforming the input to a byte slice.
func (req *Request) WriteString(s string) (int, error) {
	return req.ctx.WriteString(s)
}

// ServeFile serves the file from the path specified.
func (req *Request) ServeFile(path string) {
	fasthttp.ServeFile(req.ctx, path)
}

// Body gets a field out of the map created from the body JSON.
func (req *Request) Body(k string) interface{} {
	return req.Object[k]
}

// GetRequestHeader sets a request header value.
func (req *Request) GetRequestHeader(k string) string {
	return string(req.ctx.Request.Header.Peek(k))
}
// SetRequestHeader sets a request header value.
func (req *Request) SetRequestHeader(k, v string) {
	req.ctx.Request.Header.Set(k, v)
}

// GetResponseHeader gets a header value from the response.
func (req *Request) GetResponseHeader(k string) string {
	return string(req.ctx.Response.Header.Peek(k))
}
// SetResponseHeader sets a response header value.
func (req *Request) SetResponseHeader(k, v string) {
	req.ctx.Response.Header.Set(k, v)
}

// RawBody returns the HTTP request body.
func (req *Request) RawBody() (*http.Status, []byte) {

	b := req.ctx.PostBody()

	if b == nil {
		status, _ := http.Respond(400, "BODY IS NIL")
		return status, nil
	}

	return nil, b
}

// ReadBodyObject unmarshals the body into a map of interface{}.
func (req *Request) ReadBodyObject() *http.Status {

	status, b := req.RawBody()
	if status != nil {
		return status
	}

	if err := json.Unmarshal(b, &req.Object); err != nil {
		status, _  := http.Respond(400, err.Error())
		return status
	}

	return nil
}

// ReadBodyArray unmarshals the body into a slice of interface{}.
func (req *Request) ReadBodyArray() *http.Status {

	status, b := req.RawBody()
	if status != nil {
		return status
	}

	err := json.Unmarshal(b, &req.Array)
	if err != nil {
		status, _  := http.Respond(400, err.Error())
		return status
	}

	return nil
}

// Fail sends HTTP 500 status error.
func (req *Request) Fail() *http.Status {

	return http.Fail()
}

// Redirect redirects the http to the destination URL.
func (req *Request) Respond(args ...interface{}) *http.Status {
	status, contentType := http.Respond(args...)
	req.SetResponseHeader("Content-Type", contentType)
	return status
}

// Redirect redirects the http to the destination URL.
func (req *Request) Redirect(path string, code int) *http.Status {

	req.ctx.Redirect(path, code)

	return nil
}

// Respond calls the respond method which creates the response payload.
func (req *Request) HttpError(msg string, code int) {

	req.ctx.Error(msg, code)
	req.NewError(msg)
}
