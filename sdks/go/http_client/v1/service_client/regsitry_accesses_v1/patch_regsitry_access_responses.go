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

// PatchRegsitryAccessReader is a Reader for the PatchRegsitryAccess structure.
type PatchRegsitryAccessReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchRegsitryAccessReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchRegsitryAccessOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewPatchRegsitryAccessForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchRegsitryAccessNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchRegsitryAccessOK creates a PatchRegsitryAccessOK with default headers values
func NewPatchRegsitryAccessOK() *PatchRegsitryAccessOK {
	return &PatchRegsitryAccessOK{}
}

/*PatchRegsitryAccessOK handles this case with default header values.

A successful response.
*/
type PatchRegsitryAccessOK struct {
	Payload *service_model.V1HostAccess
}

func (o *PatchRegsitryAccessOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/{owner}/registry_accesses/{host_access.uuid}][%d] patchRegsitryAccessOK  %+v", 200, o.Payload)
}

func (o *PatchRegsitryAccessOK) GetPayload() *service_model.V1HostAccess {
	return o.Payload
}

func (o *PatchRegsitryAccessOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1HostAccess)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRegsitryAccessForbidden creates a PatchRegsitryAccessForbidden with default headers values
func NewPatchRegsitryAccessForbidden() *PatchRegsitryAccessForbidden {
	return &PatchRegsitryAccessForbidden{}
}

/*PatchRegsitryAccessForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type PatchRegsitryAccessForbidden struct {
	Payload interface{}
}

func (o *PatchRegsitryAccessForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/{owner}/registry_accesses/{host_access.uuid}][%d] patchRegsitryAccessForbidden  %+v", 403, o.Payload)
}

func (o *PatchRegsitryAccessForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *PatchRegsitryAccessForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRegsitryAccessNotFound creates a PatchRegsitryAccessNotFound with default headers values
func NewPatchRegsitryAccessNotFound() *PatchRegsitryAccessNotFound {
	return &PatchRegsitryAccessNotFound{}
}

/*PatchRegsitryAccessNotFound handles this case with default header values.

Resource does not exist.
*/
type PatchRegsitryAccessNotFound struct {
	Payload string
}

func (o *PatchRegsitryAccessNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/{owner}/registry_accesses/{host_access.uuid}][%d] patchRegsitryAccessNotFound  %+v", 404, o.Payload)
}

func (o *PatchRegsitryAccessNotFound) GetPayload() string {
	return o.Payload
}

func (o *PatchRegsitryAccessNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
