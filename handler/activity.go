package handler

import (
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

// SaveActivity             godoc
// @Summary      Store a new activity
// @Description  Takes an activity JSON and store in DB. Return saved JSON.
// @Tags         Activities
// @Produce      json
// @Param        activity  body      Activity  true  "Activity JSON"
// @Security     ApiKeyAuth
// @Success      200   {object}  Activity
// @Router       /activities [post]
func (h *DbHandler) SaveActivity(c *gin.Context) {
	Activity := Activity{}

	if err := c.ShouldBindJSON(&Activity); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := h.DB.Save(&Activity).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, Activity)
}

// GetAllBaby                godoc
// @Summary      Get all latest activities of baby
// @Description  Returns all baby's activities.
// @Tags         Activities
// @Produce      json
// @Param        baby_id  path      string  true  "search activity by baby_id"
// @Security     ApiKeyAuth
// @Success      200  {object}  Activity
// @Router       /activities/{baby_id}/latest [get]
func (h *DbHandler) GetLatestActivity(c *gin.Context) {
	Activity := []Activity{}
	baby_id := c.Param("baby_id")
	h.DB.Find(&Activity, "baby_id", baby_id)
	c.JSON(http.StatusOK, Activity)
}
