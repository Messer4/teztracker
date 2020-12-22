// Code generated by go-swagger; DO NOT EDIT.

package profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// CreateOrUpdateUserAddressHandlerFunc turns a function with the right signature into a create or update user address handler
type CreateOrUpdateUserAddressHandlerFunc func(CreateOrUpdateUserAddressParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateOrUpdateUserAddressHandlerFunc) Handle(params CreateOrUpdateUserAddressParams) middleware.Responder {
	return fn(params)
}

// CreateOrUpdateUserAddressHandler interface for that can handle valid create or update user address params
type CreateOrUpdateUserAddressHandler interface {
	Handle(CreateOrUpdateUserAddressParams) middleware.Responder
}

// NewCreateOrUpdateUserAddress creates a new http.Handler for the create or update user address operation
func NewCreateOrUpdateUserAddress(ctx *middleware.Context, handler CreateOrUpdateUserAddressHandler) *CreateOrUpdateUserAddress {
	return &CreateOrUpdateUserAddress{Context: ctx, Handler: handler}
}

/*CreateOrUpdateUserAddress swagger:route POST /v2/data/profile/address Profile createOrUpdateUserAddress

CreateOrUpdateUserAddress create or update user address API

*/
type CreateOrUpdateUserAddress struct {
	Context *middleware.Context
	Handler CreateOrUpdateUserAddressHandler
}

func (o *CreateOrUpdateUserAddress) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateOrUpdateUserAddressParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}