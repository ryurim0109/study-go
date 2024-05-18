package models

import (
	"time"
)

type ToDo struct {
	TodoId    uint `json:"id" gorm:"primaryKey"`
	Text     string    `json:"text"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}