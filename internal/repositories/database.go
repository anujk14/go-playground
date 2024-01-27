package repositories

import "github.com/jmoiron/sqlx"

type Database struct {
	DB *sqlx.DB
}

func NewDBConnection(driverName, dataSourceName string) (*Database, error) {
	db, err := sqlx.Connect(driverName, dataSourceName)

	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(10)

	database := Database{
		DB: db,
	}

	return &database, err
}
