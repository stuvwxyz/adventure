# Adventure game

### Setup
1. Use Docker Postgres DB for game  
    1. docker run -p 5432:5432 --name post-go -e POSTGRES_PASSWORD=g0Password -e POSTGRES_DB=post-go -d postgres
    1. When programs runs it will use this DB and PW to create the tables required and start the game.
1. git clone git@gitlab.com:stuvwxyz/adventure.git into Go Path directory
1. go run main.go


#### Plans for improvements 
1. Expand Character definition (Str, Dex, Con, Int, Wis, Cha)
1. Add in Inventory
    1. Worn and carried items
    1. Monies!!!
1. Expand Locations
    1. Shops in Bazar to purchase items
1. Create Encounters


