package main

import (
	"github.com/ballweera/refactor-go-shopping/controller"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/home", controller.Home)
	r.HandleFunc("/products/all", controller.GetAllProducts)
	srv := &http.Server{
		Addr:         ":9090",
		Handler:      r,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
