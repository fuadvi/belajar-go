package web

type CategoryCreateRequest struct {
	Name string `validated:"required, max=200,min=0"`
}
