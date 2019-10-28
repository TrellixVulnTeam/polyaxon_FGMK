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

package k8_s_secrets_v1

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

// NewGetK8SSecretParams creates a new GetK8SSecretParams object
// with the default values initialized.
func NewGetK8SSecretParams() *GetK8SSecretParams {
	var ()
	return &GetK8SSecretParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetK8SSecretParamsWithTimeout creates a new GetK8SSecretParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetK8SSecretParamsWithTimeout(timeout time.Duration) *GetK8SSecretParams {
	var ()
	return &GetK8SSecretParams{

		timeout: timeout,
	}
}

// NewGetK8SSecretParamsWithContext creates a new GetK8SSecretParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetK8SSecretParamsWithContext(ctx context.Context) *GetK8SSecretParams {
	var ()
	return &GetK8SSecretParams{

		Context: ctx,
	}
}

// NewGetK8SSecretParamsWithHTTPClient creates a new GetK8SSecretParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetK8SSecretParamsWithHTTPClient(client *http.Client) *GetK8SSecretParams {
	var ()
	return &GetK8SSecretParams{
		HTTPClient: client,
	}
}

/*GetK8SSecretParams contains all the parameters to send to the API endpoint
for the get k8 s secret operation typically these are written to a http.Request
*/
type GetK8SSecretParams struct {

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

// WithTimeout adds the timeout to the get k8 s secret params
func (o *GetK8SSecretParams) WithTimeout(timeout time.Duration) *GetK8SSecretParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get k8 s secret params
func (o *GetK8SSecretParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get k8 s secret params
func (o *GetK8SSecretParams) WithContext(ctx context.Context) *GetK8SSecretParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get k8 s secret params
func (o *GetK8SSecretParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get k8 s secret params
func (o *GetK8SSecretParams) WithHTTPClient(client *http.Client) *GetK8SSecretParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get k8 s secret params
func (o *GetK8SSecretParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOwner adds the owner to the get k8 s secret params
func (o *GetK8SSecretParams) WithOwner(owner string) *GetK8SSecretParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the get k8 s secret params
func (o *GetK8SSecretParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithUUID adds the uuid to the get k8 s secret params
func (o *GetK8SSecretParams) WithUUID(uuid string) *GetK8SSecretParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the get k8 s secret params
func (o *GetK8SSecretParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *GetK8SSecretParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
