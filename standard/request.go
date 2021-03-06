package jsonrouter

import 	(
	"io"
	www "net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/hjmodha/goDevice"
	//
	"github.com/jsonrouter/core/http"
	"github.com/jsonrouter/core/tree"
	"github.com/jsonrouter/logging"
	"github.com/jsonrouter/platforms/parameters"
)

type Request struct {
	*parameters.Parameters
	node *tree.Node
	res www.ResponseWriter
	r *www.Request
	path string
	method string
	Object map[string]interface{}
	Array []interface{}
}

// NewRequestObject constructs a new Request implementation for the vanilla platform.
func NewRequestObject(node *tree.Node, res www.ResponseWriter, r *www.Request) *Request {

	return &Request{
		Parameters: parameters.New(),
		node: node,
		res: res,
		r: r,
		method: r.Method,
	}
}

// Testing returns whether or not this is a test implementation.
func (req *Request) Testing() bool {
	return false
}

// Log returns the logging client.
func (req *Request) Log() logging.Logger {
	return req.node.Config.Log
}

// Res returns the 'net/http' ResponseWriter.
func (req *Request) Res() www.ResponseWriter {
	return req.res
}

// R returns the 'net/http' Request.
func (req *Request) R() interface{} {
	return req.r
}

// IsTLS returns the state of whether this request is a secure request.
func (req *Request) IsTLS() bool {

	if req.r.TLS == nil { return false }

	return req.r.TLS.HandshakeComplete
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

	return req.node.FullPath()
}

// Method returns the HTTP method of the request, e.g. POST, GET, PUT etc.
func (req *Request) Method() string {

	return req.method
}

// Device returns the device object.
func (req *Request) Device() string {
	return string(
		goDevice.GetType(
			req.R().(*www.Request),
		),
	)
}

// Writer returns the responseWriter as an io.Writer.
func (req *Request) Writer() io.Writer {

	return req.res
}

// Write calls the write method on the 'core/http' responsewriter
func (req *Request) Write(b []byte) (int, error) {
	return req.res.Write(b)
}

// WriteString calls the write method on the 'core/http' responsewriter after transforming the input to a byte slice.
func (req *Request) WriteString(s string) (int, error) {
	return req.res.Write([]byte(s))
}

// ServeFile serves the file from the path specified.
func (req *Request) ServeFile(path string) {
	www.ServeFile(req.Res(), req.R().(*www.Request), path)
}

// Body gets a field out of the map created from the body JSON.
func (req *Request) Body(k string) interface{} {
	return req.Object[k]
}

// GetRequestHeader gets a request header value.
func (req *Request) GetRequestHeader(k string) string {
	return req.r.Header.Get(k)
}
// SetRequestHeader sets a request header value.
func (req *Request) SetRequestHeader(k, v string) {
	req.r.Header.Set(k, v)
}

// GetResponseHeader gets a header value from the response.
func (req *Request) GetResponseHeader(k string) string {
	return req.res.Header().Get(k)
}
// SetResponseHeader sets a response header value.
func (req *Request) SetResponseHeader(k, v string) {
	req.res.Header().Set(k, v)
}

// RawBody returns the HTTP request body.
func (req *Request) RawBody() (*http.Status, []byte) {

	body := req.r.Body

	b, err := ioutil.ReadAll(body)
	if body != nil {
		body.Close()
	}
	if err != nil {
		status, _ := http.Respond(400, err.Error())
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

	err := json.Unmarshal(b, &req.Object)
	if err != nil {
		status, _ := http.Respond(400, err.Error())
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
		status, _ := http.Respond(400, err.Error())
		return status
	}

	return nil
}

// Fail sends HTTP 500 status error.
func (req *Request) Fail() *http.Status {
	return http.Fail()
}

// Respond calls the respond method which creates the response payload.
func (req *Request) Respond(args ...interface{}) *http.Status {
	status, contentType := http.Respond(args...)
	req.SetResponseHeader("Content-Type", contentType)
	return status
}

// Redirect redirects the http to the destination URL.
func (req *Request) Redirect(path string, code int) *http.Status {
	www.Redirect(req.res, req.r, path, code)
	return nil
}

// HttpError responds to the HTTP request with an error status.
func (req *Request) HttpError(msg string, code int) {
	www.Error(req.res, msg, code)
	req.Log().NewError(msg)
}
