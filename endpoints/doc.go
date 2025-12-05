/*
Package endpoints is the implementation of Invgate endpoints and their methods.

Each endpoint should be in its own file for clarity and readbility e.g. breaking news should be BreakingNews.go.
To expose the endpoint through the Invgo public API a Client method should be defined for each it in invgo/endpoint_methods.go.
If a new endpoint is added ensure it is added to invgo/coverage.go along with its related methods. This ensures the coverage script catches that
it has been implemented.
*/
package endpoints
