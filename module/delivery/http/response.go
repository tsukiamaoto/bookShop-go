package http

type dataResponse struct {
	Data     interface{} `json:"data"`
	NextPage string      `json:"next"`
	PrevPage string      `json:"prev"`
}
