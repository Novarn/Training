package main

import (
	"database/sql"
	"encoding/json"
	_ "encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"strconv"
)


	
	type qwerty struct {
		id   int
		Name string
	}

	
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

		db, name := connect(id)
		defer db.Close()
		json.NewEncoder(w).Encode(name)

}
	func connect(id	int) (*sql.DB, []qwerty) {
		connStr := "user=sa password=sa dbname=testdb sslmode=disable"
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			panic(err)
		}

		rows, err := db.Query("select * from t_itfb WHERE id = $1", id)
		if err != nil {
			panic(err)
		}
		defer rows.Close()
		var Name []qwerty

		for rows.Next() {
			p := qwerty{}
			err := rows.Scan(&p.id, &p.Name)
			if err != nil {
				fmt.Println(err)
				continue
			}
			Name = append(Name, p)
		}
		//Запихнул все в переменную Name
		for _, p := range Name {
			fmt.Println(p.id, p.Name)

		}
		return db, Name
	}
