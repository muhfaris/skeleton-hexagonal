package config

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

type ConfigApp struct {
	Port int
	DB   *pgx.Conn
}

func CreateConfigApp() *ConfigApp {
	db, err := initDatabase()
	if err != nil {
		panic(err)
	}

	return &ConfigApp{
		DB: db,
	}
}

func initDatabase() (*pgx.Conn, error) {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	databaseURL := fmt.Sprintf("postgres://admin123:admin123@localhost:5432/skelaton_db")
	db, err := pgx.Connect(context.Background(), databaseURL)
	if err != nil {
		return nil, err
	}

	return db, nil
}
