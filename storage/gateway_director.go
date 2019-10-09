package storage

// DirectorGateway is a type of TableGateway
type DirectorGateway struct{}

// Find retrieves the directors by id
func (gw DirectorGateway) Find(id int) (*Director, error) {
	director := new(Director)
	err := db.QueryRow("SELECT * FROM directors WHERE id = $1", id).Scan(&director.ID, &director.Name, &director.Surname,&director.Year, &director.CreatedAt)

	if err != nil {
		return nil, err
	}

	return director, nil
}

// FindAll retrieves all the directors
func (gw DirectorGateway) FindAll() ([]*Director, error) {
	rows, err := db.Query("SELECT * FROM directors")

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var directors []*Director

	for rows.Next() {
		director := new(Director)
		err = rows.Scan(&director.ID, &director.Name, &director.Surname, &director.Year, &director.CreatedAt)

		if err != nil {
			return nil, err
		}

		directors = append(directors, director)
	}

	return directors, nil
}

// Insert inserts a new row to the directors table
func (gw DirectorGateway) Insert(director *Director) (*Director, error) {
	err := db.QueryRow("INSERT INTO directors (name, surname, birth_year) VALUES ($1, $2, $3) RETURNING id",
		director.Name,
		director.Surname,
		director.Year,
	).Scan(&director.ID)

	if err != nil {
		return nil, err
	}

	return director, nil
}

// Update updates a row in directors table by id
func (gw DirectorGateway) Update(director *Director) (int64, error) {
	result, err := db.Exec("UPDATE director SET name = $2, surname = $3, birth_year = $4 WHERE id = $1",
		director.ID,
		director.Name,
		director.Surname,
		director.Year,
	)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

// Delete removes a row in table directors by id
func (gw DirectorGateway) Delete(id int) (int64, error) {
	result, err := db.Exec("DELETE FROM directors WHERE id = $1", id)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}
