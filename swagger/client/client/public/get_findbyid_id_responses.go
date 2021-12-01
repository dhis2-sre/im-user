// Code generated by go-swagger; DO NOT EDIT.

package public

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/dhis2-sre/im-users/swagger/client/models"
)

// GetFindbyidIDReader is a Reader for the GetFindbyidID structure.
type GetFindbyidIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFindbyidIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFindbyidIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetFindbyidIDOK creates a GetFindbyidIDOK with default headers values
func NewGetFindbyidIDOK() *GetFindbyidIDOK {
	return &GetFindbyidIDOK{}
}

/* GetFindbyidIDOK describes a response with status code 200, with default header values.

OK
*/
type GetFindbyidIDOK struct {
	Payload *models.DtoUser
}

func (o *GetFindbyidIDOK) Error() string {
	return fmt.Sprintf("[GET /findbyid/{id}][%d] getFindbyidIdOK  %+v", 200, o.Payload)
}
func (o *GetFindbyidIDOK) GetPayload() *models.DtoUser {
	return o.Payload
}

func (o *GetFindbyidIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DtoUser)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}