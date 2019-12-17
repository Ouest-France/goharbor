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

// NewGetRepositoriesRepoNameLabelsParams creates a new GetRepositoriesRepoNameLabelsParams object
// with the default values initialized.
func NewGetRepositoriesRepoNameLabelsParams() *GetRepositoriesRepoNameLabelsParams {
	var ()
	return &GetRepositoriesRepoNameLabelsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRepositoriesRepoNameLabelsParamsWithTimeout creates a new GetRepositoriesRepoNameLabelsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRepositoriesRepoNameLabelsParamsWithTimeout(timeout time.Duration) *GetRepositoriesRepoNameLabelsParams {
	var ()
	return &GetRepositoriesRepoNameLabelsParams{

		timeout: timeout,
	}
}

// NewGetRepositoriesRepoNameLabelsParamsWithContext creates a new GetRepositoriesRepoNameLabelsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRepositoriesRepoNameLabelsParamsWithContext(ctx context.Context) *GetRepositoriesRepoNameLabelsParams {
	var ()
	return &GetRepositoriesRepoNameLabelsParams{

		Context: ctx,
	}
}

// NewGetRepositoriesRepoNameLabelsParamsWithHTTPClient creates a new GetRepositoriesRepoNameLabelsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRepositoriesRepoNameLabelsParamsWithHTTPClient(client *http.Client) *GetRepositoriesRepoNameLabelsParams {
	var ()
	return &GetRepositoriesRepoNameLabelsParams{
		HTTPClient: client,
	}
}

/*GetRepositoriesRepoNameLabelsParams contains all the parameters to send to the API endpoint
for the get repositories repo name labels operation typically these are written to a http.Request
*/
type GetRepositoriesRepoNameLabelsParams struct {

	/*RepoName
	  The name of repository.

	*/
	RepoName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get repositories repo name labels params
func (o *GetRepositoriesRepoNameLabelsParams) WithTimeout(timeout time.Duration) *GetRepositoriesRepoNameLabelsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get repositories repo name labels params
func (o *GetRepositoriesRepoNameLabelsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get repositories repo name labels params
func (o *GetRepositoriesRepoNameLabelsParams) WithContext(ctx context.Context) *GetRepositoriesRepoNameLabelsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get repositories repo name labels params
func (o *GetRepositoriesRepoNameLabelsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get repositories repo name labels params
func (o *GetRepositoriesRepoNameLabelsParams) WithHTTPClient(client *http.Client) *GetRepositoriesRepoNameLabelsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get repositories repo name labels params
func (o *GetRepositoriesRepoNameLabelsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRepoName adds the repoName to the get repositories repo name labels params
func (o *GetRepositoriesRepoNameLabelsParams) WithRepoName(repoName string) *GetRepositoriesRepoNameLabelsParams {
	o.SetRepoName(repoName)
	return o
}

// SetRepoName adds the repoName to the get repositories repo name labels params
func (o *GetRepositoriesRepoNameLabelsParams) SetRepoName(repoName string) {
	o.RepoName = repoName
}

// WriteToRequest writes these params to a swagger request
func (o *GetRepositoriesRepoNameLabelsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param repo_name
	if err := r.SetPathParam("repo_name", o.RepoName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
