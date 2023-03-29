package entity

import "time"

type Activities struct {
	ID        int64     `json:"id" gorm:"primary_key;auto_increment"`
	Title     string    `json:"title" gorm:"size:255;not null;"`
	Email     string    `json:"email" gorm:"size:255;not null;unique"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
