package data

import "database/sql"

func NewDBConnection(url string) (*sql.DB, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	// Test the database connection.
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
