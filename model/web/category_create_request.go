package web

type CategoryRequest struct {
	Id   int    `json:"id"`
	Name string `validate:"required" json:"name"`
}
