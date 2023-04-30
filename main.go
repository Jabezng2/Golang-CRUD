package main

import (
	"fmt"
	"golang_crud/config"
	"golang_crud/controller"
	"golang_crud/helper"
	"golang_crud/repository"
	"golang_crud/router"
	"golang_crud/service"
	"net/http"
)

func main() {
	fmt.Printf("Start server")
	// database
	db := config.DatabaseConnection()

	// repository
	bookRepository := repository.NewBookRepository(db)

	// service
	bookService := service.NewBookServiceImpl(bookRepository)

	// controller
	bookController := controller.NewBookController(bookService)

	// router
	routes := router.NewRouter(bookController)

	server := http.Server{Addr: "localhost:8888", Handler: routes}

	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
