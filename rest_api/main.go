package main

import (
	"fmt"
	"log"
	"net/http"
	"rest_api/handlers"
	"rest_api/provinces"
	"rest_api/redirector"
)

type Server struct {
	store *provinces.ProvinceStore
}

func NewServer() *Server {
	store := provinces.New()
	return &Server{store: store}
}

/*TO DO
add geographical and political handling
add autentification
add YAML
add JSON
add BoltDB
add flag
add middleware
add GraphQL
*/

func main() {
	mux := defaultMux()
	pathsToUrls := map[string]string{
		"/province/": "/province/economic/",
	}
	server := handlers.NewServer()
	mux.HandleFunc("/province/economic/", server.EconomicProvinceHandler)
	mapHandler := redirector.MapHandler(pathsToUrls, mux)
	log.Fatal(http.ListenAndServe("localhost:8088", mapHandler))
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
