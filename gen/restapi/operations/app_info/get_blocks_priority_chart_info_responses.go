// Code generated by go-swagger; DO NOT EDIT.

package app_info

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/everstake/teztracker/gen/models"
)

// GetBlocksPriorityChartInfoOKCode is the HTTP code returned for type GetBlocksPriorityChartInfoOK
const GetBlocksPriorityChartInfoOKCode int = 200

/*GetBlocksPriorityChartInfoOK Application info endpoint

swagger:response getBlocksPriorityChartInfoOK
*/
type GetBlocksPriorityChartInfoOK struct {

	/*
	  In: Body
	*/
	Payload []*models.BlockPriorityChartData `json:"body,omitempty"`
}

// NewGetBlocksPriorityChartInfoOK creates GetBlocksPriorityChartInfoOK with default headers values
func NewGetBlocksPriorityChartInfoOK() *GetBlocksPriorityChartInfoOK {

	return &GetBlocksPriorityChartInfoOK{}
}

// WithPayload adds the payload to the get blocks priority chart info o k response
func (o *GetBlocksPriorityChartInfoOK) WithPayload(payload []*models.BlockPriorityChartData) *GetBlocksPriorityChartInfoOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get blocks priority chart info o k response
func (o *GetBlocksPriorityChartInfoOK) SetPayload(payload []*models.BlockPriorityChartData) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBlocksPriorityChartInfoOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.BlockPriorityChartData, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetBlocksPriorityChartInfoBadRequestCode is the HTTP code returned for type GetBlocksPriorityChartInfoBadRequest
const GetBlocksPriorityChartInfoBadRequestCode int = 400

/*GetBlocksPriorityChartInfoBadRequest Bad request

swagger:response getBlocksPriorityChartInfoBadRequest
*/
type GetBlocksPriorityChartInfoBadRequest struct {
}

// NewGetBlocksPriorityChartInfoBadRequest creates GetBlocksPriorityChartInfoBadRequest with default headers values
func NewGetBlocksPriorityChartInfoBadRequest() *GetBlocksPriorityChartInfoBadRequest {

	return &GetBlocksPriorityChartInfoBadRequest{}
}

// WriteResponse to the client
func (o *GetBlocksPriorityChartInfoBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// GetBlocksPriorityChartInfoInternalServerErrorCode is the HTTP code returned for type GetBlocksPriorityChartInfoInternalServerError
const GetBlocksPriorityChartInfoInternalServerErrorCode int = 500

/*GetBlocksPriorityChartInfoInternalServerError Internal error

swagger:response getBlocksPriorityChartInfoInternalServerError
*/
type GetBlocksPriorityChartInfoInternalServerError struct {
}

// NewGetBlocksPriorityChartInfoInternalServerError creates GetBlocksPriorityChartInfoInternalServerError with default headers values
func NewGetBlocksPriorityChartInfoInternalServerError() *GetBlocksPriorityChartInfoInternalServerError {

	return &GetBlocksPriorityChartInfoInternalServerError{}
}

// WriteResponse to the client
func (o *GetBlocksPriorityChartInfoInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
