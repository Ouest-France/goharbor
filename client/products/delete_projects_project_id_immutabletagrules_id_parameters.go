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

// NewDeleteProjectsProjectIDImmutabletagrulesIDParams creates a new DeleteProjectsProjectIDImmutabletagrulesIDParams object
// with the default values initialized.
func NewDeleteProjectsProjectIDImmutabletagrulesIDParams() *DeleteProjectsProjectIDImmutabletagrulesIDParams {
	var ()
	return &DeleteProjectsProjectIDImmutabletagrulesIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteProjectsProjectIDImmutabletagrulesIDParamsWithTimeout creates a new DeleteProjectsProjectIDImmutabletagrulesIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteProjectsProjectIDImmutabletagrulesIDParamsWithTimeout(timeout time.Duration) *DeleteProjectsProjectIDImmutabletagrulesIDParams {
	var ()
	return &DeleteProjectsProjectIDImmutabletagrulesIDParams{

		timeout: timeout,
	}
}

// NewDeleteProjectsProjectIDImmutabletagrulesIDParamsWithContext creates a new DeleteProjectsProjectIDImmutabletagrulesIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteProjectsProjectIDImmutabletagrulesIDParamsWithContext(ctx context.Context) *DeleteProjectsProjectIDImmutabletagrulesIDParams {
	var ()
	return &DeleteProjectsProjectIDImmutabletagrulesIDParams{

		Context: ctx,
	}
}

// NewDeleteProjectsProjectIDImmutabletagrulesIDParamsWithHTTPClient creates a new DeleteProjectsProjectIDImmutabletagrulesIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteProjectsProjectIDImmutabletagrulesIDParamsWithHTTPClient(client *http.Client) *DeleteProjectsProjectIDImmutabletagrulesIDParams {
	var ()
	return &DeleteProjectsProjectIDImmutabletagrulesIDParams{
		HTTPClient: client,
	}
}

/*DeleteProjectsProjectIDImmutabletagrulesIDParams contains all the parameters to send to the API endpoint
for the delete projects project ID immutabletagrules ID operation typically these are written to a http.Request
*/
type DeleteProjectsProjectIDImmutabletagrulesIDParams struct {

	/*ID
	  Immutable tag rule ID.

	*/
	ID int64
	/*ProjectID
	  Relevant project ID.

	*/
	ProjectID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete projects project ID immutabletagrules ID params
func (o *DeleteProjectsProjectIDImmutabletagrulesIDParams) WithTimeout(timeout time.Duration) *DeleteProjectsProjectIDImmutabletagrulesIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete projects project ID immutabletagrules ID params
func (o *DeleteProjectsProjectIDImmutabletagrulesIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete projects project ID immutabletagrules ID params
func (o *DeleteProjectsProjectIDImmutabletagrulesIDParams) WithContext(ctx context.Context) *DeleteProjectsProjectIDImmutabletagrulesIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete projects project ID immutabletagrules ID params
func (o *DeleteProjectsProjectIDImmutabletagrulesIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete projects project ID immutabletagrules ID params
func (o *DeleteProjectsProjectIDImmutabletagrulesIDParams) WithHTTPClient(client *http.Client) *DeleteProjectsProjectIDImmutabletagrulesIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete projects project ID immutabletagrules ID params
func (o *DeleteProjectsProjectIDImmutabletagrulesIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete projects project ID immutabletagrules ID params
func (o *DeleteProjectsProjectIDImmutabletagrulesIDParams) WithID(id int64) *DeleteProjectsProjectIDImmutabletagrulesIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete projects project ID immutabletagrules ID params
func (o *DeleteProjectsProjectIDImmutabletagrulesIDParams) SetID(id int64) {
	o.ID = id
}

// WithProjectID adds the projectID to the delete projects project ID immutabletagrules ID params
func (o *DeleteProjectsProjectIDImmutabletagrulesIDParams) WithProjectID(projectID int64) *DeleteProjectsProjectIDImmutabletagrulesIDParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the delete projects project ID immutabletagrules ID params
func (o *DeleteProjectsProjectIDImmutabletagrulesIDParams) SetProjectID(projectID int64) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteProjectsProjectIDImmutabletagrulesIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	// path param project_id
	if err := r.SetPathParam("project_id", swag.FormatInt64(o.ProjectID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
