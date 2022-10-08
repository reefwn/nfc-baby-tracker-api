package main

import (
	"fmt"
	"net/http"
	"nfc-baby-tracker-api/handler"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	_ "nfc-baby-tracker-api/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title   NFC Baby Tracker API
// @version 1.0

// @schemes https http

// @securityDefinitions.apikey ApiKeyAuth
// @in                         header
// @name                       Authorization
func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	r := setupRouter()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "api is runnig",
		})
	})
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := os.Getenv("PORT")
	r.Run(":" + port)
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	h := handler.DbHandler{}
	h.Initialize()

	r.POST("/babies", h.SaveBaby)
	r.GET("/babies/:id", h.GetBaby)
	r.GET("/babies", h.GetAllBaby)

	r.POST("activities", h.SaveActivity)
	r.GET("activities/:baby_id/latest", h.GetLatestActivity)

	return r
}
