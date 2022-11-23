package hikedb

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:root@localhost:5432/dohike?sslmode=disable"
)

var testQueries *Queries

// db connecting
func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Can't connect to DB:", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
