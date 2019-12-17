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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetReplicationExecutionsIDParams creates a new GetReplicationExecutionsIDParams object
// with the default values initialized.
func NewGetReplicationExecutionsIDParams() *GetReplicationExecutionsIDParams {
	var ()
	return &GetReplicationExecutionsIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetReplicationExecutionsIDParamsWithTimeout creates a new GetReplicationExecutionsIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetReplicationExecutionsIDParamsWithTimeout(timeout time.Duration) *GetReplicationExecutionsIDParams {
	var ()
	return &GetReplicationExecutionsIDParams{

		timeout: timeout,
	}
}

// NewGetReplicationExecutionsIDParamsWithContext creates a new GetReplicationExecutionsIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetReplicationExecutionsIDParamsWithContext(ctx context.Context) *GetReplicationExecutionsIDParams {
	var ()
	return &GetReplicationExecutionsIDParams{

		Context: ctx,
	}
}

// NewGetReplicationExecutionsIDParamsWithHTTPClient creates a new GetReplicationExecutionsIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetReplicationExecutionsIDParamsWithHTTPClient(client *http.Client) *GetReplicationExecutionsIDParams {
	var ()
	return &GetReplicationExecutionsIDParams{
		HTTPClient: client,
	}
}

/*GetReplicationExecutionsIDParams contains all the parameters to send to the API endpoint
for the get replication executions ID operation typically these are written to a http.Request
*/
type GetReplicationExecutionsIDParams struct {

	/*ID
	  The execution ID.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get replication executions ID params
func (o *GetReplicationExecutionsIDParams) WithTimeout(timeout time.Duration) *GetReplicationExecutionsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get replication executions ID params
func (o *GetReplicationExecutionsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get replication executions ID params
func (o *GetReplicationExecutionsIDParams) WithContext(ctx context.Context) *GetReplicationExecutionsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get replication executions ID params
func (o *GetReplicationExecutionsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get replication executions ID params
func (o *GetReplicationExecutionsIDParams) WithHTTPClient(client *http.Client) *GetReplicationExecutionsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get replication executions ID params
func (o *GetReplicationExecutionsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get replication executions ID params
func (o *GetReplicationExecutionsIDParams) WithID(id int64) *GetReplicationExecutionsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get replication executions ID params
func (o *GetReplicationExecutionsIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetReplicationExecutionsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
