// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetUsersCurrentParams creates a new GetUsersCurrentParams object
// with the default values initialized.
func NewGetUsersCurrentParams() *GetUsersCurrentParams {

	return &GetUsersCurrentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUsersCurrentParamsWithTimeout creates a new GetUsersCurrentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUsersCurrentParamsWithTimeout(timeout time.Duration) *GetUsersCurrentParams {

	return &GetUsersCurrentParams{

		timeout: timeout,
	}
}

// NewGetUsersCurrentParamsWithContext creates a new GetUsersCurrentParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUsersCurrentParamsWithContext(ctx context.Context) *GetUsersCurrentParams {

	return &GetUsersCurrentParams{

		Context: ctx,
	}
}

// NewGetUsersCurrentParamsWithHTTPClient creates a new GetUsersCurrentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUsersCurrentParamsWithHTTPClient(client *http.Client) *GetUsersCurrentParams {

	return &GetUsersCurrentParams{
		HTTPClient: client,
	}
}

/*GetUsersCurrentParams contains all the parameters to send to the API endpoint
for the get users current operation typically these are written to a http.Request
*/
type GetUsersCurrentParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get users current params
func (o *GetUsersCurrentParams) WithTimeout(timeout time.Duration) *GetUsersCurrentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get users current params
func (o *GetUsersCurrentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get users current params
func (o *GetUsersCurrentParams) WithContext(ctx context.Context) *GetUsersCurrentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get users current params
func (o *GetUsersCurrentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get users current params
func (o *GetUsersCurrentParams) WithHTTPClient(client *http.Client) *GetUsersCurrentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get users current params
func (o *GetUsersCurrentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetUsersCurrentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
