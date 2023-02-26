package utils

import (
	routers "labs/routers"
	structs "labs/structs"

	"github.com/gin-gonic/gin"
)

type MainRegister interface {
	DemoRouter() *gin.Engine
}

type mainRegister struct {
	router *gin.Engine
}

func NewMainRegister(router *gin.Engine) mainRegister {
	return mainRegister{
		router: router,
	}
}

func (m mainRegister) DemoRouter() *gin.Engine {
	DemoRouter := routers.NewDemoRouter
	DemoGroup := m.router.Group("/v1/demo")
	DemoData := structs.SRegisterRouter{
		Router: DemoGroup,
	}
	DemoRouter(DemoData).Start()
	return m.router
}

func (m mainRegister) GeneralRouter() *gin.Engine {
	GeneralRouter := routers.NewGeneralRouter
	GeneralGroup := m.router.Group("/v1/general")
	GeneralData := structs.SRegisterRouter{
		Router: GeneralGroup,
	}
	GeneralRouter(GeneralData).Start()
	return m.router
}
