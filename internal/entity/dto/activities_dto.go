package dto

import "time"

type ActivitiesDTO struct {
	Title     string    `json:"title"`
	Email     string    `json:"email" binding:"email"`
	CreatedAt time.Time `json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}
