package catalog

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new catalog API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for catalog API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetCatalogs gets list of all catalogs

The catalogs endpoint returns json data that represent the catalogs of
all hardware in the system.

*/
func (a *Client) GetCatalogs(params *GetCatalogsParams, authInfo runtime.ClientAuthInfoWriter) (*GetCatalogsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCatalogsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCatalogs",
		Method:             "GET",
		PathPattern:        "/catalogs",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetCatalogsReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCatalogsOK), nil
}

/*
GetCatalogsIdentifier gets list of all catalogs

The catalogs endpoint returns json data that represent the catalogs of
all hardware in the system.

*/
func (a *Client) GetCatalogsIdentifier(params *GetCatalogsIdentifierParams, authInfo runtime.ClientAuthInfoWriter) (*GetCatalogsIdentifierOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCatalogsIdentifierParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCatalogsIdentifier",
		Method:             "GET",
		PathPattern:        "/catalogs/{identifier}",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetCatalogsIdentifierReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCatalogsIdentifierOK), nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
