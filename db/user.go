package db

import "log"

func (r *DB) CreateUser(u *User) error {
	sqlStmt := `insert into users (name) values ($1)`
	_, err := r.Db.Exec(sqlStmt, u.Name)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return err
	}
	return nil
}
