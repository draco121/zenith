package controllers

import (
	"github.com/draco121/horizon/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"zenith/core"
)

type Controllers struct {
	service core.IBotService
}

func NewControllers(service core.IBotService) Controllers {
	c := Controllers{
		service: service,
	}
	return c
}

func (s *Controllers) CreateBot(c *gin.Context) {
	var bot models.Bot
	if err := c.ShouldBind(&bot); err != nil {
		c.JSON(400, err.Error())
	} else {
		res, err := s.service.CreateBot(c, &bot)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
		} else {
			c.JSON(201, res)
		}
	}
}

func (s *Controllers) UpdateBot(c *gin.Context) {
	var bot models.Bot
	if err := c.ShouldBind(&bot); err != nil {
		c.JSON(400, err.Error())
	} else {
		res, err := s.service.UpdateBot(c, &bot)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
		} else {
			c.JSON(201, res)
		}
	}
}

func (s *Controllers) DeleteBot(c *gin.Context) {
	if botId, ok := c.GetQuery("botId"); !ok {
		c.JSON(400, gin.H{
			"message": "bot id not provided",
		})
	} else {
		botId, err := primitive.ObjectIDFromHex(botId)
		if err != nil {
			c.JSON(400, err.Error())
		}
		res, err := s.service.DeleteBot(c, botId)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
		} else {
			c.JSON(201, res)
		}
	}

}

func (s *Controllers) GetBot(c *gin.Context) {
	if botId, ok := c.GetQuery("botId"); !ok {
		c.JSON(400, gin.H{
			"message": "bot id not provided",
		})
	} else {
		botId, err := primitive.ObjectIDFromHex(botId)
		if err != nil {
			c.JSON(400, err.Error())
		}
		res, err := s.service.DeleteBot(c, botId)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
		} else {
			c.JSON(201, res)
		}
	}
	if projectId, ok := c.GetQuery("projectId"); !ok {
		c.JSON(400, gin.H{
			"message": "bot id not provided",
		})
	} else {
		projectId, err := primitive.ObjectIDFromHex(projectId)
		if err != nil {
			c.JSON(400, err.Error())
		}
		res, err := s.service.DeleteBot(c, projectId)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
		} else {
			c.JSON(201, res)
		}
	}
}
