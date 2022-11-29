package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/nicola-sh/CourseProject/Hike/api"
	db "github.com/nicola-sh/CourseProject/Hike/sqlc"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:root@localhost:5432/dohike?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Can't connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("Can't start server:", err)
	}
}
