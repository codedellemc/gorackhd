package identify

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new identify API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for identify API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetNodesIdentifierObmIdentify fetches status of identify light on node through o b m if supported

Fetch status of identify light on node through OBM (if supported)

*/
func (a *Client) GetNodesIdentifierObmIdentify(params *GetNodesIdentifierObmIdentifyParams, authInfo runtime.ClientAuthInfoWriter) (*GetNodesIdentifierObmIdentifyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNodesIdentifierObmIdentifyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetNodesIdentifierObmIdentify",
		Method:             "GET",
		PathPattern:        "/nodes/{identifier}/obm/identify",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetNodesIdentifierObmIdentifyReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetNodesIdentifierObmIdentifyOK), nil
}

/*
PostNodesIdentifierObmIdentify enables or disable identify light on node through o b m if supported

Enable or disable identify light on node through OBM (if supported)

*/
func (a *Client) PostNodesIdentifierObmIdentify(params *PostNodesIdentifierObmIdentifyParams, authInfo runtime.ClientAuthInfoWriter) (*PostNodesIdentifierObmIdentifyCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostNodesIdentifierObmIdentifyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostNodesIdentifierObmIdentify",
		Method:             "POST",
		PathPattern:        "/nodes/{identifier}/obm/identify",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostNodesIdentifierObmIdentifyReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostNodesIdentifierObmIdentifyCreated), nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}