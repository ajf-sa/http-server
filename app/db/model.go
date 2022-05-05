package db

import (
	"encoding/json"
	"time"

	"github.com/alufhigi/http-server/utils"
)

type Base struct {
	PK         int        `json:"pk"`
	UUID       string     `json:"uuid"`
	Created_at time.Time  `json:"created_at"`
	Updated_at time.Time  `json:"updated_at"`
	Deleted_at *time.Time `json:"deleted_at"`
}

type Pagination struct {
	Limit int         `json:"limit"`
	Page  int         `json:"page"`
	Sort  string      `json:"sort"`
	Rows  interface{} `json:"rows"`
}

type User struct {
	Base
	Email    utils.Email    `json:"email"`
	Password utils.Password `json:"password,omitempty"`
	Name     string         `json:"name"`
	IsAdmin  bool           `json:"is_admin"`
}

type Client struct {
	Base
	Name     string `json:"name"`
	UserUUID string `json:"user_uuid"`
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
