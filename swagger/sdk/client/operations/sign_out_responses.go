// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// SignOutReader is a Reader for the SignOut structure.
type SignOutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SignOutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSignOutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewSignOutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewSignOutUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSignOutOK creates a SignOutOK with default headers values
func NewSignOutOK() *SignOutOK {
	return &SignOutOK{}
}

/* SignOutOK describes a response with status code 200, with default header values.

SignOutOK sign out o k
*/
type SignOutOK struct {
}

func (o *SignOutOK) Error() string {
	return fmt.Sprintf("[GET /signout][%d] signOutOK ", 200)
}

func (o *SignOutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSignOutUnauthorized creates a SignOutUnauthorized with default headers values
func NewSignOutUnauthorized() *SignOutUnauthorized {
	return &SignOutUnauthorized{}
}

/* SignOutUnauthorized describes a response with status code 401, with default header values.

SignOutUnauthorized sign out unauthorized
*/
type SignOutUnauthorized struct {
}

func (o *SignOutUnauthorized) Error() string {
	return fmt.Sprintf("[GET /signout][%d] signOutUnauthorized ", 401)
}

func (o *SignOutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSignOutUnsupportedMediaType creates a SignOutUnsupportedMediaType with default headers values
func NewSignOutUnsupportedMediaType() *SignOutUnsupportedMediaType {
	return &SignOutUnsupportedMediaType{}
}

/* SignOutUnsupportedMediaType describes a response with status code 415, with default header values.

SignOutUnsupportedMediaType sign out unsupported media type
*/
type SignOutUnsupportedMediaType struct {
}

func (o *SignOutUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /signout][%d] signOutUnsupportedMediaType ", 415)
}

func (o *SignOutUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}