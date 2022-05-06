package main

import (
	"strings"
	"sync"

	"github.com/alufhigi/http-server/app"
	"github.com/alufhigi/http-server/utils"
)

func main() {
	var wg sync.WaitGroup
	ports := strings.Split(utils.Config("PORT"), ",")
	dbs := strings.Split(utils.Config("DB_PATH"), ",")
	c := []app.Config{}
	for i := 0; i < len(ports); i++ {
		c = append(c, app.Config{
			Port: ports[i],
			DB:   "./storage/" + dbs[i],
		})
	}
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
