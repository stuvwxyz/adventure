package data

import (
	"Adventure/database"
	"database/sql"
	"fmt"
)

func CreateLocationsTable(db *sql.DB) {
	fmt.Println("Building Locations for your adventure")

	sqlStatement := `CREATE TABLE IF NOT EXISTS locations (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255),
		description TEXT,
		n VARCHAR(255),
		s VARCHAR(255),
		e VARCHAR(255),
		w VARCHAR(255),
		u VARCHAR(255),
		d VARCHAR(255));`

	_, err := db.Exec(sqlStatement)
	if err != nil {
		fmt.Println("Unable to create locations table")
		panic(err)
	}
	fmt.Println("Locations table created")
}

func CreateLocations(db *sql.DB) {
	database.LocationAdd(db, "Town Square",
		"You are in the Town Square. This is a wide open area where people gather.",
		"Open Bazar", "Cobblestone Road", "Dirt Road Into Woods", "Path Between Buildings", "", "")

	database.LocationAdd(db, "Open Bazar",
		"You are in an open air bazar. There are shops inside small tents and vendors selling their wares..",
		"", "Town Square", "", "", "", "")

	database.LocationAdd(db, "Cobblestone Road",
		"You head south of the Town Square following a cobble stone road",
		"Town Square", "Cobblestone Road2", "", "", "", "")
	database.LocationAdd(db, "Cobblestone Road2",
		"Continuing south on the cobble stone road you ponder what is a cobble and how it becomes stone",
		"Cobblestone Road", "Cobblestone Road3", "", "", "", "")
	database.LocationAdd(db, "Cobblestone Road3",
		"The cobble stone road branches to the east and you see a wall with rod iron gates in it.",
		"Cobblestone Road2", "", "Cobblestone RoadE", "", "", "")
	database.LocationAdd(db, "Cobblestone RoadE",
		"The cobble stone road runs up to a wall with rod iron gates in it.  Inside you can see a large cemetary.",
		"", "", "Cemetary", "Cobblestone Road3", "", "")
	database.LocationAdd(db, "Cemetary",
		"You enter a dark and creapy cemetery.  You notice the path is overgrown with roots as you try not to stumble.  There is a sturdy stone wall in front of you.",
		"", "", "", "Cobblestone RoadE", "Wall", "")
	database.LocationAdd(db, "Wall",
		"You climb the sturdy wall and look around.  Over the wall you notice many zombies milling about and trying to climb up.  The wall is too solid and tall but you tremble to think what will happen if they ever get over it.",
		"", "", "", "", "", "Cemetary")

	database.LocationAdd(db, "Dirt Road Into Woods",
		"East of Town Square the dirt road heads off towards some woods",
		"", "", "Dirt Road Into Woods2", "Town Square", "", "")
	database.LocationAdd(db, "Dirt Road Into Woods2",
		"You continue off into the woods and notice a small grove to the south",
		"", "Grove", "Dirt Road Into Woods3", "Dirt Road Into Woods", "", "")
	database.LocationAdd(db, "Grove",
		"This is a well tended grove in the woods.  You feel at peace and your body feals healther.",
		"Dirt Road Into Woods2", "", "", "", "", "")
	database.LocationAdd(db, "Dirt Road Into Woods3",
		"You continue to travel throuth the woods and notice a curve that leads south",
		"", "", "Curve", "Dirt Road Into Woods2", "", "")
	database.LocationAdd(db, "Curve",
		"The dirt path curves south into heavier woods.  It appears dark and difficult to continue",
		"", "Heavy Woods", "", "Dirt Road Into Woods3", "", "")
	database.LocationAdd(db, "Heavy Woods",
		"The woods thicken to the point where you can hardly continue.  There are fallen trees in the little path that had existed",
		"Curve", "", "", "", "Fallen Trees", "")
	database.LocationAdd(db, "Fallen Trees",
		"You climb the trees and peer over.  It appears you can go East into the DARK, DARK woods.",
		"", "", "DARK, DARK Woods", "", "", "Heavy Woods")
	database.LocationAdd(db, "DARK, DARK Woods",
		"You enter a DARK, DARK woods.  A DARK, DARK path continues east",
		"", "", "DARK, DARK path", "Fallen Trees", "", "")
	database.LocationAdd(db, "DARK, DARK path",
		"You walk down a DARK, DARK path.  There is a light glow to the east",
		"", "", "Towards the brightly lit hut", "DARK, DARK Woods", "", "")
	database.LocationAdd(db, "Towards the brightly lit hut",
		"You follow the DARK, DARK path to a brightly lit hut.  There is a hill to the south",
		"The Hut", "Towards the hill", "", "DARK, DARK path", "", "")
	database.LocationAdd(db, "The Hut",
		"You enter the hut.  There is a someone throwing ingredients into a cauldron and chanting quietly. \n She says, 'Enter and close the door, Dinner is just about ready'",
		"", "Towards the brightly lit hut", "", "", "", "")
	database.LocationAdd(db, "Towards the hill",
		"As you move towards the hill you see the full moon come out of the clouds.  \nYou hear a howl in the distance and decide there is no reason to continue into the DARK, DARK woods in the DARK, DARK night.",
		"Towards the brightly lit hut", "", "", "", "", "")


	database.LocationAdd(db, "Path Between Buildings",
		"You squeeze between the buildings.  There is little light and it feels like something with lots of legs is climbing across your cheek",
		"", "", "Town Square", "", "", "")
}