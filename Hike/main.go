package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/nicola-sh/CourseProject/Hike/api"
	db "github.com/nicola-sh/CourseProject/Hike/sqlc"
	"github.com/nicola-sh/CourseProject/Hike/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Can't load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Can't connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Can't start server:", err)
	}
}
