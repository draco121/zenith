package routes

import (
	"github.com/draco121/horizon/constants"
	"github.com/draco121/horizon/middlewares"
	"github.com/draco121/horizon/utils"
	"github.com/gin-gonic/gin"
	"zenith/controllers"
)

func RegisterRoutes(controllers controllers.Controllers, router *gin.Engine) {
	utils.Logger.Info("Registering routes...")
	v1 := router.Group("/v1")
	v1.POST("/bot", middlewares.AuthMiddleware(constants.Write), controllers.CreateBot)
	v1.GET("/bot", middlewares.AuthMiddleware(constants.Read), controllers.GetBot)
	v1.PATCH("/bot", middlewares.AuthMiddleware(constants.Write), controllers.UpdateBot)
	v1.DELETE("/bot", middlewares.AuthMiddleware(constants.Write), controllers.DeleteBot)
	utils.Logger.Info("Routes registered")
}
