// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/everstake/teztracker/gen/models"
)

// GetAccountAssetsBalancesListOKCode is the HTTP code returned for type GetAccountAssetsBalancesListOK
const GetAccountAssetsBalancesListOKCode int = 200

/*GetAccountAssetsBalancesListOK Query compatibility endpoint for account

swagger:response getAccountAssetsBalancesListOK
*/
type GetAccountAssetsBalancesListOK struct {

	/*
	  In: Body
	*/
	Payload []*models.AccountAssetBalanceRow `json:"body,omitempty"`
}

// NewGetAccountAssetsBalancesListOK creates GetAccountAssetsBalancesListOK with default headers values
func NewGetAccountAssetsBalancesListOK() *GetAccountAssetsBalancesListOK {

	return &GetAccountAssetsBalancesListOK{}
}

// WithPayload adds the payload to the get account assets balances list o k response
func (o *GetAccountAssetsBalancesListOK) WithPayload(payload []*models.AccountAssetBalanceRow) *GetAccountAssetsBalancesListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get account assets balances list o k response
func (o *GetAccountAssetsBalancesListOK) SetPayload(payload []*models.AccountAssetBalanceRow) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAccountAssetsBalancesListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.AccountAssetBalanceRow, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetAccountAssetsBalancesListBadRequestCode is the HTTP code returned for type GetAccountAssetsBalancesListBadRequest
const GetAccountAssetsBalancesListBadRequestCode int = 400

/*GetAccountAssetsBalancesListBadRequest Bad request

swagger:response getAccountAssetsBalancesListBadRequest
*/
type GetAccountAssetsBalancesListBadRequest struct {
}

// NewGetAccountAssetsBalancesListBadRequest creates GetAccountAssetsBalancesListBadRequest with default headers values
func NewGetAccountAssetsBalancesListBadRequest() *GetAccountAssetsBalancesListBadRequest {

	return &GetAccountAssetsBalancesListBadRequest{}
}

// WriteResponse to the client
func (o *GetAccountAssetsBalancesListBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// GetAccountAssetsBalancesListNotFoundCode is the HTTP code returned for type GetAccountAssetsBalancesListNotFound
const GetAccountAssetsBalancesListNotFoundCode int = 404

/*GetAccountAssetsBalancesListNotFound Not Found

swagger:response getAccountAssetsBalancesListNotFound
*/
type GetAccountAssetsBalancesListNotFound struct {
}

// NewGetAccountAssetsBalancesListNotFound creates GetAccountAssetsBalancesListNotFound with default headers values
func NewGetAccountAssetsBalancesListNotFound() *GetAccountAssetsBalancesListNotFound {

	return &GetAccountAssetsBalancesListNotFound{}
}

// WriteResponse to the client
func (o *GetAccountAssetsBalancesListNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// GetAccountAssetsBalancesListInternalServerErrorCode is the HTTP code returned for type GetAccountAssetsBalancesListInternalServerError
const GetAccountAssetsBalancesListInternalServerErrorCode int = 500

/*GetAccountAssetsBalancesListInternalServerError Internal error

swagger:response getAccountAssetsBalancesListInternalServerError
*/
type GetAccountAssetsBalancesListInternalServerError struct {
}

// NewGetAccountAssetsBalancesListInternalServerError creates GetAccountAssetsBalancesListInternalServerError with default headers values
func NewGetAccountAssetsBalancesListInternalServerError() *GetAccountAssetsBalancesListInternalServerError {

	return &GetAccountAssetsBalancesListInternalServerError{}
}

// WriteResponse to the client
func (o *GetAccountAssetsBalancesListInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}