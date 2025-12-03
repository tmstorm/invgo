package invgo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type (
	// MethodCall is used to construct a call that can be passed to an endpoints method
	MethodCall struct {
		// client is the initial client created and holds valid credentials needed
		// for every connection
		client *Client
		// Endpoint is the URl of the end point to be called in the invoked method
		Endpoint *url.URL
	}

	// InvgateError is used to construct an error received from the Invgate API
	InvgateError struct {
		Error  string `json:"error,omitempty"`
		Status int    `json:"status,omitempty"`
	}
)

// get is the internal method used for GET requests of all endpoints
func (m *MethodCall) get() ([]byte, error) {
	return methodConstructor(http.MethodGet, m, nil)
}

// post is the internal method used for POST requests of all endpoints
func (m *MethodCall) post() ([]byte, error) {
	return methodConstructor(http.MethodPost, m, nil)
}

// patch is the internal method used for PATCH requests of all endpoints
func (m *MethodCall) patch() ([]byte, error) {
	return methodConstructor(http.MethodPatch, m, nil)
}

// put is the internal method used for PUT requests of all endpoints
func (m *MethodCall) put() ([]byte, error) {
	return methodConstructor(http.MethodPut, m, nil)
}

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

	resp, err := m.client.HTTPClient.Do(req)
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
