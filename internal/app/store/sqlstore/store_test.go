package sqlstore_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "host=127.0.0.1 port=5432 user=restapi password=1234 sslmode=disable"
	}

	os.Exit(m.Run())
}
