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

// PatchDebugAttachmentReader is a Reader for the PatchDebugAttachment structure.
type PatchDebugAttachmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchDebugAttachmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPatchDebugAttachmentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPatchDebugAttachmentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPatchDebugAttachmentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewPatchDebugAttachmentConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewPatchDebugAttachmentUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewPatchDebugAttachmentServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchDebugAttachmentOK creates a PatchDebugAttachmentOK with default headers values
func NewPatchDebugAttachmentOK() *PatchDebugAttachmentOK {
	return &PatchDebugAttachmentOK{}
}

/*PatchDebugAttachmentOK handles this case with default header values.

Debug attachment modified
*/
type PatchDebugAttachmentOK struct {
	Payload *models.DebugAttachment
}

func (o *PatchDebugAttachmentOK) Error() string {
	return fmt.Sprintf("[PATCH /debugattachment/{debugAttachmentId}][%d] patchDebugAttachmentOK  %+v", 200, o.Payload)
}

func (o *PatchDebugAttachmentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DebugAttachment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchDebugAttachmentBadRequest creates a PatchDebugAttachmentBadRequest with default headers values
func NewPatchDebugAttachmentBadRequest() *PatchDebugAttachmentBadRequest {
	return &PatchDebugAttachmentBadRequest{}
}

/*PatchDebugAttachmentBadRequest handles this case with default header values.

Bad request
*/
type PatchDebugAttachmentBadRequest struct {
}

func (o *PatchDebugAttachmentBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /debugattachment/{debugAttachmentId}][%d] patchDebugAttachmentBadRequest ", 400)
}

func (o *PatchDebugAttachmentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchDebugAttachmentNotFound creates a PatchDebugAttachmentNotFound with default headers values
func NewPatchDebugAttachmentNotFound() *PatchDebugAttachmentNotFound {
	return &PatchDebugAttachmentNotFound{}
}

/*PatchDebugAttachmentNotFound handles this case with default header values.

Not found
*/
type PatchDebugAttachmentNotFound struct {
}

func (o *PatchDebugAttachmentNotFound) Error() string {
	return fmt.Sprintf("[PATCH /debugattachment/{debugAttachmentId}][%d] patchDebugAttachmentNotFound ", 404)
}

func (o *PatchDebugAttachmentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchDebugAttachmentConflict creates a PatchDebugAttachmentConflict with default headers values
func NewPatchDebugAttachmentConflict() *PatchDebugAttachmentConflict {
	return &PatchDebugAttachmentConflict{}
}

/*PatchDebugAttachmentConflict handles this case with default header values.

Conflict
*/
type PatchDebugAttachmentConflict struct {
}

func (o *PatchDebugAttachmentConflict) Error() string {
	return fmt.Sprintf("[PATCH /debugattachment/{debugAttachmentId}][%d] patchDebugAttachmentConflict ", 409)
}

func (o *PatchDebugAttachmentConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchDebugAttachmentUnprocessableEntity creates a PatchDebugAttachmentUnprocessableEntity with default headers values
func NewPatchDebugAttachmentUnprocessableEntity() *PatchDebugAttachmentUnprocessableEntity {
	return &PatchDebugAttachmentUnprocessableEntity{}
}

/*PatchDebugAttachmentUnprocessableEntity handles this case with default header values.

Invalid input
*/
type PatchDebugAttachmentUnprocessableEntity struct {
}

func (o *PatchDebugAttachmentUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PATCH /debugattachment/{debugAttachmentId}][%d] patchDebugAttachmentUnprocessableEntity ", 422)
}

func (o *PatchDebugAttachmentUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchDebugAttachmentServiceUnavailable creates a PatchDebugAttachmentServiceUnavailable with default headers values
func NewPatchDebugAttachmentServiceUnavailable() *PatchDebugAttachmentServiceUnavailable {
	return &PatchDebugAttachmentServiceUnavailable{}
}

/*PatchDebugAttachmentServiceUnavailable handles this case with default header values.

Service Unavailable
*/
type PatchDebugAttachmentServiceUnavailable struct {
}

func (o *PatchDebugAttachmentServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /debugattachment/{debugAttachmentId}][%d] patchDebugAttachmentServiceUnavailable ", 503)
}

func (o *PatchDebugAttachmentServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
