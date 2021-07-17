package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net/http"
	"rest_api/auth"
	"rest_api/handlers"
	"rest_api/middleware"
	"rest_api/provinces"
	"rest_api/redirector"

	"github.com/boltdb/bolt"
)

type Server struct {
	store *provinces.ProvinceStore
}

func NewServer() *Server {
	store := provinces.New()
	return &Server{store: store}
}

var (
	db *bolt.DB
)

/*TO DO
add GraphQL
*/

func main() {
	defer db.Close()
	auth.SetCredentials()

	addr := flag.String("port", ":4001", "port server running")
	certFile := flag.String("cert", "cert.pem", "cert file")
	keyFile := flag.String("key", "key.pem", "key.pem file")
	yml := flag.String("yml", "provinces.yml", "path to yaml file with redirection instructions")

	mux := defaultMux()
	pathsToUrls := map[string]string{
		"/province/": "/province/economic/",
	}
	server := handlers.NewServer()
	mux.HandleFunc("/province/economic/", server.EconomicProvinceHandler)

	midHandler := middleware.Logging(mux)
	midHandler = middleware.PanicRecovery(midHandler)
	mapHandler := redirector.MapHandler(pathsToUrls, midHandler)
	yamlHandler, err := redirector.YAMLHandler(*yml, mapHandler)
	if err != nil {
		panic(err)
	}

	srv := &http.Server{
		Addr:    *addr,
		Handler: yamlHandler,
		TLSConfig: &tls.Config{
			MinVersion:               tls.VersionTLS13,
			PreferServerCipherSuites: true,
		},
	}

	log.Printf("Starting server on %s", *addr)
	srv_err := srv.ListenAndServeTLS(*certFile, *keyFile)
	log.Fatal(srv_err)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
