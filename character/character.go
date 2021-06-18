package character

import (
	"fmt"
	"math/rand"
	"time"
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
	Strength	int
	Dexterity	int
	Constitution int
	Intelligence int
	Wisdom		int
	Charisma	int
}

var Player Character

func GenerateCharacter() {
	fmt.Print("What shall you be known as?: ")
	_, _ = fmt.Scan(&Player.Name)

	fmt.Print("How old ist ", Player.Name, "?: ")
	_, _ = fmt.Scan(&Player.Age)

	fmt.Print("Wouldst thou like to play male, female, non-binary?: ")
	_, _ = fmt.Scan(&Player.Gender)

	fmt.Print("What class shall thee be?: ")
	_, _ = fmt.Scan(&Player.Class)

	fmt.Print("What race shall thee be?: ")
	_, _ = fmt.Scan(&Player.Race)

	rand.Seed(time.Now().UnixNano())

	Player.HitPoints = randomInt(2, 10)
	Player.Strength = (randomInt(1, 6) + randomInt(1,6) + randomInt(1, 6))
	Player.Dexterity = (randomInt(1, 6) + randomInt(1,6) + randomInt(1, 6))
	Player.Constitution = (randomInt(1, 6) + randomInt(1,6) + randomInt(1, 6))
	Player.Intelligence = (randomInt(1, 6) + randomInt(1,6) + randomInt(1, 6))
	Player.Wisdom = (randomInt(1, 6) + randomInt(1,6) + randomInt(1, 6))
	Player.Charisma = (randomInt(1, 6) + randomInt(1,6) + randomInt(1, 6))

	DisplayCharacter()
}

func DisplayCharacter() {
	fmt.Println("\n\n\nHello, ", Player.Name)
	fmt.Println("You are ", Player.Age, " years old.")
	fmt.Println("You identify as ", Player.Gender)
	fmt.Println("Class: ", Player.Class)
	fmt.Println("You have ", Player.HitPoints, "hit points.")
	fmt.Println("\nYour statistics are:")
	fmt.Println("Strength:", Player.Strength )
	fmt.Println("Dexterity:", Player.Dexterity )
	fmt.Println("Constitution:", Player.Constitution )
	fmt.Println("Intelligence:", Player.Intelligence )
	fmt.Println("Wisdom:", Player.Wisdom )
	fmt.Println("Charisma:", Player.Charisma )
}

func InventoryCharacter() {
	fmt.Println("\n\nYou are currently carring nothing")
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

