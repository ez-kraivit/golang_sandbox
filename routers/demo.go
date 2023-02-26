package routers

import (
	"labs/controllers"
	"labs/repository"
	"labs/services"
	structs "labs/structs"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DemoRouter interface {
	Start(data demoRouter) gin.IRoutes
}

type demoRouter struct {
	router *gin.RouterGroup
	db     *gorm.DB
}

func NewDemoRouter(data structs.SRegisterRouter) demoRouter {
	return demoRouter{
		router: data.Router,
	}
}

func (d demoRouter) Start() gin.IRoutes {

	usersRepository := repository.NewUsersRepository(d.db)

	demoService := services.NewDemoService(usersRepository)
	demoController := controllers.NewDemoController(demoService)

	d.router.GET("", demoController.Get)

	return d.router
}
