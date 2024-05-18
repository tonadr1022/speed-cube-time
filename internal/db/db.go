package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/cenkalti/backoff"
	_ "github.com/cockroachdb/cockroach-go/v2/crdb"
	_ "github.com/lib/pq"
)

// initializes the database connection and creates tables if not initialized
func InitDB(dsn string) (*sql.DB, error) {
	fmt.Println("DSN: " + dsn)
	var (
		db  *sql.DB
		err error
	)

	pgConnString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGDATABASE"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
	)

	openDB := func() error {
		db, err = sql.Open("postgres", pgConnString)
		return err
	}

	// if err = openDB(); err != nil {
	// 	return nil, err
	// }
	// open the connection, retry retires it, especially useful when
	// docker compose starts the server before the DB is ready for connections
	err = backoff.Retry(openDB, backoff.NewExponentialBackOff())
	if err != nil {
		return nil, err
	}

	// _, err = db.Exec(`
	//    CREATE TABLE IF NOT EXISTS "user" (
	//        id TEXT PRIMARY KEY,
	//        username VARCHAR(255) UNIQUE NOT NULL,
	//        password VARCHAR(255) NOT NULL,
	//        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	//    );
	//    `)
	// if err != nil {
	// 	return nil, err
	// }
	//
	// _, err = db.Exec(`
	//    CREATE TABLE IF NOT EXISTS sessions (
	//        id TEXT PRIMARY KEY,
	//        name VARCHAR(255) NOT NULL,
	//        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	//        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	//        user_id TEXT,
	//        FOREIGN KEY (user_id) REFERENCES "user"(id) ON DELETE CASCADE
	//    );
	// `)
	// if err != nil {
	// 	return nil, err
	// }

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
