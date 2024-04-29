package main

type ReviewData struct {
	Ads              []AdsItem     `json:"ads"`
	ExploreMoreItems []interface{} `json:"explore_more_items"`
	NextPages        []string      `json:"next_pages"`
	Results          []SearchItem  `json:"results"`
}

type ReviewResult struct {
	Data   ReviewData `json:"data"`
	Status string     `json:"status"`
	URL    string     `json:"url"`
}
