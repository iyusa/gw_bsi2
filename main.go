package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ussidev/permata_trx/common"
	"github.com/ussidev/permata_trx/controller"
	"github.com/ussidev/permata_trx/model"
)

func main() {
	common.LoadConfig("permata_trx.yaml")
	common.Ensure(model.Initialize())
	common.Ensure(model.InitAuth())

	// setup gin
	if common.Config.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()

	router.GET("/", home)

	service := router.Group("/service")
	service.Use(controller.AuthMiddleware())

	service.POST("/balance", controller.Balance)

	port := common.Config.HTTPPort
	fmt.Printf("Starting Permata-Out HTTP Server version %s on %d\n", common.Version, port)

	err := router.Run(fmt.Sprintf(":%d", port))
	common.Ensure(err)
}

func home(c *gin.Context) {
	c.JSON(200, gin.H{"status": "Permata-Out Server is running...", "version": common.Version, "server": "HTTP"})
}
