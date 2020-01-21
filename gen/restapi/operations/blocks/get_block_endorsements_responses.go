// Code generated by go-swagger; DO NOT EDIT.

package blocks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	models "github.com/bullblock-io/tezTracker/gen/models"
)

// GetBlockEndorsementsOKCode is the HTTP code returned for type GetBlockEndorsementsOK
const GetBlockEndorsementsOKCode int = 200

/*GetBlockEndorsementsOK Endpoint for endorsements of a block by hash or level

swagger:response getBlockEndorsementsOK
*/
type GetBlockEndorsementsOK struct {
	/*The total number of data entries.

	 */
	XTotalCount int64 `json:"X-Total-Count"`

	/*
	  In: Body
	*/
	Payload []*models.OperationsRow `json:"body,omitempty"`
}

// NewGetBlockEndorsementsOK creates GetBlockEndorsementsOK with default headers values
func NewGetBlockEndorsementsOK() *GetBlockEndorsementsOK {

	return &GetBlockEndorsementsOK{}
}

// WithXTotalCount adds the xTotalCount to the get block endorsements o k response
func (o *GetBlockEndorsementsOK) WithXTotalCount(xTotalCount int64) *GetBlockEndorsementsOK {
	o.XTotalCount = xTotalCount
	return o
}

// SetXTotalCount sets the xTotalCount to the get block endorsements o k response
func (o *GetBlockEndorsementsOK) SetXTotalCount(xTotalCount int64) {
	o.XTotalCount = xTotalCount
}

// WithPayload adds the payload to the get block endorsements o k response
func (o *GetBlockEndorsementsOK) WithPayload(payload []*models.OperationsRow) *GetBlockEndorsementsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get block endorsements o k response
func (o *GetBlockEndorsementsOK) SetPayload(payload []*models.OperationsRow) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBlockEndorsementsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header X-Total-Count

	xTotalCount := swag.FormatInt64(o.XTotalCount)
	if xTotalCount != "" {
		rw.Header().Set("X-Total-Count", xTotalCount)
	}

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.OperationsRow, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetBlockEndorsementsBadRequestCode is the HTTP code returned for type GetBlockEndorsementsBadRequest
const GetBlockEndorsementsBadRequestCode int = 400

/*GetBlockEndorsementsBadRequest Bad request

swagger:response getBlockEndorsementsBadRequest
*/
type GetBlockEndorsementsBadRequest struct {
}

// NewGetBlockEndorsementsBadRequest creates GetBlockEndorsementsBadRequest with default headers values
func NewGetBlockEndorsementsBadRequest() *GetBlockEndorsementsBadRequest {

	return &GetBlockEndorsementsBadRequest{}
}

// WriteResponse to the client
func (o *GetBlockEndorsementsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// GetBlockEndorsementsNotFoundCode is the HTTP code returned for type GetBlockEndorsementsNotFound
const GetBlockEndorsementsNotFoundCode int = 404

/*GetBlockEndorsementsNotFound Not Found

swagger:response getBlockEndorsementsNotFound
*/
type GetBlockEndorsementsNotFound struct {
}

// NewGetBlockEndorsementsNotFound creates GetBlockEndorsementsNotFound with default headers values
func NewGetBlockEndorsementsNotFound() *GetBlockEndorsementsNotFound {

	return &GetBlockEndorsementsNotFound{}
}

// WriteResponse to the client
func (o *GetBlockEndorsementsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// GetBlockEndorsementsInternalServerErrorCode is the HTTP code returned for type GetBlockEndorsementsInternalServerError
const GetBlockEndorsementsInternalServerErrorCode int = 500

/*GetBlockEndorsementsInternalServerError Internal error

swagger:response getBlockEndorsementsInternalServerError
*/
type GetBlockEndorsementsInternalServerError struct {
}

// NewGetBlockEndorsementsInternalServerError creates GetBlockEndorsementsInternalServerError with default headers values
func NewGetBlockEndorsementsInternalServerError() *GetBlockEndorsementsInternalServerError {

	return &GetBlockEndorsementsInternalServerError{}
}

// WriteResponse to the client
func (o *GetBlockEndorsementsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
