package invgo

import (
	"encoding/json"

	"github.com/tmstorm/invgo/internal/utils"
)

// CategoriesMethods is used to call CategoriesMethods
type CategoriesMethods MethodCall

// Categories is used to get all categories for the current Invgate instance
// See https://releases.invgate.com/service-desk/api/#categories
func (c *Client) Categories() *CategoriesMethods {
	ep := c.APIURL.JoinPath("/categories")
	return &CategoriesMethods{
		client:   c,
		Endpoint: ep,
	}
}

// CategoriesGetResponse is used to map a category from the Categories GET method
type CategoriesGetResponse struct {
	ID               string `json:"id,omitempty"` // NOTE: API documentation says this is of type int but is delivered as type string
	ParentCategoryID int    `json:"parent_category_id,omitempty"`
	Name             string `json:"name,omitempty"`
}

type CategoriesGetParams struct {
	ID int `url:"id"`
}

// Get method returns an array of categories in the current Invgate instance
// If id == 0 all IDs will be provided
// See https://releases.invgate.com/service-desk/api/#categories-GET
func (cat *CategoriesMethods) Get(p CategoriesGetParams) ([]CategoriesGetResponse, error) {
	err := checkScopes(cat.client.CurrentScopes, CategoriesGet)
	if err != nil {
		return []CategoriesGetResponse{}, err
	}

	q, err := utils.StructToQuery(p)
	if err != nil {
		return nil, err
	}
	cat.Endpoint.RawQuery = q.Encode()

	m := MethodCall(*cat)
	resp, err := m.get()
	if err != nil {
		return []CategoriesGetResponse{}, err
	}

	var d []CategoriesGetResponse
	err = json.Unmarshal(resp, &d)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
