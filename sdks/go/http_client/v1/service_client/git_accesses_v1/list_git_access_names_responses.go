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

package git_accesses_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	service_model "github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// ListGitAccessNamesReader is a Reader for the ListGitAccessNames structure.
type ListGitAccessNamesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListGitAccessNamesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListGitAccessNamesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewListGitAccessNamesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListGitAccessNamesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListGitAccessNamesOK creates a ListGitAccessNamesOK with default headers values
func NewListGitAccessNamesOK() *ListGitAccessNamesOK {
	return &ListGitAccessNamesOK{}
}

/*ListGitAccessNamesOK handles this case with default header values.

A successful response.
*/
type ListGitAccessNamesOK struct {
	Payload *service_model.V1ListHostAccessesResponse
}

func (o *ListGitAccessNamesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/git_accesses/names][%d] listGitAccessNamesOK  %+v", 200, o.Payload)
}

func (o *ListGitAccessNamesOK) GetPayload() *service_model.V1ListHostAccessesResponse {
	return o.Payload
}

func (o *ListGitAccessNamesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ListHostAccessesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListGitAccessNamesForbidden creates a ListGitAccessNamesForbidden with default headers values
func NewListGitAccessNamesForbidden() *ListGitAccessNamesForbidden {
	return &ListGitAccessNamesForbidden{}
}

/*ListGitAccessNamesForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type ListGitAccessNamesForbidden struct {
	Payload interface{}
}

func (o *ListGitAccessNamesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/git_accesses/names][%d] listGitAccessNamesForbidden  %+v", 403, o.Payload)
}

func (o *ListGitAccessNamesForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ListGitAccessNamesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListGitAccessNamesNotFound creates a ListGitAccessNamesNotFound with default headers values
func NewListGitAccessNamesNotFound() *ListGitAccessNamesNotFound {
	return &ListGitAccessNamesNotFound{}
}

/*ListGitAccessNamesNotFound handles this case with default header values.

Resource does not exist.
*/
type ListGitAccessNamesNotFound struct {
	Payload string
}

func (o *ListGitAccessNamesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/git_accesses/names][%d] listGitAccessNamesNotFound  %+v", 404, o.Payload)
}

func (o *ListGitAccessNamesNotFound) GetPayload() string {
	return o.Payload
}

func (o *ListGitAccessNamesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
