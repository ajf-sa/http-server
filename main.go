package main

import "github.com/alufhigi/netServer/app"

func main() {
	app := app.New()
	app.Routers()
	app.Run()

}
