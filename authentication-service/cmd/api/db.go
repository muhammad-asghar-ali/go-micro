package main

import (
	"database/sql"
	"log"
	"os"
	"time"
)

var count int64

func openDB(dns string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dns)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func connectDB() *sql.DB {
	dns := os.Getenv("DNS")

	for {
		connection, err := openDB(dns)
		if err != nil {
			log.Println("postgres not yet ready ...")
			count++
		} else {
			log.Println("postgres ready ...")
			return connection
		}

		if count > 10 {
			log.Println(err)
			return nil
		}

		log.Println("backing off for 2 seconds ...")
		time.Sleep(2 * time.Second)
		continue
	}
}
