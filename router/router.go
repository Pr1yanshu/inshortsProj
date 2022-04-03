package router

import (
	"github.com/gin-gonic/gin"
	"inshortsProj/controller"
)

func AddCovidRoutes(group *gin.RouterGroup) {
	group.GET("healthCheck", controller.GetHealthStatus)
	group.GET("getCovidData", controller.GetCovidData)
	group.POST("updateCovidData", controller.UpdateCovidData)
}
