package main

import (
	"database/sql"
	"log"

	"github.com/cznic/ql"
	"golang.org/x/net/context"
)

func setUp(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(`
	CREATE TABLE note (
	  id BIGINT 
	  ,title STRING
	  ,body STRING
	  ,created_at STRING
	  ,updated_at STRING
	);
	`)
	if err != nil {
		return err
	}
	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}

func tearDown(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(`
	DROP TABLE note;
	`)
	if err != nil {
		return err
	}
	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}

func main() {
	ql.RegisterDriver()
	db, err := sql.Open("ql", "./db/ql.db")
	defer db.Close()
	if err != nil {
		log.Fatalf("failed to open db: %s", err)
	}
	rootCtx := context.Background()
	rootCtx = context.WithValue(rootCtx, "db", db)

	if err = setUp(db); err != nil {
		log.Fatalf("failed to create table: %s", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("failed to ping db: %s", err)
	}
	if err = tearDown(db); err != nil {
		log.Fatalf("failed to drop table: %s", err)
	}
}
