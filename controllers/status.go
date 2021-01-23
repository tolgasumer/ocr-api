package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/otiai10/gosseract/v2"
)

const version = "0.1"

// Status ...
func Status(c *gin.Context) {
	langs, err := gosseract.GetAvailableLanguages()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	client := gosseract.NewClient()
	defer client.Close()
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello!",
		"version": version,
		"tesseract": gin.H{
			"version":   client.Version(),
			"languages": langs,
		},
	})
}
