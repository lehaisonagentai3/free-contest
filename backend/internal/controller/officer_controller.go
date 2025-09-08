package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lehaisonagentai3/free-contest/backend/internal/model"
	_ "github.com/lehaisonagentai3/free-contest/backend/internal/model"
	"github.com/lehaisonagentai3/free-contest/backend/internal/service"
)

type OfficerController struct {
	contestService *service.ContestService
}

func NewOfficerController(contestService *service.ContestService) *OfficerController {
	return &OfficerController{
		contestService: contestService,
	}
}

type ListOfficerResponse struct {
	Data    []*model.Officer `json:"data"`
	Count   int              `json:"count"`
	Message string           `json:"message"`
	Status  string           `json:"status"`
}

type OfficerResponse struct {
	Data    *model.Officer `json:"data"`
	Message string         `json:"message"`
	Status  string         `json:"status"`
}

// GetAllOfficers godoc
// @Summary Get all officers
// @Description Retrieves all officers in the system with their unit information
// @Tags Officers
// @Accept json
// @Produce json
// @Success 200 {array} ListOfficerResponse "List of officers with unit information"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/v1/officers [get]
func (oc *OfficerController) GetAllOfficers(c *gin.Context) {
	officers := oc.contestService.GetAllOfficers()
	c.JSON(http.StatusOK, ListOfficerResponse{
		Data:    officers,
		Count:   len(officers),
		Message: "Officers retrieved successfully",
		Status:  "success",
	})
}

// GetOfficerByID godoc
// @Summary Get officer by ID
// @Description Retrieves a specific officer by their ID with unit information
// @Tags Officers
// @Accept json
// @Produce json
// @Param id path int true "Officer ID"
// @Success 200 {object} OfficerResponse "Officer information with unit details"
// @Failure 400 {object} map[string]string "Bad request - invalid officer ID"
// @Failure 404 {object} map[string]string "Officer not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/v1/officers/{id} [get]
func (oc *OfficerController) GetOfficerByID(c *gin.Context) {
	// Get officer ID from URL parameter
	officerIDStr := c.Param("id")
	officerID, err := strconv.Atoi(officerIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid officer ID format",
		})
		return
	}

	// Get officer from service
	officer, err := oc.contestService.GetOfficerByID(officerID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, OfficerResponse{
		Data:    officer,
		Message: "Officer retrieved successfully",
		Status:  "success",
	})
}
