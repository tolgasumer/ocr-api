package main

import (
	"log"
	"os"

	"ocr-api/controllers"

	"github.com/gin-gonic/gin"
)

var logger *log.Logger

func main() {

	router := gin.Default()
	// API
	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.GET("/", controllers.Index)
			v1.GET("/status", controllers.Status)
			v1.POST("/base64", controllers.Base64)
			v1.POST("/file", controllers.FileUpload)
		}
	}

	router.MaxMultipartMemory = 8 << 20 // 8 MiB

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalln("Required env `PORT` is not specified.")
	}
	log.Printf("listening on port %s", port)

	router.Run(":" + port)

}
