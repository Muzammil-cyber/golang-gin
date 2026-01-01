package service

import "github.com/muzammil-cyber/golang-gin/entity"

type VideoService interface {
	Save(entity.Video) entity.Video
	GetAll() []entity.Video
}

type videoService struct {
	videos []entity.Video
}

func New() VideoService {
	return &videoService{}
}

func (s *videoService) Save(video entity.Video) entity.Video {
	s.videos = append(s.videos, video)
	return video
}

func (s *videoService) GetAll() []entity.Video {
	return s.videos
}
