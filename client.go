// Package invgo
package invgo

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/oauth2/clientcredentials"
)

type (
	// Invgate is used to set the configuration options for connection to an Invgate API instance
	Invgate struct {
		BaseURL      string      `json:"base_url,omitempty"`
		TokenURL     string      `json:"token_url,omitempty"`
		ClientID     string      `json:"client_id,omitempty"`
		ClientSecret string      `json:"client_secret,omitempty"`
		Scopes       []ScopeType `json:"scopes,omitempty"`
	}

	// Client is used to build a connection with an Invgate api instance
	Client struct {
		// HTTPClient is used to access the underlying *http.Client
		HTTPClient *http.Client
		// CurrentScopes are used to keep track of the allowed scopes when making requests
		CurrentScopes []ScopeType
		// APIURL is used to se the BaseURL URL for connecting to an API Invgate instance
		APIURL *url.URL
	}

	ClientCore interface {
		DoRequest(req *http.Request) (*http.Response, error)
	}
)

// Define the base path for the invgate API
var invgateAPIPath = "/api/v1"

// New authenticates with the Invgate using the provided options and returns the client for API calls.
func New(cfg *Invgate) (*Client, error) {
	// check that at least one scop has been provided
	if len(cfg.Scopes) == 0 {
		return nil, errors.New("no scopes were provided")
	}

	// Set config for creating a connection with Invgate
	scopes := createScopes(cfg.Scopes)
	cred := &clientcredentials.Config{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		TokenURL:     cfg.TokenURL,
		Scopes:       scopes,
	}

	// Parse base url given to ensure it is not malformed
	baseURL := strings.TrimSuffix(cfg.BaseURL, "/")
	apiURL, err := url.Parse(baseURL + invgateAPIPath)
	if err != nil {
		return nil, err
	}

	// Create a client for future use
	client := &Client{
		HTTPClient:    cred.Client(context.Background()),
		CurrentScopes: cfg.Scopes,
		APIURL:        apiURL,
	}

	return client, nil
}
