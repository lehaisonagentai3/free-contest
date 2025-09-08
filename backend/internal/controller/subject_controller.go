package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lehaisonagentai3/free-contest/backend/internal/model"
	_ "github.com/lehaisonagentai3/free-contest/backend/internal/model"
	"github.com/lehaisonagentai3/free-contest/backend/internal/service"
)

type SubjectController struct {
	contestService *service.ContestService
}

func NewSubjectController(contestService *service.ContestService) *SubjectController {
	return &SubjectController{
		contestService: contestService,
	}
}

type ListSubjectResponse struct {
	Data    []*model.Subject `json:"data,omitempty"`
	Count   int              `json:"count,omitempty"`
	Status  string           `json:"status,omitempty"`
	Message string           `json:"message,omitempty"`
}

// GetAllSubjects returns all subjects without questions and chapters
// @Summary Get all subjects
// @Description Get list of all subjects with basic information only (no questions or chapters)
// @Tags subjects
// @Accept json
// @Produce json
// @Success 200 {object} ListSubjectResponse "success response with subjects list"
// @Failure 500 {object} map[string]interface{} "internal server error"
// @Router /api/v1/subjects [get]
func (c *SubjectController) GetAllSubjects(ctx *gin.Context) {
	subjects := c.contestService.GetAllSubjects()

	ctx.JSON(http.StatusOK, ListSubjectResponse{
		Data:    subjects,
		Count:   len(subjects),
		Status:  "success",
		Message: "Subjects retrieved successfully",
	})

}
