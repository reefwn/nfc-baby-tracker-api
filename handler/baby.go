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

// SaveBaby             godoc
// @Summary      Store a new baby
// @Description  Takes a baby JSON and store in DB. Return saved JSON.
// @Tags         Babies
// @Produce      json
// @Param        baby  body      Baby  true  "Baby JSON"
// @Success      200   {object}  Baby
// @Router       /babies [post]
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

// GetAllBaby                godoc
// @Summary      Get all babies
// @Description  Returns all babies.
// @Tags         Babies
// @Produce      json
// @Success      200  {object}  Baby
// @Router       /babies [get]
func (h *DbHandler) GetAllBaby(c *gin.Context) {
	Babys := []Baby{}
	h.DB.Find(&Babys)
	c.JSON(http.StatusOK, Babys)
}

// GetBaby                godoc
// @Summary      Get single baby by id
// @Description  Returns the baby whose id value matches the id.
// @Tags         Babies
// @Produce      json
// @Param        id  path      string  true  "search baby by id"
// @Success      200  {object}  Baby
// @Router       /babies/{id} [get]
func (h *DbHandler) GetBaby(c *gin.Context) {
	id := c.Param("id")
	Baby := Baby{}

	if err := h.DB.Find(&Baby, "id", id).Error; err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, Baby)
}
