package main

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"log"
)

func main() {
	m, err := migrate.New(
		"file://cmd/migrations/schema/",
		"postgres://localhost:5432/postgres?sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	erra := m.Steps(2)

	if erra != nil {
		log.Fatal(err)
	}
}

//func main() {
//	db, err := sql.Open("postgres", "")
//
//	if err != nil {
//		log.Fatal("1", err)
//	}
//	driver, err := postgres.WithInstance(db, &postgres.Config{})
//	if err != nil {
//		log.Fatal("2", err)
//	}
//	m, err := migrate.NewWithDatabaseInstance(
//		"",
//		"postgres", driver)
//
//	if err != nil {
//		log.Fatal("3", err)
//	}
//
//	erra := m.Up() // or m.Step(2) if you want to explicitly set the number of migrations to run
//
//	if erra != nil {
//		log.Fatal("4", err)
//	}
//}
