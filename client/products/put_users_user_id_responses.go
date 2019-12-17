// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PutUsersUserIDReader is a Reader for the PutUsersUserID structure.
type PutUsersUserIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutUsersUserIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutUsersUserIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutUsersUserIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutUsersUserIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutUsersUserIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutUsersUserIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutUsersUserIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutUsersUserIDOK creates a PutUsersUserIDOK with default headers values
func NewPutUsersUserIDOK() *PutUsersUserIDOK {
	return &PutUsersUserIDOK{}
}

/*PutUsersUserIDOK handles this case with default header values.

Updated user's profile successfully.
*/
type PutUsersUserIDOK struct {
}

func (o *PutUsersUserIDOK) Error() string {
	return fmt.Sprintf("[PUT /users/{user_id}][%d] putUsersUserIdOK ", 200)
}

func (o *PutUsersUserIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutUsersUserIDBadRequest creates a PutUsersUserIDBadRequest with default headers values
func NewPutUsersUserIDBadRequest() *PutUsersUserIDBadRequest {
	return &PutUsersUserIDBadRequest{}
}

/*PutUsersUserIDBadRequest handles this case with default header values.

Invalid user ID.
*/
type PutUsersUserIDBadRequest struct {
}

func (o *PutUsersUserIDBadRequest) Error() string {
	return fmt.Sprintf("[PUT /users/{user_id}][%d] putUsersUserIdBadRequest ", 400)
}

func (o *PutUsersUserIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutUsersUserIDUnauthorized creates a PutUsersUserIDUnauthorized with default headers values
func NewPutUsersUserIDUnauthorized() *PutUsersUserIDUnauthorized {
	return &PutUsersUserIDUnauthorized{}
}

/*PutUsersUserIDUnauthorized handles this case with default header values.

User need to log in first.
*/
type PutUsersUserIDUnauthorized struct {
}

func (o *PutUsersUserIDUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /users/{user_id}][%d] putUsersUserIdUnauthorized ", 401)
}

func (o *PutUsersUserIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutUsersUserIDForbidden creates a PutUsersUserIDForbidden with default headers values
func NewPutUsersUserIDForbidden() *PutUsersUserIDForbidden {
	return &PutUsersUserIDForbidden{}
}

/*PutUsersUserIDForbidden handles this case with default header values.

User does not have permission of admin role.
*/
type PutUsersUserIDForbidden struct {
}

func (o *PutUsersUserIDForbidden) Error() string {
	return fmt.Sprintf("[PUT /users/{user_id}][%d] putUsersUserIdForbidden ", 403)
}

func (o *PutUsersUserIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutUsersUserIDNotFound creates a PutUsersUserIDNotFound with default headers values
func NewPutUsersUserIDNotFound() *PutUsersUserIDNotFound {
	return &PutUsersUserIDNotFound{}
}

/*PutUsersUserIDNotFound handles this case with default header values.

User ID does not exist.
*/
type PutUsersUserIDNotFound struct {
}

func (o *PutUsersUserIDNotFound) Error() string {
	return fmt.Sprintf("[PUT /users/{user_id}][%d] putUsersUserIdNotFound ", 404)
}

func (o *PutUsersUserIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutUsersUserIDInternalServerError creates a PutUsersUserIDInternalServerError with default headers values
func NewPutUsersUserIDInternalServerError() *PutUsersUserIDInternalServerError {
	return &PutUsersUserIDInternalServerError{}
}

/*PutUsersUserIDInternalServerError handles this case with default header values.

Unexpected internal errors.
*/
type PutUsersUserIDInternalServerError struct {
}

func (o *PutUsersUserIDInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /users/{user_id}][%d] putUsersUserIdInternalServerError ", 500)
}

func (o *PutUsersUserIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
