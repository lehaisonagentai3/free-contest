package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/lehaisonagentai3/free-contest/backend/internal/config"
	"github.com/lehaisonagentai3/free-contest/backend/internal/controller"
	"github.com/lehaisonagentai3/free-contest/backend/internal/service"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/lehaisonagentai3/free-contest/backend/docs" // Import swagger docs
)

// @title Free Contest API
// @version 1.0
// @description This is a free contest API server for managing officer tests and contests.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8298
func main() {
	// Load configuration
	conf, err := config.LoadAppConfig("config.json")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize service
	contestService, err := service.NewContestService(conf)
	if err != nil {
		log.Fatalf("Failed to initialize contest service: %v", err)
	}

	contest := contestService.GetContestInfo()
	fmt.Println("Contest loaded successfully!")
	fmt.Printf("Contest ID: %d\n", contest.ID)
	fmt.Printf("Contest Name: %s\n", contest.Name)
	fmt.Printf("Contest Folder Path: %s\n", contest.FolderPath)
	fmt.Println("Total Subjects:", len(contest.Subjects))
	for _, subject := range contest.Subjects {
		fmt.Printf("Subject ID: %d, Name: %s, Description: %s, Total Chapters: %d, Test time: %d, Total Questions In Test: %d\n", subject.ID, subject.Name, subject.Description, len(subject.Chapters), subject.TestTime, subject.NumQuestionTest)
		for _, chapter := range subject.Chapters {
			fmt.Printf("Chapter ID: %d, Name: %s, NumQuestionTest: %d, Total question: %d\n", chapter.ID, chapter.Name, chapter.NumQuestionTest, chapter.TotalQuestions)
		}
	}
	fmt.Println("Contest service initialized successfully!")

	// Initialize controllers
	testController := controller.NewTestController(contestService)
	unitController := controller.NewUnitController(contestService)
	officerController := controller.NewOfficerController(contestService)
	subjectController := controller.NewSubjectController(contestService)

	// Initialize Gin router
	router := gin.Default()

	// Configure CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // In production, specify exact origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// API routes
	v1 := router.Group("/api/v1")
	{
		tests := v1.Group("/tests")
		{
			tests.GET("/officer-subject", testController.GetSubjectTestForOfficer)
			tests.POST("/start", testController.StartTest)
			tests.POST("/submit", testController.SubmitTest)
		}

		// Units routes
		v1.GET("/units", unitController.GetAllUnits)

		// Officers routes
		v1.GET("/officers", officerController.GetAllOfficers)
		v1.GET("/officers/:id", officerController.GetOfficerByID)

		// Subjects routes
		v1.GET("/subjects", subjectController.GetAllSubjects)
	}

	// Swagger documentation route
	router.GET("/v1/api/free-contest/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "OK",
			"message": "Free Contest API is running",
		})
	})

	// Start the server
	log.Println("Starting server on port 8298...")
	log.Println("Swagger documentation available at: http://localhost:8298/v1/api/free-contest/swagger/index.html")

	if err := router.Run(":8298"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
