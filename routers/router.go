package routers

import (
	"github.com/gin-gonic/gin"
	"pokemon/setting"
	"pokemon/service"
)

func InitRouter() (*gin.Engine) {
	gin.SetMode(setting.RunMode)

	router := gin.Default()

	v1 := router.Group("/api/v1")

	// 技能
	skillGroup := v1.Group("/skill")
	{
		skillGroup.POST("/", service.CreateSkill)
		skillGroup.GET("/", service.FetchAllSkill)
		skillGroup.GET("/:id", service.FetchOneSkill)
		skillGroup.PUT("/:id", service.UpdateSkill)
		skillGroup.DELETE("/:id", service.DeleteSkill)
	}

	// 属性
	propGroup := v1.Group("/prop")
	{
		propGroup.POST("/", service.CreateProp)
		propGroup.GET("/", service.FetchAllProp)
		propGroup.GET("/:id", service.FetchOneProp)
		propGroup.PUT("/:id", service.UpdateProp)
		propGroup.DELETE("/:id", service.DeleteProp)
	}

	// 类型
	typeGroup := v1.Group("/type")
	{
		typeGroup.POST("/", service.CreateType)
		typeGroup.GET("/", service.FetchAllType)
		typeGroup.GET("/:id", service.FetchOneType)
		typeGroup.PUT("/:id", service.UpdateType)
		typeGroup.DELETE("/:id", service.DeleteType)
	}

	// 特性
	specialGroup := v1.Group("/special")
	{
		specialGroup.POST("/", service.CreateSpecial)
		specialGroup.GET("/", service.FetchAllSpecial)
		specialGroup.GET("/:id", service.FetchOneSpecial)
		specialGroup.PUT("/:id", service.UpdateSpecial)
		specialGroup.DELETE("/:id", service.DeleteSpecial)
	}
	router.Run()

	return router
}