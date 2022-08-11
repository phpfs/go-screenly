package main

import (
	"errors"
	"fmt"
	"github.com/phpfs/go-screenly/screenly"
	"github.com/phpfs/go-screenly/screenly/client/screens"
)

// Interactive documentation: https://developer.screenlyapp.com
func main() {
	// Fill in your API token as retrieved from Screenly
	api := screenly.NewScreenlyAPI("my-secret-api-key")
	// Use params.Id or similar to set request properties
	params := screens.NewScreensListParams()
	res, err := api.Screens.ScreensList(params, nil)

	if err != nil {
		// Errors can be checked against default error types
		if errors.Is(err, screens.NewScreensListUnauthorized()) {
			fmt.Printf("Unauthorized! Please check your API token.")
		} else {
			fmt.Printf("An unexpected error occured: %v", err)
		}
	} else {
		if len(res.Payload) == 0 {
			fmt.Println("No screens found!")
		}
		// Results are given as slice...
		for _, screen := range res.Payload {
			// ...and properly typed within
			fmt.Printf("%s - Status: %s\n", *screen.Name, *screen.Status)
		}
	}
}
