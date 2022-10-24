// Code generated by go-swagger; DO NOT EDIT.

package http_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetPanicHandlerFunc turns a function with the right signature into a get panic handler
type GetPanicHandlerFunc func(GetPanicParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetPanicHandlerFunc) Handle(params GetPanicParams) middleware.Responder {
	return fn(params)
}

// GetPanicHandler interface for that can handle valid get panic params
type GetPanicHandler interface {
	Handle(GetPanicParams) middleware.Responder
}

// NewGetPanic creates a new http.Handler for the get panic operation
func NewGetPanic(ctx *middleware.Context, handler GetPanicHandler) *GetPanic {
	return &GetPanic{Context: ctx, Handler: handler}
}

/*
	GetPanic swagger:route GET /panic HTTP API getPanic

# Panic

crashes the process with exit code 255
*/
type GetPanic struct {
	Context *middleware.Context
	Handler GetPanicHandler
}

func (o *GetPanic) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetPanicParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
