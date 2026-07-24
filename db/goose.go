package db

import (
	"database/sql"

	"github.com/pressly/goose/v3"
)

func Migrate(db *sql.DB, dir string) error {
	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	return goose.Up(
		db,
		dir,
	)
}

func Rollback(db *sql.DB, dir string) error {
	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	return goose.Down(
		db,
		dir,
	)
}
