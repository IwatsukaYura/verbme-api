package model

import "time"

type Word struct {
	ID				uint   `json:"id" gorm:"primaryKey"`
	Word			string `json:"word"`
	Meaning		string `json:"meaning"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User			User	`json:"user" gorm:"foreignKey:UserID; constraint:OnDelete:CASCADE;"`
	UserID    uint   `json:"user_id"`
	Diary			Diary	`json:"diary" gorm:"foreignKey:DiaryID;"`
	DiaryID   uint   `json:"diary_id"`
	Completed bool   `json:"completed"`
}

type WordResponse struct {
	ID				uint   `json:"id" gorm:"primaryKey"`
	Word			string `json:"word"`
	Meaning		string `json:"meaning"`
	UpdatedAt time.Time `json:"updated_at"`
	Completed bool   `json:"completed"`
}