// Code generated by go-swagger; DO NOT EDIT.

package kubernetes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetHealthzOKCode is the HTTP code returned for type GetHealthzOK
const GetHealthzOKCode int = 200

/*
GetHealthzOK OK

swagger:response getHealthzOK
*/
type GetHealthzOK struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewGetHealthzOK creates GetHealthzOK with default headers values
func NewGetHealthzOK() *GetHealthzOK {

	return &GetHealthzOK{}
}

// WithPayload adds the payload to the get healthz o k response
func (o *GetHealthzOK) WithPayload(payload interface{}) *GetHealthzOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get healthz o k response
func (o *GetHealthzOK) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetHealthzOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}