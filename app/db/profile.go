package db

import "log"

func (r *DB) CrateTableProfile() error {
	_, err := r.Db.Exec(`
		CREATE TABLE IF NOT EXISTS profiles (
			pk INTEGER AUTOINCREMENT UNIQUE ,
			uuid TEXT NOT NULL UNIQUE,
			user_uuid TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP ,
			deleted_at TIMESTAMP,
			PRIMARY KEY (pk,uuid,user_uuid) ,
			FOREIGN KEY (user_uuid) REFERENCES users(uuid)  ON UPDATE SET NULL ON DELETE SET NULL
		)`)
	if err != nil {
		log.Printf("%s\n", err)
	}
	return nil
}
