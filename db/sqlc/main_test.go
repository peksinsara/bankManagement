package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
)

const (
	dbDriverName = "postgres"
	dbSourceName = "postgresql://root:secret@localhost:5432/bankManagement?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriverName, dbSourceName)
	if err != nil {
		log.Fatal("Failed to open connect to database: ", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
