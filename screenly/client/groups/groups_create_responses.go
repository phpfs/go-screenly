// Code generated by go-swagger; DO NOT EDIT.

package groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/phpfs/go-screenly/screenly/models"
)

// GroupsCreateReader is a Reader for the GroupsCreate structure.
type GroupsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GroupsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewGroupsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGroupsCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGroupsCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGroupsCreateCreated creates a GroupsCreateCreated with default headers values
func NewGroupsCreateCreated() *GroupsCreateCreated {
	return &GroupsCreateCreated{}
}

/* GroupsCreateCreated describes a response with status code 201, with default header values.

GroupsCreateCreated groups create created
*/
type GroupsCreateCreated struct {
	Payload *models.GroupRead
}

func (o *GroupsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /groups/][%d] groupsCreateCreated  %+v", 201, o.Payload)
}
func (o *GroupsCreateCreated) GetPayload() *models.GroupRead {
	return o.Payload
}

func (o *GroupsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GroupRead)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGroupsCreateBadRequest creates a GroupsCreateBadRequest with default headers values
func NewGroupsCreateBadRequest() *GroupsCreateBadRequest {
	return &GroupsCreateBadRequest{}
}

/* GroupsCreateBadRequest describes a response with status code 400, with default header values.

You sent a malformed or bad request.
*/
type GroupsCreateBadRequest struct {
}

func (o *GroupsCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /groups/][%d] groupsCreateBadRequest ", 400)
}

func (o *GroupsCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGroupsCreateUnauthorized creates a GroupsCreateUnauthorized with default headers values
func NewGroupsCreateUnauthorized() *GroupsCreateUnauthorized {
	return &GroupsCreateUnauthorized{}
}

/* GroupsCreateUnauthorized describes a response with status code 401, with default header values.

You provided invalid credentials.
*/
type GroupsCreateUnauthorized struct {
}

func (o *GroupsCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /groups/][%d] groupsCreateUnauthorized ", 401)
}

func (o *GroupsCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
