package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alufhigi/http-server/app"
	"github.com/alufhigi/http-server/utils"
	"github.com/google/uuid"
)

//go test -v -run  Test_App_Router -count=1
func Test_App_Router(t *testing.T) {
	app, err := app.New(app.Config{
		Port: utils.Config("PORT"),
		DB:   utils.Config("DB_URL"),
	})
	if err != nil {
		panic(err)
	}
	app.Routers()
	// app.Run()
	testHomePage(t, app)
	testRegister(t, app)

}

func testHomePage(t *testing.T, app *app.Server) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	app.Index(w, req)
	res := w.Result()
	defer res.Body.Close()
	_, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}

}

func testRegister(t *testing.T, app *app.Server) {
	jsonValue, _ := json.Marshal(map[string]interface{}{"email": uuid.New().String() + "@gmail.com", "password": "Pass@123"})
	req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	app.Register(w, req)
	res := w.Result()
	defer res.Body.Close()
	_, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)

	}
}
