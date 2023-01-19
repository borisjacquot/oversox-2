package http

import (
	"encoding/json"
	"fmt"
	net_http "net/http"

	"github.com/borisjacquot/oversox-2/internal/client/overfast"
)

type httpClient struct {
	url    string
	client *net_http.Client
}

func NewClient(url string) overfast.Client {
	client := &net_http.Client{}

	return &httpClient{
		url:    url,
		client: client,
	}
}

func (c *httpClient) SearchPlayer(player string) (overfast.SearchPlayerResponse, error) {
	url := fmt.Sprintf("%s/players?name=%s", c.url, player)

	req, err := net_http.NewRequest("GET", url, nil)
	if err != nil {
		return overfast.SearchPlayerResponse{}, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return overfast.SearchPlayerResponse{}, err
	}

	body := overfast.SearchPlayerResponse{}
	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		return overfast.SearchPlayerResponse{}, err
	}

	fmt.Println(resp.StatusCode)

	return body, nil
}
