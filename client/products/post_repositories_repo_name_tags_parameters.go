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

// NewPostRepositoriesRepoNameTagsParams creates a new PostRepositoriesRepoNameTagsParams object
// with the default values initialized.
func NewPostRepositoriesRepoNameTagsParams() *PostRepositoriesRepoNameTagsParams {
	var ()
	return &PostRepositoriesRepoNameTagsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRepositoriesRepoNameTagsParamsWithTimeout creates a new PostRepositoriesRepoNameTagsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRepositoriesRepoNameTagsParamsWithTimeout(timeout time.Duration) *PostRepositoriesRepoNameTagsParams {
	var ()
	return &PostRepositoriesRepoNameTagsParams{

		timeout: timeout,
	}
}

// NewPostRepositoriesRepoNameTagsParamsWithContext creates a new PostRepositoriesRepoNameTagsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRepositoriesRepoNameTagsParamsWithContext(ctx context.Context) *PostRepositoriesRepoNameTagsParams {
	var ()
	return &PostRepositoriesRepoNameTagsParams{

		Context: ctx,
	}
}

// NewPostRepositoriesRepoNameTagsParamsWithHTTPClient creates a new PostRepositoriesRepoNameTagsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostRepositoriesRepoNameTagsParamsWithHTTPClient(client *http.Client) *PostRepositoriesRepoNameTagsParams {
	var ()
	return &PostRepositoriesRepoNameTagsParams{
		HTTPClient: client,
	}
}

/*PostRepositoriesRepoNameTagsParams contains all the parameters to send to the API endpoint
for the post repositories repo name tags operation typically these are written to a http.Request
*/
type PostRepositoriesRepoNameTagsParams struct {

	/*RepoName
	  Relevant repository name.

	*/
	RepoName string
	/*Request
	  Request to give source image and target tag.

	*/
	Request *models.RetagReq

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post repositories repo name tags params
func (o *PostRepositoriesRepoNameTagsParams) WithTimeout(timeout time.Duration) *PostRepositoriesRepoNameTagsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post repositories repo name tags params
func (o *PostRepositoriesRepoNameTagsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post repositories repo name tags params
func (o *PostRepositoriesRepoNameTagsParams) WithContext(ctx context.Context) *PostRepositoriesRepoNameTagsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post repositories repo name tags params
func (o *PostRepositoriesRepoNameTagsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post repositories repo name tags params
func (o *PostRepositoriesRepoNameTagsParams) WithHTTPClient(client *http.Client) *PostRepositoriesRepoNameTagsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post repositories repo name tags params
func (o *PostRepositoriesRepoNameTagsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRepoName adds the repoName to the post repositories repo name tags params
func (o *PostRepositoriesRepoNameTagsParams) WithRepoName(repoName string) *PostRepositoriesRepoNameTagsParams {
	o.SetRepoName(repoName)
	return o
}

// SetRepoName adds the repoName to the post repositories repo name tags params
func (o *PostRepositoriesRepoNameTagsParams) SetRepoName(repoName string) {
	o.RepoName = repoName
}

// WithRequest adds the request to the post repositories repo name tags params
func (o *PostRepositoriesRepoNameTagsParams) WithRequest(request *models.RetagReq) *PostRepositoriesRepoNameTagsParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the post repositories repo name tags params
func (o *PostRepositoriesRepoNameTagsParams) SetRequest(request *models.RetagReq) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *PostRepositoriesRepoNameTagsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param repo_name
	if err := r.SetPathParam("repo_name", o.RepoName); err != nil {
		return err
	}

	if o.Request != nil {
		if err := r.SetBodyParam(o.Request); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
