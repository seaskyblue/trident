// Code generated by go-swagger; DO NOT EDIT.

package networking

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// FcPortCollectionGetReader is a Reader for the FcPortCollectionGet structure.
type FcPortCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FcPortCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFcPortCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFcPortCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFcPortCollectionGetOK creates a FcPortCollectionGetOK with default headers values
func NewFcPortCollectionGetOK() *FcPortCollectionGetOK {
	return &FcPortCollectionGetOK{}
}

/* FcPortCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type FcPortCollectionGetOK struct {
	Payload *models.FcPortResponse
}

func (o *FcPortCollectionGetOK) Error() string {
	return fmt.Sprintf("[GET /network/fc/ports][%d] fcPortCollectionGetOK  %+v", 200, o.Payload)
}
func (o *FcPortCollectionGetOK) GetPayload() *models.FcPortResponse {
	return o.Payload
}

func (o *FcPortCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FcPortResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFcPortCollectionGetDefault creates a FcPortCollectionGetDefault with default headers values
func NewFcPortCollectionGetDefault(code int) *FcPortCollectionGetDefault {
	return &FcPortCollectionGetDefault{
		_statusCode: code,
	}
}

/* FcPortCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type FcPortCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the fc port collection get default response
func (o *FcPortCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *FcPortCollectionGetDefault) Error() string {
	return fmt.Sprintf("[GET /network/fc/ports][%d] fc_port_collection_get default  %+v", o._statusCode, o.Payload)
}
func (o *FcPortCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *FcPortCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}