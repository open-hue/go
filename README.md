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
  resp, _ := c.GetLightByIdWithResponse(context.Background(), "MyLightID")
  lights := (*resp.JSON200).Data
  if len(lights) > 0 {
    fmt.Printf("Light with ID \"%s\" is called \"%s\"\n", lights[0].Id, lights[0].Metadata.Name)
  } else {
    fmt.Println("No lights found")
  }
}
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
build       Builds client code
spec/update Updates go dependencies
deps/update Refreshes the OpenAPI specification from its source
generate    Generates client code from OpenAPI specification
help        Lists help commands
```

### Updating Spec

The spec is managed within [open-hue/spec][spec] repository. This repository has a git sub-module with a static reference to prevent creating inconsistent builds. All updadtes to the spec must be intentionally pulled. The idea is to eventually pin static releases from the specification repository.

To upddate the internal sub-module reference:

1. `make spec/update`
1. `make generate`

Check for diffs, if there are any relevant changes it means the spec changed and thus the code must be regenerated & recompiled.

[spec]: http://github.com/open-hue/spec
