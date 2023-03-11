package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "80"

type Config struct{}

func main() {
	app := Config{}
	log.Printf("Starting broker service at http://127.0.0.1:%s\n", webPort)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.router(),
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
