package model

import (
	"time"
)

type Common struct {
	ID        uint       `json:"id" gorm:"primarykey;autoIncrement"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"index"`
}
