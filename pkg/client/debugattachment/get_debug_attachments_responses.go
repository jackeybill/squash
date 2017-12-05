// Code generated by go-swagger; DO NOT EDIT.

package debugattachment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/solo-io/squash/pkg/models"
)

// GetDebugAttachmentsReader is a Reader for the GetDebugAttachments structure.
type GetDebugAttachmentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDebugAttachmentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDebugAttachmentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 408:
		result := NewGetDebugAttachmentsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewGetDebugAttachmentsUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDebugAttachmentsOK creates a GetDebugAttachmentsOK with default headers values
func NewGetDebugAttachmentsOK() *GetDebugAttachmentsOK {
	return &GetDebugAttachmentsOK{}
}

/*GetDebugAttachmentsOK handles this case with default header values.

OK
*/
type GetDebugAttachmentsOK struct {
	Payload models.GetDebugAttachmentsOKBody
}

func (o *GetDebugAttachmentsOK) Error() string {
	return fmt.Sprintf("[GET /debugattachment][%d] getDebugAttachmentsOK  %+v", 200, o.Payload)
}

func (o *GetDebugAttachmentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDebugAttachmentsRequestTimeout creates a GetDebugAttachmentsRequestTimeout with default headers values
func NewGetDebugAttachmentsRequestTimeout() *GetDebugAttachmentsRequestTimeout {
	return &GetDebugAttachmentsRequestTimeout{}
}

/*GetDebugAttachmentsRequestTimeout handles this case with default header values.

Request timed out
*/
type GetDebugAttachmentsRequestTimeout struct {
}

func (o *GetDebugAttachmentsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /debugattachment][%d] getDebugAttachmentsRequestTimeout ", 408)
}

func (o *GetDebugAttachmentsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDebugAttachmentsUnprocessableEntity creates a GetDebugAttachmentsUnprocessableEntity with default headers values
func NewGetDebugAttachmentsUnprocessableEntity() *GetDebugAttachmentsUnprocessableEntity {
	return &GetDebugAttachmentsUnprocessableEntity{}
}

/*GetDebugAttachmentsUnprocessableEntity handles this case with default header values.

Validation exception
*/
type GetDebugAttachmentsUnprocessableEntity struct {
}

func (o *GetDebugAttachmentsUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /debugattachment][%d] getDebugAttachmentsUnprocessableEntity ", 422)
}

func (o *GetDebugAttachmentsUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
