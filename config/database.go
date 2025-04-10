package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func ConnectToDatabase(){
		var err error

	    dsn := "postgres://postgres:123456@localhost:5432/crud_app"
		
		DB, err = sql.Open("pgx", dsn)

		if err != nil {
			log.Fatalf("Could not connect to the database: %v", err)
		}
	
		fmt.Println("Database connected!")
}