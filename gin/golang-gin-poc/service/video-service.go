package service

import "github.com/beautytiger/go-playground/gin/golang-gin-poc/entity"

type VideoService interface {
	Save(video entity.Video) entity.Video
	FindAll() []entity.Video
}

type videoService struct {
	video []entity.Video
}

func New() VideoService {
	return &videoService{}
}

func (s *videoService) Save(video entity.Video) entity.Video {
	s.video = append(s.video, video)
	return video
}

func (s *videoService) FindAll() []entity.Video {
	return s.video
}
