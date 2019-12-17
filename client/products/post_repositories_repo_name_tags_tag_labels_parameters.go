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

// NewPostRepositoriesRepoNameTagsTagLabelsParams creates a new PostRepositoriesRepoNameTagsTagLabelsParams object
// with the default values initialized.
func NewPostRepositoriesRepoNameTagsTagLabelsParams() *PostRepositoriesRepoNameTagsTagLabelsParams {
	var ()
	return &PostRepositoriesRepoNameTagsTagLabelsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRepositoriesRepoNameTagsTagLabelsParamsWithTimeout creates a new PostRepositoriesRepoNameTagsTagLabelsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRepositoriesRepoNameTagsTagLabelsParamsWithTimeout(timeout time.Duration) *PostRepositoriesRepoNameTagsTagLabelsParams {
	var ()
	return &PostRepositoriesRepoNameTagsTagLabelsParams{

		timeout: timeout,
	}
}

// NewPostRepositoriesRepoNameTagsTagLabelsParamsWithContext creates a new PostRepositoriesRepoNameTagsTagLabelsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRepositoriesRepoNameTagsTagLabelsParamsWithContext(ctx context.Context) *PostRepositoriesRepoNameTagsTagLabelsParams {
	var ()
	return &PostRepositoriesRepoNameTagsTagLabelsParams{

		Context: ctx,
	}
}

// NewPostRepositoriesRepoNameTagsTagLabelsParamsWithHTTPClient creates a new PostRepositoriesRepoNameTagsTagLabelsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostRepositoriesRepoNameTagsTagLabelsParamsWithHTTPClient(client *http.Client) *PostRepositoriesRepoNameTagsTagLabelsParams {
	var ()
	return &PostRepositoriesRepoNameTagsTagLabelsParams{
		HTTPClient: client,
	}
}

/*PostRepositoriesRepoNameTagsTagLabelsParams contains all the parameters to send to the API endpoint
for the post repositories repo name tags tag labels operation typically these are written to a http.Request
*/
type PostRepositoriesRepoNameTagsTagLabelsParams struct {

	/*Label
	  Only the ID property is required.

	*/
	Label *models.Label
	/*RepoName
	  The name of repository.

	*/
	RepoName string
	/*Tag
	  The tag of the image.

	*/
	Tag string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post repositories repo name tags tag labels params
func (o *PostRepositoriesRepoNameTagsTagLabelsParams) WithTimeout(timeout time.Duration) *PostRepositoriesRepoNameTagsTagLabelsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post repositories repo name tags tag labels params
func (o *PostRepositoriesRepoNameTagsTagLabelsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post repositories repo name tags tag labels params
func (o *PostRepositoriesRepoNameTagsTagLabelsParams) WithContext(ctx context.Context) *PostRepositoriesRepoNameTagsTagLabelsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post repositories repo name tags tag labels params
func (o *PostRepositoriesRepoNameTagsTagLabelsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post repositories repo name tags tag labels params
func (o *PostRepositoriesRepoNameTagsTagLabelsParams) WithHTTPClient(client *http.Client) *PostRepositoriesRepoNameTagsTagLabelsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post repositories repo name tags tag labels params
func (o *PostRepositoriesRepoNameTagsTagLabelsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLabel adds the label to the post repositories repo name tags tag labels params
func (o *PostRepositoriesRepoNameTagsTagLabelsParams) WithLabel(label *models.Label) *PostRepositoriesRepoNameTagsTagLabelsParams {
	o.SetLabel(label)
	return o
}

// SetLabel adds the label to the post repositories repo name tags tag labels params
func (o *PostRepositoriesRepoNameTagsTagLabelsParams) SetLabel(label *models.Label) {
	o.Label = label
}

// WithRepoName adds the repoName to the post repositories repo name tags tag labels params
func (o *PostRepositoriesRepoNameTagsTagLabelsParams) WithRepoName(repoName string) *PostRepositoriesRepoNameTagsTagLabelsParams {
	o.SetRepoName(repoName)
	return o
}

// SetRepoName adds the repoName to the post repositories repo name tags tag labels params
func (o *PostRepositoriesRepoNameTagsTagLabelsParams) SetRepoName(repoName string) {
	o.RepoName = repoName
}

// WithTag adds the tag to the post repositories repo name tags tag labels params
func (o *PostRepositoriesRepoNameTagsTagLabelsParams) WithTag(tag string) *PostRepositoriesRepoNameTagsTagLabelsParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the post repositories repo name tags tag labels params
func (o *PostRepositoriesRepoNameTagsTagLabelsParams) SetTag(tag string) {
	o.Tag = tag
}

// WriteToRequest writes these params to a swagger request
func (o *PostRepositoriesRepoNameTagsTagLabelsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Label != nil {
		if err := r.SetBodyParam(o.Label); err != nil {
			return err
		}
	}

	// path param repo_name
	if err := r.SetPathParam("repo_name", o.RepoName); err != nil {
		return err
	}

	// path param tag
	if err := r.SetPathParam("tag", o.Tag); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
