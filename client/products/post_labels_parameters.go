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

	models "github.com/Ouest-France/goharbor/models"
)

// NewPostLabelsParams creates a new PostLabelsParams object
// with the default values initialized.
func NewPostLabelsParams() *PostLabelsParams {
	var ()
	return &PostLabelsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLabelsParamsWithTimeout creates a new PostLabelsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLabelsParamsWithTimeout(timeout time.Duration) *PostLabelsParams {
	var ()
	return &PostLabelsParams{

		timeout: timeout,
	}
}

// NewPostLabelsParamsWithContext creates a new PostLabelsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLabelsParamsWithContext(ctx context.Context) *PostLabelsParams {
	var ()
	return &PostLabelsParams{

		Context: ctx,
	}
}

// NewPostLabelsParamsWithHTTPClient creates a new PostLabelsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLabelsParamsWithHTTPClient(client *http.Client) *PostLabelsParams {
	var ()
	return &PostLabelsParams{
		HTTPClient: client,
	}
}

/*PostLabelsParams contains all the parameters to send to the API endpoint
for the post labels operation typically these are written to a http.Request
*/
type PostLabelsParams struct {

	/*Label
	  The json object of label.

	*/
	Label *models.Label

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post labels params
func (o *PostLabelsParams) WithTimeout(timeout time.Duration) *PostLabelsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post labels params
func (o *PostLabelsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post labels params
func (o *PostLabelsParams) WithContext(ctx context.Context) *PostLabelsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post labels params
func (o *PostLabelsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post labels params
func (o *PostLabelsParams) WithHTTPClient(client *http.Client) *PostLabelsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post labels params
func (o *PostLabelsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLabel adds the label to the post labels params
func (o *PostLabelsParams) WithLabel(label *models.Label) *PostLabelsParams {
	o.SetLabel(label)
	return o
}

// SetLabel adds the label to the post labels params
func (o *PostLabelsParams) SetLabel(label *models.Label) {
	o.Label = label
}

// WriteToRequest writes these params to a swagger request
func (o *PostLabelsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Label != nil {
		if err := r.SetBodyParam(o.Label); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}