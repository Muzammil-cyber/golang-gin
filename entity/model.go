package entity

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	CreatedAt time.Time      `json:"created_at,omitempty" xml:"created_at" form:"created_at" swaggerignore:"true"`
	UpdatedAt time.Time      `json:"updated_at,omitempty" xml:"updated_at" form:"updated_at" swaggerignore:"true"`
	DeletedAt gorm.DeletedAt `json:"-" xml:"-" form:"-" gorm:"index"`
}
