// Code generated by go-swagger; DO NOT EDIT.

package kubernetes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetReadyzHandlerFunc turns a function with the right signature into a get readyz handler
type GetReadyzHandlerFunc func(GetReadyzParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetReadyzHandlerFunc) Handle(params GetReadyzParams) middleware.Responder {
	return fn(params)
}

// GetReadyzHandler interface for that can handle valid get readyz params
type GetReadyzHandler interface {
	Handle(GetReadyzParams) middleware.Responder
}

// NewGetReadyz creates a new http.Handler for the get readyz operation
func NewGetReadyz(ctx *middleware.Context, handler GetReadyzHandler) *GetReadyz {
	return &GetReadyz{Context: ctx, Handler: handler}
}

/*
	GetReadyz swagger:route GET /readyz Kubernetes getReadyz

# Readiness check

used by Kubernetes readiness probe
*/
type GetReadyz struct {
	Context *middleware.Context
	Handler GetReadyzHandler
}

func (o *GetReadyz) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetReadyzParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
