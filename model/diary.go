package model

import "time"

type Diary struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User			User	`json:"user" gorm:"foreignKey:UserID; constraint:OnDelete:CASCADE;"`
	UserID    uint   `json:"user_id"`
}

type DiaryResponse struct {
	ID				uint   `json:"id" gorm:"primaryKey"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	UpdatedAt time.Time `json:"updated_at"`
}