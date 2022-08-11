// Code generated by go-swagger; DO NOT EDIT.

package playlists

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/phpfs/go-screenly/screenly/models"
)

// PlaylistsUpdateReader is a Reader for the PlaylistsUpdate structure.
type PlaylistsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PlaylistsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPlaylistsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPlaylistsUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPlaylistsUpdateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPlaylistsUpdateOK creates a PlaylistsUpdateOK with default headers values
func NewPlaylistsUpdateOK() *PlaylistsUpdateOK {
	return &PlaylistsUpdateOK{}
}

/* PlaylistsUpdateOK describes a response with status code 200, with default header values.

PlaylistsUpdateOK playlists update o k
*/
type PlaylistsUpdateOK struct {
	Payload *models.PlaylistRead
}

func (o *PlaylistsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /playlists/{id}/][%d] playlistsUpdateOK  %+v", 200, o.Payload)
}
func (o *PlaylistsUpdateOK) GetPayload() *models.PlaylistRead {
	return o.Payload
}

func (o *PlaylistsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PlaylistRead)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPlaylistsUpdateBadRequest creates a PlaylistsUpdateBadRequest with default headers values
func NewPlaylistsUpdateBadRequest() *PlaylistsUpdateBadRequest {
	return &PlaylistsUpdateBadRequest{}
}

/* PlaylistsUpdateBadRequest describes a response with status code 400, with default header values.

You sent a malformed or bad request.
*/
type PlaylistsUpdateBadRequest struct {
}

func (o *PlaylistsUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /playlists/{id}/][%d] playlistsUpdateBadRequest ", 400)
}

func (o *PlaylistsUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPlaylistsUpdateUnauthorized creates a PlaylistsUpdateUnauthorized with default headers values
func NewPlaylistsUpdateUnauthorized() *PlaylistsUpdateUnauthorized {
	return &PlaylistsUpdateUnauthorized{}
}

/* PlaylistsUpdateUnauthorized describes a response with status code 401, with default header values.

You provided invalid credentials.
*/
type PlaylistsUpdateUnauthorized struct {
}

func (o *PlaylistsUpdateUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /playlists/{id}/][%d] playlistsUpdateUnauthorized ", 401)
}

func (o *PlaylistsUpdateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}