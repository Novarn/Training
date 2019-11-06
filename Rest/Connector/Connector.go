package Connector

import (
	"database/sql"
	"fmt"
)

type qwerty struct {
	id   int
	Name string
}

func Connect(id	int) (*sql.DB, []qwerty) {
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
