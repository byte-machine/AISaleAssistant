package parser_controllers

import (
	"AISale/services/parsing"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ParsePhones(c *gin.Context) {
	err := parsing.ParsePhonesCSV("/app/files/excel/phones.csv")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "completed"})
}
