package database

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgresInfo struct {
	connectionString string
	database         *sqlx.DB
}

func NewPostgresConnection(user, password, db, host string) (*PostgresInfo, error) {
	pg := &PostgresInfo{
		connectionString: fmt.Sprintf("user=%s dbname=%s sslmode=disable password=%s host=%s", user, db, password, host),
	}

	// does connect need to be open instead
	dbConn, err := sqlx.Connect("postgres", pg.connectionString)
	if err != nil {
		return nil, err
	}
	if err := dbConn.Ping(); err != nil {
		return nil, err
	} else {
		log.Printf("Connected to database: %s", db)
	}

	pg.database = dbConn
	return pg, nil
}

func (pg *PostgresInfo) Close() error {
	return pg.database.Close()
}
