package main

import (
	"log"
	"net/http"
)


func main() {
	PORT := "3000"
	api := &Api{
		addr: ":" + PORT,
	}

	// NOTE: this is the router instance (Mux is nothing but a fancy word for a router)
	mux := http.NewServeMux()
	// NOTE: this is the server instance
	srv := &http.Server{
		Addr:    api.addr,
		Handler: mux,
	}

	mux.HandleFunc("GET /users", api.getUserHanlder)
	mux.HandleFunc("POST /users", api.createUserHanlder)

	log.Printf("listening on PORT %s", PORT)
	log.Fatal(srv.ListenAndServe())
}
