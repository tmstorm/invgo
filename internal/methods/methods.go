// Package methods
package methods

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/tmstorm/invgo/scopes"
)

type (
	// MethodCall is used to construct a call that can be passed to an endpoints method
	MethodCall struct {
		// Client is the initial client created and holds valid credentials needed
		// for every connection
		Client *Client
		// Endpoint is the URl of the end point to be called in the invoked method
		Endpoint *url.URL
	}

	// Client is used to build a connection with an Invgate api instance
	Client struct {
		// HTTPClient is used to access the underlying *http.Client
		HTTPClient *http.Client
		// CurrentScopes are used to keep track of the allowed scopes when making requests
		CurrentScopes []scopes.ScopeType
		// APIURL is used to se the BaseURL URL for connecting to an API Invgate instance
		APIURL *url.URL
	}

	// InvgateError is used to construct an error received from the Invgate API
	InvgateError struct {
		Error  string `json:"error,omitempty"`
		Status int    `json:"status,omitempty"`
	}
)

// RemoteGet is the underlying GET method called when making a GET request to Invgate
func (m *MethodCall) RemoteGet() ([]byte, error) { return m.get() }

// get is the internal method used for GET requests of all endpoints
func (m *MethodCall) get() ([]byte, error) {
	return methodConstructor(http.MethodGet, m, nil)
}

// RemotePost is the underlying POST method called when making a POST request to Invgate
func (m *MethodCall) RemotePost() ([]byte, error) { return m.post() }

// post is the internal method used for POST requests of all endpoints
func (m *MethodCall) post() ([]byte, error) {
	return methodConstructor(http.MethodPost, m, nil)
}

// RemotePatch is the underlying PATCH method called when making a PATCH request to Invgate
func (m *MethodCall) RemotePatch() ([]byte, error) { return m.patch() }

// patch is the internal method used for PATCH requests of all endpoints
func (m *MethodCall) patch() ([]byte, error) {
	return methodConstructor(http.MethodPatch, m, nil)
}

// RemotePut is the underlying PUT method called when making a PUT request to Invgate
func (m *MethodCall) RemotePut() ([]byte, error) { return m.put() }

// put is the internal method used for PUT requests of all endpoints
func (m *MethodCall) put() ([]byte, error) {
	return methodConstructor(http.MethodPut, m, nil)
}

// RemoteDelete is the underlying DELETE method called when making a DELETE request to Invgate
func (m *MethodCall) RemoteDelete() ([]byte, error) { return m.delete() }

// delete is the internal method used for DELETE requests of all endpoints
func (m *MethodCall) delete() ([]byte, error) {
	return methodConstructor(http.MethodDelete, m, nil)
}

// methodConstructor is used to build and call all internal methods the the Invgate API
func methodConstructor(methodType string, m *MethodCall, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(methodType, m.Endpoint.String(), body)
	if err != nil {
		return nil, err
	}

	resp, err := m.Client.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	return checkErrorResponse(resp)
}

// checkErrorResponse is used to check for errors from the Invgate API
func checkErrorResponse(r *http.Response) ([]byte, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	if r.StatusCode != 200 {
		var m InvgateError
		err = json.Unmarshal(body, &m)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%v status: %v", m.Error, m.Status)
	}
	return body, nil
}
