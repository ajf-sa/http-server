package main

import (
	"github.com/alufhigi/netServer/app"
	"github.com/alufhigi/netServer/db"
	"github.com/alufhigi/netServer/utils"
)

func main() {
	db, err := db.New("sqlite3", utils.Config("DB_PATH"))
	if err != nil {
		panic(err)
	}
	defer db.Db.Close()
	app := app.New(db)
	app.Routers()
	app.Run()

}
