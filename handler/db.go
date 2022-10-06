package handler

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbHandler struct {
	DB *gorm.DB
}

func (h *DbHandler) Initialize() {
	db, err := gorm.Open(postgres.Open("postgresql://user:example@0.0.0.0:5432/baby_tracker"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&Baby{})
	db.AutoMigrate(&Activity{})

	h.DB = db
}
