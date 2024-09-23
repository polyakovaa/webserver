package storage

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// instance of storage
type Storage struct {
	config *Config
	db     *sql.DB
}

func New(config *Config) *Storage {
	return &Storage{
		config: config,
	}
}

// open connection
func (storage *Storage) Open() error {
	db, err := sql.Open("postgres", storage.config.DatabaseURI)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	storage.db = db
	log.Println("Database connction created succesfully!")
	return nil
}

// close connection
func (storage *Storage) Close() {
	storage.db.Close()
}
