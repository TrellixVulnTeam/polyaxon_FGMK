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

// NewPatchSearchParams creates a new PatchSearchParams object
// with the default values initialized.
func NewPatchSearchParams() *PatchSearchParams {
	var ()
	return &PatchSearchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchSearchParamsWithTimeout creates a new PatchSearchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchSearchParamsWithTimeout(timeout time.Duration) *PatchSearchParams {
	var ()
	return &PatchSearchParams{

		timeout: timeout,
	}
}

// NewPatchSearchParamsWithContext creates a new PatchSearchParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchSearchParamsWithContext(ctx context.Context) *PatchSearchParams {
	var ()
	return &PatchSearchParams{

		Context: ctx,
	}
}

// NewPatchSearchParamsWithHTTPClient creates a new PatchSearchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchSearchParamsWithHTTPClient(client *http.Client) *PatchSearchParams {
	var ()
	return &PatchSearchParams{
		HTTPClient: client,
	}
}

/*PatchSearchParams contains all the parameters to send to the API endpoint
for the patch search operation typically these are written to a http.Request
*/
type PatchSearchParams struct {

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
	/*SearchUUID
	  UUID

	*/
	SearchUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch search params
func (o *PatchSearchParams) WithTimeout(timeout time.Duration) *PatchSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch search params
func (o *PatchSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch search params
func (o *PatchSearchParams) WithContext(ctx context.Context) *PatchSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch search params
func (o *PatchSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch search params
func (o *PatchSearchParams) WithHTTPClient(client *http.Client) *PatchSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch search params
func (o *PatchSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch search params
func (o *PatchSearchParams) WithBody(body *service_model.V1Search) *PatchSearchParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch search params
func (o *PatchSearchParams) SetBody(body *service_model.V1Search) {
	o.Body = body
}

// WithOwner adds the owner to the patch search params
func (o *PatchSearchParams) WithOwner(owner string) *PatchSearchParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the patch search params
func (o *PatchSearchParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithProject adds the project to the patch search params
func (o *PatchSearchParams) WithProject(project string) *PatchSearchParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the patch search params
func (o *PatchSearchParams) SetProject(project string) {
	o.Project = project
}

// WithSearchUUID adds the searchUUID to the patch search params
func (o *PatchSearchParams) WithSearchUUID(searchUUID string) *PatchSearchParams {
	o.SetSearchUUID(searchUUID)
	return o
}

// SetSearchUUID adds the searchUuid to the patch search params
func (o *PatchSearchParams) SetSearchUUID(searchUUID string) {
	o.SearchUUID = searchUUID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param search.uuid
	if err := r.SetPathParam("search.uuid", o.SearchUUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
