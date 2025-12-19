package main

import (
	"fmt"
	"log"
	"math-app/internal/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("math_app.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	var task models.Task
	result := db.First(&task, 73)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	fmt.Printf("Task ID: %d\n", task.ID)
	fmt.Printf("Task Title: %s\n", task.Title)
	fmt.Printf("Task Code:\n%s\n", task.Code)
}
