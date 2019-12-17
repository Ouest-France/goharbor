// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PostRepositoriesRepoNameLabelsReader is a Reader for the PostRepositoriesRepoNameLabels structure.
type PostRepositoriesRepoNameLabelsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRepositoriesRepoNameLabelsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostRepositoriesRepoNameLabelsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPostRepositoriesRepoNameLabelsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostRepositoriesRepoNameLabelsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostRepositoriesRepoNameLabelsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRepositoriesRepoNameLabelsOK creates a PostRepositoriesRepoNameLabelsOK with default headers values
func NewPostRepositoriesRepoNameLabelsOK() *PostRepositoriesRepoNameLabelsOK {
	return &PostRepositoriesRepoNameLabelsOK{}
}

/*PostRepositoriesRepoNameLabelsOK handles this case with default header values.

Successfully.
*/
type PostRepositoriesRepoNameLabelsOK struct {
}

func (o *PostRepositoriesRepoNameLabelsOK) Error() string {
	return fmt.Sprintf("[POST /repositories/{repo_name}/labels][%d] postRepositoriesRepoNameLabelsOK ", 200)
}

func (o *PostRepositoriesRepoNameLabelsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostRepositoriesRepoNameLabelsUnauthorized creates a PostRepositoriesRepoNameLabelsUnauthorized with default headers values
func NewPostRepositoriesRepoNameLabelsUnauthorized() *PostRepositoriesRepoNameLabelsUnauthorized {
	return &PostRepositoriesRepoNameLabelsUnauthorized{}
}

/*PostRepositoriesRepoNameLabelsUnauthorized handles this case with default header values.

Unauthorized.
*/
type PostRepositoriesRepoNameLabelsUnauthorized struct {
}

func (o *PostRepositoriesRepoNameLabelsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /repositories/{repo_name}/labels][%d] postRepositoriesRepoNameLabelsUnauthorized ", 401)
}

func (o *PostRepositoriesRepoNameLabelsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostRepositoriesRepoNameLabelsForbidden creates a PostRepositoriesRepoNameLabelsForbidden with default headers values
func NewPostRepositoriesRepoNameLabelsForbidden() *PostRepositoriesRepoNameLabelsForbidden {
	return &PostRepositoriesRepoNameLabelsForbidden{}
}

/*PostRepositoriesRepoNameLabelsForbidden handles this case with default header values.

Forbidden. User should have write permisson for the repository to perform the action.
*/
type PostRepositoriesRepoNameLabelsForbidden struct {
}

func (o *PostRepositoriesRepoNameLabelsForbidden) Error() string {
	return fmt.Sprintf("[POST /repositories/{repo_name}/labels][%d] postRepositoriesRepoNameLabelsForbidden ", 403)
}

func (o *PostRepositoriesRepoNameLabelsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostRepositoriesRepoNameLabelsNotFound creates a PostRepositoriesRepoNameLabelsNotFound with default headers values
func NewPostRepositoriesRepoNameLabelsNotFound() *PostRepositoriesRepoNameLabelsNotFound {
	return &PostRepositoriesRepoNameLabelsNotFound{}
}

/*PostRepositoriesRepoNameLabelsNotFound handles this case with default header values.

Resource not found.
*/
type PostRepositoriesRepoNameLabelsNotFound struct {
}

func (o *PostRepositoriesRepoNameLabelsNotFound) Error() string {
	return fmt.Sprintf("[POST /repositories/{repo_name}/labels][%d] postRepositoriesRepoNameLabelsNotFound ", 404)
}

func (o *PostRepositoriesRepoNameLabelsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
