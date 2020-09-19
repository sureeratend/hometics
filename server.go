package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("hello test: I'm Gopher")

	r := mux.NewRouter()
	r.HandleFunc("/pair-device", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"status":"active"}`))
	}).Methods(http.MethodPost)

	server := http.Server{
		Addr:    "127.0.0.1:2009",
		Handler: r,
	}

	log.Println("starting.....")
	log.Fatal(server.ListenAndServe())
}
