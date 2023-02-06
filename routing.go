package main

import (
	"database/sql"
    "net/http"

	"github.com/Euvaz/Backstage-Hive/logger"
	"github.com/gin-gonic/gin"
)

func registerRoutes (router *gin.Engine, db *sql.DB) {
    router.GET("/drones", func(ctx *gin.Context) {
        logger.Info("Handling /drones")
        ctx.JSON(http.StatusOK, gin.H {
            "foo": "bar",
        })
    })
}
