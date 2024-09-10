//go:build ignore

package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	migrate "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func getOperation() string {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Please provide an operation. Either up or down")
		os.Exit(1)
	}
	operation := args[1]
	if operation != "up" && operation != "down" {
		fmt.Println("Invalid operation. Please provide either up or down")
		os.Exit(1)
	}
	return operation
}

func getDatasource() string {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dataSource := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)
	return dataSource
}

func main() {
	operation := getOperation()
	dataSource := getDatasource()
	db, err := sql.Open("postgres", dataSource)
	if err != nil {
		panic(err)
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		panic(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://./migrations",
		"postgres", driver)
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}
	if operation == "up" {
		err = m.Up()
	} else {
		err = m.Down()
	}
	if err != nil {
		tx.Rollback()
		panic(err)
	}
}
