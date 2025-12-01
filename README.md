# Invgo

A Go module for integrating with the Invgate API, supporting per-instance authentication and API access.

>[!NOTE]
>This is under active development. Not all endpoints are implemented. See [API Coverage](./API_COVERAGE.md) for a detailed list of what is currently covered.

## Features

* 🌐 Connect to multiple Invgate instances
* 🔐 Handle authentication securely
* 📦 Simplified API client for Invgate endpoints

- [Install](#install)
- [Usage](#usage)
- [Configure](#configure)
    - [Scopes](#scopes)
- [Contributing](#contributing)

## Install

```bash
go get github.com/tmstorm/invgo
```

## Usage

```go
package main

import (
        "github.com/tmstorm/invgo"
        "log"
)

func main() {
    // Set Invgate scopes
    // these must be defined before creating the client to
    // ensure it is given the correct access permissions to the API
    scopes := []invgo.ScopeType{
        invgo.BreakingNewsAll,
    }

    // Create the client
    client, err := invgate.New(&invgo.Invgate{
        BaseURL: "https://invgate-instance.com",
        TokenURL: "https://invgate-instance.com/oauth/token",
        ClientID: "client_id",
        ClientSecret: "client_secret",
        Scopes: scopes,
    })
    if err != nil {
        log.Fatalf("unable to create Invgate client: %s", err.Error())
    }

    version, err := client.ServiceDeskVersion().Get()
    if err != nil {
        log.Printf("could not get service desk version: %s", err.Error())
    } else {
        log.Printf("Users: %s", version)
    }
}
```

## Configure

The `New(cfg *Invgate)` function accepts an `Invgate` struct with:
* `BaseURL`: the root URL of your Invgate instance
* `TokenURL`: the OAuth2 token URL
* `ClientID`: your API client ID
* `ClientSecret`: your API client secret
* `Scopes`: a slice of ScopeType representing required permissions

`New` authenticates with Invgate using OAuth2 client credentials and returns a configured Client ready for API calls.

## Scopes

Invgate requires **scopes** for API access in the format:

```
api.{version}.{endpoint}.{subresource}:{method}
```

Example: `api.v1.breakingnews.all:get`

Scopes are defined in `invgo.ScopeType` constants. All supported scope types can be found in `scopes.go`. 

>[!NOTE]
>Ingvate does not allow a generic 'all' for an endpoints methods, so you must add the corresponding scope for every method you intend on using for that endpoint.
>However if an endpoint only accepts one method I have shortened the scope name to indicate that it only supports one method.

```go
import "github.com/tmstorm/invgo"

main() {
   scopes := []invgo.ScopeType{
        invgo.BreakingNewsAll, // BreakingNewsAll only supports the GET method
        invgo.BreakingNewsGet, // BreakingNews is its own endpoint that supports GET, POST, and PUT
        invgo.BreakingNewsPost,
        invgo.BreakingNewsPut,
    }
}
```

Scopes must be included in your `cfg.Scopes` array at client creation, and are checked at runtime before API calls

## Contributing

If you implement an endpoint please add it to the coverage.go and run the coverage script to update the API_COVERAGE.md before making a PR.
```bash
go run ./scripts/coverage_report.go

```
