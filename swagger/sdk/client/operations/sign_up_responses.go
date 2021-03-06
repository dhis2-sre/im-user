// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/dhis2-sre/im-user/swagger/sdk/models"
)

// SignUpReader is a Reader for the SignUp structure.
type SignUpReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SignUpReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewSignUpCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSignUpBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewSignUpUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSignUpCreated creates a SignUpCreated with default headers values
func NewSignUpCreated() *SignUpCreated {
	return &SignUpCreated{}
}

/* SignUpCreated describes a response with status code 201, with default header values.

User
*/
type SignUpCreated struct {
	Payload *models.User
}

func (o *SignUpCreated) Error() string {
	return fmt.Sprintf("[POST /users][%d] signUpCreated  %+v", 201, o.Payload)
}
func (o *SignUpCreated) GetPayload() *models.User {
	return o.Payload
}

func (o *SignUpCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSignUpBadRequest creates a SignUpBadRequest with default headers values
func NewSignUpBadRequest() *SignUpBadRequest {
	return &SignUpBadRequest{}
}

/* SignUpBadRequest describes a response with status code 400, with default header values.

SignUpBadRequest sign up bad request
*/
type SignUpBadRequest struct {
}

func (o *SignUpBadRequest) Error() string {
	return fmt.Sprintf("[POST /users][%d] signUpBadRequest ", 400)
}

func (o *SignUpBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSignUpUnsupportedMediaType creates a SignUpUnsupportedMediaType with default headers values
func NewSignUpUnsupportedMediaType() *SignUpUnsupportedMediaType {
	return &SignUpUnsupportedMediaType{}
}

/* SignUpUnsupportedMediaType describes a response with status code 415, with default header values.

SignUpUnsupportedMediaType sign up unsupported media type
*/
type SignUpUnsupportedMediaType struct {
}

func (o *SignUpUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /users][%d] signUpUnsupportedMediaType ", 415)
}

func (o *SignUpUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
