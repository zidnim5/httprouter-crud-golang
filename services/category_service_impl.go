package services

import (
	"api-test/exception"
	"api-test/helper"
	"api-test/model/domain"
	"api-test/model/web"
	"api-test/repository"
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	validate           *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository, DB *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{CategoryRepository: categoryRepository, DB: DB, validate: validate}
}

func (c *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryRequest) web.CategoryResponse {

	err := c.validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := c.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Id:   0,
		Name: request.Name,
	}

	category = c.CategoryRepository.Save(ctx, tx, category)
	return helper.ToCategoryResponse(category)

}

func (c *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	err := c.validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := c.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Id:   request.Id,
		Name: request.Name,
	}

	category = c.CategoryRepository.Update(ctx, tx, category)

	return helper.ToCategoryResponse(category)

}

func (c *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	tx, err := c.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := c.CategoryRepository.FindId(ctx, tx, categoryId)
	helper.PanicIfError(err)

	c.CategoryRepository.Delete(ctx, tx, category)
}

func (c *CategoryServiceImpl) FindId(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := c.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := c.CategoryRepository.FindId(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToCategoryResponse(category)
}

func (c *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, err := c.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories := c.CategoryRepository.FindAll(ctx, tx)

	return helper.ToCategoryResponses(categories)

}
