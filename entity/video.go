package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Person represents a person entity
type Person struct {
	ID    uuid.UUID `json:"id,omitempty" xml:"id" form:"id" gorm:"type:text;primaryKey" example:"123e4567-e89b-12d3-a456-426614174000" swaggerignore:"true"` // Person ID (auto-generated, omit in create requests)
	Name  string    `json:"name" xml:"name" form:"name" binding:"required,min=2,max=50" example:"John Doe"`                             // Person name (2-50 characters)
	Age   int       `json:"age" xml:"age" form:"age" binding:"gte=0,lte=120" example:"30"`                                              // Person age (0-120)
	Email string    `json:"email" xml:"email" form:"email" binding:"required,email" example:"john.doe@example.com"`                     // Person email
	Model
}

// BeforeCreate hook to generate UUID before creating a Person
func (p *Person) BeforeCreate(tx *gorm.DB) error {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	return nil
}

// Video represents a video entity
type Video struct {
	ID          uuid.UUID `json:"id,omitempty" xml:"id" form:"id" gorm:"type:text;primaryKey" example:"123e4567-e89b-12d3-a456-426614174001" swaggerignore:"true"`                // Video ID (auto-generated, omit in create requests)
	Title       string    `json:"title" xml:"title" form:"title" binding:"min=3,max=100" gorm:"type:varchar(100)" example:"Introduction to Golang"`          // Video title (3-100 characters)
	Description string    `json:"description" xml:"description" form:"description" binding:"max=500" gorm:"type:varchar(500)" example:"Learn Golang basics"` // Video description (max 500 characters)
	URL         string    `json:"url" xml:"url" form:"url" binding:"required,url" gorm:"type:varchar(255)" example:"https://www.youtube.com/watch?v=abc"`    // Video URL
	Author      Person    `json:"author" xml:"author" form:"author" binding:"required" gorm:"foreignKey:AuthorID;references:ID"`                             // Video author
	AuthorID    uuid.UUID `json:"-" xml:"-" form:"-" gorm:"type:text"`                                                                                       // Author ID (foreign key)
	Model
}

// BeforeCreate hook to generate UUID before creating a Video
func (v *Video) BeforeCreate(tx *gorm.DB) error {
	if v.ID == uuid.Nil {
		v.ID = uuid.New()
	}
	return nil
}
