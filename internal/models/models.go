package models

type Lesson struct {
	ID            uint   `gorm:"primaryKey"`
	Grade         int    `gorm:"not null"`
	Title         string `gorm:"not null"`
	CanvaEmbedURL string `gorm:"not null"`
	Description   string // Post-refactor: Ensure this matches existing schema
	Tasks         []Task `gorm:"constraint:OnDelete:CASCADE;"`
}

type Task struct {
	ID          uint `gorm:"primaryKey"`
	LessonID    uint `gorm:"not null"`
	Title       string
	Description string
	Code        string `gorm:"type:text;not null"`
	Order       int
}
