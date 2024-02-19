package main

import (
	"authentication/data"
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

const webPort = "80"

type (
	Config struct {
		DB     *sql.DB
		Models data.Models
	}
)

func main() {
	conn := connectDB()
	if conn == nil {
		log.Panic("can't connect to database")
	}
	app := Config{
		DB:     conn,
		Models: data.New(conn),
	}

	log.Println("start authentication server at port ", webPort)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Panic(err)
	}
}
