package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Activity struct {
	Id        uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	BabyId    uuid.UUID `json:"baby_id" gorm:"type:uuid;not null;" form:"baby_id" binding:"required"`
	Type      string    `json:"type" sql:"type:activity_type" gorm:"not null" form:"type" binding:"required"`
	Time      time.Time `json:"time" gorm:"not null" form:"time" binding:"required"`
	CreatedAt time.Time `json:"created_at" gorm:"default:now();not null"`
}

func (h *DbHandler) SaveActivity(c *gin.Context) {
	Activity := Activity{}

	if err := c.ShouldBindJSON(&Activity); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	fmt.Println("baby_id", &Activity.BabyId)
	if err := h.DB.Save(&Activity).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, Activity)
}
