package main

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	_ "github.com/lib/pq"
)

func TestPostgresWithContainer(t *testing.T) {
	ctx := context.Background()

	
pgContainer, err := postgres.RunContainer(ctx,
		postgres.WithDatabase("testdb"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("password"),
	)
	require.NoError(t, err)
	defer pgContainer.Terminate(ctx)

	
	connStr, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
	require.NoError(t, err)

	
	time.Sleep(2 * time.Second)

	
	db, err := sql.Open("postgres", connStr)
	require.NoError(t, err)
	defer db.Close()

	err = db.Ping()
	require.NoError(t, err)

	fmt.Println("Connected to TEST Postgres ")


	require.NoError(t, createTable(db))

	
	_, err = db.Exec("INSERT INTO product(name, price) VALUES($1,$2)", "Phone", 999)
	require.NoError(t, err)

	
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM product").Scan(&count)
	require.NoError(t, err)

	require.Equal(t, 1, count)
}

func createTable(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS product (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		price NUMERIC NOT NULL
	)`)
	return err
}
