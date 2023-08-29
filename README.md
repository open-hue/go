# [open-hue](https://github.com/open-hue)/go

Community-driven Go API Client for Philips Hue.

## Disclaimer

The code is mostly auto-generated from the [OpenAPI Spec][spec] and it is currently under active development.

## Usage

It is quite early stage but it is functional to the single endpoint implemented.

```go
package main

import (
  "fmt"
  "context"
  hue "github.com/open-hue/go"
)

func main() {
  c, _ := hue.NewAuthenticatedClientWithResponses("https://10.0.0.23", "MyAppKey")
  resp, _ := c.GetLightsWithResponse(context.Background())
  lights := (*resp.JSON200).Data
  for _, l := range lights {
    fmt.Printf("Light {ID:\"%s\", Name: \"%s\"}\n", l.Id, l.Metadata.Name)
  }
}
```

### Built-in examples

There are examples for each endpoint in the `examples/` folder.

Most assume environment variables for configuration to simplify things, e.g.:

```sh
export HUE_SERVER="https://10.20.30.40"
export HUE_APPLICATION_KEY="this-is-clearly-not-valid"

go run examples/lights/all/main.go
```

Results in:

```sh
Light {ID:"c8sb0f94-50c7-20fb-9070-a063a91bfc00", Name: "Hue lamp 1"}
Light {ID:"71bje066-5ded-22e9-9ekd-5f5f40e6a8ae", Name: "Hue lamp 2"}
Light {ID:"8039b831-d937-282a-83f5-0bd3ea97dac3", Name: "Hue lamp 3"}
Light {ID:"5be4a6f4-4504-2c64-b7f6-dfe3004ff7f3", Name: "LED Strip 1"}
```

## Local development

### Directory

```
├── LICENSE       - Self-explanatory
├── Makefile      - Aliases for internal relevant commands
├── README.md     - This document
├── client.gen.go - Generated golang code
├── client.go     - Configuration
├── go.mod        - Go mod
├── go.sum        - Go dependencies lock
└── spec/         - Sub-module from open-hue/spec
    └── spec.yaml - OpenAPI Spec
```

### Makefile

```sh
help         Lists help commands
go/build     Builds client code
go/generate  Generates client code from OpenAPI specification
go/update    Updates go dependencies
spec/update  Refreshes the OpenAPI specification from its source
```

### Updating Spec

The spec is managed within [open-hue/spec][spec] repository. This repository has a git sub-module with a static reference to prevent creating inconsistent builds. All updadtes to the spec must be intentionally pulled. The idea is to eventually pin static releases from the specification repository.

To upddate the internal sub-module reference:

1. `make spec/update`
1. `make go/generate`

Check for diffs, if there are any relevant changes it means the spec changed and thus the code must be regenerated & recompiled.

[spec]: http://github.com/open-hue/spec
