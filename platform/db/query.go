package db

import (
	"database/sql"
	"demo/postgres"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"os"
)

func Query() *postgres.Queries {
	db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	//defer func(db *sql.DB) {
	//	err := db.Close()
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}(db)

	query := postgres.New(db)

	return query
}
