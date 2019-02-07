package jsonrouter

import 	(
	"io"
	www "net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/hjmodha/goDevice"
	"github.com/golangdaddy/go.uuid"
	//
	"github.com/jsonrouter/validation"
	"github.com/jsonrouter/core/http"
	"github.com/jsonrouter/core/tree"
	"github.com/jsonrouter/logging"
)

type Request struct {
	node *tree.Node
	res www.ResponseWriter
	r *www.Request
	path string
	method string
	params map[string]interface{}
	bodyParams map[string]interface{}
	Object map[string]interface{}
	Array []interface{}
}

func NewRequestObject(node *tree.Node, res www.ResponseWriter, r *www.Request) *Request {

	return &Request{
		node: node,
		res: res,
		r: r,
		method: r.Method,
		params: node.RequestParams,
		Object: validation.Object{},
		Array: validation.Array{},
	}
}

func (req *Request) Testing() bool {
	return false
}

func (req *Request) UID() (string, error) {

	uid, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	return uid.String(), nil
}

func (req *Request) Log() logging.Logger {
	return req.node.Config.Log
}

func (req *Request) Res() www.ResponseWriter {
	return req.res
}

func (req *Request) R() interface{} {
	return req.r
}

func (req *Request) IsTLS() bool {

	if req.r.TLS == nil { return false }

	return req.r.TLS.HandshakeComplete
}

func (req *Request) BodyArray() []interface{} {

	return req.Array
}

func (req *Request) BodyObject() map[string]interface{} {

	return req.Object
}

func (req *Request) FullPath() string {

	return req.node.FullPath()
}

func (req *Request) Method() string {

	return req.method
}

func (req *Request) Node() *tree.Node {

	return req.node
}

func (req *Request) Device() string {
	return string(
		goDevice.GetType(
			req.R().(*www.Request),
		),
	)
}

func (req *Request) Writer() io.Writer {

	return req.res
}

func (req *Request) Write(b []byte) {

	req.res.Write(b)
}

func (req *Request) WriteString(s string) {

	req.res.Write([]byte(s))
}

func (req *Request) ServeFile(path string) {

	www.ServeFile(req.Res(), req.R().(*www.Request), path)
}

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

	err = json.Unmarshal(b, &req.Object); if err != nil { return http.Respond(400, err.Error()) }

	return nil
}

func (req *Request) ReadBodyArray() *http.Status {

	body := req.r.Body

	b, err := ioutil.ReadAll(body)

	if body != nil { body.Close() }

	if err != nil { return http.Respond(400, err.Error()) }

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
