// Code generated by go-swagger; DO NOT EDIT.

package node_pool_config_items

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new node pool config items API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for node pool config items API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AddOrUpdateNodePoolConfigItem adds update config item

Add/update a configuration item unique to the node pool.
*/
func (a *Client) AddOrUpdateNodePoolConfigItem(params *AddOrUpdateNodePoolConfigItemParams, authInfo runtime.ClientAuthInfoWriter) (*AddOrUpdateNodePoolConfigItemOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddOrUpdateNodePoolConfigItemParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addOrUpdateNodePoolConfigItem",
		Method:             "PUT",
		PathPattern:        "/kubernetes-clusters/{cluster_id}/node-pools/{node_pool_name}/config-items/{config_key}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddOrUpdateNodePoolConfigItemReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddOrUpdateNodePoolConfigItemOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addOrUpdateNodePoolConfigItem: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteNodePoolConfigItem deletes config item

Deletes config item.
*/
func (a *Client) DeleteNodePoolConfigItem(params *DeleteNodePoolConfigItemParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteNodePoolConfigItemNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteNodePoolConfigItemParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteNodePoolConfigItem",
		Method:             "DELETE",
		PathPattern:        "/kubernetes-clusters/{cluster_id}/node-pools/{node_pool_name}/config-items/{config_key}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteNodePoolConfigItemReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteNodePoolConfigItemNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteNodePoolConfigItem: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}