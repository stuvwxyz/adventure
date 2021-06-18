package actions

import (
	"Adventure/character"
	"Adventure/database"
	"database/sql"
	"fmt"
)

func DisplayLocation(location database.Location) {
	fmt.Println("\n\n\n")
	fmt.Println(location.Description)
	fmt.Println("Directions available are: ")
	if location.N != "" {
		fmt.Println("N - ", location.N)
	}
	if location.S != "" {
		fmt.Println("S - ", location.S)
	}
	if location.E != "" {
		fmt.Println("E - ", location.E)
	}
	if location.W != "" {
		fmt.Println("W - ", location.W)
	}
	if location.U != "" {
		fmt.Println("U - ", location.U)
	}
	if location.D != "" {
		fmt.Println("D - ", location.D)
	}
	fmt.Println("I - Inventory, C - Character Info, G - Save Game")
}

func MoveLocations(db *sql.DB, move string, location database.Location) database.Location {
	NewName := location.Name

	if move == "N" && location.N != "" {
		fmt.Println("\n\nYou travel North")
		NewName = location.N
	} else if move == "S" && location.S != "" {
		fmt.Println("\n\nYou travel South")
		NewName = location.S
	} else if move == "E" && location.E != "" {
		fmt.Println("\n\nYou travel East")
		NewName = location.E
	} else if move == "W" && location.W != "" {
		fmt.Println("\n\nYou travel West")
		NewName = location.W
	} else if move == "U" && location.U != "" {
		fmt.Println("\n\nYou travel Up")
		NewName = location.U
	} else if move == "D" && location.D != "" {
		fmt.Println("\n\nYou travel Down")
		NewName = location.D
	} else if (move == "I") {
		character.InventoryCharacter()
	} else if (move == "C") {
		character.DisplayCharacter()
	} else if (move == "G") {
		database.CharacterSave(db)
	} else {
		fmt.Println("\n\n\n\nThat choice is invalid, please choose again.")
	}

	newLocation := database.LocationGet(db, NewName)
	return newLocation
}