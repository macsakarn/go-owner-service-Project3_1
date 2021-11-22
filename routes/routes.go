package routes

import (
	"accounts/controller"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func getPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "8888"
		fmt.Println("No Port In Heroku/ Server Run : " + port)
	}
	return ":" + port
}
func StartGin() {
	router := gin.Default()

	onwer := router.Group("/owner")
	{
		onwer.GET("/info", controller.GetInfo())
		onwer.PUT("/info", controller.UpdateInfo())
		onwer.GET("/account", controller.GetAccounts())
		onwer.POST("/account", controller.InsertAccounts())
		onwer.PUT("/account", controller.UpdateAccount())
		onwer.DELETE("/account", controller.DeleteAccounts())
	}

	router.Run(getPort())
}
