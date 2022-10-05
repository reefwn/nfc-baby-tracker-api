package main

import (
	"log"
	"nfc-baby-tracker-api/handler"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	r := setupRouter()
	port, _ := viper.Get("PORT").(string)
	r.Run(port)
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	h := handler.DbHandler{}
	h.Initialize()

	r.POST("/baby", h.SaveBaby)
	r.GET("/baby/:id", h.GetBaby)
	r.GET("/babies", h.GetAllBaby)

	r.POST("activity", h.SaveActivity)

	return r
}
