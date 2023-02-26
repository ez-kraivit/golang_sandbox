package routers

import (
	structs "labs/structs"

	"github.com/gin-gonic/gin"
)

type GeneralRouter interface {
	Start(data demoRouter) gin.IRoutes
}

type generalRouter struct {
	router *gin.RouterGroup
}

func NewGeneralRouter(data structs.SRegisterRouter) generalRouter {
	return generalRouter{
		router: data.Router,
	}
}

func (d generalRouter) Start() gin.IRoutes {

	d.router.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "Hello General",
		})
	})
	return d.router
}
