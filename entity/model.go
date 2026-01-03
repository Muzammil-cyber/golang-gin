package entity

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	CreatedAt time.Time      `json:"created_at" xml:"created_at" form:"created_at"`
	UpdatedAt time.Time      `json:"updated_at" xml:"updated_at" form:"updated_at" `
	DeletedAt gorm.DeletedAt `json:"-" xml:"-" form:"-" gorm:"index"`
}
