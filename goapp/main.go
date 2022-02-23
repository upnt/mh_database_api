package main

import (
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)


func main() {
    dbconf  := "root:rsQlQy4S12YY0@tcp(localhost:3306)/mh_database?charset=utf8mb4"
    db, err := sql.Open("mysql", dbconf)

    defer db.Close()

    if err != nil {
        fmt.Println(err.Error())
    }

    err = db.Ping()

    if err != nil {
        fmt.Println(err.Error())
        return
    } else {
        fmt.Println("connect: connection success")
    }
}
