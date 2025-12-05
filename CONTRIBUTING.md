# Adding a new endpoint

1. Create `invgo/endpoints/endpoint_name.go`
2. Define:
    ```go
    type (
        NewEnpointMethods struct { methods.MethodCall }
        NewEndpointGetResponse struct {
            Message string `json:message`
        }
        NewEnpointGetParams struct {
            ID string `url:id,required`
        }
    )

    //Get for NewEndpoint
    // See [Link to Invgate API docs for new endpoint]
    func (c *NewEnpointMethods) Get(p NewEnpointGetParams) (NewEndpointGetResponse, error) {
        r := NewEndpointGetResponse{}

        // Ensure required scope is set before request is made.
        // The only acception to this is if it is and Attribute endpoint.
        // Since they all call the same Get method this must be set in
        // endpoint_methods.go when creating the Invgo client method.
        c.RequiredScope = scopes.NewEndpointGet

        // Construct url params
        q, err := utils.StructToQuery(p)
        if err != nil {
            return r, err
        }
        c.Endpoint.RawQuery = q.Encode()

        // Send Request to Invgate
        resp, err := c.RemoteGet()
        if err != nil {
            return r, err
        }

        // Map results to NewEndpointGetResponse
        err = json.Unmarshal(resp, &r)
        if err != nil {
            return r, err
        }
        return r, nil
    }

    ```
3. Add to `invgo/endpoint_methods.go`:
    ```go
    func (c *Client) NewEndpoint() *endpoints.NewEnpointMethods {
        return NewPublicMethod[endpoints.NewEnpointMethods](c, "/newendpoint")
    }
    ```
4. Add to `coverage.go`
5. Run coverage script:
    ```
    go run ./scripts/coverage_report.go
    ```
