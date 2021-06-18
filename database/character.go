package database

import (
	"Adventure/character"
	"database/sql"
	"fmt"
)

func CharacterSave(db *sql.DB) {
	sqlStatement := `
		INSERT INTO characters (name, age, gender, class, race, location, hit_points, strength, dexterity, constitution, intelligence, wisdom, charisma)
		VALUES ($1, $2, $3, $4 ,$5, $6, $7, $8, $9, $10, $11, $12, $13)
		RETURNING id`
	id := 0
	err := db.QueryRow(sqlStatement, character.Player.Name, character.Player.Age, character.Player.Gender, character.Player.Class,
		character.Player.Race, character.Player.Location, character.Player.HitPoints, character.Player.Strength,
		character.Player.Dexterity, character.Player.Constitution, character.Player.Intelligence, character.Player.Wisdom,
		character.Player.Charisma).Scan(&id)
	if err != nil {
		fmt.Println("Location not added")
		panic(err)
	}
	//fmt.Println("New record is: ", id)
}
