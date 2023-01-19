package overfast

type Client interface {
	SearchPlayer(string) (SearchPlayerResponse, error)
}

//structs here
type SearchPlayerRequest struct {
	Name     string `json:"name"`
	Privacy  string `json:"privacy"`
	OrdereBy string `json:"order_by"`
	Offset   string `json:"offset"`
	Limit    string `json:"limit"`
}

type SearchPlayerResponse struct {
	Total   int `json:"name"`
	Results []struct {
		PlayerID  string `json:"player_id"`
		Name      string `json:"name"`
		Privacy   string `json:"privacy"`
		CareerURL string `json:"career_url"`
	} `json:"results"`
	// for debug
	Detail []struct {
		Loc  []string `json:"location"`
		Msg  string   `json:"msg"`
		Type string   `json:"type"`
	} `json:"detail"`
}
