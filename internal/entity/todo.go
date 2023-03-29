package entity

import "time"

type Todo struct {
	ID              int64     `json:"id" gorm:"primary_key;auto_increment"`
	ActivityGroupID int64     `json:"activity_group_id" gorm:"not null;" sql:"type:int REFERENCES activities(id)"`
	Title           string    `json:"title" gorm:"size:255;not null;"`
	IsActive        bool      `json:"is_active" gorm:"not null; default:true"`
	Priority        string    `json:"priority" gorm:"size:255;not null;"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}
