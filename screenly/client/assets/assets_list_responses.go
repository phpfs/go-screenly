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

// AssetsListReader is a Reader for the AssetsList structure.
type AssetsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AssetsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAssetsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAssetsListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAssetsListOK creates a AssetsListOK with default headers values
func NewAssetsListOK() *AssetsListOK {
	return &AssetsListOK{}
}

/* AssetsListOK describes a response with status code 200, with default header values.

AssetsListOK assets list o k
*/
type AssetsListOK struct {
	Payload *models.Asset
}

func (o *AssetsListOK) Error() string {
	return fmt.Sprintf("[GET /assets/][%d] assetsListOK  %+v", 200, o.Payload)
}
func (o *AssetsListOK) GetPayload() *models.Asset {
	return o.Payload
}

func (o *AssetsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Asset)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAssetsListUnauthorized creates a AssetsListUnauthorized with default headers values
func NewAssetsListUnauthorized() *AssetsListUnauthorized {
	return &AssetsListUnauthorized{}
}

/* AssetsListUnauthorized describes a response with status code 401, with default header values.

You provided invalid credentials.
*/
type AssetsListUnauthorized struct {
}

func (o *AssetsListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /assets/][%d] assetsListUnauthorized ", 401)
}

func (o *AssetsListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}