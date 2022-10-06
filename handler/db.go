package handler

import (
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbHandler struct {
	DB *gorm.DB
}

func (h *DbHandler) Initialize() {
	dsn := viper.Get("DB_DSN").(string)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&Baby{})
	db.AutoMigrate(&Activity{})

	h.DB = db
}
