package dto

import "github.com/muzammil-cyber/golang-gin/entity"

// VideoCreateRequest represents the payload to create a video
// IDs and timestamps are omitted; they are generated server-side.
type VideoCreateRequest struct {
	Title       string        `json:"title" binding:"min=3,max=100" example:"Introduction to Golang"`
	Description string        `json:"description" binding:"max=500" example:"Learn Golang basics"`
	URL         string        `json:"url" binding:"required,url" example:"https://www.youtube.com/watch?v=abc"`
	Author      entity.Person `json:"author" binding:"required"`
}
