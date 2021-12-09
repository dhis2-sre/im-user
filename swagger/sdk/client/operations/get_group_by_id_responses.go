// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/dhis2-sre/im-users/swagger/sdk/models"
)

// GetGroupByIDReader is a Reader for the GetGroupByID structure.
type GetGroupByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGroupByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGroupByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetGroupByIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetGroupByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetGroupByIDUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGroupByIDOK creates a GetGroupByIDOK with default headers values
func NewGetGroupByIDOK() *GetGroupByIDOK {
	return &GetGroupByIDOK{}
}

/* GetGroupByIDOK describes a response with status code 200, with default header values.

Group
*/
type GetGroupByIDOK struct {
	Payload *models.Group
}

func (o *GetGroupByIDOK) Error() string {
	return fmt.Sprintf("[GET /groups/{id}][%d] getGroupByIdOK  %+v", 200, o.Payload)
}
func (o *GetGroupByIDOK) GetPayload() *models.Group {
	return o.Payload
}

func (o *GetGroupByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Group)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupByIDForbidden creates a GetGroupByIDForbidden with default headers values
func NewGetGroupByIDForbidden() *GetGroupByIDForbidden {
	return &GetGroupByIDForbidden{}
}

/* GetGroupByIDForbidden describes a response with status code 403, with default header values.

GetGroupByIDForbidden get group by Id forbidden
*/
type GetGroupByIDForbidden struct {
}

func (o *GetGroupByIDForbidden) Error() string {
	return fmt.Sprintf("[GET /groups/{id}][%d] getGroupByIdForbidden ", 403)
}

func (o *GetGroupByIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetGroupByIDNotFound creates a GetGroupByIDNotFound with default headers values
func NewGetGroupByIDNotFound() *GetGroupByIDNotFound {
	return &GetGroupByIDNotFound{}
}

/* GetGroupByIDNotFound describes a response with status code 404, with default header values.

GetGroupByIDNotFound get group by Id not found
*/
type GetGroupByIDNotFound struct {
}

func (o *GetGroupByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /groups/{id}][%d] getGroupByIdNotFound ", 404)
}

func (o *GetGroupByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetGroupByIDUnsupportedMediaType creates a GetGroupByIDUnsupportedMediaType with default headers values
func NewGetGroupByIDUnsupportedMediaType() *GetGroupByIDUnsupportedMediaType {
	return &GetGroupByIDUnsupportedMediaType{}
}

/* GetGroupByIDUnsupportedMediaType describes a response with status code 415, with default header values.

GetGroupByIDUnsupportedMediaType get group by Id unsupported media type
*/
type GetGroupByIDUnsupportedMediaType struct {
}

func (o *GetGroupByIDUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /groups/{id}][%d] getGroupByIdUnsupportedMediaType ", 415)
}

func (o *GetGroupByIDUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
