// Code generated by go-swagger; DO NOT EDIT.

package blocks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/bullblock-io/tezTracker/gen/models"
)

// GetBlockOKCode is the HTTP code returned for type GetBlockOK
const GetBlockOKCode int = 200

/*GetBlockOK Query compatibility endpoint for block by hash

swagger:response getBlockOK
*/
type GetBlockOK struct {

	/*
	  In: Body
	*/
	Payload *models.BlockResult `json:"body,omitempty"`
}

// NewGetBlockOK creates GetBlockOK with default headers values
func NewGetBlockOK() *GetBlockOK {

	return &GetBlockOK{}
}

// WithPayload adds the payload to the get block o k response
func (o *GetBlockOK) WithPayload(payload *models.BlockResult) *GetBlockOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get block o k response
func (o *GetBlockOK) SetPayload(payload *models.BlockResult) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBlockOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetBlockNotFoundCode is the HTTP code returned for type GetBlockNotFound
const GetBlockNotFoundCode int = 404

/*GetBlockNotFound Not Found

swagger:response getBlockNotFound
*/
type GetBlockNotFound struct {
}

// NewGetBlockNotFound creates GetBlockNotFound with default headers values
func NewGetBlockNotFound() *GetBlockNotFound {

	return &GetBlockNotFound{}
}

// WriteResponse to the client
func (o *GetBlockNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// GetBlockInternalServerErrorCode is the HTTP code returned for type GetBlockInternalServerError
const GetBlockInternalServerErrorCode int = 500

/*GetBlockInternalServerError Internal error

swagger:response getBlockInternalServerError
*/
type GetBlockInternalServerError struct {
}

// NewGetBlockInternalServerError creates GetBlockInternalServerError with default headers values
func NewGetBlockInternalServerError() *GetBlockInternalServerError {

	return &GetBlockInternalServerError{}
}

// WriteResponse to the client
func (o *GetBlockInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
