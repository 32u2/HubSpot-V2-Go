package main

import (
	"fmt"

	"github.com/retgits/hubspot/client"
)

func main() {
	hubspot := client.NewClient().WithAPIKey("<myAccessToken>")
	fmt.Println(hubspot.APIKey)
}
