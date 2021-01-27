// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/everstake/teztracker/gen/models"
)

// GetBakersVotingOKCode is the HTTP code returned for type GetBakersVotingOK
const GetBakersVotingOKCode int = 200

/*GetBakersVotingOK Get bakers voting data

swagger:response getBakersVotingOK
*/
type GetBakersVotingOK struct {

	/*
	  In: Body
	*/
	Payload *models.BakersVoting `json:"body,omitempty"`
}

// NewGetBakersVotingOK creates GetBakersVotingOK with default headers values
func NewGetBakersVotingOK() *GetBakersVotingOK {

	return &GetBakersVotingOK{}
}

// WithPayload adds the payload to the get bakers voting o k response
func (o *GetBakersVotingOK) WithPayload(payload *models.BakersVoting) *GetBakersVotingOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get bakers voting o k response
func (o *GetBakersVotingOK) SetPayload(payload *models.BakersVoting) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBakersVotingOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetBakersVotingBadRequestCode is the HTTP code returned for type GetBakersVotingBadRequest
const GetBakersVotingBadRequestCode int = 400

/*GetBakersVotingBadRequest Bad request

swagger:response getBakersVotingBadRequest
*/
type GetBakersVotingBadRequest struct {
}

// NewGetBakersVotingBadRequest creates GetBakersVotingBadRequest with default headers values
func NewGetBakersVotingBadRequest() *GetBakersVotingBadRequest {

	return &GetBakersVotingBadRequest{}
}

// WriteResponse to the client
func (o *GetBakersVotingBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// GetBakersVotingNotFoundCode is the HTTP code returned for type GetBakersVotingNotFound
const GetBakersVotingNotFoundCode int = 404

/*GetBakersVotingNotFound Not Found

swagger:response getBakersVotingNotFound
*/
type GetBakersVotingNotFound struct {
}

// NewGetBakersVotingNotFound creates GetBakersVotingNotFound with default headers values
func NewGetBakersVotingNotFound() *GetBakersVotingNotFound {

	return &GetBakersVotingNotFound{}
}

// WriteResponse to the client
func (o *GetBakersVotingNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// GetBakersVotingInternalServerErrorCode is the HTTP code returned for type GetBakersVotingInternalServerError
const GetBakersVotingInternalServerErrorCode int = 500

/*GetBakersVotingInternalServerError Internal error

swagger:response getBakersVotingInternalServerError
*/
type GetBakersVotingInternalServerError struct {
}

// NewGetBakersVotingInternalServerError creates GetBakersVotingInternalServerError with default headers values
func NewGetBakersVotingInternalServerError() *GetBakersVotingInternalServerError {

	return &GetBakersVotingInternalServerError{}
}

// WriteResponse to the client
func (o *GetBakersVotingInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}