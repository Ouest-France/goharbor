// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteRepositoriesRepoNameReader is a Reader for the DeleteRepositoriesRepoName structure.
type DeleteRepositoriesRepoNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRepositoriesRepoNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteRepositoriesRepoNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteRepositoriesRepoNameBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteRepositoriesRepoNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteRepositoriesRepoNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteRepositoriesRepoNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 412:
		result := NewDeleteRepositoriesRepoNamePreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteRepositoriesRepoNameOK creates a DeleteRepositoriesRepoNameOK with default headers values
func NewDeleteRepositoriesRepoNameOK() *DeleteRepositoriesRepoNameOK {
	return &DeleteRepositoriesRepoNameOK{}
}

/*DeleteRepositoriesRepoNameOK handles this case with default header values.

Delete successfully.
*/
type DeleteRepositoriesRepoNameOK struct {
}

func (o *DeleteRepositoriesRepoNameOK) Error() string {
	return fmt.Sprintf("[DELETE /repositories/{repo_name}][%d] deleteRepositoriesRepoNameOK ", 200)
}

func (o *DeleteRepositoriesRepoNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRepositoriesRepoNameBadRequest creates a DeleteRepositoriesRepoNameBadRequest with default headers values
func NewDeleteRepositoriesRepoNameBadRequest() *DeleteRepositoriesRepoNameBadRequest {
	return &DeleteRepositoriesRepoNameBadRequest{}
}

/*DeleteRepositoriesRepoNameBadRequest handles this case with default header values.

Invalid repo_name.
*/
type DeleteRepositoriesRepoNameBadRequest struct {
}

func (o *DeleteRepositoriesRepoNameBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /repositories/{repo_name}][%d] deleteRepositoriesRepoNameBadRequest ", 400)
}

func (o *DeleteRepositoriesRepoNameBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRepositoriesRepoNameUnauthorized creates a DeleteRepositoriesRepoNameUnauthorized with default headers values
func NewDeleteRepositoriesRepoNameUnauthorized() *DeleteRepositoriesRepoNameUnauthorized {
	return &DeleteRepositoriesRepoNameUnauthorized{}
}

/*DeleteRepositoriesRepoNameUnauthorized handles this case with default header values.

Unauthorized.
*/
type DeleteRepositoriesRepoNameUnauthorized struct {
}

func (o *DeleteRepositoriesRepoNameUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /repositories/{repo_name}][%d] deleteRepositoriesRepoNameUnauthorized ", 401)
}

func (o *DeleteRepositoriesRepoNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRepositoriesRepoNameForbidden creates a DeleteRepositoriesRepoNameForbidden with default headers values
func NewDeleteRepositoriesRepoNameForbidden() *DeleteRepositoriesRepoNameForbidden {
	return &DeleteRepositoriesRepoNameForbidden{}
}

/*DeleteRepositoriesRepoNameForbidden handles this case with default header values.

Forbidden.
*/
type DeleteRepositoriesRepoNameForbidden struct {
}

func (o *DeleteRepositoriesRepoNameForbidden) Error() string {
	return fmt.Sprintf("[DELETE /repositories/{repo_name}][%d] deleteRepositoriesRepoNameForbidden ", 403)
}

func (o *DeleteRepositoriesRepoNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRepositoriesRepoNameNotFound creates a DeleteRepositoriesRepoNameNotFound with default headers values
func NewDeleteRepositoriesRepoNameNotFound() *DeleteRepositoriesRepoNameNotFound {
	return &DeleteRepositoriesRepoNameNotFound{}
}

/*DeleteRepositoriesRepoNameNotFound handles this case with default header values.

Repository not found.
*/
type DeleteRepositoriesRepoNameNotFound struct {
}

func (o *DeleteRepositoriesRepoNameNotFound) Error() string {
	return fmt.Sprintf("[DELETE /repositories/{repo_name}][%d] deleteRepositoriesRepoNameNotFound ", 404)
}

func (o *DeleteRepositoriesRepoNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRepositoriesRepoNamePreconditionFailed creates a DeleteRepositoriesRepoNamePreconditionFailed with default headers values
func NewDeleteRepositoriesRepoNamePreconditionFailed() *DeleteRepositoriesRepoNamePreconditionFailed {
	return &DeleteRepositoriesRepoNamePreconditionFailed{}
}

/*DeleteRepositoriesRepoNamePreconditionFailed handles this case with default header values.

Precondition Failed.
*/
type DeleteRepositoriesRepoNamePreconditionFailed struct {
}

func (o *DeleteRepositoriesRepoNamePreconditionFailed) Error() string {
	return fmt.Sprintf("[DELETE /repositories/{repo_name}][%d] deleteRepositoriesRepoNamePreconditionFailed ", 412)
}

func (o *DeleteRepositoriesRepoNamePreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}