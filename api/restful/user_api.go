package restful

import (
	"github.com/gin-gonic/gin"
	"steward/businese"
)

func LoadRoute(engine *gin.Engine) {
	var basePath = "/steward/api/v1"
	bp := engine.Group(basePath)
	userR := bp.Group("/user")
	{
		userR.POST("/add", businese.NewUserController().Add)
		userR.DELETE("/delete", businese.NewUserController().Delete)
		userR.GET("/find", businese.NewUserController().Find)
		userR.GET("/pageList", businese.NewUserController().PageList)
		userR.PUT("/update", businese.NewUserController().Update)
	}
	deviceR := bp.Group("/device")
	{
		deviceR.POST("/add")
		deviceR.DELETE("/delete")
		deviceR.GET("/find")
		deviceR.GET("/pageList")
		deviceR.PUT("/update")
	}
}
