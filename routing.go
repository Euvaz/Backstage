package main

import (
	"database/sql"
    "encoding/json"
	"io/ioutil"
	_ "net/http"

	"github.com/Euvaz/Backstage-Hive/logger"
	"github.com/gin-gonic/gin"
)

type Token struct {
    Key string  `json:"key"`
}

func registerRoutes (router *gin.Engine, db *sql.DB) {
    router.GET("/drones/:name", func(ctx *gin.Context) {
        logger.Info("Handling GET /drones")
    })

    router.POST("/drones/:name", func(ctx *gin.Context) {
        logger.Info("Handling POST /drones")
        name := ctx.Param("name")
        logger.Info(name)

        jsonData, err := ioutil.ReadAll(ctx.Request.Body)
        if err != nil {
            logger.Error(err.Error())
        }

        var token Token
        err = json.Unmarshal(jsonData, &token)
        if err != nil {
            logger.Fatal(err.Error())
        }
        logger.Info(string(token.Key))
    })
}
