package main

import (
	"database/sql"
	"io/ioutil"
    "fmt"
	_ "net/http"

	"github.com/Euvaz/Backstage-Hive/logger"
	"github.com/gin-gonic/gin"
)

func registerRoutes (router *gin.Engine, db *sql.DB) {
    router.GET("/drones", func(ctx *gin.Context) {
        logger.Info("Handling GET /drones")
    })

    router.POST("/drones", func(ctx *gin.Context) {
        logger.Info("Handling POST /drones")
        jsonData, err := ioutil.ReadAll(ctx.Request.Body)
        if err != nil {
            logger.Error(err.Error())
        }
        fmt.Println(string(jsonData))
    })
}
