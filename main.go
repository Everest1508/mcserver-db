package main

import (
	"net/http"

	"github.com/everest1508/mcserver-db/constants"
	"github.com/everest1508/mcserver-db/db"
	"github.com/everest1508/mcserver-db/models"
	"github.com/everest1508/mcserver-db/services"
	utils "github.com/everest1508/mcserver-db/utils/api"
	"github.com/gin-gonic/gin"
)

func init() {
	db.InitSqlite3()
	db.DB.AutoMigrate(&models.Server{})
}

func pongHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "pong")
}

func main() {
	client := utils.NewAPIClient(constants.BASE_URL)
	services.NewCronJob(client).StartCron()

	r := gin.Default()
	r.GET("/ping", pongHandler)
	r.GET("/servers", func(ctx *gin.Context) {
		var servers []models.Server
		db.DB.Find(&servers)
		ctx.JSON(200, servers)
	},
	)

	r.Run()
}
