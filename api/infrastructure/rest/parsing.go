package rest

import (
	"AISale/api/infrastructure/controllers/parser_controllers"
	"AISale/config"
	"github.com/gin-gonic/gin"
)

func ParsingRoutes(router *gin.Engine, settings config.Settings) {
	productGroup := router.Group("chat")
	{
		productGroup.POST("/parse_phones", parser_controllers.ParsePhones)
	}
}
