package hikedb

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/nicola-sh/CourseProject/Hike/util"
)

var testQueries *Queries
var testDB *sql.DB

// db connecting
func TestMain(m *testing.M) {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Can't load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Can't connect to DB:", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
