package base

type Page struct {
	PageSize  int `json:"pageSize"`
	PageIndex int `json:"pageIndex"`
}

type Sort struct {
	SortBy    string `json:"sortBy"`
	SortOrder string `json:"sortOrder"`
}
