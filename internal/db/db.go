package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/cenkalti/backoff"
	_ "github.com/cockroachdb/cockroach-go/v2/crdb"
	"github.com/lib/pq"
	"github.com/tonadr1022/speed-cube-time/internal/apperrors"
)

// initializes the database connection and creates tables if not initialized
func InitDB() (*sql.DB, error) {
	var (
		db  *sql.DB
		err error
	)

	// pgConnString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
	// 	os.Getenv("PGHOST"),
	// 	os.Getenv("PGPORT"),
	// 	os.Getenv("PGDATABASE"),
	// 	os.Getenv("PGUSER"),
	// 	os.Getenv("PGPASSWORD"),
	// )

	// fmt.Printf("conn string \n%v\n", pgConnString)
	fmt.Printf("\nDB URL: %s\n", os.Getenv("DATABASE_URL"))
	openDB := func() error {
		db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
		return err
	}

	// open the connection, retry retires it, especially useful when
	// docker compose starts the server before the DB is ready for connections
	err = backoff.Retry(openDB, backoff.NewExponentialBackOff())
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func TransformError(err error) error {
	if err == nil {
		return nil
	}
	if dbErr, ok := err.(*pq.Error); ok {
		if dbErr.Code == "23505" {
			return apperrors.ErrAlreadyExists
		}
	}
	return err
}

type Scanner interface {
	Scan(dest ...any) error
}
