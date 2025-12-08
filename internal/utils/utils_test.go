package utils_test

import (
	"net/url"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tmstorm/invgo/internal/utils"
)

type Test struct {
	ID      int   `url:"id"`
	UserIDs []int `url:"user_ids"`
	Nested  Nested
}

type Nested struct {
	Locations []string `url:"locations"`
}

func TestStructToQuery(t *testing.T) {
	a := assert.New(t)

	u, err := url.Parse("http://example.com")
	a.NoError(err)

	fake := Test{
		ID:      1,
		UserIDs: []int{1, 2},
		Nested: Nested{
			Locations: []string{"home", "work"},
		},
	}

	q, err := utils.StructToQuery(fake)
	a.NoError(err)

	u.RawQuery = q.Encode()

	for k, v := range q {
		switch k {
		case "id":
			a.Contains(v, strconv.Itoa(fake.ID))
		case "user_ids[0]":
			a.Contains(v, strconv.Itoa(fake.UserIDs[0]))
		case "user_ids[1]":
			a.Contains(v, strconv.Itoa(fake.UserIDs[1]))
		case "locations[0]":
			a.Contains(v, fake.Nested.Locations[0])
		case "locations[1]":
			a.Contains(v, fake.Nested.Locations[1])
		default:
			a.Fail("key, value pair not found - key: %s, value: %s", k, v)
		}
	}
}

func TestParseURL(t *testing.T) {
	a := assert.New(t)

	// secure
	u, err := utils.ParseURL("https://secure.com", "/api/v1", false)
	a.NoError(err)
	a.Equal("https", u.Scheme)

	// no scheme provided missing url leading slash
	u, err = utils.ParseURL("secure.com", "api/v1", false)
	a.NoError(err)
	a.Equal("https", u.Scheme)

	// unsecure and don't allow http
	u, err = utils.ParseURL("http://unsecure.com", "/api/v1", false)
	a.NoError(err)
	a.Equal("https", u.Scheme)

	// unsecure and allow http
	u, err = utils.ParseURL("http://unsecure.com", "/api/v1", true)
	a.NoError(err)
	a.Equal("http", u.Scheme)
}
