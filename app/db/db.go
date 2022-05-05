package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	Db *sql.DB
}

func New(driverName, dataSourceName string) (*DB, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	r := &DB{db}
	_, err = db.Exec("PRAGMA foreign_keys=ON;")
	if err != nil {
		return nil, err
	}
	var enabled bool
	err = db.QueryRow("PRAGMA foreign_keys;").Scan(&enabled)
	if err != nil {
		return nil, err
	}

	r.CreateTableUser()
	r.CrateTableProfile()
	r.CreateTableClient()
	r.CreateTableRole()
	r.CreateTAblePermission()
	r.CreateTableRolePermission()

	return r, nil
}
