package db

import (
	"errors"
	"log"

	"github.com/alufhigi/http-server/utils"
)

func (r *DB) CreateTableUser() error {
	_, err := r.Db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY,
			email TEXT NOT NULL,
			password TEXT NOT NULL,
			name TEXT NOT NULL
		)
	`)
	if err != nil {
		log.Printf("%q: %s\n", err)
	}
	return nil
}
func (r *DB) CreateUser(u *User) error {
	if r.IsUser(u.Email) {
		return errors.New("User already exists")
	}
	sqlStmt := `insert into users (email,password,name) values ($1,$2,$3)`
	_, err := r.Db.Exec(sqlStmt, u.Email, u.Password, u.Name)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return err
	}
	return nil
}
func (r *DB) FindOneUserByID(id int) (*User, error) {
	var u User
	sqlStmt := `SELECT id,email,password,name FROM users WHERE id=$1`
	err := r.Db.QueryRow(sqlStmt, id).Scan(&u.Id, &u.Email, &u.Password, &u.Name)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return nil, err
	}
	return &u, nil
}

func (r *DB) FindOneUserByEmail(e utils.Email) (*User, error) {
	u := new(User)
	sqlStmt := `select id,email,password,name from users where email = $1`
	row := r.Db.QueryRow(sqlStmt, e)
	err := row.Scan(&u.Id, &u.Email, &u.Password, &u.Name)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return nil, err
	}
	return u, nil
}

func (r *DB) IsUser(e utils.Email) bool {
	i := new(int)
	sqlStmt := `select 1 from users where email = $1`
	row := r.Db.QueryRow(sqlStmt, e)
	err := row.Scan(&i)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return false
	}
	return true
}
