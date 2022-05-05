package db

func (r *DB) CreateTableClient() error {
	_, err := r.Db.Exec(`
		CREATE TABLE IF NOT EXISTS clients (
			pk INTEGER   NOT NULL UNIQUE ,
			uuid TEXT NOT NULL UNIQUE ,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP ,
			deleted_at TIMESTAMP ,
			name TEXT NOT NULL,
			user_uuid text NOT NULL,
			PRIMARY KEY (pk,uuid,user_uuid) ,
			FOREIGN KEY (user_uuid) REFERENCES users(id) ON UPDATE SET NULL ON DELETE SET NULL
		)
	`)
	if err != nil {
		return err
	}
	return nil
}

func (r *DB) CreateOneClient(client *Client) error {
	_, err := r.Db.Exec(`
		INSERT INTO clients (uuid,name,user_uuid)
		VALUES ($1,$2,$3)
	`, client.UUID, client.Name, client.UserUUID)
	if err != nil {
		return err
	}
	return nil
}

func (r *DB) FindAllClient() ([]Client, error) {
	var clients []Client
	rows, err := r.Db.Query(`
		SELECT pk,uuid,name,user_uuid FROM clients
	`)
	if err != nil {

		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var client Client
		err := rows.Scan(&client.PK, &client.UUID, &client.Name, &client.UserUUID)
		if err != nil {

			return nil, err
		}
		clients = append(clients, client)
	}
	return clients, nil
}

func (r *DB) FindOneClientByID(uuid string) (*Client, error) {
	var client Client
	sqlStmt := `SELECT pk,uuid,name,user_uuid FROM clients WHERE uuid=$1`
	err := r.Db.QueryRow(sqlStmt, uuid).Scan(&client.PK, &client.UUID, &client.Name, &client.UserUUID)
	if err != nil {
		return nil, err
	}
	return &client, nil
}
