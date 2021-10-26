package controller

import (
	"api-test/helper"
	"api-test/model/web"
	"api-test/services"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	service services.CategoryService
}

func NewCategoryController(services services.CategoryService) CategoryController {
	return &CategoryControllerImpl{service: services}
}

func (c *CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {

	categoryCreateRequest := web.CategoryRequest{}
	helper.WriteToRequestBody(request, &categoryCreateRequest)

	categoryResponse := c.service.Create(request.Context(), categoryCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (c *CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {

	categoryUpdateRequest := web.CategoryUpdateRequest{}
	helper.WriteToRequestBody(request, &categoryUpdateRequest)

	categoryResponse := c.service.Update(request.Context(), categoryUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (c *CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	categoryId := param.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	c.service.Delete(request.Context(), id)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (c *CategoryControllerImpl) FindId(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	categoryId := param.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryResponse := c.service.FindId(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)

}

func (c *CategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	categoryResponse := c.service.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
