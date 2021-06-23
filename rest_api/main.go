package main

import (
	"fmt"
	"log"
	"net/http"
	"rest_api/handlers"
	"rest_api/server"
)

func main() {
	mux := defaultMux()
	server := server.NewServer()
	mux.HandleFunc("/province/", handlers.Server.provinceHandler)
	log.Fatal(http.ListenAndServe("localhost:8088", mux))
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
