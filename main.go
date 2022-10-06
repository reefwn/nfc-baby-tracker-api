package main

import (
	"log"
	"nfc-baby-tracker-api/handler"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

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
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	r := setupRouter()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port, _ := viper.Get("PORT").(string)
	r.Run(port)
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
