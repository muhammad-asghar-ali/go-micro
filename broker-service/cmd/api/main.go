package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "80"

type (
	Config struct{}
)

func main() {

	app := Config{}

	log.Println("start broker server at port ", webPort)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Panic(err)
	}
}
