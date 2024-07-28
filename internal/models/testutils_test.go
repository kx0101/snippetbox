package models

import (
	"database/sql"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func newTestDB(t *testing.T) *sql.DB {
	err := godotenv.Load("../../.env")
	if err != nil {
		t.Fatal(err)
	}

	db, err := sql.Open("mysql", os.Getenv("MYSQL_DSN_TEST"))
	if err != nil {
		t.Fatal(err)
	}

	script, err := os.ReadFile("./testdata/setup.sql")
	if err != nil {
		t.Fatal(err)
	}

	_, err = db.Exec(string(script))
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		script, err := os.ReadFile("./testdata/teardown.sql")
		if err != nil {
			t.Fatal(err)
		}

		_, err = db.Exec(string(script))
		if err != nil {
			t.Fatal(err)
		}

		db.Close()
	})

	return db
}
