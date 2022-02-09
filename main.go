package main

import (
	"api-test/app"
	"api-test/controller"
	"api-test/exception"
	"api-test/helper"
	"api-test/repository"
	"api-test/services"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var (
	validate *validator.Validate
)

func main() {

	DB := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := services.NewCategoryService(categoryRepository, DB, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
