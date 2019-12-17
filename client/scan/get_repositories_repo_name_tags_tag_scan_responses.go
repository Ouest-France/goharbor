// Code generated by go-swagger; DO NOT EDIT.

package scan

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/Ouest-France/goharbor/models"
)

// GetRepositoriesRepoNameTagsTagScanReader is a Reader for the GetRepositoriesRepoNameTagsTagScan structure.
type GetRepositoriesRepoNameTagsTagScanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRepositoriesRepoNameTagsTagScanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRepositoriesRepoNameTagsTagScanOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetRepositoriesRepoNameTagsTagScanUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRepositoriesRepoNameTagsTagScanForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRepositoriesRepoNameTagsTagScanNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRepositoriesRepoNameTagsTagScanInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRepositoriesRepoNameTagsTagScanOK creates a GetRepositoriesRepoNameTagsTagScanOK with default headers values
func NewGetRepositoriesRepoNameTagsTagScanOK() *GetRepositoriesRepoNameTagsTagScanOK {
	return &GetRepositoriesRepoNameTagsTagScanOK{}
}

/*GetRepositoriesRepoNameTagsTagScanOK handles this case with default header values.

The report details of the specified artifact identified by the repo_name and tag.
*/
type GetRepositoriesRepoNameTagsTagScanOK struct {
	Payload *models.Report
}

func (o *GetRepositoriesRepoNameTagsTagScanOK) Error() string {
	return fmt.Sprintf("[GET /repositories/{repo_name}/tags/{tag}/scan][%d] getRepositoriesRepoNameTagsTagScanOK  %+v", 200, o.Payload)
}

func (o *GetRepositoriesRepoNameTagsTagScanOK) GetPayload() *models.Report {
	return o.Payload
}

func (o *GetRepositoriesRepoNameTagsTagScanOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Report)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRepositoriesRepoNameTagsTagScanUnauthorized creates a GetRepositoriesRepoNameTagsTagScanUnauthorized with default headers values
func NewGetRepositoriesRepoNameTagsTagScanUnauthorized() *GetRepositoriesRepoNameTagsTagScanUnauthorized {
	return &GetRepositoriesRepoNameTagsTagScanUnauthorized{}
}

/*GetRepositoriesRepoNameTagsTagScanUnauthorized handles this case with default header values.

Unauthorized request
*/
type GetRepositoriesRepoNameTagsTagScanUnauthorized struct {
}

func (o *GetRepositoriesRepoNameTagsTagScanUnauthorized) Error() string {
	return fmt.Sprintf("[GET /repositories/{repo_name}/tags/{tag}/scan][%d] getRepositoriesRepoNameTagsTagScanUnauthorized ", 401)
}

func (o *GetRepositoriesRepoNameTagsTagScanUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRepositoriesRepoNameTagsTagScanForbidden creates a GetRepositoriesRepoNameTagsTagScanForbidden with default headers values
func NewGetRepositoriesRepoNameTagsTagScanForbidden() *GetRepositoriesRepoNameTagsTagScanForbidden {
	return &GetRepositoriesRepoNameTagsTagScanForbidden{}
}

/*GetRepositoriesRepoNameTagsTagScanForbidden handles this case with default header values.

Request is not allowed
*/
type GetRepositoriesRepoNameTagsTagScanForbidden struct {
}

func (o *GetRepositoriesRepoNameTagsTagScanForbidden) Error() string {
	return fmt.Sprintf("[GET /repositories/{repo_name}/tags/{tag}/scan][%d] getRepositoriesRepoNameTagsTagScanForbidden ", 403)
}

func (o *GetRepositoriesRepoNameTagsTagScanForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRepositoriesRepoNameTagsTagScanNotFound creates a GetRepositoriesRepoNameTagsTagScanNotFound with default headers values
func NewGetRepositoriesRepoNameTagsTagScanNotFound() *GetRepositoriesRepoNameTagsTagScanNotFound {
	return &GetRepositoriesRepoNameTagsTagScanNotFound{}
}

/*GetRepositoriesRepoNameTagsTagScanNotFound handles this case with default header values.

The target artifact is not found
*/
type GetRepositoriesRepoNameTagsTagScanNotFound struct {
}

func (o *GetRepositoriesRepoNameTagsTagScanNotFound) Error() string {
	return fmt.Sprintf("[GET /repositories/{repo_name}/tags/{tag}/scan][%d] getRepositoriesRepoNameTagsTagScanNotFound ", 404)
}

func (o *GetRepositoriesRepoNameTagsTagScanNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRepositoriesRepoNameTagsTagScanInternalServerError creates a GetRepositoriesRepoNameTagsTagScanInternalServerError with default headers values
func NewGetRepositoriesRepoNameTagsTagScanInternalServerError() *GetRepositoriesRepoNameTagsTagScanInternalServerError {
	return &GetRepositoriesRepoNameTagsTagScanInternalServerError{}
}

/*GetRepositoriesRepoNameTagsTagScanInternalServerError handles this case with default header values.

Internal server error happened
*/
type GetRepositoriesRepoNameTagsTagScanInternalServerError struct {
}

func (o *GetRepositoriesRepoNameTagsTagScanInternalServerError) Error() string {
	return fmt.Sprintf("[GET /repositories/{repo_name}/tags/{tag}/scan][%d] getRepositoriesRepoNameTagsTagScanInternalServerError ", 500)
}

func (o *GetRepositoriesRepoNameTagsTagScanInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
