package db

import (
	"errors"
	"log"

	"github.com/alufhigi/http-server/utils"
	"github.com/google/uuid"
)

func (r *DB) CreateTableUser() error {
	_, err := r.Db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			pk INTEGER auto_increment ,
			uuid TEXT NOT NULL ,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP ,
			deleted_at TIMESTAMP ,
			email TEXT NOT NULL,
			password TEXT NOT NULL,
			name TEXT NOT NULL,
			is_admin BOOLEAN NOT NULL DEFAULT FALSE,
			CONSTRAINT pk_users PRIMARY KEY (pk,uuid)
		)
	`)
	if err != nil {
		log.Printf("%s\n", err)
	}
	return nil
}
func (r *DB) CreateUser(u *User) error {
	if r.IsUser(u.Email) {
		return errors.New("User already exists")
	}
	u.UUID = uuid.New().String()
	sqlStmt := `insert into users (uuid,email,password,name) values ($1,$2,$3,$4)`
	_, err := r.Db.Exec(sqlStmt, u.UUID, u.Email, u.Password, u.Name)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return err
	}
	return nil
}
func (r *DB) FindOneUserByID(uuid string) (*User, error) {
	var u User

	sqlStmt := `SELECT uuid,email,password,name,is_admin FROM users WHERE uuid=$1`
	err := r.Db.QueryRow(sqlStmt, uuid).Scan(&u.UUID, &u.Email, &u.Password, &u.Name, &u.IsAdmin)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return nil, err
	}
	return &u, nil
}

func (r *DB) FindOneUserByEmail(e utils.Email) (*User, error) {
	u := new(User)
	sqlStmt := `select uuid,email,password,name,is_admin from users where email = $1`
	row := r.Db.QueryRow(sqlStmt, e)
	err := row.Scan(&u.UUID, &u.Email, &u.Password, &u.Name, &u.IsAdmin)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return nil, err
	}
	return u, nil
}

func (r *DB) FindAllUser(p *Pagination) ([]User, error) {
	var users []User

	sqlStmt := `select uuid,email,name,is_admin from users order by pk desc limit $1 offset $2`
	rows, err := r.Db.Query(sqlStmt, p.Limit, p.Page)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var u User
		err := rows.Scan(&u.UUID, &u.Email, &u.Name, &u.IsAdmin)
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

func (r *DB) DeleteUser(uuid string) error {
	sqlStmt := `delete from users where uuid=$1`
	_, err := r.Db.Exec(sqlStmt, uuid)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return err
	}
	return nil
}

func (r *DB) UpdateUserEmail(uuid string, e utils.Email) error {
	sqlStmt := `update users set email=$1 where uuid=$2`
	_, err := r.Db.Exec(sqlStmt, e, uuid)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return err
	}
	return nil
}

func (r *DB) UpdateUserName(u *User) error {
	sqlStmt := `update users set name=$1 where uuid=$2`
	_, err := r.Db.Exec(sqlStmt, u.Name, u.UUID)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return err
	}
	return nil
}
func (r *DB) UpdateUserPassword(u *User) error {
	sqlStmt := `update users set password=$1 where uuid=$2`
	_, err := r.Db.Exec(sqlStmt, u.Password, u.UUID)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return err
	}
	return nil
}

func (r *DB) UpdateUserAdmin(u *User) error {
	sqlStmt := `update users set is_admin=$1 where uuid=$2`
	_, err := r.Db.Exec(sqlStmt, u.IsAdmin, u.UUID)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return err
	}
	return nil
}
