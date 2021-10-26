package services

import (
	"api-test/model/web"
	"context"
)

type CategoryService interface {
	Create(ctx context.Context, request web.CategoryRequest) web.CategoryResponse
	Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse
	Delete(ctx context.Context, categoryId int)
	FindId(ctx context.Context, categoryId int) web.CategoryResponse
	FindAll(ctx context.Context) []web.CategoryResponse
}
