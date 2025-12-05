package endpoints

import (
	"encoding/json"

	"github.com/tmstorm/invgo/internal/methods"
	"github.com/tmstorm/invgo/internal/utils"
	"github.com/tmstorm/invgo/scopes"
)

type (
	// CategoriesMethods is used to call CategoriesMethods
	CategoriesMethods struct{ methods.MethodCall }

	// CategoriesGetResponse is used to map a category from the Categories GET method
	CategoriesGetResponse struct {
		ID               string `json:"id,omitempty"` // NOTE: API documentation says this is of type int but is delivered as type string
		ParentCategoryID int    `json:"parent_category_id,omitempty"`
		Name             string `json:"name,omitempty"`
	}

	CategoriesGetParams struct {
		ID int `url:"id"`
	}
)

// Get for Categories
// Requires scope: CategoriesGet
// If id == 0 all IDs will be provided
// See https://releases.invgate.com/service-desk/api/#categories-GET
func (cat *CategoriesMethods) Get(p CategoriesGetParams) ([]CategoriesGetResponse, error) {
	cat.RequiredScope = scopes.CategoriesGet

	q, err := utils.StructToQuery(p)
	if err != nil {
		return nil, err
	}
	cat.Endpoint.RawQuery = q.Encode()

	resp, err := cat.RemoteGet()
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
