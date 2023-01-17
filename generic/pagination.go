package generic

type AdvancedTradePaginationResponse struct {
	HasNext bool   `json:"has_next"`
	Cursor  string `json:"cursor"`
	Size    int64  `json:"size"`
}
type AdvancedTradePaginationRequest struct {
	Limit  int32  `url:"limit,omitempty"`
	Cursor string `url:"cursor,omitempty"`
}

type SignInWithCoinbasePaginationRequest struct {
	EndingBefore  string `url:"ending_before,omitempty"`
	StartingAfter string `url:"starting_after,omitempty"`
	Limit         int32  `url:"limit,omitempty"`
	Order         string `url:"order,omitempty"`
}

type SignInWithCoinbasePaginationResponse struct {
	EndingBefore  string `json:"ending_before"`
	StartingAfter string `json:"starting_after"`
	Limit         int32  `json:"limit"`
	Order         string `json:"order"`
	PreviousUri   string `json:"previous_uri"`
	NextUri       string `json:"next_uri"`
}
