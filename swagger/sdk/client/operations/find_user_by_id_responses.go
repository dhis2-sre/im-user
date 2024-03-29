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

// FindUserByIDReader is a Reader for the FindUserByID structure.
type FindUserByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindUserByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindUserByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewFindUserByIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewFindUserByIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewFindUserByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewFindUserByIDUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewFindUserByIDOK creates a FindUserByIDOK with default headers values
func NewFindUserByIDOK() *FindUserByIDOK {
	return &FindUserByIDOK{}
}

/*
FindUserByIDOK describes a response with status code 200, with default header values.

User
*/
type FindUserByIDOK struct {
	Payload *models.User
}

// IsSuccess returns true when this find user by Id o k response has a 2xx status code
func (o *FindUserByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find user by Id o k response has a 3xx status code
func (o *FindUserByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find user by Id o k response has a 4xx status code
func (o *FindUserByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find user by Id o k response has a 5xx status code
func (o *FindUserByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find user by Id o k response a status code equal to that given
func (o *FindUserByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the find user by Id o k response
func (o *FindUserByIDOK) Code() int {
	return 200
}

func (o *FindUserByIDOK) Error() string {
	return fmt.Sprintf("[GET /users/{id}][%d] findUserByIdOK  %+v", 200, o.Payload)
}

func (o *FindUserByIDOK) String() string {
	return fmt.Sprintf("[GET /users/{id}][%d] findUserByIdOK  %+v", 200, o.Payload)
}

func (o *FindUserByIDOK) GetPayload() *models.User {
	return o.Payload
}

func (o *FindUserByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindUserByIDUnauthorized creates a FindUserByIDUnauthorized with default headers values
func NewFindUserByIDUnauthorized() *FindUserByIDUnauthorized {
	return &FindUserByIDUnauthorized{}
}

/*
FindUserByIDUnauthorized describes a response with status code 401, with default header values.

FindUserByIDUnauthorized find user by Id unauthorized
*/
type FindUserByIDUnauthorized struct {
}

// IsSuccess returns true when this find user by Id unauthorized response has a 2xx status code
func (o *FindUserByIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this find user by Id unauthorized response has a 3xx status code
func (o *FindUserByIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find user by Id unauthorized response has a 4xx status code
func (o *FindUserByIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this find user by Id unauthorized response has a 5xx status code
func (o *FindUserByIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this find user by Id unauthorized response a status code equal to that given
func (o *FindUserByIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the find user by Id unauthorized response
func (o *FindUserByIDUnauthorized) Code() int {
	return 401
}

func (o *FindUserByIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /users/{id}][%d] findUserByIdUnauthorized ", 401)
}

func (o *FindUserByIDUnauthorized) String() string {
	return fmt.Sprintf("[GET /users/{id}][%d] findUserByIdUnauthorized ", 401)
}

func (o *FindUserByIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFindUserByIDForbidden creates a FindUserByIDForbidden with default headers values
func NewFindUserByIDForbidden() *FindUserByIDForbidden {
	return &FindUserByIDForbidden{}
}

/*
FindUserByIDForbidden describes a response with status code 403, with default header values.

FindUserByIDForbidden find user by Id forbidden
*/
type FindUserByIDForbidden struct {
}

// IsSuccess returns true when this find user by Id forbidden response has a 2xx status code
func (o *FindUserByIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this find user by Id forbidden response has a 3xx status code
func (o *FindUserByIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find user by Id forbidden response has a 4xx status code
func (o *FindUserByIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this find user by Id forbidden response has a 5xx status code
func (o *FindUserByIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this find user by Id forbidden response a status code equal to that given
func (o *FindUserByIDForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the find user by Id forbidden response
func (o *FindUserByIDForbidden) Code() int {
	return 403
}

func (o *FindUserByIDForbidden) Error() string {
	return fmt.Sprintf("[GET /users/{id}][%d] findUserByIdForbidden ", 403)
}

func (o *FindUserByIDForbidden) String() string {
	return fmt.Sprintf("[GET /users/{id}][%d] findUserByIdForbidden ", 403)
}

func (o *FindUserByIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFindUserByIDNotFound creates a FindUserByIDNotFound with default headers values
func NewFindUserByIDNotFound() *FindUserByIDNotFound {
	return &FindUserByIDNotFound{}
}

/*
FindUserByIDNotFound describes a response with status code 404, with default header values.

FindUserByIDNotFound find user by Id not found
*/
type FindUserByIDNotFound struct {
}

// IsSuccess returns true when this find user by Id not found response has a 2xx status code
func (o *FindUserByIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this find user by Id not found response has a 3xx status code
func (o *FindUserByIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find user by Id not found response has a 4xx status code
func (o *FindUserByIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this find user by Id not found response has a 5xx status code
func (o *FindUserByIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this find user by Id not found response a status code equal to that given
func (o *FindUserByIDNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the find user by Id not found response
func (o *FindUserByIDNotFound) Code() int {
	return 404
}

func (o *FindUserByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /users/{id}][%d] findUserByIdNotFound ", 404)
}

func (o *FindUserByIDNotFound) String() string {
	return fmt.Sprintf("[GET /users/{id}][%d] findUserByIdNotFound ", 404)
}

func (o *FindUserByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFindUserByIDUnsupportedMediaType creates a FindUserByIDUnsupportedMediaType with default headers values
func NewFindUserByIDUnsupportedMediaType() *FindUserByIDUnsupportedMediaType {
	return &FindUserByIDUnsupportedMediaType{}
}

/*
FindUserByIDUnsupportedMediaType describes a response with status code 415, with default header values.

FindUserByIDUnsupportedMediaType find user by Id unsupported media type
*/
type FindUserByIDUnsupportedMediaType struct {
}

// IsSuccess returns true when this find user by Id unsupported media type response has a 2xx status code
func (o *FindUserByIDUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this find user by Id unsupported media type response has a 3xx status code
func (o *FindUserByIDUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find user by Id unsupported media type response has a 4xx status code
func (o *FindUserByIDUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this find user by Id unsupported media type response has a 5xx status code
func (o *FindUserByIDUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this find user by Id unsupported media type response a status code equal to that given
func (o *FindUserByIDUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

// Code gets the status code for the find user by Id unsupported media type response
func (o *FindUserByIDUnsupportedMediaType) Code() int {
	return 415
}

func (o *FindUserByIDUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /users/{id}][%d] findUserByIdUnsupportedMediaType ", 415)
}

func (o *FindUserByIDUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /users/{id}][%d] findUserByIdUnsupportedMediaType ", 415)
}

func (o *FindUserByIDUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
