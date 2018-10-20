package data

import (
	"database/sql"
	"fmt"
)

func CreateCharactersTable(db *sql.DB) {
	//fmt.Println("Building Locations for your adventure")

	sqlStatement := `CREATE TABLE IF NOT EXISTS characters (
		id SERIAL,
		name VARCHAR(255)PRIMARY KEY,
		age INTEGER,
		gender VARCHAR(255),
		class VARCHAR(255),
		race VARCHAR(255),
		location VARCHAR(255),
		hit_points VARCHAR(255),
		strength VARCHAR(255),
		dexterity VARCHAR(255),
		constitution VARCHAR(255),
		intelligence VARCHAR(255),
		wisdom VARCHAR(255),
		charisma VARCHAR(255));`

	_, err := db.Exec(sqlStatement)
	if err != nil {
		fmt.Println("Unable to create locations table")
		panic(err)
	}
	//fmt.Println("Locations table created")
}
