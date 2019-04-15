package parameters

import (
	"sync"
)

func New() *Parameters {
	return &Parameters{
		params: map[string]interface{}{},
		bodyParams: map[string]interface{}{},
	}
}

type Parameters struct {
	params map[string]interface{}
	bodyParams map[string]interface{}
	sync.RWMutex
}

// Param gets a variable that has been stored in the params object.
// This could be an arguement from the request path, or have other vars stored there for random access.
func (self *Parameters) Param(k string) interface{} {
	self.RLock()
	defer self.RUnlock()
	return self.params[k]
}
// Params returns the params object.
// This object is intended to be used for storing path parameters.
func (self *Parameters) Params() map[string]interface{} {
	self.RLock()
	defer self.RUnlock()
	return self.params
}
// SetParam sets a value from the params object.
func (self *Parameters) SetParam(k string, v interface{}) {
	self.Lock()
	defer self.Unlock()
	self.params[k] = v
}
// SetParam replaces the params object with the supplied map.
func (self *Parameters) SetParams(m map[string]interface{}) {
	self.Lock()
	defer self.Unlock()
	self.params = m
}

// BodyParam gets a variable that has been stored in the bodyparams object.
func (self *Parameters) BodyParam(k string) interface{} {
	self.RLock()
	defer self.RUnlock()
	return self.bodyParams[k]
}
// BodyParam returns the bodyparams object.
func (self *Parameters) BodyParams() map[string]interface{} {
	self.RLock()
	defer self.RUnlock()
	return self.bodyParams
}
// SetBodyParam sets a value from the params object.
func (self *Parameters) SetBodyParam(k string, v interface{}) {
	self.Lock()
	defer self.Unlock()
	self.bodyParams[k] = v
}
// SetBodyParams sets a value from the bodyparams object.
func (self *Parameters) SetBodyParams(m map[string]interface{}) {
	self.Lock()
	defer self.Unlock()
	self.bodyParams = m
}
