package store

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"

	"github.com/logikone/sunday-funday/store/migrate"
)

func Connect(driver, dataSource string) (*sqlx.DB, error) {
	db, err := sql.Open(driver, dataSource)

	if err != nil {
		return nil, err
	}

	dbx := sqlx.NewDb(db, driver)

	if err := pingDatabase(dbx); err != nil {
		return nil, err
	}

	if err := setupDatabase(dbx); err != nil {
		return nil, err
	}

	return dbx, nil
}

func Must(db *sqlx.DB, err error) *sqlx.DB {
	if err != nil {
		panic(err)
	}

	return db
}

func pingDatabase(db *sqlx.DB) (err error) {
	for i := 0; i < 30; i++ {
		err = db.Ping()
		if err == nil {
			return
		}
		time.Sleep(time.Second)
	}
	return
}

func setupDatabase(db *sqlx.DB) error {
	return migrate.Migrate(db.DB)
}
