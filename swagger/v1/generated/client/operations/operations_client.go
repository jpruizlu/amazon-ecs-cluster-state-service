// Copyright 2016-2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetInstance Get instance using cluster name and instance ARN
*/
func (a *Client) GetInstance(params *GetInstanceParams) (*GetInstanceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInstanceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetInstance",
		Method:             "GET",
		PathPattern:        "/instances/{cluster}/{arn}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetInstanceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetInstanceOK), nil

}

/*
GetTask Get task using cluster name and task ARN
*/
func (a *Client) GetTask(params *GetTaskParams) (*GetTaskOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTaskParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetTask",
		Method:             "GET",
		PathPattern:        "/tasks/{cluster}/{arn}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetTaskReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTaskOK), nil

}

/*
ListInstances Lists all instances, after applying filters if any
*/
func (a *Client) ListInstances(params *ListInstancesParams) (*ListInstancesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListInstancesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListInstances",
		Method:             "GET",
		PathPattern:        "/instances",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListInstancesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListInstancesOK), nil

}

/*
ListTasks Lists all tasks, after applying filters if any
*/
func (a *Client) ListTasks(params *ListTasksParams) (*ListTasksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListTasksParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListTasks",
		Method:             "GET",
		PathPattern:        "/tasks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListTasksReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListTasksOK), nil

}

/*
StreamInstances Streams all instances
*/
func (a *Client) StreamInstances(params *StreamInstancesParams, writer io.Writer) (*StreamInstancesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStreamInstancesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "StreamInstances",
		Method:             "GET",
		PathPattern:        "/stream/instances",
		ProducesMediaTypes: []string{"application/octet-stream"},
		ConsumesMediaTypes: []string{"application/octet-stream"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &StreamInstancesReader{formats: a.formats, writer: writer},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*StreamInstancesOK), nil

}

/*
StreamTasks Streams all tasks
*/
func (a *Client) StreamTasks(params *StreamTasksParams, writer io.Writer) (*StreamTasksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStreamTasksParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "StreamTasks",
		Method:             "GET",
		PathPattern:        "/stream/tasks",
		ProducesMediaTypes: []string{"application/octet-stream"},
		ConsumesMediaTypes: []string{"application/octet-stream"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &StreamTasksReader{formats: a.formats, writer: writer},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*StreamTasksOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}