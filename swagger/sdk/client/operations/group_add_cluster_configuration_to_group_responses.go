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

// GroupAddClusterConfigurationToGroupReader is a Reader for the GroupAddClusterConfigurationToGroup structure.
type GroupAddClusterConfigurationToGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GroupAddClusterConfigurationToGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewGroupAddClusterConfigurationToGroupCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGroupAddClusterConfigurationToGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGroupAddClusterConfigurationToGroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGroupAddClusterConfigurationToGroupUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGroupAddClusterConfigurationToGroupCreated creates a GroupAddClusterConfigurationToGroupCreated with default headers values
func NewGroupAddClusterConfigurationToGroupCreated() *GroupAddClusterConfigurationToGroupCreated {
	return &GroupAddClusterConfigurationToGroupCreated{}
}

/* GroupAddClusterConfigurationToGroupCreated describes a response with status code 201, with default header values.

Group
*/
type GroupAddClusterConfigurationToGroupCreated struct {
	Payload *models.Group
}

func (o *GroupAddClusterConfigurationToGroupCreated) Error() string {
	return fmt.Sprintf("[POST /groups/{groupId}/cluster-configuration][%d] groupAddClusterConfigurationToGroupCreated  %+v", 201, o.Payload)
}
func (o *GroupAddClusterConfigurationToGroupCreated) GetPayload() *models.Group {
	return o.Payload
}

func (o *GroupAddClusterConfigurationToGroupCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Group)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGroupAddClusterConfigurationToGroupBadRequest creates a GroupAddClusterConfigurationToGroupBadRequest with default headers values
func NewGroupAddClusterConfigurationToGroupBadRequest() *GroupAddClusterConfigurationToGroupBadRequest {
	return &GroupAddClusterConfigurationToGroupBadRequest{}
}

/* GroupAddClusterConfigurationToGroupBadRequest describes a response with status code 400, with default header values.

GroupAddClusterConfigurationToGroupBadRequest group add cluster configuration to group bad request
*/
type GroupAddClusterConfigurationToGroupBadRequest struct {
}

func (o *GroupAddClusterConfigurationToGroupBadRequest) Error() string {
	return fmt.Sprintf("[POST /groups/{groupId}/cluster-configuration][%d] groupAddClusterConfigurationToGroupBadRequest ", 400)
}

func (o *GroupAddClusterConfigurationToGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGroupAddClusterConfigurationToGroupForbidden creates a GroupAddClusterConfigurationToGroupForbidden with default headers values
func NewGroupAddClusterConfigurationToGroupForbidden() *GroupAddClusterConfigurationToGroupForbidden {
	return &GroupAddClusterConfigurationToGroupForbidden{}
}

/* GroupAddClusterConfigurationToGroupForbidden describes a response with status code 403, with default header values.

GroupAddClusterConfigurationToGroupForbidden group add cluster configuration to group forbidden
*/
type GroupAddClusterConfigurationToGroupForbidden struct {
}

func (o *GroupAddClusterConfigurationToGroupForbidden) Error() string {
	return fmt.Sprintf("[POST /groups/{groupId}/cluster-configuration][%d] groupAddClusterConfigurationToGroupForbidden ", 403)
}

func (o *GroupAddClusterConfigurationToGroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGroupAddClusterConfigurationToGroupUnsupportedMediaType creates a GroupAddClusterConfigurationToGroupUnsupportedMediaType with default headers values
func NewGroupAddClusterConfigurationToGroupUnsupportedMediaType() *GroupAddClusterConfigurationToGroupUnsupportedMediaType {
	return &GroupAddClusterConfigurationToGroupUnsupportedMediaType{}
}

/* GroupAddClusterConfigurationToGroupUnsupportedMediaType describes a response with status code 415, with default header values.

GroupAddClusterConfigurationToGroupUnsupportedMediaType group add cluster configuration to group unsupported media type
*/
type GroupAddClusterConfigurationToGroupUnsupportedMediaType struct {
}

func (o *GroupAddClusterConfigurationToGroupUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /groups/{groupId}/cluster-configuration][%d] groupAddClusterConfigurationToGroupUnsupportedMediaType ", 415)
}

func (o *GroupAddClusterConfigurationToGroupUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}