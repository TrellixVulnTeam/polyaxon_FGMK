// Copyright 2019 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by go-swagger; DO NOT EDIT.

package search_v1

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

	service_model "github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// NewCreateSearchParams creates a new CreateSearchParams object
// with the default values initialized.
func NewCreateSearchParams() *CreateSearchParams {
	var ()
	return &CreateSearchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateSearchParamsWithTimeout creates a new CreateSearchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateSearchParamsWithTimeout(timeout time.Duration) *CreateSearchParams {
	var ()
	return &CreateSearchParams{

		timeout: timeout,
	}
}

// NewCreateSearchParamsWithContext creates a new CreateSearchParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateSearchParamsWithContext(ctx context.Context) *CreateSearchParams {
	var ()
	return &CreateSearchParams{

		Context: ctx,
	}
}

// NewCreateSearchParamsWithHTTPClient creates a new CreateSearchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateSearchParamsWithHTTPClient(client *http.Client) *CreateSearchParams {
	var ()
	return &CreateSearchParams{
		HTTPClient: client,
	}
}

/*CreateSearchParams contains all the parameters to send to the API endpoint
for the create search operation typically these are written to a http.Request
*/
type CreateSearchParams struct {

	/*Body
	  Artifact store body

	*/
	Body *service_model.V1Search
	/*Owner
	  Owner of the namespace

	*/
	Owner string
	/*Project
	  Project under namesapce

	*/
	Project string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create search params
func (o *CreateSearchParams) WithTimeout(timeout time.Duration) *CreateSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create search params
func (o *CreateSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create search params
func (o *CreateSearchParams) WithContext(ctx context.Context) *CreateSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create search params
func (o *CreateSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create search params
func (o *CreateSearchParams) WithHTTPClient(client *http.Client) *CreateSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create search params
func (o *CreateSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create search params
func (o *CreateSearchParams) WithBody(body *service_model.V1Search) *CreateSearchParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create search params
func (o *CreateSearchParams) SetBody(body *service_model.V1Search) {
	o.Body = body
}

// WithOwner adds the owner to the create search params
func (o *CreateSearchParams) WithOwner(owner string) *CreateSearchParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the create search params
func (o *CreateSearchParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithProject adds the project to the create search params
func (o *CreateSearchParams) WithProject(project string) *CreateSearchParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the create search params
func (o *CreateSearchParams) SetProject(project string) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *CreateSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	// path param project
	if err := r.SetPathParam("project", o.Project); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
