package response

type Page struct {
	Items any  `json:"items"`
	Meta  Meta `json:"meta"`
}

type Meta struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
	Total int `json:"total"`
	Pages int `json:"pages"`
}
