package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Index ...
func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello!",
		"version": version,
	})
}
