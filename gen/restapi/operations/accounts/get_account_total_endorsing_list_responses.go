// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/everstake/teztracker/gen/models"
)

// GetAccountTotalEndorsingListOKCode is the HTTP code returned for type GetAccountTotalEndorsingListOK
const GetAccountTotalEndorsingListOKCode int = 200

/*GetAccountTotalEndorsingListOK Query compatibility endpoint for account baking

swagger:response getAccountTotalEndorsingListOK
*/
type GetAccountTotalEndorsingListOK struct {

	/*
	  In: Body
	*/
	Payload *models.AccountEndorsingRow `json:"body,omitempty"`
}

// NewGetAccountTotalEndorsingListOK creates GetAccountTotalEndorsingListOK with default headers values
func NewGetAccountTotalEndorsingListOK() *GetAccountTotalEndorsingListOK {

	return &GetAccountTotalEndorsingListOK{}
}

// WithPayload adds the payload to the get account total endorsing list o k response
func (o *GetAccountTotalEndorsingListOK) WithPayload(payload *models.AccountEndorsingRow) *GetAccountTotalEndorsingListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get account total endorsing list o k response
func (o *GetAccountTotalEndorsingListOK) SetPayload(payload *models.AccountEndorsingRow) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAccountTotalEndorsingListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAccountTotalEndorsingListBadRequestCode is the HTTP code returned for type GetAccountTotalEndorsingListBadRequest
const GetAccountTotalEndorsingListBadRequestCode int = 400

/*GetAccountTotalEndorsingListBadRequest Bad request

swagger:response getAccountTotalEndorsingListBadRequest
*/
type GetAccountTotalEndorsingListBadRequest struct {
}

// NewGetAccountTotalEndorsingListBadRequest creates GetAccountTotalEndorsingListBadRequest with default headers values
func NewGetAccountTotalEndorsingListBadRequest() *GetAccountTotalEndorsingListBadRequest {

	return &GetAccountTotalEndorsingListBadRequest{}
}

// WriteResponse to the client
func (o *GetAccountTotalEndorsingListBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// GetAccountTotalEndorsingListNotFoundCode is the HTTP code returned for type GetAccountTotalEndorsingListNotFound
const GetAccountTotalEndorsingListNotFoundCode int = 404

/*GetAccountTotalEndorsingListNotFound Not Found

swagger:response getAccountTotalEndorsingListNotFound
*/
type GetAccountTotalEndorsingListNotFound struct {
}

// NewGetAccountTotalEndorsingListNotFound creates GetAccountTotalEndorsingListNotFound with default headers values
func NewGetAccountTotalEndorsingListNotFound() *GetAccountTotalEndorsingListNotFound {

	return &GetAccountTotalEndorsingListNotFound{}
}

// WriteResponse to the client
func (o *GetAccountTotalEndorsingListNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
