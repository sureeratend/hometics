package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Pair struct {
	DeviceID int64
	UserID   int64
}

// git init
// git add server.go go.mod
// git commit -m "[N] Initial project"
func main() {
	fmt.Println("hello test: I'm Gopher")

	r := mux.NewRouter()
	r.HandleFunc("/pair-device", PairDeviceHandler).Methods(http.MethodPost)

	addr := fmt.Sprintf("0.0.0.0:%s", os.Getenv("PORT"))
	fmt.Println("test addr:", addr)
	server := http.Server{
		Addr:    addr,
		Handler: r,
	}

	log.Println("starting.....")
	log.Fatal(server.ListenAndServe())
}

func PairDeviceHandler(w http.ResponseWriter, r *http.Request) {
	var p Pair
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		fmt.Printf("fdsjkfghsd\n")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	defer r.Body.Close()
	fmt.Printf("pair: %#v\n", p)
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	//db, err := sql.Open("postgres", "postgres://cfpbmgmh:4RM8d4XhNM9zD3GqjSPp5K9e7REh7STF@satao.db.elephantsql.com:5432/cfpbmgmh")

	if err != nil {
		log.Fatal("connetc to database fail error", err, os.Getenv("DATABASE_URL"))
	}
	defer db.Close()

	createTb := `CREATE TABLE IF NOT EXISTS pair( DEVICE_ID INTEGER NOT NULL, USER_ID INTEGER NOT NULL);`

	_, err = db.Exec(createTb)
	if err != nil {
		log.Fatal("cant 't connect to database err:", err)
	}
	fmt.Println("Create database success")

	_, err = db.Exec("INSERT INTO pairs VALUES($1,$2);", p.DeviceID, p.UserID)
	if err != nil {
		fmt.Printf("fdsjkfghsd\n")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.Write([]byte(`{"status":"active"}`))
}
