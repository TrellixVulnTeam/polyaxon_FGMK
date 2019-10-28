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

package regsitry_accesses_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	service_model "github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// UpdateRegsitryAccessReader is a Reader for the UpdateRegsitryAccess structure.
type UpdateRegsitryAccessReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRegsitryAccessReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateRegsitryAccessOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateRegsitryAccessForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateRegsitryAccessNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateRegsitryAccessOK creates a UpdateRegsitryAccessOK with default headers values
func NewUpdateRegsitryAccessOK() *UpdateRegsitryAccessOK {
	return &UpdateRegsitryAccessOK{}
}

/*UpdateRegsitryAccessOK handles this case with default header values.

A successful response.
*/
type UpdateRegsitryAccessOK struct {
	Payload *service_model.V1HostAccess
}

func (o *UpdateRegsitryAccessOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/{owner}/registry_accesses/{host_access.uuid}][%d] updateRegsitryAccessOK  %+v", 200, o.Payload)
}

func (o *UpdateRegsitryAccessOK) GetPayload() *service_model.V1HostAccess {
	return o.Payload
}

func (o *UpdateRegsitryAccessOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1HostAccess)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRegsitryAccessForbidden creates a UpdateRegsitryAccessForbidden with default headers values
func NewUpdateRegsitryAccessForbidden() *UpdateRegsitryAccessForbidden {
	return &UpdateRegsitryAccessForbidden{}
}

/*UpdateRegsitryAccessForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type UpdateRegsitryAccessForbidden struct {
	Payload interface{}
}

func (o *UpdateRegsitryAccessForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v1/{owner}/registry_accesses/{host_access.uuid}][%d] updateRegsitryAccessForbidden  %+v", 403, o.Payload)
}

func (o *UpdateRegsitryAccessForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateRegsitryAccessForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRegsitryAccessNotFound creates a UpdateRegsitryAccessNotFound with default headers values
func NewUpdateRegsitryAccessNotFound() *UpdateRegsitryAccessNotFound {
	return &UpdateRegsitryAccessNotFound{}
}

/*UpdateRegsitryAccessNotFound handles this case with default header values.

Resource does not exist.
*/
type UpdateRegsitryAccessNotFound struct {
	Payload string
}

func (o *UpdateRegsitryAccessNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v1/{owner}/registry_accesses/{host_access.uuid}][%d] updateRegsitryAccessNotFound  %+v", 404, o.Payload)
}

func (o *UpdateRegsitryAccessNotFound) GetPayload() string {
	return o.Payload
}

func (o *UpdateRegsitryAccessNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
