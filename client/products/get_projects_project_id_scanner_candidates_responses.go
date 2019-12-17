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

// GetProjectsProjectIDScannerCandidatesReader is a Reader for the GetProjectsProjectIDScannerCandidates structure.
type GetProjectsProjectIDScannerCandidatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProjectsProjectIDScannerCandidatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProjectsProjectIDScannerCandidatesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetProjectsProjectIDScannerCandidatesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetProjectsProjectIDScannerCandidatesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetProjectsProjectIDScannerCandidatesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetProjectsProjectIDScannerCandidatesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetProjectsProjectIDScannerCandidatesOK creates a GetProjectsProjectIDScannerCandidatesOK with default headers values
func NewGetProjectsProjectIDScannerCandidatesOK() *GetProjectsProjectIDScannerCandidatesOK {
	return &GetProjectsProjectIDScannerCandidatesOK{}
}

/*GetProjectsProjectIDScannerCandidatesOK handles this case with default header values.

A list of scanner registrations.
*/
type GetProjectsProjectIDScannerCandidatesOK struct {
	Payload []*models.ScannerRegistration
}

func (o *GetProjectsProjectIDScannerCandidatesOK) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/scanner/candidates][%d] getProjectsProjectIdScannerCandidatesOK  %+v", 200, o.Payload)
}

func (o *GetProjectsProjectIDScannerCandidatesOK) GetPayload() []*models.ScannerRegistration {
	return o.Payload
}

func (o *GetProjectsProjectIDScannerCandidatesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectsProjectIDScannerCandidatesBadRequest creates a GetProjectsProjectIDScannerCandidatesBadRequest with default headers values
func NewGetProjectsProjectIDScannerCandidatesBadRequest() *GetProjectsProjectIDScannerCandidatesBadRequest {
	return &GetProjectsProjectIDScannerCandidatesBadRequest{}
}

/*GetProjectsProjectIDScannerCandidatesBadRequest handles this case with default header values.

Bad project ID or query parameters
*/
type GetProjectsProjectIDScannerCandidatesBadRequest struct {
}

func (o *GetProjectsProjectIDScannerCandidatesBadRequest) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/scanner/candidates][%d] getProjectsProjectIdScannerCandidatesBadRequest ", 400)
}

func (o *GetProjectsProjectIDScannerCandidatesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProjectsProjectIDScannerCandidatesUnauthorized creates a GetProjectsProjectIDScannerCandidatesUnauthorized with default headers values
func NewGetProjectsProjectIDScannerCandidatesUnauthorized() *GetProjectsProjectIDScannerCandidatesUnauthorized {
	return &GetProjectsProjectIDScannerCandidatesUnauthorized{}
}

/*GetProjectsProjectIDScannerCandidatesUnauthorized handles this case with default header values.

Unauthorized request
*/
type GetProjectsProjectIDScannerCandidatesUnauthorized struct {
}

func (o *GetProjectsProjectIDScannerCandidatesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/scanner/candidates][%d] getProjectsProjectIdScannerCandidatesUnauthorized ", 401)
}

func (o *GetProjectsProjectIDScannerCandidatesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProjectsProjectIDScannerCandidatesForbidden creates a GetProjectsProjectIDScannerCandidatesForbidden with default headers values
func NewGetProjectsProjectIDScannerCandidatesForbidden() *GetProjectsProjectIDScannerCandidatesForbidden {
	return &GetProjectsProjectIDScannerCandidatesForbidden{}
}

/*GetProjectsProjectIDScannerCandidatesForbidden handles this case with default header values.

Request is not allowed
*/
type GetProjectsProjectIDScannerCandidatesForbidden struct {
}

func (o *GetProjectsProjectIDScannerCandidatesForbidden) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/scanner/candidates][%d] getProjectsProjectIdScannerCandidatesForbidden ", 403)
}

func (o *GetProjectsProjectIDScannerCandidatesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProjectsProjectIDScannerCandidatesInternalServerError creates a GetProjectsProjectIDScannerCandidatesInternalServerError with default headers values
func NewGetProjectsProjectIDScannerCandidatesInternalServerError() *GetProjectsProjectIDScannerCandidatesInternalServerError {
	return &GetProjectsProjectIDScannerCandidatesInternalServerError{}
}

/*GetProjectsProjectIDScannerCandidatesInternalServerError handles this case with default header values.

Internal server error happened
*/
type GetProjectsProjectIDScannerCandidatesInternalServerError struct {
}

func (o *GetProjectsProjectIDScannerCandidatesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/scanner/candidates][%d] getProjectsProjectIdScannerCandidatesInternalServerError ", 500)
}

func (o *GetProjectsProjectIDScannerCandidatesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
