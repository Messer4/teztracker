// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetPublicBakersListHandlerFunc turns a function with the right signature into a get public bakers list handler
type GetPublicBakersListHandlerFunc func(GetPublicBakersListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetPublicBakersListHandlerFunc) Handle(params GetPublicBakersListParams) middleware.Responder {
	return fn(params)
}

// GetPublicBakersListHandler interface for that can handle valid get public bakers list params
type GetPublicBakersListHandler interface {
	Handle(GetPublicBakersListParams) middleware.Responder
}

// NewGetPublicBakersList creates a new http.Handler for the get public bakers list operation
func NewGetPublicBakersList(ctx *middleware.Context, handler GetPublicBakersListHandler) *GetPublicBakersList {
	return &GetPublicBakersList{Context: ctx, Handler: handler}
}

/*GetPublicBakersList swagger:route GET /v2/data/{platform}/{network}/public_bakers Accounts getPublicBakersList

GetPublicBakersList get public bakers list API

*/
type GetPublicBakersList struct {
	Context *middleware.Context
	Handler GetPublicBakersListHandler
}

func (o *GetPublicBakersList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetPublicBakersListParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}