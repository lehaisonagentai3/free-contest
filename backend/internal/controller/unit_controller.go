package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lehaisonagentai3/free-contest/backend/internal/model"
	"github.com/lehaisonagentai3/free-contest/backend/internal/service"
)

type UnitController struct {
	contestService *service.ContestService
}

func NewUnitController(contestService *service.ContestService) *UnitController {
	return &UnitController{
		contestService: contestService,
	}
}

type ListUnitResponse struct {
	Data    []string `json:"data,omitempty"`
	Count   int      `json:"count,omitempty"`
	Status  string   `json:"status,omitempty"`
	Message string   `json:"message,omitempty"`
}

// GetAllUnits godoc
// @Summary Get all units
// @Description Retrieves all units in the system
// @Tags Units
// @Accept json
// @Produce json
// @Success 200 {array} ListUnitResponse "List of units"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/v1/units [get]
func (uc *UnitController) GetAllUnits(c *gin.Context) {
	units := uc.contestService.GetAllUnits()
	c.JSON(http.StatusOK, ListUnitResponse{
		Data:    units,
		Count:   len(units),
		Status:  "success",
		Message: "Units retrieved successfully",
	})

}
