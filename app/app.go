package main

import (
	"sync"

	"github.com/alufhigi/http-server/app"
)

func main() {
	var wg sync.WaitGroup
	c := []app.Config{}
	c = append(c, app.Config{Port: "5555", DB: "./storage/app.db"})
	c = append(c, app.Config{Port: "5556", DB: "./storage/app1.db"})
	c = append(c, app.Config{Port: "5557", DB: "./storage/app2.db"})
	c = append(c, app.Config{Port: "5558", DB: "./storage/app3.db"})
	c = append(c, app.Config{Port: "5559", DB: "./storage/app4.db"})
	for _, v := range c {
		wg.Add(len(c))
		go func(v app.Config) {

			app := Serve(v)
			app.Run()
		}(v)
	}
	defer wg.Done()
	wg.Wait()
}

func Serve(c app.Config) *app.Server {
	app, err := app.New(c)
	if err != nil {
		panic(err)
	}
	app.Routers()
	return app

}
