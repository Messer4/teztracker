// Code generated by go-swagger; DO NOT EDIT.

package voting

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetCurrentVotingPeriodReader is a Reader for the GetCurrentVotingPeriod structure.
type GetCurrentVotingPeriodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCurrentVotingPeriodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCurrentVotingPeriodOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetCurrentVotingPeriodInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCurrentVotingPeriodOK creates a GetCurrentVotingPeriodOK with default headers values
func NewGetCurrentVotingPeriodOK() *GetCurrentVotingPeriodOK {
	return &GetCurrentVotingPeriodOK{}
}

/*GetCurrentVotingPeriodOK handles this case with default header values.

Endpoint for current period kind
*/
type GetCurrentVotingPeriodOK struct {
	Payload string
}

func (o *GetCurrentVotingPeriodOK) Error() string {
	return fmt.Sprintf("[GET /chains/{network}/blocks/{block}/votes/current_period_kind][%d] getCurrentVotingPeriodOK  %+v", 200, o.Payload)
}

func (o *GetCurrentVotingPeriodOK) GetPayload() string {
	return o.Payload
}

func (o *GetCurrentVotingPeriodOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCurrentVotingPeriodInternalServerError creates a GetCurrentVotingPeriodInternalServerError with default headers values
func NewGetCurrentVotingPeriodInternalServerError() *GetCurrentVotingPeriodInternalServerError {
	return &GetCurrentVotingPeriodInternalServerError{}
}

/*GetCurrentVotingPeriodInternalServerError handles this case with default header values.

Internal error
*/
type GetCurrentVotingPeriodInternalServerError struct {
}

func (o *GetCurrentVotingPeriodInternalServerError) Error() string {
	return fmt.Sprintf("[GET /chains/{network}/blocks/{block}/votes/current_period_kind][%d] getCurrentVotingPeriodInternalServerError ", 500)
}

func (o *GetCurrentVotingPeriodInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
