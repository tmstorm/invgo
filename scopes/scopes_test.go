package scopes_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tmstorm/invgo/scopes"
)

func TestCreateScopes(t *testing.T) {
	a := assert.New(t)
	scps := []scopes.ScopeType{
		scopes.BreakingNewsAll,
		scopes.BreakingNewsAttributesStatus,
		scopes.BreakingNewsGet,
		scopes.CategoriesGet,
	}

	str := scopes.CreateScopes(scps)

	for _, v := range str {
		a.Contains(scps, scopes.ScopeType(v))
	}
}

func TestCheckScopes(t *testing.T) {
	a := assert.New(t)
	scps := []scopes.ScopeType{
		scopes.BreakingNewsAll,
		scopes.BreakingNewsAttributesStatus,
		scopes.BreakingNewsGet,
		scopes.CategoriesGet,
	}

	for _, v := range scps {
		err := scopes.CheckScopes(scps, v)
		a.NoError(err)
	}

	err := scopes.CheckScopes(scps, scopes.HelpDesksGet)
	a.Error(err)
}
