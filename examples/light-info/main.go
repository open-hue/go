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
	lightID := os.Getenv("HUE_LIGHT_ID")

	// TODO Move to client boilerplate code
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	c, err := hue.NewAuthenticatedClientWithResponses(server, applicationKey)
	if err != nil {
		panic(err)
	}
	resp, err := c.GetLightByIdWithResponse(context.Background(), lightID)
	if err != nil {
		panic(err)
	}
	lights := (*resp.JSON200).Data
	if len(lights) > 0 {
		fmt.Printf("Light with ID \"%s\" is called \"%s\"\n", lights[0].Id, lights[0].Metadata.Name)
	} else {
		fmt.Println("No lights found")
	}
}
