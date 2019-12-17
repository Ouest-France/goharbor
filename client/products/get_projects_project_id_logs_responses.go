// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/Ouest-France/goharbor/models"
)

// GetProjectsProjectIDLogsReader is a Reader for the GetProjectsProjectIDLogs structure.
type GetProjectsProjectIDLogsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProjectsProjectIDLogsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProjectsProjectIDLogsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetProjectsProjectIDLogsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetProjectsProjectIDLogsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetProjectsProjectIDLogsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetProjectsProjectIDLogsOK creates a GetProjectsProjectIDLogsOK with default headers values
func NewGetProjectsProjectIDLogsOK() *GetProjectsProjectIDLogsOK {
	return &GetProjectsProjectIDLogsOK{}
}

/*GetProjectsProjectIDLogsOK handles this case with default header values.

Get access log successfully.
*/
type GetProjectsProjectIDLogsOK struct {
	/*Link refers to the previous page and next page
	 */
	Link string
	/*The total count of access logs
	 */
	XTotalCount int64

	Payload []*models.AccessLog
}

func (o *GetProjectsProjectIDLogsOK) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/logs][%d] getProjectsProjectIdLogsOK  %+v", 200, o.Payload)
}

func (o *GetProjectsProjectIDLogsOK) GetPayload() []*models.AccessLog {
	return o.Payload
}

func (o *GetProjectsProjectIDLogsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Link
	o.Link = response.GetHeader("Link")

	// response header X-Total-Count
	xTotalCount, err := swag.ConvertInt64(response.GetHeader("X-Total-Count"))
	if err != nil {
		return errors.InvalidType("X-Total-Count", "header", "int64", response.GetHeader("X-Total-Count"))
	}
	o.XTotalCount = xTotalCount

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectsProjectIDLogsBadRequest creates a GetProjectsProjectIDLogsBadRequest with default headers values
func NewGetProjectsProjectIDLogsBadRequest() *GetProjectsProjectIDLogsBadRequest {
	return &GetProjectsProjectIDLogsBadRequest{}
}

/*GetProjectsProjectIDLogsBadRequest handles this case with default header values.

Illegal format of provided ID value.
*/
type GetProjectsProjectIDLogsBadRequest struct {
}

func (o *GetProjectsProjectIDLogsBadRequest) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/logs][%d] getProjectsProjectIdLogsBadRequest ", 400)
}

func (o *GetProjectsProjectIDLogsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProjectsProjectIDLogsUnauthorized creates a GetProjectsProjectIDLogsUnauthorized with default headers values
func NewGetProjectsProjectIDLogsUnauthorized() *GetProjectsProjectIDLogsUnauthorized {
	return &GetProjectsProjectIDLogsUnauthorized{}
}

/*GetProjectsProjectIDLogsUnauthorized handles this case with default header values.

User need to log in first.
*/
type GetProjectsProjectIDLogsUnauthorized struct {
}

func (o *GetProjectsProjectIDLogsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/logs][%d] getProjectsProjectIdLogsUnauthorized ", 401)
}

func (o *GetProjectsProjectIDLogsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProjectsProjectIDLogsInternalServerError creates a GetProjectsProjectIDLogsInternalServerError with default headers values
func NewGetProjectsProjectIDLogsInternalServerError() *GetProjectsProjectIDLogsInternalServerError {
	return &GetProjectsProjectIDLogsInternalServerError{}
}

/*GetProjectsProjectIDLogsInternalServerError handles this case with default header values.

Unexpected internal errors.
*/
type GetProjectsProjectIDLogsInternalServerError struct {
}

func (o *GetProjectsProjectIDLogsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/logs][%d] getProjectsProjectIdLogsInternalServerError ", 500)
}

func (o *GetProjectsProjectIDLogsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
