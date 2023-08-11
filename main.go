package main

import (
	"arieansyah/golang-restful-api/app"
	"arieansyah/golang-restful-api/controller"
	"arieansyah/golang-restful-api/exception"
	"arieansyah/golang-restful-api/helper"
	"arieansyah/golang-restful-api/middleware"
	"arieansyah/golang-restful-api/repository"
	"arieansyah/golang-restful-api/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db := app.NewDB()
	validate := validator.New()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
