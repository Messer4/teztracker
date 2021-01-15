// Code generated by go-swagger; DO NOT EDIT.

package profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// VerifyEmailTokenHandlerFunc turns a function with the right signature into a verify email token handler
type VerifyEmailTokenHandlerFunc func(VerifyEmailTokenParams) middleware.Responder

// Handle executing the request and returning a response
func (fn VerifyEmailTokenHandlerFunc) Handle(params VerifyEmailTokenParams) middleware.Responder {
	return fn(params)
}

// VerifyEmailTokenHandler interface for that can handle valid verify email token params
type VerifyEmailTokenHandler interface {
	Handle(VerifyEmailTokenParams) middleware.Responder
}

// NewVerifyEmailToken creates a new http.Handler for the verify email token operation
func NewVerifyEmailToken(ctx *middleware.Context, handler VerifyEmailTokenHandler) *VerifyEmailToken {
	return &VerifyEmailToken{Context: ctx, Handler: handler}
}

/*VerifyEmailToken swagger:route POST /v2/data/profile/verify/email/token Profile verifyEmailToken

VerifyEmailToken verify email token API

*/
type VerifyEmailToken struct {
	Context *middleware.Context
	Handler VerifyEmailTokenHandler
}

func (o *VerifyEmailToken) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewVerifyEmailTokenParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}