package catalogs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new catalogs API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for catalogs API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetNodesIdentifierCatalogs fetches catalog of specified node

Fetch catalog of specified node

*/
func (a *Client) GetNodesIdentifierCatalogs(params *GetNodesIdentifierCatalogsParams, authInfo runtime.ClientAuthInfoWriter) (*GetNodesIdentifierCatalogsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNodesIdentifierCatalogsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetNodesIdentifierCatalogs",
		Method:             "GET",
		PathPattern:        "/nodes/{identifier}/catalogs",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetNodesIdentifierCatalogsReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetNodesIdentifierCatalogsOK), nil
}

/*
GetNodesIdentifierCatalogsSource fetches catalog of specified node for given source

Fetch catalog of specified node for given source

*/
func (a *Client) GetNodesIdentifierCatalogsSource(params *GetNodesIdentifierCatalogsSourceParams, authInfo runtime.ClientAuthInfoWriter) (*GetNodesIdentifierCatalogsSourceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNodesIdentifierCatalogsSourceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetNodesIdentifierCatalogsSource",
		Method:             "GET",
		PathPattern:        "/nodes/{identifier}/catalogs/{source}",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetNodesIdentifierCatalogsSourceReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetNodesIdentifierCatalogsSourceOK), nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
