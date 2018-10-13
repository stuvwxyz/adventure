package database

import (
	"database/sql"
	"fmt"
)

type Location struct {
	ID			int
	Name		string
	Description string
	N  			string
	S     		string
	E			string
	W			string
	U			string
	D			string
}

func LocationAdd(db *sql.DB, name string, description string, n string, s string, e string,
	w string, u string, d string) {
	sqlStatement := `
		INSERT INTO locations (name, description, n, s, e, w, u, d)
		VALUES ($1, $2, $3, $4 ,$5, $6, $7, $8)
		RETURNING id`
	id := 0
	err := db.QueryRow(sqlStatement, name, description, n, s, e, w, u, d).Scan(&id)
	if err != nil {
		fmt.Println("Location not added")
		panic(err)
	}
	fmt.Println("New record is: ", id)
}

func LocationGet(db *sql.DB, name string) Location {
	sqlStatement := `SELECT * from locations WHERE name=$1`
	var location Location
	row := db.QueryRow(sqlStatement, name)
	err := row.Scan(&location.ID, &location.Name, &location.Description, &location.N, &location.S,
		&location.E, &location.W, &location.U, &location.D)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were selected for ", name)
	case nil:
		return location
	default:
		panic(err)
	}
	return location
}
