package main

import (
	"fmt"
	"golang_crud/helper"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	fmt.Printf("Start server")

	routes := httprouter.New()

	server := http.Server{Addr: "localhost:8888", Handler: routes}

	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
