package main

import (
	"github.com/alufhigi/http-server/app"
	"github.com/alufhigi/http-server/db"
	"github.com/alufhigi/http-server/utils"
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
