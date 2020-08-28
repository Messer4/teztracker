// Code generated by go-swagger; DO NOT EDIT.

package assets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetAssetsListHandlerFunc turns a function with the right signature into a get assets list handler
type GetAssetsListHandlerFunc func(GetAssetsListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAssetsListHandlerFunc) Handle(params GetAssetsListParams) middleware.Responder {
	return fn(params)
}

// GetAssetsListHandler interface for that can handle valid get assets list params
type GetAssetsListHandler interface {
	Handle(GetAssetsListParams) middleware.Responder
}

// NewGetAssetsList creates a new http.Handler for the get assets list operation
func NewGetAssetsList(ctx *middleware.Context, handler GetAssetsListHandler) *GetAssetsList {
	return &GetAssetsList{Context: ctx, Handler: handler}
}

/*GetAssetsList swagger:route GET /v2/data/{network}/assets Assets getAssetsList

GetAssetsList get assets list API

*/
type GetAssetsList struct {
	Context *middleware.Context
	Handler GetAssetsListHandler
}

func (o *GetAssetsList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetAssetsListParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}