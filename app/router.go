package app

import (
	"api-test/controller"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(categoryController controller.CategoryController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindId)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	return router

}
