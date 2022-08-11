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

// PlaylistsPartialUpdateReader is a Reader for the PlaylistsPartialUpdate structure.
type PlaylistsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PlaylistsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPlaylistsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPlaylistsPartialUpdateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPlaylistsPartialUpdateOK creates a PlaylistsPartialUpdateOK with default headers values
func NewPlaylistsPartialUpdateOK() *PlaylistsPartialUpdateOK {
	return &PlaylistsPartialUpdateOK{}
}

/* PlaylistsPartialUpdateOK describes a response with status code 200, with default header values.

PlaylistsPartialUpdateOK playlists partial update o k
*/
type PlaylistsPartialUpdateOK struct {
	Payload *models.PlaylistRead
}

func (o *PlaylistsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /playlists/{id}/][%d] playlistsPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *PlaylistsPartialUpdateOK) GetPayload() *models.PlaylistRead {
	return o.Payload
}

func (o *PlaylistsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PlaylistRead)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPlaylistsPartialUpdateUnauthorized creates a PlaylistsPartialUpdateUnauthorized with default headers values
func NewPlaylistsPartialUpdateUnauthorized() *PlaylistsPartialUpdateUnauthorized {
	return &PlaylistsPartialUpdateUnauthorized{}
}

/* PlaylistsPartialUpdateUnauthorized describes a response with status code 401, with default header values.

You provided invalid credentials.
*/
type PlaylistsPartialUpdateUnauthorized struct {
}

func (o *PlaylistsPartialUpdateUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /playlists/{id}/][%d] playlistsPartialUpdateUnauthorized ", 401)
}

func (o *PlaylistsPartialUpdateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}