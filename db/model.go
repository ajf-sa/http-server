package db

import (
	"encoding/json"

	"github.com/alufhigi/http-server/utils"
)

type Pagination struct {
	Limit int         `json:"limit"`
	Page  int         `json:"page"`
	Sort  string      `json:"sort"`
	Rows  interface{} `json:"rows"`
}

type User struct {
	Id       int            `json:"id"`
	Email    utils.Email    `json:"email"`
	Password utils.Password `json:"password,omitempty"`
	Name     string         `json:"name"`
}

func (u *User) UnmarshalJSON(data []byte) error {
	type user User
	var u2 user
	if err := json.Unmarshal(data, &u2); err != nil {
		return err
	}
	*u = User(u2)
	return nil
}

func (u *User) Validate() error {
	if e := u.Email.Validate(); e != nil {
		return e
	}
	if p := u.Password.Validate(); p != nil {
		return p
	}
	return nil
}
