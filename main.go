package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/ballweera/refactor-go-shopping/controller"
	"github.com/ballweera/refactor-go-shopping/product"
	"github.com/ballweera/refactor-go-shopping/user"
	"github.com/gorilla/mux"
)

func main() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connection to finish - e.g. 15s or 1m")
	flag.Parse()

	r := mux.NewRouter()
	r.HandleFunc("/home", controller.Home)

	productCtrl := product.NewProductController()
	r.HandleFunc("/products/all", productCtrl.GetAllProducts)

	userCtrl := user.NewUserController()
	r.HandleFunc("/register", userCtrl.Register)
	srv := &http.Server{
		Addr:         ":9090",
		Handler:      r,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  1 * time.Minute,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	srv.Shutdown(ctx)
	log.Println("Shutting down...")
	os.Exit(0)
}
