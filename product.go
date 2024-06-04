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
