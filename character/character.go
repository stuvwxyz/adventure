package character

import (
	"fmt"
)

type Character struct {
	ID	int
	Name string
	Age int
	Gender string
	Class string
	Race string
	Location string
	HitPoints	int
}

var Player Character

func GenerateCharacter() {
	fmt.Print("What shall you be known as?: ")
	_, _ = fmt.Scan(&Player.Name)

	fmt.Print("How old ist thou?: ")
	_, _ = fmt.Scan(&Player.Age)

	fmt.Print("Wouldst thou like to play male, female, non-binary?: ")
	_, _ = fmt.Scan(&Player.Gender)

	fmt.Print("What class shall thee be?: ")
	_, _ = fmt.Scan(&Player.Class)

	fmt.Print("What race shall thee be?: ")
	_, _ = fmt.Scan(&Player.Race)


	fmt.Println(Player.Name)
	fmt.Println(Player.Age)
	fmt.Println(Player.Gender)
	fmt.Println(Player.Class)
}

