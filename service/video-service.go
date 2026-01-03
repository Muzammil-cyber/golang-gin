package service

import (
	"github.com/muzammil-cyber/golang-gin/entity"
	"github.com/muzammil-cyber/golang-gin/repository"
)

type VideoService interface {
	Save(entity.Video) (entity.Video, error)
	GetAll() ([]entity.Video, error)
	GetByID(string) (*entity.Video, error)
	Update(entity.Video) (entity.Video, error)
	Delete(string) error
}

type videoService struct {
	videos repository.VideoRepository
}

func New(repo repository.VideoRepository) VideoService {
	return &videoService{
		videos: repo,
	}
}

func (s *videoService) Save(video entity.Video) (entity.Video, error) {
	err := s.videos.Save(&video)
	return video, err
}

func (s *videoService) GetAll() ([]entity.Video, error) {
	return s.videos.FindAll()
}

func (s *videoService) GetByID(id string) (*entity.Video, error) {
	return s.videos.FindByID(id)
}

func (s *videoService) Update(video entity.Video) (entity.Video, error) {
	err := s.videos.Update(&video)
	return video, err
}

func (s *videoService) Delete(id string) error {
	return s.videos.Delete(id)
}
