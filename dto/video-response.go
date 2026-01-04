package dto

import "github.com/muzammil-cyber/golang-gin/entity"

// VideoResponse represents a single video response
type VideoResponse struct {
	entity.Video
}

// VideosResponse represents a list of videos response
type VideosResponse struct {
	Videos []entity.Video `json:"videos"` // List of videos
}

// MessageResponse represents a success message response
type MessageResponse struct {
	Message string `json:"message" example:"Video deleted successfully"` // Success message
}
