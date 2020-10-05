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

	"go.uber.org/zap"
)

// git init
// git add server.go go.mod
// git commit -m "[N] Initial project"
func main() {
	fmt.Println("hello test: I'm Gopher")

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	create := createPairDevice(db)
	r.Handle("/pair-device", PairDeviceHandler(create)).Methods(http.MethodPost)

	addr := fmt.Sprintf("0.0.0.0:%s", os.Getenv("PORT"))
	fmt.Println("test addr:", addr)
	server := http.Server{
		Addr:    addr,
		Handler: r,
	}

	log.Println("starting.....")
	log.Fatal(server.ListenAndServe())
}

type Pair struct {
	DeviceID int64
	UserID   int64
}

func PairDeviceHandler(device Device) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var p Pair
		err := json.NewDecoder(r.Body).Decode(&p)

		l := zap.NewExample()
		l = l.With(zap.Namespace("hometic"), zap.String("I'm goher", "fdskfkdj"))
		l.Info("test")

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err.Error())
			return
		}

		defer r.Body.Close()

		fmt.Printf("pair: %#v\n", p)

		err = device.Pair(p)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(err.Error())
			return
		}
		w.Write([]byte(`{"status":"active"}`))
	}
}

type Device interface {
	Pair(p Pair) error
}

type createPairDeviceFunc func(p Pair) error

func (fn createPairDeviceFunc) Pair(p Pair) error {
	return fn(p)
}
func createPairDevice(db *sql.DB) createPairDeviceFunc {
	return func(p Pair) error {
		_, err := db.Exec("INSERT INTO pair VALUES ($1,$2)", p.DeviceID, p.UserID)

		return err
	}
}
