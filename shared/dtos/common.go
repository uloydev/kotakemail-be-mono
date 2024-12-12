package dtos

type PaginationMeta struct {
	Count    int    `json:"count"`
	Limit    int    `json:"limit"`
	Page     int    `json:"page,omitempty"`
	NextPage string `json:"next_page,omitempty"`
	PrevPage string `json:"prev_page,omitempty"`
	Cursor   string `json:"cursor,omitempty"`
}
