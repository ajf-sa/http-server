package db

import "log"

func (r *DB) CrateTableProfile() error {
	_, err := r.Db.Exec(`
		CREATE TABLE IF NOT EXISTS profiles (
			pk INTEGER auto_increment ,
			uuid TEXT NOT NULL,
			user_uuid TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP ,
			deleted_at TIMESTAMP,
			CONSTRAINT pk_profiles PRIMARY KEY (pk,uuid)
			CONSTRAINT fk_column FOREIGN KEY (user_uuid) REFERENCES users(uuid)
		)`)
	if err != nil {
		log.Printf("%s\n", err)
	}
	return nil
}
