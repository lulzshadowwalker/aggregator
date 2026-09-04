package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/lulzshadowwalker/aggregator/internal/config"
)

var Connection *sql.DB

func init() {
	connection, err := sql.Open("postgres", config.Instance.Database.URL)
	if err != nil {
		panic(fmt.Errorf("failed to connect to database: %w", err))
	}

	if err := connection.Ping(); err != nil {
		panic(fmt.Errorf("failed to ping database: %w", err))
	}

	Connection = connection
}
