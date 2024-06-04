package main

import (
	"database/sql"
	"fmt"
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

    CreateProductTable(db)

    productToInsert := Product{
        Name: "Book",
        Price: 15.55,
        Available: true,
    }
    var pk int = InsertProduct(db, productToInsert)
    fmt.Println(pk)
}
