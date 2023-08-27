package web

type CategoryUpdateRequest struct {
	Id   int    `validated:"required"`
	Name string `validated:"required, max=200,min=0"`
}
