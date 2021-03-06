// Code generated by go-swagger; DO NOT EDIT.

package blocks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetBakingRightsHandlerFunc turns a function with the right signature into a get baking rights handler
type GetBakingRightsHandlerFunc func(GetBakingRightsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetBakingRightsHandlerFunc) Handle(params GetBakingRightsParams) middleware.Responder {
	return fn(params)
}

// GetBakingRightsHandler interface for that can handle valid get baking rights params
type GetBakingRightsHandler interface {
	Handle(GetBakingRightsParams) middleware.Responder
}

// NewGetBakingRights creates a new http.Handler for the get baking rights operation
func NewGetBakingRights(ctx *middleware.Context, handler GetBakingRightsHandler) *GetBakingRights {
	return &GetBakingRights{Context: ctx, Handler: handler}
}

/*GetBakingRights swagger:route GET /v2/data/{platform}/{network}/baking_rights Blocks getBakingRights

GetBakingRights get baking rights API

*/
type GetBakingRights struct {
	Context *middleware.Context
	Handler GetBakingRightsHandler
}

func (o *GetBakingRights) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetBakingRightsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
