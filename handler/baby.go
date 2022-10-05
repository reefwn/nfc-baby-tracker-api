package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Baby struct {
	Id        uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
	DOB       time.Time `json:"dob" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"default:now();not null"`
}

func (h *DbHandler) SaveBaby(c *gin.Context) {
	Baby := Baby{}

	if err := c.ShouldBindJSON(&Baby); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := h.DB.Save(&Baby).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, Baby)
}

func (h *DbHandler) GetAllBaby(c *gin.Context) {
	Babys := []Baby{}
	h.DB.Find(&Babys)
	c.JSON(http.StatusOK, Babys)
}

func (h *DbHandler) GetBaby(c *gin.Context) {
	id := c.Param("id")
	Baby := Baby{}

	if err := h.DB.Find(&Baby, "id", id).Error; err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, Baby)
}
