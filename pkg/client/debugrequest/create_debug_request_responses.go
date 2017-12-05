// Code generated by go-swagger; DO NOT EDIT.

package debugrequest

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/solo-io/squash/pkg/models"
)

// CreateDebugRequestReader is a Reader for the CreateDebugRequest structure.
type CreateDebugRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateDebugRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateDebugRequestCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateDebugRequestCreated creates a CreateDebugRequestCreated with default headers values
func NewCreateDebugRequestCreated() *CreateDebugRequestCreated {
	return &CreateDebugRequestCreated{}
}

/*CreateDebugRequestCreated handles this case with default header values.

OK
*/
type CreateDebugRequestCreated struct {
	Payload *models.DebugRequest
}

func (o *CreateDebugRequestCreated) Error() string {
	return fmt.Sprintf("[POST /debugrequest][%d] createDebugRequestCreated  %+v", 201, o.Payload)
}

func (o *CreateDebugRequestCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DebugRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
