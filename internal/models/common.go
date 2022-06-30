package model

import (
	"time"
)

type Common struct {
	ID        int        `json:"id" gorm:"autoIncrement;primaryKey;`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"index"`
}
