package db

func (r *DB) CreateTableRole() error {
	_, err := r.Db.Exec(`
		CREATE TABLE IF NOT EXISTS roles (
			pk INTEGER AUTOINCREMENT UNIQUE ,
			uuid TEXT NOT NULL UNIQUE ,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP ,
			deleted_at TIMESTAMP ,
			name TEXT,
			slug TEXT NOT NULL,
			PRIMARY KEY (pk,uuid)
		)
	`)
	if err != nil {
		return err
	}
	return nil
}

func (r *DB) CreateTAblePermission() error {
	_, err := r.Db.Exec(`
		CREATE TABLE IF NOT EXISTS permissions (
			pk INTEGER AUTOINCREMENT UNIQUE ,
			uuid TEXT NOT NULL UNIQUE ,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP ,
			deleted_at TIMESTAMP ,
			name TEXT,
			slug TEXT NOT NULL,
			PRIMARY KEY (pk,uuid)
		)
	`)
	if err != nil {
		return err
	}
	return nil
}

func (r *DB) CreateTableRolePermission() error {
	_, err := r.Db.Exec(`
		CREATE TABLE IF NOT EXISTS role_permissions (
			pk INTEGER AUTOINCREMENT UNIQUE ,
			uuid TEXT NOT NULL UNIQUE ,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP ,
			deleted_at TIMESTAMP ,
			role_uuid TEXT NOT NULL,
			permission_uuid TEXT NOT NULL,
			PRIMARY KEY (pk,uuid,role_uuid,permission_uuid) ,
			FOREIGN KEY (role_uuid) REFERENCES roles(uuid) ON UPDATE SET NULL ON DELETE SET NULL ,
			FOREIGN KEY (permission_uuid) REFERENCES permissions(uuid) ON UPDATE SET NULL ON DELETE SET NULL
		)
	`)
	if err != nil {
		return err
	}
	return nil
}
