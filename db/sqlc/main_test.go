package db

import (
	"database/sql"
	"log"
	"os"
	"simplebank/utils"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := utils.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	dbDriver := config.DBDriver
	dbSource := config.DBSource

	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	if err = testDB.Ping(); err != nil {
		log.Fatal("cannot ping db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
