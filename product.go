package main

import (
	"database/sql"
	"log"
)

func CreateProductTable(db *sql.DB) {
    var query string = `
        create table if not exists product (
            id serial primary key,
            name varchar(100) not null,
            price numeric(6, 2) not null,
            available boolean,
            created_at timestamp default now()
        );
    `

    _, err := db.Exec(query)
    if err != nil {
        log.Fatal(err)
    }
}

func InsertProduct(db *sql.DB, product Product) int {
    var query string = `
        insert into product(name, price, available)
        values($1, $2, $3)
        returning id;
    `
    
    var pk int
    err := db.QueryRow(query, product.Name, product.Price, product.Available).Scan(&pk)
    if err != nil {
        log.Fatal(err)
    }

    return pk
}

func SelectProduct(db *sql.DB, pk int) Product {
    product := Product{}

    query := `select name, price, available from product where id = $1;`

    err := db.QueryRow(query, pk).Scan(&product.Name, &product.Price, &product.Available)
    if err != nil {
        if err == sql.ErrNoRows {
            log.Fatal("No rows found with id: ", pk)
        }
        log.Fatal(err)
    }

    return product
}

func SelectMultipleProducts(db *sql.DB) []Product {
    products := []Product{}

    query := `select name, price, available from product`

    rows, err := db.Query(query)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    product := Product{}
    for rows.Next() {
        err = rows.Scan(&product.Name, &product.Price, &product.Available)
        if err != nil {
            log.Fatal(err)
        }
        products = append(products, product)
    }

    return products
}
