package helper

import (
	"api-test/model/domain"
	"api-test/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToCategoryResponses(category []domain.Category) []web.CategoryResponse {
	var response []web.CategoryResponse
	for _, item_data := range category {
		response = append(response, ToCategoryResponse(item_data))
	}

	return response
}
