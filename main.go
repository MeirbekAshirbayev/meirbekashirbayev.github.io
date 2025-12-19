package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"math-app/internal/database"
	"math-app/internal/handlers"
	"math-app/internal/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()

	r := gin.Default()

	// Apply Global Security Middleware
	r.Use(middleware.RateLimitMiddleware())
	r.Use(middleware.SecurityHeadersMiddleware())

	// Register Custom Functions BEFORE LoadHTMLGlob
	r.SetFuncMap(template.FuncMap{
		"safe": func(s string) template.HTML {
			return template.HTML(s)
		},
	})

	r.LoadHTMLGlob("templates/*")

	// Public Routes
	r.GET("/", handlers.HomeHandler)
	r.GET("/grade/:num", handlers.GradeListHandler)
	r.GET("/lesson/:id", handlers.LessonHandler)
	r.GET("/task/:id", handlers.TaskHandler)
	r.GET("/privacy", func(c *gin.Context) {
		c.HTML(http.StatusOK, "privacy.html", nil)
	})
	r.GET("/terms", func(c *gin.Context) {
		c.HTML(http.StatusOK, "terms.html", nil)
	})

	// Admin Routes
	adminUser := os.Getenv("ADMIN_USER")
	if adminUser == "" {
		adminUser = "aselsyrym"
	}
	adminPass := os.Getenv("ADMIN_PASS")
	if adminPass == "" {
		adminPass = "ushermiko1234!"
		fmt.Println("WARNING: Using default admin password. Set ADMIN_PASS environment variable in production!")
	}

	admin := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		adminUser: adminPass,
	}))

	admin.Use(middleware.CsrfMiddleware())

	{
		admin.GET("/", handlers.AdminDashboardHandler)
		admin.GET("/add", handlers.AdminAddFormHandler)
		admin.POST("/add", handlers.AdminAddHandler)
		admin.GET("/edit/:id", handlers.AdminEditFormHandler)
		admin.POST("/edit/:id", handlers.AdminEditHandler)
		admin.POST("/delete/:id", handlers.AdminDeleteHandler)
		// Export Static Site
		admin.POST("/export", handlers.AdminExportHandler)
	}

	fmt.Println("Server running on http://localhost:8080")

	// Preview Server (for static export)
	go func() {
		rStatic := gin.Default()
		// Serve the export directory at root
		rStatic.Static("/", "./site_export")
		fmt.Println("Static Preview running on http://localhost:8081")
		rStatic.Run(":8081")
	}()

	r.Run(":8080")
}
