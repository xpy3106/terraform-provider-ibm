// Code generated by go-swagger; DO NOT EDIT.

package compute

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

// NewGetImagesIDTagsParams creates a new GetImagesIDTagsParams object
// with the default values initialized.
func NewGetImagesIDTagsParams() *GetImagesIDTagsParams {
	var ()
	return &GetImagesIDTagsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetImagesIDTagsParamsWithTimeout creates a new GetImagesIDTagsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetImagesIDTagsParamsWithTimeout(timeout time.Duration) *GetImagesIDTagsParams {
	var ()
	return &GetImagesIDTagsParams{

		timeout: timeout,
	}
}

// NewGetImagesIDTagsParamsWithContext creates a new GetImagesIDTagsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetImagesIDTagsParamsWithContext(ctx context.Context) *GetImagesIDTagsParams {
	var ()
	return &GetImagesIDTagsParams{

		Context: ctx,
	}
}

// NewGetImagesIDTagsParamsWithHTTPClient creates a new GetImagesIDTagsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetImagesIDTagsParamsWithHTTPClient(client *http.Client) *GetImagesIDTagsParams {
	var ()
	return &GetImagesIDTagsParams{
		HTTPClient: client,
	}
}

/*GetImagesIDTagsParams contains all the parameters to send to the API endpoint
for the get images ID tags operation typically these are written to a http.Request
*/
type GetImagesIDTagsParams struct {

	/*ID
	  The image id

	*/
	ID string
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get images ID tags params
func (o *GetImagesIDTagsParams) WithTimeout(timeout time.Duration) *GetImagesIDTagsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get images ID tags params
func (o *GetImagesIDTagsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get images ID tags params
func (o *GetImagesIDTagsParams) WithContext(ctx context.Context) *GetImagesIDTagsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get images ID tags params
func (o *GetImagesIDTagsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get images ID tags params
func (o *GetImagesIDTagsParams) WithHTTPClient(client *http.Client) *GetImagesIDTagsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get images ID tags params
func (o *GetImagesIDTagsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get images ID tags params
func (o *GetImagesIDTagsParams) WithID(id string) *GetImagesIDTagsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get images ID tags params
func (o *GetImagesIDTagsParams) SetID(id string) {
	o.ID = id
}

// WithVersion adds the version to the get images ID tags params
func (o *GetImagesIDTagsParams) WithVersion(version string) *GetImagesIDTagsParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get images ID tags params
func (o *GetImagesIDTagsParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *GetImagesIDTagsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// query param version
	qrVersion := o.Version
	qVersion := qrVersion
	if qVersion != "" {
		if err := r.SetQueryParam("version", qVersion); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}