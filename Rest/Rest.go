package main

import (
	"Rest/Connector"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

	func main() {
		router := mux.NewRouter()
		router.HandleFunc("/Name/{Name}/{id:[0-9]+}", trouble)
		http.Handle("/", router)

		fmt.Println("Server is listening...")
		http.ListenAndServe(":8181", nil)

	} 

func trouble(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println("Get info about id #", id)
	}

		db, name := Connector.Connect(id)
		defer db.Close()

		json.NewEncoder(w).Encode(name)

}



