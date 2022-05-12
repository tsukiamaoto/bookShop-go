package model

type Query struct {
	Limit    string `json:"limit" form:"limit"`
	SortType string `json:"sort_type" form:"sort_type"`
	Order    string `json:"order" form:"order"`
	Cursor   Cursor
}

type Cursor struct {
	NextPage string `json:"next" form:"next"`
	PrevPage string `json:"prev" form:"prev"`
}
