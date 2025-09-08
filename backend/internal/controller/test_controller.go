package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lehaisonagentai3/free-contest/backend/internal/model"
	_ "github.com/lehaisonagentai3/free-contest/backend/internal/model"
	"github.com/lehaisonagentai3/free-contest/backend/internal/service"
)

type TestController struct {
	contestService *service.ContestService
}

func NewTestController(contestService *service.ContestService) *TestController {
	return &TestController{
		contestService: contestService,
	}
}

type TestResponse struct {
	Data    *model.Test `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	Status  string      `json:"status,omitempty"`
}

type SubmissionResponse struct {
	Data    *model.Submission `json:"data,omitempty"`
	Message string            `json:"message,omitempty"`
	Status  string            `json:"status,omitempty"`
}

// GetSubjectTestForOfficer godoc
// @Summary Get a test for an officer in a specific subject
// @Description Creates and returns a test for an officer in a specific subject with random questions
// @Tags Tests
// @Accept json
// @Produce json
// @Param officerID query int true "Officer ID"
// @Param subjectID query int true "Subject ID"
// @Success 200 {object} TestResponse "Test created successfully or existing test returned"
// @Failure 400 {object} map[string]string "Bad request - missing or invalid parameters"
// @Failure 404 {object} map[string]string "Officer or subject not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/v1/tests/officer-subject [get]
func (tc *TestController) GetSubjectTestForOfficer(c *gin.Context) {
	// Get query parameters
	officerIDStr := c.Query("officerID")
	subjectIDStr := c.Query("subjectID")

	// Validate parameters
	if officerIDStr == "" || subjectIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Both officerID and subjectID query parameters are required",
		})
		return
	}

	// Parse officerID
	officerID, err := strconv.Atoi(officerIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid officerID: must be a valid integer",
		})
		return
	}

	// Parse subjectID
	subjectID, err := strconv.Atoi(subjectIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid subjectID: must be a valid integer",
		})
		return
	}

	// Call service to get the test
	test, err := tc.contestService.GetSubjectTestForOfficer(officerID, subjectID)
	if err != nil {
		// Handle different types of errors
		switch err.Error() {
		case "officer not found":
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
		case "subject not found":
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
		return
	}

	// Return the test
	c.JSON(http.StatusOK, TestResponse{
		Data:    test,
		Message: "Test retrieved successfully",
		Status:  "success",
	})
}

// StartTest godoc
// @Summary Start a test for an officer
// @Description Starts a test by setting the start time for the specified officer and test
// @Tags Tests
// @Accept json
// @Produce json
// @Param officerID query int true "Officer ID"
// @Param testID query int true "Test ID"
// @Success 200 {object} TestResponse "Test started successfully"
// @Failure 400 {object} map[string]string "Bad request - missing or invalid parameters"
// @Failure 404 {object} map[string]string "Test not found"
// @Failure 409 {object} map[string]string "Test already started"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/v1/tests/start [post]
func (tc *TestController) StartTest(c *gin.Context) {
	// Get query parameters
	officerIDStr := c.Query("officerID")
	testIDStr := c.Query("testID")

	// Validate parameters
	if officerIDStr == "" || testIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Both officerID and testID query parameters are required",
		})
		return
	}

	// Parse officerID
	officerID, err := strconv.Atoi(officerIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid officerID: must be a valid integer",
		})
		return
	}

	// Parse testID
	testID, err := strconv.Atoi(testIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid testID: must be a valid integer",
		})
		return
	}

	// Call service to start the test
	test, err := tc.contestService.StartTest(officerID, testID)
	if err != nil {
		// Handle different types of errors
		switch err.Error() {
		case "no tests found":
			c.JSON(http.StatusNotFound, gin.H{
				"error": "No tests found",
			})
		case "no tests found for officer":
			c.JSON(http.StatusNotFound, gin.H{
				"error": "No tests found for officer",
			})
		case "test not found":
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Test not found",
			})
		case "test already started":
			c.JSON(http.StatusConflict, gin.H{
				"error": "Test already started",
			})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
		return
	}

	// Return the started test
	c.JSON(http.StatusOK, TestResponse{
		Data:    test,
		Message: "Test started successfully",
		Status:  "success",
	})
}

// SubmitTest godoc
// @Summary Submit test answers
// @Description Submit answers for a test and get the calculated score
// @Tags Tests
// @Accept json
// @Produce json
// @Param officerID query int true "Officer ID"
// @Param testID query int true "Test ID"
// @Param submission body map[string]string true "Question ID to answer mapping"
// @Success 200 {object} SubmissionResponse "Test submitted successfully with score"
// @Failure 400 {object} map[string]string "Bad request - missing or invalid parameters"
// @Failure 404 {object} map[string]string "Test not found"
// @Failure 409 {object} map[string]string "Test already submitted or not started"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/v1/tests/submit [post]
func (tc *TestController) SubmitTest(c *gin.Context) {
	// Get query parameters
	officerIDStr := c.Query("officerID")
	testIDStr := c.Query("testID")

	// Validate parameters
	if officerIDStr == "" || testIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Both officerID and testID query parameters are required",
		})
		return
	}

	// Parse officerID
	officerID, err := strconv.Atoi(officerIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid officerID: must be a valid integer",
		})
		return
	}

	// Parse testID
	testID, err := strconv.Atoi(testIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid testID: must be a valid integer",
		})
		return
	}

	// Parse request body for answers
	var answers map[string]string
	if err := c.ShouldBindJSON(&answers); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body: must be a JSON object with question IDs as keys and answers as values",
		})
		return
	}

	// Validate that answers is not empty
	if len(answers) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Answers cannot be empty",
		})
		return
	}

	// Call service to submit the test
	submission, err := tc.contestService.SubmitTest(officerID, testID, answers)
	if err != nil {
		// Handle different types of errors
		switch err.Error() {
		case "no tests found":
			c.JSON(http.StatusNotFound, gin.H{
				"error": "No tests found",
			})
		case "no tests found for officer":
			c.JSON(http.StatusNotFound, gin.H{
				"error": "No tests found for officer",
			})
		case "test not found":
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Test not found",
			})
		case "test has not been started yet":
			c.JSON(http.StatusConflict, gin.H{
				"error": "Test has not been started yet",
			})
		case "test has already been submitted":
			c.JSON(http.StatusConflict, gin.H{
				"error": "Test has already been submitted",
			})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
		return
	}

	// Return the submission result
	c.JSON(http.StatusOK, SubmissionResponse{
		Data:    submission,
		Message: "Submission successful",
		Status:  "success",
	})
}
