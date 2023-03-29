package dto

import "time"

type TodoDTO struct {
	Title           string    `json:"title"`
	ActivityGroupID int       `json:"activity_group_id"`
	IsActive        bool      `json:"is_active"`
	Status          string    `json:"status"`
	Priority        string    `json:"priority"`
	CreatedAt       time.Time `json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt       time.Time `json:"updatedAt" gorm:"TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}
