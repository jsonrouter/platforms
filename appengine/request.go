package jsonrouter

import 	(
	"io"
	"sync"
	www "net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/hjmodha/goDevice"
	//
	"google.golang.org/appengine"
	"github.com/jsonrouter/core/http"
	"github.com/jsonrouter/logging"
	"github.com/jsonrouter/logging/ae"
	"github.com/jsonrouter/core/tree"
	"github.com/golangdaddy/go.uuid"
	)

type Request struct {
	config *tree.Config
	path string
	Node *tree.Node
	method string
	res www.ResponseWriter
	r *www.Request
	params map[string]interface{}
	bodyParams map[string]interface{}
	Object map[string]interface{}
	Array []interface{}
	logClient logging.Logger
	sync.RWMutex
}

// NewRequestObject constructs a new Request implementation for the App Engine platform.
func NewRequestObject(node *tree.Node, res www.ResponseWriter, r *www.Request) *Request {

	return &Request{
		config:			node.Config,
		Node:			node,
		res:		  	res,
		r: 				r,
		params:			node.RequestParams,
		bodyParams:		map[string]interface{}{},
		method:			r.Method,
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

	ctx := appengine.NewContext(req.r)

	req.Lock()
	defer req.Unlock()

	if req.logClient == nil {
		req.logClient = logs.NewClient(appengine.AppID(ctx), ctx).NewLogger()
	}
	return req.logClient
}

// Res returns the 'net/http' ResponseWriter.
func (req *Request) Res() www.ResponseWriter {

	return req.res
}

// Res returns the 'net/http' Request.
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

	if len(req.path) == 0 {
		req.path = req.Node.FullPath()
	}

	return req.path
}

// Method returns the HTTP method of the request, e.g. POST, GET, PUT etc
func (req *Request) Method() string {

	return req.method
}

// Device returns the device object
func (req *Request) Device() string {
	r := req.R().(*www.Request)
	return string(goDevice.GetType(r))
}

// Writer returns the response
func (req *Request) Writer() io.Writer {
	return req.res
}

// Writer calls the write method on the 'core/http' responsewriter
func (req *Request) Write(b []byte) {

	req.res.Write(b)
}

// Writer calls the write method on the 'core/http' responsewriter after transforming the input to a byte slice.
func (req *Request) WriteString(s string) {

	req.res.Write([]byte(s))
}

// ServeFile serves the file from the path specified.
func (req *Request) ServeFile(path string) {

	www.ServeFile(req.Res(), req.R().(*www.Request), path)
}

// Body gets a field out of the map created from the body JSON.
func (req *Request) Body(k string) interface{} {

	return req.Object[k]
}


func (req *Request) Param(k string) interface{} { return req.params[k] }
func (req *Request) Params() map[string]interface{} { return req.params }
func (req *Request) SetParam(k string, v interface{}) { req.params[k] = v }
func (req *Request) SetParams(m map[string]interface{}) { req.params = m }

func (req *Request) BodyParam(k string) interface{} { return req.bodyParams[k] }
func (req *Request) BodyParams() map[string]interface{} { return req.bodyParams }
func (req *Request) SetBodyParam(k string, v interface{}) { req.bodyParams[k] = v }
func (req *Request) SetBodyParams(m map[string]interface{}) { req.bodyParams = m }

func (req *Request) GetResponseHeader(k string) string {
	header, ok := req.res.Header()[k]
	if !ok || len(header) == 0 {
		return ""
	}
	return req.res.Header()[k][0]
}

func (req *Request) SetResponseHeader(k, v string) {
	req.res.Header().Set(k, v)
}

func (req *Request) GetRequestHeader(k string) string {
	return req.r.Header.Get(k)
}

func (req *Request) SetRequestHeader(k, v string) {
	req.r.Header.Set(k, v)
}

func (req *Request) RawBody() (*http.Status, []byte) {

	body := req.r.Body

	b, err := ioutil.ReadAll(body)
	if body != nil { body.Close() }
	if err != nil { return http.Respond(400, err.Error()), nil }

	return nil, b
}

func (req *Request) ReadBodyObject() *http.Status {

	body := req.r.Body

	b, err := ioutil.ReadAll(body)
	if body != nil { body.Close() }
	if err != nil { return http.Respond(400, err.Error()) }

	req.Object = map[string]interface{}{}
	err = json.Unmarshal(b, &req.Object); if err != nil { return http.Respond(400, err.Error()) }

	return nil
}

func (req *Request) ReadBodyArray() *http.Status {

	body := req.r.Body

	b, err := ioutil.ReadAll(body)
	if body != nil { body.Close() }
	if err != nil { return http.Respond(400, err.Error()) }

	req.Array = []interface{}{}
	err = json.Unmarshal(b, &req.Array); if err != nil { return http.Respond(400, err.Error()) }

	return nil
}

func (req *Request) Fail() *http.Status {

	return http.Fail()
}

func (req *Request) Respond(args ...interface{}) *http.Status {

	return http.Respond(args...)
}

func (req *Request) Redirect(path string, code int) *http.Status {

	www.Redirect(req.res, req.r, path, code)

	return nil
}

func (req *Request) HttpError(msg string, code int) {

	www.Error(req.res, msg, code)
	req.Log().NewError(msg)
}
