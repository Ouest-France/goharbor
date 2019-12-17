// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/Ouest-France/goharbor/models"
)

// GetReplicationExecutionsIDTasksReader is a Reader for the GetReplicationExecutionsIDTasks structure.
type GetReplicationExecutionsIDTasksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReplicationExecutionsIDTasksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetReplicationExecutionsIDTasksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetReplicationExecutionsIDTasksBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetReplicationExecutionsIDTasksUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetReplicationExecutionsIDTasksForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetReplicationExecutionsIDTasksNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetReplicationExecutionsIDTasksInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetReplicationExecutionsIDTasksOK creates a GetReplicationExecutionsIDTasksOK with default headers values
func NewGetReplicationExecutionsIDTasksOK() *GetReplicationExecutionsIDTasksOK {
	return &GetReplicationExecutionsIDTasksOK{}
}

/*GetReplicationExecutionsIDTasksOK handles this case with default header values.

Success.
*/
type GetReplicationExecutionsIDTasksOK struct {
	Payload []*models.ReplicationTask
}

func (o *GetReplicationExecutionsIDTasksOK) Error() string {
	return fmt.Sprintf("[GET /replication/executions/{id}/tasks][%d] getReplicationExecutionsIdTasksOK  %+v", 200, o.Payload)
}

func (o *GetReplicationExecutionsIDTasksOK) GetPayload() []*models.ReplicationTask {
	return o.Payload
}

func (o *GetReplicationExecutionsIDTasksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReplicationExecutionsIDTasksBadRequest creates a GetReplicationExecutionsIDTasksBadRequest with default headers values
func NewGetReplicationExecutionsIDTasksBadRequest() *GetReplicationExecutionsIDTasksBadRequest {
	return &GetReplicationExecutionsIDTasksBadRequest{}
}

/*GetReplicationExecutionsIDTasksBadRequest handles this case with default header values.

Bad request.
*/
type GetReplicationExecutionsIDTasksBadRequest struct {
}

func (o *GetReplicationExecutionsIDTasksBadRequest) Error() string {
	return fmt.Sprintf("[GET /replication/executions/{id}/tasks][%d] getReplicationExecutionsIdTasksBadRequest ", 400)
}

func (o *GetReplicationExecutionsIDTasksBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetReplicationExecutionsIDTasksUnauthorized creates a GetReplicationExecutionsIDTasksUnauthorized with default headers values
func NewGetReplicationExecutionsIDTasksUnauthorized() *GetReplicationExecutionsIDTasksUnauthorized {
	return &GetReplicationExecutionsIDTasksUnauthorized{}
}

/*GetReplicationExecutionsIDTasksUnauthorized handles this case with default header values.

User need to login first.
*/
type GetReplicationExecutionsIDTasksUnauthorized struct {
}

func (o *GetReplicationExecutionsIDTasksUnauthorized) Error() string {
	return fmt.Sprintf("[GET /replication/executions/{id}/tasks][%d] getReplicationExecutionsIdTasksUnauthorized ", 401)
}

func (o *GetReplicationExecutionsIDTasksUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetReplicationExecutionsIDTasksForbidden creates a GetReplicationExecutionsIDTasksForbidden with default headers values
func NewGetReplicationExecutionsIDTasksForbidden() *GetReplicationExecutionsIDTasksForbidden {
	return &GetReplicationExecutionsIDTasksForbidden{}
}

/*GetReplicationExecutionsIDTasksForbidden handles this case with default header values.

User has no privilege for the operation.
*/
type GetReplicationExecutionsIDTasksForbidden struct {
}

func (o *GetReplicationExecutionsIDTasksForbidden) Error() string {
	return fmt.Sprintf("[GET /replication/executions/{id}/tasks][%d] getReplicationExecutionsIdTasksForbidden ", 403)
}

func (o *GetReplicationExecutionsIDTasksForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetReplicationExecutionsIDTasksNotFound creates a GetReplicationExecutionsIDTasksNotFound with default headers values
func NewGetReplicationExecutionsIDTasksNotFound() *GetReplicationExecutionsIDTasksNotFound {
	return &GetReplicationExecutionsIDTasksNotFound{}
}

/*GetReplicationExecutionsIDTasksNotFound handles this case with default header values.

Resource requested does not exist.
*/
type GetReplicationExecutionsIDTasksNotFound struct {
}

func (o *GetReplicationExecutionsIDTasksNotFound) Error() string {
	return fmt.Sprintf("[GET /replication/executions/{id}/tasks][%d] getReplicationExecutionsIdTasksNotFound ", 404)
}

func (o *GetReplicationExecutionsIDTasksNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetReplicationExecutionsIDTasksInternalServerError creates a GetReplicationExecutionsIDTasksInternalServerError with default headers values
func NewGetReplicationExecutionsIDTasksInternalServerError() *GetReplicationExecutionsIDTasksInternalServerError {
	return &GetReplicationExecutionsIDTasksInternalServerError{}
}

/*GetReplicationExecutionsIDTasksInternalServerError handles this case with default header values.

Unexpected internal errors.
*/
type GetReplicationExecutionsIDTasksInternalServerError struct {
}

func (o *GetReplicationExecutionsIDTasksInternalServerError) Error() string {
	return fmt.Sprintf("[GET /replication/executions/{id}/tasks][%d] getReplicationExecutionsIdTasksInternalServerError ", 500)
}

func (o *GetReplicationExecutionsIDTasksInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
