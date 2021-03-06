// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PostProjectsProjectIDImmutabletagrulesReader is a Reader for the PostProjectsProjectIDImmutabletagrules structure.
type PostProjectsProjectIDImmutabletagrulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostProjectsProjectIDImmutabletagrulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostProjectsProjectIDImmutabletagrulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostProjectsProjectIDImmutabletagrulesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostProjectsProjectIDImmutabletagrulesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostProjectsProjectIDImmutabletagrulesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostProjectsProjectIDImmutabletagrulesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostProjectsProjectIDImmutabletagrulesOK creates a PostProjectsProjectIDImmutabletagrulesOK with default headers values
func NewPostProjectsProjectIDImmutabletagrulesOK() *PostProjectsProjectIDImmutabletagrulesOK {
	return &PostProjectsProjectIDImmutabletagrulesOK{}
}

/*PostProjectsProjectIDImmutabletagrulesOK handles this case with default header values.

Add the immutable tag rule successfully.
*/
type PostProjectsProjectIDImmutabletagrulesOK struct {
}

func (o *PostProjectsProjectIDImmutabletagrulesOK) Error() string {
	return fmt.Sprintf("[POST /projects/{project_id}/immutabletagrules][%d] postProjectsProjectIdImmutabletagrulesOK ", 200)
}

func (o *PostProjectsProjectIDImmutabletagrulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostProjectsProjectIDImmutabletagrulesBadRequest creates a PostProjectsProjectIDImmutabletagrulesBadRequest with default headers values
func NewPostProjectsProjectIDImmutabletagrulesBadRequest() *PostProjectsProjectIDImmutabletagrulesBadRequest {
	return &PostProjectsProjectIDImmutabletagrulesBadRequest{}
}

/*PostProjectsProjectIDImmutabletagrulesBadRequest handles this case with default header values.

Illegal format of provided ID value.
*/
type PostProjectsProjectIDImmutabletagrulesBadRequest struct {
}

func (o *PostProjectsProjectIDImmutabletagrulesBadRequest) Error() string {
	return fmt.Sprintf("[POST /projects/{project_id}/immutabletagrules][%d] postProjectsProjectIdImmutabletagrulesBadRequest ", 400)
}

func (o *PostProjectsProjectIDImmutabletagrulesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostProjectsProjectIDImmutabletagrulesUnauthorized creates a PostProjectsProjectIDImmutabletagrulesUnauthorized with default headers values
func NewPostProjectsProjectIDImmutabletagrulesUnauthorized() *PostProjectsProjectIDImmutabletagrulesUnauthorized {
	return &PostProjectsProjectIDImmutabletagrulesUnauthorized{}
}

/*PostProjectsProjectIDImmutabletagrulesUnauthorized handles this case with default header values.

User need to log in first.
*/
type PostProjectsProjectIDImmutabletagrulesUnauthorized struct {
}

func (o *PostProjectsProjectIDImmutabletagrulesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /projects/{project_id}/immutabletagrules][%d] postProjectsProjectIdImmutabletagrulesUnauthorized ", 401)
}

func (o *PostProjectsProjectIDImmutabletagrulesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostProjectsProjectIDImmutabletagrulesForbidden creates a PostProjectsProjectIDImmutabletagrulesForbidden with default headers values
func NewPostProjectsProjectIDImmutabletagrulesForbidden() *PostProjectsProjectIDImmutabletagrulesForbidden {
	return &PostProjectsProjectIDImmutabletagrulesForbidden{}
}

/*PostProjectsProjectIDImmutabletagrulesForbidden handles this case with default header values.

User have no permission to get immutable tag rule of the project.
*/
type PostProjectsProjectIDImmutabletagrulesForbidden struct {
}

func (o *PostProjectsProjectIDImmutabletagrulesForbidden) Error() string {
	return fmt.Sprintf("[POST /projects/{project_id}/immutabletagrules][%d] postProjectsProjectIdImmutabletagrulesForbidden ", 403)
}

func (o *PostProjectsProjectIDImmutabletagrulesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostProjectsProjectIDImmutabletagrulesInternalServerError creates a PostProjectsProjectIDImmutabletagrulesInternalServerError with default headers values
func NewPostProjectsProjectIDImmutabletagrulesInternalServerError() *PostProjectsProjectIDImmutabletagrulesInternalServerError {
	return &PostProjectsProjectIDImmutabletagrulesInternalServerError{}
}

/*PostProjectsProjectIDImmutabletagrulesInternalServerError handles this case with default header values.

Internal server errors.
*/
type PostProjectsProjectIDImmutabletagrulesInternalServerError struct {
}

func (o *PostProjectsProjectIDImmutabletagrulesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /projects/{project_id}/immutabletagrules][%d] postProjectsProjectIdImmutabletagrulesInternalServerError ", 500)
}

func (o *PostProjectsProjectIDImmutabletagrulesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
