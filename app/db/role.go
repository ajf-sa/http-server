package db

func (r *DB) CreateTableRole() error {
	_, err := r.Db.Exec(`
		CREATE TABLE IF NOT EXISTS roles (
			pk INTEGER auto_increment ,
			uuid TEXT NOT NULL ,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP ,
			deleted_at TIMESTAMP ,
			name TEXT,
			slug TEXT NOT NULL,
			CONSTRAINT pk_roles PRIMARY KEY (pk,uuid)
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
			pk INTEGER auto_increment ,
			uuid TEXT NOT NULL ,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP ,
			deleted_at TIMESTAMP ,
			name TEXT,
			slug TEXT NOT NULL,
			CONSTRAINT pk_permissions PRIMARY KEY (pk,uuid)
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
			pk INTEGER auto_increment ,
			uuid TEXT NOT NULL ,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP ,
			deleted_at TIMESTAMP ,
			role_uuid TEXT NOT NULL,
			permission_uuid TEXT NOT NULL,
			CONSTRAINT pk_role_permissions PRIMARY KEY (pk,uuid)
			CONSTRAINT fk_role FOREIGN KEY (role_uuid) REFERENCES roles(uuid)
			CONSTRAINT fk_permission FOREIGN KEY (permission_uuid) REFERENCES permissions(uuid)
		)
	`)
	if err != nil {
		return err
	}
	return nil
}
