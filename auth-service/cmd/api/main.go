package main

import (
	"auth/data"
	"database/sql"
)

const webPort = "80"

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {

}
