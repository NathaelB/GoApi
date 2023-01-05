package main

import (
	"github.com/jmoiron/sqlx"
)

type Store interface {
	Connect() error
	Disconnect() error
	GetPeoples() ([]Person, error)
}

type dbStore struct {
	db *sqlx.DB
}

func (dbStore *dbStore) Connect() (err error) {
	dbStore.db, err = sqlx.Connect("postgres", "user=nathael dbname=go password=nathael sslmode=disable")

	if err != nil {
		return err
	}
	return nil
}

func (dbStore *dbStore) Disconnect() (err error) {
	if err = dbStore.db.Close(); err != nil {
		return err
	}
	return nil
}

func (dbStore *dbStore) GetPeoples() ([]Person, error) {
	var people []Person
	if err := dbStore.db.Select(&people, "SELECT * FROM people"); err != nil {
		return nil, err
	}
	return people, nil
}
