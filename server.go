package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// git init
// git add server.go go.mod
// git commit -m "[N] Initial project"
func main() {
	fmt.Println("hello test: I'm Gopher")

	r := mux.NewRouter()
	r.HandleFunc("/pair-device", PairDeviceHandler).Methods(http.MethodPost)

	server := http.Server{
		Addr:    "127.0.0.1:2009",
		Handler: r,
	}

	log.Println("starting.....")
	log.Fatal(server.ListenAndServe())
}

func PairDeviceHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"status":"active"}`))
}
