package invgo_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tmstorm/invgo"
	"github.com/tmstorm/invgo/scopes"
)

func TestInvgo(t *testing.T) {
	a := assert.New(t)
	cfg := &invgo.Invgate{
		BaseURL:      "http://test.invgate-instance.com",
		TokenURL:     "https://test.invgate-instance.com/oauth/token",
		ClientID:     "12345",
		ClientSecret: "clientSecret",
		AllowHTTP:    true,
		Scopes:       []scopes.ScopeType{scopes.BreakingNewsAll},
	}

	c, err := invgo.New(cfg)
	a.NoError(err)
	a.Equal("http", c.APIURL.Scheme)

	a.Equal(c.APIURL.String(), cfg.BaseURL+invgo.InvgateAPIPath)
	a.Contains(c.CurrentScopes, scopes.BreakingNewsAll)

	cfg.AllowHTTP = false
	cNext, err := invgo.New(cfg)
	a.NoError(err)
	a.Equal("https", cNext.APIURL.Scheme)
}
