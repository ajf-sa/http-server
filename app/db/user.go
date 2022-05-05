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
			name TEXT NOT NULL,
			is_admin BOOLEAN NOT NULL DEFAULT FALSE
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

	sqlStmt := `SELECT id,email,password,name,is_admin FROM users WHERE id=$1`
	err := r.Db.QueryRow(sqlStmt, id).Scan(&u.ID, &u.Email, &u.Password, &u.Name, &u.IsAdmin)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return nil, err
	}
	return &u, nil
}

func (r *DB) FindOneUserByEmail(e utils.Email) (*User, error) {
	u := new(User)
	sqlStmt := `select id,email,password,name,is_admin from users where email = $1`
	row := r.Db.QueryRow(sqlStmt, e)
	err := row.Scan(&u.ID, &u.Email, &u.Password, &u.Name, &u.IsAdmin)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return nil, err
	}
	return u, nil
}

func (r *DB) FindAllUser(p *Pagination) ([]User, error) {
	var users []User
	sqlStmt := `select id,email,name,is_admin from users order by id desc limit $1 offset $2`
	rows, err := r.Db.Query(sqlStmt, p.Limit, p.Page)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var u User
		err := rows.Scan(&u.ID, &u.Email, &u.Name, &u.IsAdmin)
		if err != nil {
			log.Printf("%q: %s\n", err, sqlStmt)
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
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

func (r *DB) DeleteUser(id int) error {
	sqlStmt := `delete from users where id=$1`
	_, err := r.Db.Exec(sqlStmt, id)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return err
	}
	return nil
}

func (r *DB) UpdateUserEmail(id int, e utils.Email) error {
	sqlStmt := `update users set email=$1 where id=$2`
	_, err := r.Db.Exec(sqlStmt, e, id)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return err
	}
	return nil
}

func (r *DB) UpdateUserName(u *User) error {
	sqlStmt := `update users set name=$1 where id=$2`
	_, err := r.Db.Exec(sqlStmt, u.Name, u.ID)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return err
	}
	return nil
}
func (r *DB) UpdateUserPassword(u *User) error {
	sqlStmt := `update users set password=$1 where id=$2`
	_, err := r.Db.Exec(sqlStmt, u.Password, u.ID)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return err
	}
	return nil
}

func (r *DB) UpdateUserAdmin(u *User) error {
	sqlStmt := `update users set is_admin=$1 where id=$2`
	_, err := r.Db.Exec(sqlStmt, u.IsAdmin, u.ID)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return err
	}
	return nil
}
