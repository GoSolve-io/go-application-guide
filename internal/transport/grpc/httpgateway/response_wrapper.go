package httpgateway

import (
	"net/http"
)

// ResponseWrapper wraps the response writer and allows the middleware to retrieve the return code
type ResponseWrapper struct {
	StatusCode int
	response   http.ResponseWriter
}

// NewResponseWrapper returns a new wrapper with the response
func NewResponseWrapper(response http.ResponseWriter) *ResponseWrapper {
	return &ResponseWrapper{
		response: response,
	}
}

// Header implements the ResponseWriter interface
func (r ResponseWrapper) Header() http.Header {
	return r.response.Header()
}

// Write implements the ResponseWriter interface
func (r ResponseWrapper) Write(bytes []byte) (int, error) {
	return r.response.Write(bytes)
}

// WriteHeader implements the ResponseWriter interface
func (r ResponseWrapper) WriteHeader(statusCode int) {
	r.StatusCode = statusCode
	r.response.WriteHeader(statusCode)
}
