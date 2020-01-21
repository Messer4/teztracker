// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	models "github.com/bullblock-io/tezTracker/gen/models"
)

// GetAccountDelegatorsOKCode is the HTTP code returned for type GetAccountDelegatorsOK
const GetAccountDelegatorsOKCode int = 200

/*GetAccountDelegatorsOK Endpoint for account delegators

swagger:response getAccountDelegatorsOK
*/
type GetAccountDelegatorsOK struct {
	/*The total number of data entries.

	 */
	XTotalCount int64 `json:"X-Total-Count"`

	/*
	  In: Body
	*/
	Payload []*models.AccountsRow `json:"body,omitempty"`
}

// NewGetAccountDelegatorsOK creates GetAccountDelegatorsOK with default headers values
func NewGetAccountDelegatorsOK() *GetAccountDelegatorsOK {

	return &GetAccountDelegatorsOK{}
}

// WithXTotalCount adds the xTotalCount to the get account delegators o k response
func (o *GetAccountDelegatorsOK) WithXTotalCount(xTotalCount int64) *GetAccountDelegatorsOK {
	o.XTotalCount = xTotalCount
	return o
}

// SetXTotalCount sets the xTotalCount to the get account delegators o k response
func (o *GetAccountDelegatorsOK) SetXTotalCount(xTotalCount int64) {
	o.XTotalCount = xTotalCount
}

// WithPayload adds the payload to the get account delegators o k response
func (o *GetAccountDelegatorsOK) WithPayload(payload []*models.AccountsRow) *GetAccountDelegatorsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get account delegators o k response
func (o *GetAccountDelegatorsOK) SetPayload(payload []*models.AccountsRow) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAccountDelegatorsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header X-Total-Count

	xTotalCount := swag.FormatInt64(o.XTotalCount)
	if xTotalCount != "" {
		rw.Header().Set("X-Total-Count", xTotalCount)
	}

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.AccountsRow, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetAccountDelegatorsBadRequestCode is the HTTP code returned for type GetAccountDelegatorsBadRequest
const GetAccountDelegatorsBadRequestCode int = 400

/*GetAccountDelegatorsBadRequest Bad request

swagger:response getAccountDelegatorsBadRequest
*/
type GetAccountDelegatorsBadRequest struct {
}

// NewGetAccountDelegatorsBadRequest creates GetAccountDelegatorsBadRequest with default headers values
func NewGetAccountDelegatorsBadRequest() *GetAccountDelegatorsBadRequest {

	return &GetAccountDelegatorsBadRequest{}
}

// WriteResponse to the client
func (o *GetAccountDelegatorsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// GetAccountDelegatorsInternalServerErrorCode is the HTTP code returned for type GetAccountDelegatorsInternalServerError
const GetAccountDelegatorsInternalServerErrorCode int = 500

/*GetAccountDelegatorsInternalServerError Internal error

swagger:response getAccountDelegatorsInternalServerError
*/
type GetAccountDelegatorsInternalServerError struct {
}

// NewGetAccountDelegatorsInternalServerError creates GetAccountDelegatorsInternalServerError with default headers values
func NewGetAccountDelegatorsInternalServerError() *GetAccountDelegatorsInternalServerError {

	return &GetAccountDelegatorsInternalServerError{}
}

// WriteResponse to the client
func (o *GetAccountDelegatorsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
