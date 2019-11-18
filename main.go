package main

import (
	"log"
	"net/http"
	"time"

	"github.com/ballweera/refactor-go-shopping/controller"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/home", controller.Home)

	productCtrl := controller.NewProductController()
	r.HandleFunc("/products/all", productCtrl.GetAllProducts)
	srv := &http.Server{
		Addr:         ":9090",
		Handler:      r,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
