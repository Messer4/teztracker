// Code generated by go-swagger; DO NOT EDIT.

package profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetUserNotesHandlerFunc turns a function with the right signature into a get user notes handler
type GetUserNotesHandlerFunc func(GetUserNotesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetUserNotesHandlerFunc) Handle(params GetUserNotesParams) middleware.Responder {
	return fn(params)
}

// GetUserNotesHandler interface for that can handle valid get user notes params
type GetUserNotesHandler interface {
	Handle(GetUserNotesParams) middleware.Responder
}

// NewGetUserNotes creates a new http.Handler for the get user notes operation
func NewGetUserNotes(ctx *middleware.Context, handler GetUserNotesHandler) *GetUserNotes {
	return &GetUserNotes{Context: ctx, Handler: handler}
}

/*GetUserNotes swagger:route GET /v2/data/profile/notes Profile getUserNotes

GetUserNotes get user notes API

*/
type GetUserNotes struct {
	Context *middleware.Context
	Handler GetUserNotesHandler
}

func (o *GetUserNotes) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetUserNotesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}