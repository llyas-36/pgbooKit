package pgbookit

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func DB() *pgxpool.Pool {

	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))// RUN .env file in the terminal or use export DATABASE_URL=url-link

	if err != nil{
		fmt.Fprintf(os.Stderr, "Unable to create")
		os.Exit(1)
	}

	return dbpool
}
