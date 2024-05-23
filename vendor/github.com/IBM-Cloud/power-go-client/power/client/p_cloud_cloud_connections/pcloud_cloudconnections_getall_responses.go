// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_cloud_connections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudCloudconnectionsGetallReader is a Reader for the PcloudCloudconnectionsGetall structure.
type PcloudCloudconnectionsGetallReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudCloudconnectionsGetallReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudCloudconnectionsGetallOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudCloudconnectionsGetallBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudCloudconnectionsGetallUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudCloudconnectionsGetallForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudCloudconnectionsGetallNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPcloudCloudconnectionsGetallRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudCloudconnectionsGetallInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections] pcloud.cloudconnections.getall", response, response.Code())
	}
}

// NewPcloudCloudconnectionsGetallOK creates a PcloudCloudconnectionsGetallOK with default headers values
func NewPcloudCloudconnectionsGetallOK() *PcloudCloudconnectionsGetallOK {
	return &PcloudCloudconnectionsGetallOK{}
}

/*
PcloudCloudconnectionsGetallOK describes a response with status code 200, with default header values.

OK
*/
type PcloudCloudconnectionsGetallOK struct {
	Payload *models.CloudConnections
}

// IsSuccess returns true when this pcloud cloudconnections getall o k response has a 2xx status code
func (o *PcloudCloudconnectionsGetallOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud cloudconnections getall o k response has a 3xx status code
func (o *PcloudCloudconnectionsGetallOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudconnections getall o k response has a 4xx status code
func (o *PcloudCloudconnectionsGetallOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud cloudconnections getall o k response has a 5xx status code
func (o *PcloudCloudconnectionsGetallOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudconnections getall o k response a status code equal to that given
func (o *PcloudCloudconnectionsGetallOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pcloud cloudconnections getall o k response
func (o *PcloudCloudconnectionsGetallOK) Code() int {
	return 200
}

func (o *PcloudCloudconnectionsGetallOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections][%d] pcloudCloudconnectionsGetallOK  %+v", 200, o.Payload)
}

func (o *PcloudCloudconnectionsGetallOK) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections][%d] pcloudCloudconnectionsGetallOK  %+v", 200, o.Payload)
}

func (o *PcloudCloudconnectionsGetallOK) GetPayload() *models.CloudConnections {
	return o.Payload
}

func (o *PcloudCloudconnectionsGetallOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudConnections)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsGetallBadRequest creates a PcloudCloudconnectionsGetallBadRequest with default headers values
func NewPcloudCloudconnectionsGetallBadRequest() *PcloudCloudconnectionsGetallBadRequest {
	return &PcloudCloudconnectionsGetallBadRequest{}
}

/*
PcloudCloudconnectionsGetallBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudCloudconnectionsGetallBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudconnections getall bad request response has a 2xx status code
func (o *PcloudCloudconnectionsGetallBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudconnections getall bad request response has a 3xx status code
func (o *PcloudCloudconnectionsGetallBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudconnections getall bad request response has a 4xx status code
func (o *PcloudCloudconnectionsGetallBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudconnections getall bad request response has a 5xx status code
func (o *PcloudCloudconnectionsGetallBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudconnections getall bad request response a status code equal to that given
func (o *PcloudCloudconnectionsGetallBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pcloud cloudconnections getall bad request response
func (o *PcloudCloudconnectionsGetallBadRequest) Code() int {
	return 400
}

func (o *PcloudCloudconnectionsGetallBadRequest) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections][%d] pcloudCloudconnectionsGetallBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudCloudconnectionsGetallBadRequest) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections][%d] pcloudCloudconnectionsGetallBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudCloudconnectionsGetallBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsGetallBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsGetallUnauthorized creates a PcloudCloudconnectionsGetallUnauthorized with default headers values
func NewPcloudCloudconnectionsGetallUnauthorized() *PcloudCloudconnectionsGetallUnauthorized {
	return &PcloudCloudconnectionsGetallUnauthorized{}
}

/*
PcloudCloudconnectionsGetallUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudCloudconnectionsGetallUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudconnections getall unauthorized response has a 2xx status code
func (o *PcloudCloudconnectionsGetallUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudconnections getall unauthorized response has a 3xx status code
func (o *PcloudCloudconnectionsGetallUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudconnections getall unauthorized response has a 4xx status code
func (o *PcloudCloudconnectionsGetallUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudconnections getall unauthorized response has a 5xx status code
func (o *PcloudCloudconnectionsGetallUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudconnections getall unauthorized response a status code equal to that given
func (o *PcloudCloudconnectionsGetallUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud cloudconnections getall unauthorized response
func (o *PcloudCloudconnectionsGetallUnauthorized) Code() int {
	return 401
}

func (o *PcloudCloudconnectionsGetallUnauthorized) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections][%d] pcloudCloudconnectionsGetallUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudCloudconnectionsGetallUnauthorized) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections][%d] pcloudCloudconnectionsGetallUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudCloudconnectionsGetallUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsGetallUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsGetallForbidden creates a PcloudCloudconnectionsGetallForbidden with default headers values
func NewPcloudCloudconnectionsGetallForbidden() *PcloudCloudconnectionsGetallForbidden {
	return &PcloudCloudconnectionsGetallForbidden{}
}

/*
PcloudCloudconnectionsGetallForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudCloudconnectionsGetallForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudconnections getall forbidden response has a 2xx status code
func (o *PcloudCloudconnectionsGetallForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudconnections getall forbidden response has a 3xx status code
func (o *PcloudCloudconnectionsGetallForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudconnections getall forbidden response has a 4xx status code
func (o *PcloudCloudconnectionsGetallForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudconnections getall forbidden response has a 5xx status code
func (o *PcloudCloudconnectionsGetallForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudconnections getall forbidden response a status code equal to that given
func (o *PcloudCloudconnectionsGetallForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pcloud cloudconnections getall forbidden response
func (o *PcloudCloudconnectionsGetallForbidden) Code() int {
	return 403
}

func (o *PcloudCloudconnectionsGetallForbidden) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections][%d] pcloudCloudconnectionsGetallForbidden  %+v", 403, o.Payload)
}

func (o *PcloudCloudconnectionsGetallForbidden) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections][%d] pcloudCloudconnectionsGetallForbidden  %+v", 403, o.Payload)
}

func (o *PcloudCloudconnectionsGetallForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsGetallForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsGetallNotFound creates a PcloudCloudconnectionsGetallNotFound with default headers values
func NewPcloudCloudconnectionsGetallNotFound() *PcloudCloudconnectionsGetallNotFound {
	return &PcloudCloudconnectionsGetallNotFound{}
}

/*
PcloudCloudconnectionsGetallNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudCloudconnectionsGetallNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudconnections getall not found response has a 2xx status code
func (o *PcloudCloudconnectionsGetallNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudconnections getall not found response has a 3xx status code
func (o *PcloudCloudconnectionsGetallNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudconnections getall not found response has a 4xx status code
func (o *PcloudCloudconnectionsGetallNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudconnections getall not found response has a 5xx status code
func (o *PcloudCloudconnectionsGetallNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudconnections getall not found response a status code equal to that given
func (o *PcloudCloudconnectionsGetallNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the pcloud cloudconnections getall not found response
func (o *PcloudCloudconnectionsGetallNotFound) Code() int {
	return 404
}

func (o *PcloudCloudconnectionsGetallNotFound) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections][%d] pcloudCloudconnectionsGetallNotFound  %+v", 404, o.Payload)
}

func (o *PcloudCloudconnectionsGetallNotFound) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections][%d] pcloudCloudconnectionsGetallNotFound  %+v", 404, o.Payload)
}

func (o *PcloudCloudconnectionsGetallNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsGetallNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsGetallRequestTimeout creates a PcloudCloudconnectionsGetallRequestTimeout with default headers values
func NewPcloudCloudconnectionsGetallRequestTimeout() *PcloudCloudconnectionsGetallRequestTimeout {
	return &PcloudCloudconnectionsGetallRequestTimeout{}
}

/*
PcloudCloudconnectionsGetallRequestTimeout describes a response with status code 408, with default header values.

Request Timeout
*/
type PcloudCloudconnectionsGetallRequestTimeout struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudconnections getall request timeout response has a 2xx status code
func (o *PcloudCloudconnectionsGetallRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudconnections getall request timeout response has a 3xx status code
func (o *PcloudCloudconnectionsGetallRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudconnections getall request timeout response has a 4xx status code
func (o *PcloudCloudconnectionsGetallRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudconnections getall request timeout response has a 5xx status code
func (o *PcloudCloudconnectionsGetallRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudconnections getall request timeout response a status code equal to that given
func (o *PcloudCloudconnectionsGetallRequestTimeout) IsCode(code int) bool {
	return code == 408
}

// Code gets the status code for the pcloud cloudconnections getall request timeout response
func (o *PcloudCloudconnectionsGetallRequestTimeout) Code() int {
	return 408
}

func (o *PcloudCloudconnectionsGetallRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections][%d] pcloudCloudconnectionsGetallRequestTimeout  %+v", 408, o.Payload)
}

func (o *PcloudCloudconnectionsGetallRequestTimeout) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections][%d] pcloudCloudconnectionsGetallRequestTimeout  %+v", 408, o.Payload)
}

func (o *PcloudCloudconnectionsGetallRequestTimeout) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsGetallRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsGetallInternalServerError creates a PcloudCloudconnectionsGetallInternalServerError with default headers values
func NewPcloudCloudconnectionsGetallInternalServerError() *PcloudCloudconnectionsGetallInternalServerError {
	return &PcloudCloudconnectionsGetallInternalServerError{}
}

/*
PcloudCloudconnectionsGetallInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudCloudconnectionsGetallInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudconnections getall internal server error response has a 2xx status code
func (o *PcloudCloudconnectionsGetallInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudconnections getall internal server error response has a 3xx status code
func (o *PcloudCloudconnectionsGetallInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudconnections getall internal server error response has a 4xx status code
func (o *PcloudCloudconnectionsGetallInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud cloudconnections getall internal server error response has a 5xx status code
func (o *PcloudCloudconnectionsGetallInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud cloudconnections getall internal server error response a status code equal to that given
func (o *PcloudCloudconnectionsGetallInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud cloudconnections getall internal server error response
func (o *PcloudCloudconnectionsGetallInternalServerError) Code() int {
	return 500
}

func (o *PcloudCloudconnectionsGetallInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections][%d] pcloudCloudconnectionsGetallInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudCloudconnectionsGetallInternalServerError) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections][%d] pcloudCloudconnectionsGetallInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudCloudconnectionsGetallInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsGetallInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
