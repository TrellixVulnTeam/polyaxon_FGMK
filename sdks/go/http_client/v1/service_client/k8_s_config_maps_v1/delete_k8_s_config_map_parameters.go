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

package k8_s_config_maps_v1

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

// NewDeleteK8SConfigMapParams creates a new DeleteK8SConfigMapParams object
// with the default values initialized.
func NewDeleteK8SConfigMapParams() *DeleteK8SConfigMapParams {
	var ()
	return &DeleteK8SConfigMapParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteK8SConfigMapParamsWithTimeout creates a new DeleteK8SConfigMapParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteK8SConfigMapParamsWithTimeout(timeout time.Duration) *DeleteK8SConfigMapParams {
	var ()
	return &DeleteK8SConfigMapParams{

		timeout: timeout,
	}
}

// NewDeleteK8SConfigMapParamsWithContext creates a new DeleteK8SConfigMapParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteK8SConfigMapParamsWithContext(ctx context.Context) *DeleteK8SConfigMapParams {
	var ()
	return &DeleteK8SConfigMapParams{

		Context: ctx,
	}
}

// NewDeleteK8SConfigMapParamsWithHTTPClient creates a new DeleteK8SConfigMapParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteK8SConfigMapParamsWithHTTPClient(client *http.Client) *DeleteK8SConfigMapParams {
	var ()
	return &DeleteK8SConfigMapParams{
		HTTPClient: client,
	}
}

/*DeleteK8SConfigMapParams contains all the parameters to send to the API endpoint
for the delete k8 s config map operation typically these are written to a http.Request
*/
type DeleteK8SConfigMapParams struct {

	/*Owner
	  Owner of the namespace

	*/
	Owner string
	/*UUID
	  Unique integer identifier of the entity

	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete k8 s config map params
func (o *DeleteK8SConfigMapParams) WithTimeout(timeout time.Duration) *DeleteK8SConfigMapParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete k8 s config map params
func (o *DeleteK8SConfigMapParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete k8 s config map params
func (o *DeleteK8SConfigMapParams) WithContext(ctx context.Context) *DeleteK8SConfigMapParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete k8 s config map params
func (o *DeleteK8SConfigMapParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete k8 s config map params
func (o *DeleteK8SConfigMapParams) WithHTTPClient(client *http.Client) *DeleteK8SConfigMapParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete k8 s config map params
func (o *DeleteK8SConfigMapParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOwner adds the owner to the delete k8 s config map params
func (o *DeleteK8SConfigMapParams) WithOwner(owner string) *DeleteK8SConfigMapParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the delete k8 s config map params
func (o *DeleteK8SConfigMapParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithUUID adds the uuid to the delete k8 s config map params
func (o *DeleteK8SConfigMapParams) WithUUID(uuid string) *DeleteK8SConfigMapParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the delete k8 s config map params
func (o *DeleteK8SConfigMapParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteK8SConfigMapParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
