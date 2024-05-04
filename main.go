package main

import (
	"github.com/draco121/botmanagerservice/controllers"
	"github.com/draco121/botmanagerservice/core"
	"github.com/draco121/botmanagerservice/repository"
	"github.com/draco121/botmanagerservice/routes"
	"github.com/draco121/common/database"
	"github.com/draco121/common/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

func RunApp() {
	utils.Logger.Info("starting botmanagerservice...")
	client := database.NewMongoDatabase(os.Getenv("MONGODB_URI"))
	db := client.Database("botmanagerservice")
	repo := repository.NewBotRepository(db)
	service := core.NewBotService(client, repo)
	controller := controllers.NewControllers(service)
	router := gin.New()
	router.Use(gin.LoggerWithWriter(utils.Logger.Out))
	routes.RegisterRoutes(controller, router)
	utils.Logger.Info("started botmanagerservice...")
	err := router.Run()
	if err != nil {
		utils.Logger.Fatal(err.Error())
		return
	}
}
func main() {
	_ = godotenv.Load()
	RunApp()
}
