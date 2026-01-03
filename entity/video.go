package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Person struct {
	ID    uuid.UUID `json:"id" xml:"id" form:"id" gorm:"type:text;primaryKey"`
	Name  string    `json:"name" xml:"name" form:"name" binding:"required,min=2,max=50"`
	Age   int       `json:"age" xml:"age" form:"age" binding:"gte=0,lte=120"`
	Email string    `json:"email" xml:"email" form:"email" binding:"required,email"`
	Model
}

// BeforeCreate hook to generate UUID before creating a Person
func (p *Person) BeforeCreate(tx *gorm.DB) error {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	return nil
}

type Video struct {
	ID          uuid.UUID `json:"id" xml:"id" form:"id" gorm:"type:text;primaryKey"`
	Title       string    `json:"title" xml:"title" form:"title" binding:"min=3,max=100" gorm:"type:varchar(100)"`
	Description string    `json:"description" xml:"description" form:"description" binding:"max=500" gorm:"type:varchar(500)"`
	URL         string    `json:"url" xml:"url" form:"url" binding:"required,url" gorm:"type:varchar(255)"`
	Author      Person    `json:"author" xml:"author" form:"author" binding:"required" gorm:"foreignKey:AuthorID;references:ID"`
	AuthorID    uuid.UUID `json:"-" xml:"-" form:"-" gorm:"type:text"`
	Model
}

// BeforeCreate hook to generate UUID before creating a Video
func (v *Video) BeforeCreate(tx *gorm.DB) error {
	if v.ID == uuid.Nil {
		v.ID = uuid.New()
	}
	return nil
}
