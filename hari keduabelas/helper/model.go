package helper

import (
	"belajar-resfull-api/model/domain"
	"belajar-resfull-api/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToCategoryReponses(categories []domain.Category) []web.CategoryResponse {
	var categoryReponse []web.CategoryResponse
	for _, category := range categories {
		categoryReponse = append(categoryReponse, ToCategoryResponse(category))
	}
	return categoryReponse
}
