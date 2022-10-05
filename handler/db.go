package handler

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbHandler struct {
	DB *gorm.DB
}

func (h *DbHandler) Initialize() {
	db, err := gorm.Open(postgres.Open("postgresql://user:example@0.0.0.0:5432/baby_tracker"))
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&Baby{})
	db.AutoMigrate(&Activity{})

	h.DB = db
}
