package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "net/http"
    "github.com/jinzhu/gorm"
    _ "github.com/go-sql-driver/mysql"
)

type Monster struct {
    ID      string `json:"id"`
    NameEng string `json:"name_eng"`
    NameJpn string `json:"name_jpn"`
}

func main() {
    router := gin.Default()

    db, err := connectDB()
    if err != nil {
        panic(err.Error())
    } else {
        fmt.Println("Connected")
    }

    monsters := []*Monster{}

    tx := db.Table("monster").Begin()
    err = tx.Find(&monsters).Commit().Error

    if err != nil {
        panic(err.Error())
    } else {
        fmt.Println("get all monsters")
    }

    router.GET("/", func(ctx *gin.Context){
        ctx.JSON(http.StatusOK, monsters)
    })

    router.Run(":8080")
}

func connectDB() (database *gorm.DB, err error) {
    DBMS     := "mysql"

    USER     := "root"
    PASSWORD := "rsQlQy4S12YY0"
    PROTOCOL := "tcp(mysql:3306)"
    DBNAME   := "mh"

    CONNECT := USER + ":" + PASSWORD
    CONNECT += "@" + PROTOCOL + "/" + DBNAME
    CONNECT += "?charset=utf8&&parseTime=true&loc=Asia%2fTokyo"

    return gorm.Open(DBMS, CONNECT)
}
