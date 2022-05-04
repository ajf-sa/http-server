package main

import (
	"github.com/alufhigi/netServer/app"
	"github.com/alufhigi/netServer/db"
)

func main() {
	db, err := db.New("sqlite3", "./app.db")
	if err != nil {
		panic(err)
	}
	defer db.Db.Close()
	app := app.New(db)
	app.Routers()
	app.Run()

}
