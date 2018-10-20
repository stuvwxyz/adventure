package main

import (
	"Adventure/actions"
	"Adventure/character"
	"Adventure/data"
	"Adventure/database"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"strings"
)

//Constants for database connection
const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "g0Password"
	dbname = "adventure"
)

// Define Global variables
var db *sql.DB

func main() {
	// Define variables
	var err error
	// Set up Database connection
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user,password, dbname)

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Error on Open")
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("Error establishing connection to DB")
		panic(err)
	}

	fmt.Println("Successfully connected to the DB")

	// Ask player for the character information
	data.CreateCharactersTable(db)
	character.GenerateCharacter()
	database.CharacterAdd(db)

	// Initialize the Locations Table
	data.CreateLocationsTable(db)
	data.CreateLocations(db)

	//Start at Town square
	CurrentLocation := database.LocationGet(db, "Town Square")

	for {
		actions.DisplayLocation(CurrentLocation)

		var Move string
		fmt.Print("Action: ")
		_, _ = fmt.Scan(&Move)

		CurrentLocation = actions.MoveLocations(db, strings.ToUpper(Move), CurrentLocation)
	}
}