package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"

	hue "github.com/open-hue/go"
)

func main() {
	server := os.Getenv("HUE_SERVER")
	applicationKey := os.Getenv("HUE_APPLICATION_KEY")

	// TODO Move to client boilerplate code
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	c, err := hue.NewAuthenticatedClientWithResponses(server, applicationKey)
	if err != nil {
		panic(err)
	}
	resp, err := c.GetLightsWithResponse(context.Background())
	if err != nil {
		panic(err)
	}
	lights := (*resp.JSON200).Data
	for _, l := range lights {
		fmt.Printf("Light {ID:\"%s\", Name: \"%s\"}\n", l.Id, l.Metadata.Name)
	}
}
