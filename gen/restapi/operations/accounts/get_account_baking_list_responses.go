// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	models "github.com/everstake/teztracker/gen/models"
)

// GetAccountBakingListOKCode is the HTTP code returned for type GetAccountBakingListOK
const GetAccountBakingListOKCode int = 200

/*GetAccountBakingListOK Query compatibility endpoint for account baking

swagger:response getAccountBakingListOK
*/
type GetAccountBakingListOK struct {
	/*The total number of data entries.

	 */
	XTotalCount int64 `json:"X-Total-Count"`

	/*
	  In: Body
	*/
	Payload []*models.AccountBakingRow `json:"body,omitempty"`
}

// NewGetAccountBakingListOK creates GetAccountBakingListOK with default headers values
func NewGetAccountBakingListOK() *GetAccountBakingListOK {

	return &GetAccountBakingListOK{}
}

// WithXTotalCount adds the xTotalCount to the get account baking list o k response
func (o *GetAccountBakingListOK) WithXTotalCount(xTotalCount int64) *GetAccountBakingListOK {
	o.XTotalCount = xTotalCount
	return o
}

// SetXTotalCount sets the xTotalCount to the get account baking list o k response
func (o *GetAccountBakingListOK) SetXTotalCount(xTotalCount int64) {
	o.XTotalCount = xTotalCount
}

// WithPayload adds the payload to the get account baking list o k response
func (o *GetAccountBakingListOK) WithPayload(payload []*models.AccountBakingRow) *GetAccountBakingListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get account baking list o k response
func (o *GetAccountBakingListOK) SetPayload(payload []*models.AccountBakingRow) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAccountBakingListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header X-Total-Count

	xTotalCount := swag.FormatInt64(o.XTotalCount)
	if xTotalCount != "" {
		rw.Header().Set("X-Total-Count", xTotalCount)
	}

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.AccountBakingRow, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetAccountBakingListBadRequestCode is the HTTP code returned for type GetAccountBakingListBadRequest
const GetAccountBakingListBadRequestCode int = 400

/*GetAccountBakingListBadRequest Bad request

swagger:response getAccountBakingListBadRequest
*/
type GetAccountBakingListBadRequest struct {
}

// NewGetAccountBakingListBadRequest creates GetAccountBakingListBadRequest with default headers values
func NewGetAccountBakingListBadRequest() *GetAccountBakingListBadRequest {

	return &GetAccountBakingListBadRequest{}
}

// WriteResponse to the client
func (o *GetAccountBakingListBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// GetAccountBakingListNotFoundCode is the HTTP code returned for type GetAccountBakingListNotFound
const GetAccountBakingListNotFoundCode int = 404

/*GetAccountBakingListNotFound Not Found

swagger:response getAccountBakingListNotFound
*/
type GetAccountBakingListNotFound struct {
}

// NewGetAccountBakingListNotFound creates GetAccountBakingListNotFound with default headers values
func NewGetAccountBakingListNotFound() *GetAccountBakingListNotFound {

	return &GetAccountBakingListNotFound{}
}

// WriteResponse to the client
func (o *GetAccountBakingListNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
