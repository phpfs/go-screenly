// Code generated by go-swagger; DO NOT EDIT.

package assets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/phpfs/go-screenly/screenly/models"
)

// AssetsReadReader is a Reader for the AssetsRead structure.
type AssetsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AssetsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAssetsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAssetsReadUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAssetsReadOK creates a AssetsReadOK with default headers values
func NewAssetsReadOK() *AssetsReadOK {
	return &AssetsReadOK{}
}

/* AssetsReadOK describes a response with status code 200, with default header values.

AssetsReadOK assets read o k
*/
type AssetsReadOK struct {
	Payload *models.Asset
}

func (o *AssetsReadOK) Error() string {
	return fmt.Sprintf("[GET /assets/{id}/][%d] assetsReadOK  %+v", 200, o.Payload)
}
func (o *AssetsReadOK) GetPayload() *models.Asset {
	return o.Payload
}

func (o *AssetsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Asset)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAssetsReadUnauthorized creates a AssetsReadUnauthorized with default headers values
func NewAssetsReadUnauthorized() *AssetsReadUnauthorized {
	return &AssetsReadUnauthorized{}
}

/* AssetsReadUnauthorized describes a response with status code 401, with default header values.

You provided invalid credentials.
*/
type AssetsReadUnauthorized struct {
}

func (o *AssetsReadUnauthorized) Error() string {
	return fmt.Sprintf("[GET /assets/{id}/][%d] assetsReadUnauthorized ", 401)
}

func (o *AssetsReadUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
