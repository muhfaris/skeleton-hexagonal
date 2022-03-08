package config

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ConfigApp struct {
	Port   int
	DB     *pgx.Conn
	Client *mongo.Client
}

func CreateConfigApp() *ConfigApp {
	db, err := initDatabaseMySQL()
	if err != nil {
		panic(err)
	}

	client, err := initDatabaseMongoDB()
	if err != nil {
		panic(err)
	}

	return &ConfigApp{
		DB:     db,
		Client: client,
	}
}

func initDatabaseMySQL() (*pgx.Conn, error) {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	databaseURL := fmt.Sprintf("postgres://admin123:admin123@localhost:5432/skelaton_db")
	db, err := pgx.Connect(context.Background(), databaseURL)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func initDatabaseMongoDB() (*mongo.Client, error) {
	credential := options.Credential{
		Username: "admin123",
		Password: "admin123",
	}

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(credential)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Connect(context.Background())
	if err != nil {
		return nil, err
	}

	return client, nil
}
