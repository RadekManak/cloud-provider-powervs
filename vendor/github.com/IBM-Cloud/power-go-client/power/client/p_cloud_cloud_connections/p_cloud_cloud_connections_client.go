// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_cloud_connections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new p cloud cloud connections API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for p cloud cloud connections API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	PcloudCloudconnectionsDelete(params *PcloudCloudconnectionsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudCloudconnectionsDeleteOK, *PcloudCloudconnectionsDeleteAccepted, error)

	PcloudCloudconnectionsGet(params *PcloudCloudconnectionsGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudCloudconnectionsGetOK, error)

	PcloudCloudconnectionsGetall(params *PcloudCloudconnectionsGetallParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudCloudconnectionsGetallOK, error)

	PcloudCloudconnectionsNetworksDelete(params *PcloudCloudconnectionsNetworksDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudCloudconnectionsNetworksDeleteOK, *PcloudCloudconnectionsNetworksDeleteAccepted, error)

	PcloudCloudconnectionsNetworksPut(params *PcloudCloudconnectionsNetworksPutParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudCloudconnectionsNetworksPutOK, *PcloudCloudconnectionsNetworksPutAccepted, error)

	PcloudCloudconnectionsPost(params *PcloudCloudconnectionsPostParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudCloudconnectionsPostOK, *PcloudCloudconnectionsPostCreated, *PcloudCloudconnectionsPostAccepted, error)

	PcloudCloudconnectionsPut(params *PcloudCloudconnectionsPutParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudCloudconnectionsPutOK, *PcloudCloudconnectionsPutAccepted, error)

	PcloudCloudconnectionsVirtualprivatecloudsGetall(params *PcloudCloudconnectionsVirtualprivatecloudsGetallParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudCloudconnectionsVirtualprivatecloudsGetallOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  PcloudCloudconnectionsDelete deletes a cloud connection
*/
func (a *Client) PcloudCloudconnectionsDelete(params *PcloudCloudconnectionsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudCloudconnectionsDeleteOK, *PcloudCloudconnectionsDeleteAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudCloudconnectionsDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "pcloud.cloudconnections.delete",
		Method:             "DELETE",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudCloudconnectionsDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *PcloudCloudconnectionsDeleteOK:
		return value, nil, nil
	case *PcloudCloudconnectionsDeleteAccepted:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for p_cloud_cloud_connections: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PcloudCloudconnectionsGet gets a cloud connection s state information
*/
func (a *Client) PcloudCloudconnectionsGet(params *PcloudCloudconnectionsGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudCloudconnectionsGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudCloudconnectionsGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "pcloud.cloudconnections.get",
		Method:             "GET",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudCloudconnectionsGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PcloudCloudconnectionsGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for pcloud.cloudconnections.get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PcloudCloudconnectionsGetall gets all cloud connections in this cloud instance
*/
func (a *Client) PcloudCloudconnectionsGetall(params *PcloudCloudconnectionsGetallParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudCloudconnectionsGetallOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudCloudconnectionsGetallParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "pcloud.cloudconnections.getall",
		Method:             "GET",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudCloudconnectionsGetallReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PcloudCloudconnectionsGetallOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for pcloud.cloudconnections.getall: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PcloudCloudconnectionsNetworksDelete deletes a network from a cloud connection
*/
func (a *Client) PcloudCloudconnectionsNetworksDelete(params *PcloudCloudconnectionsNetworksDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudCloudconnectionsNetworksDeleteOK, *PcloudCloudconnectionsNetworksDeleteAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudCloudconnectionsNetworksDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "pcloud.cloudconnections.networks.delete",
		Method:             "DELETE",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}/networks/{network_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudCloudconnectionsNetworksDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *PcloudCloudconnectionsNetworksDeleteOK:
		return value, nil, nil
	case *PcloudCloudconnectionsNetworksDeleteAccepted:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for p_cloud_cloud_connections: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PcloudCloudconnectionsNetworksPut adds a network to the cloud connection
*/
func (a *Client) PcloudCloudconnectionsNetworksPut(params *PcloudCloudconnectionsNetworksPutParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudCloudconnectionsNetworksPutOK, *PcloudCloudconnectionsNetworksPutAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudCloudconnectionsNetworksPutParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "pcloud.cloudconnections.networks.put",
		Method:             "PUT",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}/networks/{network_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudCloudconnectionsNetworksPutReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *PcloudCloudconnectionsNetworksPutOK:
		return value, nil, nil
	case *PcloudCloudconnectionsNetworksPutAccepted:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for p_cloud_cloud_connections: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PcloudCloudconnectionsPost creates a new cloud connection
*/
func (a *Client) PcloudCloudconnectionsPost(params *PcloudCloudconnectionsPostParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudCloudconnectionsPostOK, *PcloudCloudconnectionsPostCreated, *PcloudCloudconnectionsPostAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudCloudconnectionsPostParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "pcloud.cloudconnections.post",
		Method:             "POST",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudCloudconnectionsPostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, nil, err
	}
	switch value := result.(type) {
	case *PcloudCloudconnectionsPostOK:
		return value, nil, nil, nil
	case *PcloudCloudconnectionsPostCreated:
		return nil, value, nil, nil
	case *PcloudCloudconnectionsPostAccepted:
		return nil, nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for p_cloud_cloud_connections: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PcloudCloudconnectionsPut updates a cloud connection
*/
func (a *Client) PcloudCloudconnectionsPut(params *PcloudCloudconnectionsPutParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudCloudconnectionsPutOK, *PcloudCloudconnectionsPutAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudCloudconnectionsPutParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "pcloud.cloudconnections.put",
		Method:             "PUT",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudCloudconnectionsPutReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *PcloudCloudconnectionsPutOK:
		return value, nil, nil
	case *PcloudCloudconnectionsPutAccepted:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for p_cloud_cloud_connections: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PcloudCloudconnectionsVirtualprivatecloudsGetall gets all virtual private cloud connections in this cloud instance
*/
func (a *Client) PcloudCloudconnectionsVirtualprivatecloudsGetall(params *PcloudCloudconnectionsVirtualprivatecloudsGetallParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudCloudconnectionsVirtualprivatecloudsGetallOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudCloudconnectionsVirtualprivatecloudsGetallParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "pcloud.cloudconnections.virtualprivateclouds.getall",
		Method:             "GET",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections-virtual-private-clouds",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudCloudconnectionsVirtualprivatecloudsGetallReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PcloudCloudconnectionsVirtualprivatecloudsGetallOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for pcloud.cloudconnections.virtualprivateclouds.getall: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
