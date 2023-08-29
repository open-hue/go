package client

//go:generate oapi-codegen --package=client -generate=client,types -o ./client.gen.go spec/spec.yaml
import (
	sp "github.com/deepmap/oapi-codegen/pkg/securityprovider"
)

// Context
// Most of the code is properly generated, except for
// * Authentication Boilerplate — "application key" header authentication options are not generated somehow, the methods below simplify for client users
// * (TODO) SSL Certificates — Since Hue Bridges run locally, they do not have valid SSL certificates. To cover for that we have a few methods that allow skipping certificacte validation

// ApplicationKeyHeader is used to authenticate requests from the client to Philips Hue Bridge
const ApplicationKeyHeader = "hue-application-key"

// Creates an authenticacted client with all required options
func NewAuthenticatedClient(server string, applicationKey string) (*Client, error) {
	p, err := sp.NewSecurityProviderApiKey("header", ApplicationKeyHeader, applicationKey)
	if err != nil {
		return nil, err
	}

	return NewClient(server, WithRequestEditorFn(p.Intercept))
}

// Creates an authenticacted client with responses and with all required options
func NewAuthenticatedClientWithResponses(server string, applicationKey string) (*ClientWithResponses, error) {
	p, err := sp.NewSecurityProviderApiKey("header", ApplicationKeyHeader, applicationKey)
	if err != nil {
		return nil, err
	}

	return NewClientWithResponses(server, WithRequestEditorFn(p.Intercept))
}
