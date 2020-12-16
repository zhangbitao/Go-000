package data

import "database/sql"

type DB struct {
	db *sql.DB
}

func NewDB(db *sql.DB) *DB {
	return &DB{
		db: db,
	}
}

func (db *DB) SaveQA(question string, answers []string) {
	// save data to db
	return
}

func (db *DB) Close() error {
	return db.db.Close()
}
