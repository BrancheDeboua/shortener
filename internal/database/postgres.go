package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
)

type PostgresConnector struct {
	connStr string
}

func NewPostgresConnector(connStr string) *PostgresConnector {
	return &PostgresConnector{
		connStr: connStr,
	}
}

func (pg *PostgresConnector) Connect() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), pg.connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
		return &pgx.Conn{}, err
	}

	return conn, nil
}
