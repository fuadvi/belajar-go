package main

import (
	"belajar-resfull-api/helper"
	"belajar-resfull-api/middleware"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:8080",
		Handler: authMiddleware,
	}
}

func main() {
	//db := app.NewDb()
	//validate := validator.New()
	//categoryRepository := repository.NewCategoryRepository()
	//categoryService := service.NewCategoryService(categoryRepository, db, validate)
	//categoryController := controller.NewCategoryController(categoryService)
	//
	//router := app.NewRouter(categoryController)
	//authMiddleWare := middleware.NewAuthMiddleware(router)
	//server := NewServer(authMiddleWare)

	server := InitializedServer()

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
