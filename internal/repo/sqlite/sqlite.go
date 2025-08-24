package sqlite

import (
	"database/sql"
	"fmt"
)

type Repo struct {
	db *sql.DB
}

func New(repoPath string) (*Repo, error) {
	const op = "repo.sqlite.New"

	db, err := sql.Open("sqlite3", repoPath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	// Когда-то миграции будут исследованы на 100%

	stmt, err := db.Prepare(`

	CREATE TABLE IF NOT EXISTS roles (
    	id SERIAL PRIMARY KEY,
    	name VARCHAR(50) NOT NULL UNIQUE,
    	description TEXT
	);

	CREATE TABLE IF NOT EXISTS users (
    	id SERIAL PRIMARY KEY,
    	username VARCHAR(100) NOT NULL UNIQUE,
    	password_hash TEXT NOT NULL,
    	role_id INTEGER NOT NULL REFERENCES roles(id),
    	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS forms (
    	id SERIAL PRIMARY KEY,
    	title TEXT NOT NULL,
    	description TEXT,
    	creator_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS form_fields (
    	id SERIAL PRIMARY KEY,
    	form_id INTEGER NOT NULL REFERENCES forms(id) ON DELETE CASCADE,
    	label TEXT NOT NULL,
    	field_type VARCHAR(50) NOT NULL,
    	options JSONB,
    	required BOOLEAN DEFAULT false,
    	display_order INTEGER DEFAULT 0
	);

	CREATE TABLE IF NOT EXISTS responses (
    	id SERIAL PRIMARY KEY,
    	form_id INTEGER NOT NULL REFERENCES forms(id) ON DELETE CASCADE,
    	user_id INTEGER REFERENCES users(id),
    	submitted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);


	CREATE TABLE IF NOT EXISTS answer_values (
    	id SERIAL PRIMARY KEY,
    	response_id INTEGER NOT NULL REFERENCES responses(id) ON DELETE CASCADE,
    	field_id INTEGER NOT NULL REFERENCES form_fields(id),
    	value TEXT
	);

	`)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Repo{db: db}, nil
}
