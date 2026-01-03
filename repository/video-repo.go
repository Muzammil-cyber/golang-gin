package repository

import (
	"github.com/muzammil-cyber/golang-gin/database/sqlite"
	"github.com/muzammil-cyber/golang-gin/entity"
	"gorm.io/gorm"
)

type VideoRepository interface {
	Save(video *entity.Video) error
	Update(video *entity.Video) error
	FindByID(id string) (*entity.Video, error)
	FindAll() ([]entity.Video, error)
	Delete(id string) error
}

type videoRepository struct {
	// db connection or any other dependencies can be added here
	db *gorm.DB
}

func NewVideoRepository() VideoRepository {
	sqliteDB, err := sqlite.NewSQLiteDB()
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}
	// Migrate the Video and Person schema
	err = sqliteDB.GetDB().AutoMigrate(&entity.Person{}, &entity.Video{})
	if err != nil {
		panic("Failed to migrate database schema: " + err.Error())
	}

	return &videoRepository{
		db: sqliteDB.GetDB(),
	}
}

// Implement the methods of VideoRepository interface here
func (r *videoRepository) Save(video *entity.Video) error {
	return r.db.Create(video).Error
}

func (r *videoRepository) Update(video *entity.Video) error {
	return r.db.Save(video).Error
}

func (r *videoRepository) FindByID(id string) (*entity.Video, error) {
	var video entity.Video
	if err := r.db.Preload("Author").First(&video, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &video, nil
}

func (r *videoRepository) FindAll() ([]entity.Video, error) {
	var videos []entity.Video
	if err := r.db.Preload("Author").Find(&videos).Error; err != nil {
		return nil, err
	}
	return videos, nil
}

func (r *videoRepository) Delete(id string) error {
	return r.db.Delete(&entity.Video{}, "id = ?", id).Error
}
