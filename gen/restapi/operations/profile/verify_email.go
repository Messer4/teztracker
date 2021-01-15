// Code generated by go-swagger; DO NOT EDIT.

package profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// VerifyEmailHandlerFunc turns a function with the right signature into a verify email handler
type VerifyEmailHandlerFunc func(VerifyEmailParams) middleware.Responder

// Handle executing the request and returning a response
func (fn VerifyEmailHandlerFunc) Handle(params VerifyEmailParams) middleware.Responder {
	return fn(params)
}

// VerifyEmailHandler interface for that can handle valid verify email params
type VerifyEmailHandler interface {
	Handle(VerifyEmailParams) middleware.Responder
}

// NewVerifyEmail creates a new http.Handler for the verify email operation
func NewVerifyEmail(ctx *middleware.Context, handler VerifyEmailHandler) *VerifyEmail {
	return &VerifyEmail{Context: ctx, Handler: handler}
}

/*VerifyEmail swagger:route POST /v2/data/profile/verify/email Profile verifyEmail

VerifyEmail verify email API

*/
type VerifyEmail struct {
	Context *middleware.Context
	Handler VerifyEmailHandler
}

func (o *VerifyEmail) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewVerifyEmailParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}