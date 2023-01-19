package main

import (
	"fmt"

	overfast_http "github.com/borisjacquot/oversox-2/internal/client/overfast/http"
)

func main() {
	url := "https://overfast-api.tekrop.fr"

	ofClient := overfast_http.NewClient(url)

	result, err := ofClient.SearchPlayer("LemonAdd")
	fmt.Printf("%v, %s", result, err)

}
