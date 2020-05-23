package main

import (
	"TestMS/product-api/handlers"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main()  {
	log.Println("http server running")
	l := log.New(os.Stdout, "product", log.LstdFlags)
	ph := handlers.NewProducts(l)
	sm := mux.NewRouter()
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", ph.GetProducts)
	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProduct)
	putRouter.Use(ph.MiddlewarValidProduct)
	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", ph.AddProduct)
	postRouter.Use(ph.MiddlewarValidProduct)

	server := http.Server{
		Addr:              ":8080",
		Handler:           sm,
		TLSConfig:         nil,
		ReadTimeout:       1 * time.Second,
		ReadHeaderTimeout: 1 * time.Second,
		WriteTimeout:      1 * time.Second,
		IdleTimeout:       5 * time.Second,
	}


	go func() {
		err := server.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <- sigChan
	l.Println("Receive terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 3 * time.Second)
	server.Shutdown(tc)

	fmt.Println("sdf")
}
