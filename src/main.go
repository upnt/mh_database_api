package main

import (
    "github.com/gin-gonic/gin"
    "github.com/upnt/mh_database_api/src/controller"
)


func main() {
    router := gin.Default()

    router.GET("/", controller.IndexGET)
    router.Run(":8080")
}
