package invgo

import (
	"context"
	"errors"

	"github.com/tmstorm/invgo/internal/methods"
	"github.com/tmstorm/invgo/internal/utils"
	"github.com/tmstorm/invgo/scopes"
	"golang.org/x/oauth2/clientcredentials"
)

type (
	// Invgate is used to set the configuration options for connection to an Invgate API instance.
	Invgate struct {
		// BaseURL defines the url of the Invgate instance.
		BaseURL string
		// TokenURL defines the oAuth2 token endpoint for authentication with instance.
		TokenURL string
		// ClientID defines the ClientID of the Invgate instance.
		ClientID string
		// ClientSecret defines the secret created in Invgate to allow Invgo to connect to the instance.
		ClientSecret string
		// AllowHTTP defines if the connection should be allowed to let the base url keep the scheme http or not.
		// WARNING: This should only be set to true in testing or dev environments.
		AllowHTTP bool
		// Scopes defines which scopes will be requested when requestion the token from the Invgate instance.
		// If a scope is not defined here the client will be denied access to its endpoint on future requests.
		Scopes []scopes.ScopeType
	}

	// Client implements methods.Client for use to connect with Invgate.
	Client methods.Client
)

// InvgateAPIPath defines the base path for the Invgate API.
// WARNING: This should not be changed unless you know for sure it will not break Invgo.
var InvgateAPIPath = "/api/v1"

// New authenticates with the Invgate using the provided options and returns the client for API calls.
func New(cfg *Invgate) (*Client, error) {
	// check that at least one scop has been provided
	if len(cfg.Scopes) == 0 {
		return nil, errors.New("no scopes were provided")
	}

	// Set config for creating a connection with Invgate
	scopes := scopes.CreateScopes(cfg.Scopes)
	cred := &clientcredentials.Config{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		TokenURL:     cfg.TokenURL,
		Scopes:       scopes,
	}

	// Parse base url given to ensure it is not malformed
	apiURL, err := utils.ParseURL(cfg.BaseURL, InvgateAPIPath, cfg.AllowHTTP)
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
