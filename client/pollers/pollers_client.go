package pollers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new pollers API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for pollers API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeletePollersIdentifier deletes the specified poller

delete the specified poller

*/
func (a *Client) DeletePollersIdentifier(params *DeletePollersIdentifierParams, authInfo runtime.ClientAuthInfoWriter) (*DeletePollersIdentifierNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePollersIdentifierParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeletePollersIdentifier",
		Method:             "DELETE",
		PathPattern:        "/pollers/{identifier}",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeletePollersIdentifierReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeletePollersIdentifierNoContent), nil
}

/*
GetNodesIdentifierPollers fetches all pollers for specified node

Fetch all pollers for specified node

*/
func (a *Client) GetNodesIdentifierPollers(params *GetNodesIdentifierPollersParams, authInfo runtime.ClientAuthInfoWriter) (*GetNodesIdentifierPollersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNodesIdentifierPollersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetNodesIdentifierPollers",
		Method:             "GET",
		PathPattern:        "/nodes/{identifier}/pollers",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetNodesIdentifierPollersReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetNodesIdentifierPollersOK), nil
}

/*
GetPollers gets list of all pollers

get list of all pollers

*/
func (a *Client) GetPollers(params *GetPollersParams, authInfo runtime.ClientAuthInfoWriter) (*GetPollersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPollersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPollers",
		Method:             "GET",
		PathPattern:        "/pollers",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPollersReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPollersOK), nil
}

/*
GetPollersIdentifier gets specifics of the specified poller

Get specifics of the specified poller

*/
func (a *Client) GetPollersIdentifier(params *GetPollersIdentifierParams, authInfo runtime.ClientAuthInfoWriter) (*GetPollersIdentifierOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPollersIdentifierParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPollersIdentifier",
		Method:             "GET",
		PathPattern:        "/pollers/{identifier}",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPollersIdentifierReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPollersIdentifierOK), nil
}

/*
GetPollersLibrary gets list of possible library pollers

get list of possible library pollers

*/
func (a *Client) GetPollersLibrary(params *GetPollersLibraryParams, authInfo runtime.ClientAuthInfoWriter) (*GetPollersLibraryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPollersLibraryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPollersLibrary",
		Method:             "GET",
		PathPattern:        "/pollers/library",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPollersLibraryReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPollersLibraryOK), nil
}

/*
GetPollersLibraryIdentifier gets a single library poller

get a single library poller

*/
func (a *Client) GetPollersLibraryIdentifier(params *GetPollersLibraryIdentifierParams, authInfo runtime.ClientAuthInfoWriter) (*GetPollersLibraryIdentifierOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPollersLibraryIdentifierParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPollersLibraryIdentifier",
		Method:             "GET",
		PathPattern:        "/pollers/library/{identifier}",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPollersLibraryIdentifierReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPollersLibraryIdentifierOK), nil
}

/*
PatchPollersIdentifier patches specifics of the specified poller

patch specifics of the specified poller

*/
func (a *Client) PatchPollersIdentifier(params *PatchPollersIdentifierParams, authInfo runtime.ClientAuthInfoWriter) (*PatchPollersIdentifierOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchPollersIdentifierParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchPollersIdentifier",
		Method:             "PATCH",
		PathPattern:        "/pollers/{identifier}",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PatchPollersIdentifierReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchPollersIdentifierOK), nil
}

/*
PostPollers creates a poller

create a poller

*/
func (a *Client) PostPollers(params *PostPollersParams, authInfo runtime.ClientAuthInfoWriter) (*PostPollersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostPollersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostPollers",
		Method:             "POST",
		PathPattern:        "/pollers",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostPollersReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostPollersOK), nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
