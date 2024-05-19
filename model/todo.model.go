package models

import (
	"time"
)

type ToDo struct {
	Id    uint `json:"todoId" gorm:"primaryKey"`
	Text     string    `json:"text"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	IsDel int64 `json:"isDel"`
}