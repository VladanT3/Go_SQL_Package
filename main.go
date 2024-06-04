package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
    _ "github.com/lib/pq"
)

func main() {
    godotenv.Load(".env")
    var dbConn string = os.Getenv("DB_CONN")

    db, err := sql.Open("postgres", dbConn)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }

    createProductTable(db)
}

func createProductTable(db *sql.DB) {
    var query string = `
        create table if not exists product (
            id serial primary key,
            name varchar(100) not null,
            price numeric(6, 2) not null,
            available boolean,
            created_at timestamp default now()
        )
    `

    _, err := db.Exec(query)
    if err != nil {
        log.Fatal(err)
    }
}
