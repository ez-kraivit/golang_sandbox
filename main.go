package main

import (
	"fmt"
	utils "labs/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	Cors := utils.CorsFunc
	JsonLoggerMiddleware := utils.JsonLoggerMiddleware

	router := gin.Default()

	router.Use(Cors())
	router.Use(JsonLoggerMiddleware())

	mainRegister := utils.NewMainRegister(router)
	mainRegister.DemoRouter()
	mainRegister.GeneralRouter()

	fmt.Println("Server running on port 0.0.0.0:50230")
	router.Run("0.0.0.0:50230")

}
