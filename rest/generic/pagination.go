package generic

type Pagination struct {
	HasNext bool   `json:"has_next"`
	Cursor  string `json:"cursor"`
	Size    int64  `json:"size"`
}
