package db

import (
	"encoding/json"

	"github.com/alufhigi/http-server/utils"
)

type User struct {
	Id       int            `json:"id"`
	Email    utils.Email    `json:"email"`
	Password utils.Password `json:"password"`
	Name     string         `json:"name"`
}

func (u *User) UmarshalJSON(data []byte) error {
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
