package main

import (
	"api-test/app"
	"api-test/controller"
	"api-test/exception"
	"api-test/helper"
	"api-test/repository"
	"api-test/services"
	"net/http"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

var (
	validate *validator.Validate
)

func main() {
	// todo: logrus for logs
	file, _ := os.OpenFile("storages/logs/application-"+time.Now().Format("02-01-2006")+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logrus.SetOutput(file)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.DebugLevel)

	//  todo: init
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
