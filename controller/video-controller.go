package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/muzammil-cyber/golang-gin/dto"
	"github.com/muzammil-cyber/golang-gin/entity"
	"github.com/muzammil-cyber/golang-gin/service"
	"github.com/muzammil-cyber/golang-gin/utils"
	"github.com/muzammil-cyber/golang-gin/validators"
)

type VideoController interface {
	Save(ctx *gin.Context)
	GetAll(ctx *gin.Context) []entity.Video
	ShowAll(ctx *gin.Context)
	GetByID(ctx *gin.Context) entity.Video
	Update(ctx *gin.Context) entity.Video
	Delete(ctx *gin.Context) error
}

type controller struct {
	videoService service.VideoService
}

var validate *validator.Validate

func New(videoService service.VideoService) VideoController {
	validate = validator.New()
	validate.RegisterValidation("is-idx", validators.IsIdx)

	return &controller{
		videoService: videoService,
	}
}

// Save godoc
// @Summary Create a new video
// @Description Create a new video entry with associated author information. Requires JWT authentication.
// @Tags Videos
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token" default(Bearer <token>)
// @Param video body dto.VideoCreateRequest true "Video object with nested author information"
// @Success 200 {object} entity.Video "Successfully created video with generated ID"
// @Failure 400 {object} dto.ValidationErrorResponse "Invalid request format or validation errors"
// @Failure 401 {object} dto.ErrorResponse "Unauthorized - valid JWT token required"
// @Failure 500 {object} dto.ErrorResponse "Internal server error while saving video"
// @Security BearerAuth
// @Router /api/videos [post]
func (c *controller) Save(ctx *gin.Context) {
	var video dto.VideoCreateRequest
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ValidationErrorResponse{Errors: utils.FormatValidationError(err)})
		return
	}
	err = validate.Struct(video)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ValidationErrorResponse{Errors: utils.FormatValidationError(err)})
		return
	}
	savedVideo, err := c.videoService.Save(video)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}
	ctx.JSON(200, savedVideo)
}

// GetAll godoc
// @Summary Get all videos
// @Description Retrieve a complete list of all videos with their author information. Requires JWT authentication.
// @Tags Videos
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token" default(Bearer <token>)
// @Success 200 {array} entity.Video "List of all videos with author details"
// @Failure 401 {object} dto.ErrorResponse "Unauthorized - valid JWT token required"
// @Failure 500 {object} dto.ErrorResponse "Internal server error while fetching videos"
// @Security BearerAuth
// @Router /api/videos [get]
func (c *controller) GetAll(ctx *gin.Context) []entity.Video {
	videos, err := c.videoService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return nil
	}
	return videos
}

// ShowAll godoc
// @Summary Show all videos (HTML view)
// @Description Display all videos in an HTML template for browser viewing. This endpoint is public and does not require authentication.
// @Tags Views
// @Accept html
// @Produce html
// @Success 200 {string} html "HTML page displaying all videos"
// @Failure 500 {object} dto.ErrorResponse "Internal server error while loading videos"
// @Router /view/ [get]
func (c *controller) ShowAll(ctx *gin.Context) {
	videos, err := c.videoService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"title":  "Video List",
		"videos": videos,
	})
}

// GetByID godoc
// @Summary Get video by ID
// @Description Retrieve detailed information for a specific video by its unique ID. Requires JWT authentication.
// @Tags Videos
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token" default(Bearer <token>)
// @Param id path string true "Video UUID" format(uuid)
// @Success 200 {object} entity.Video "Video details with author information"
// @Failure 401 {object} dto.ErrorResponse "Unauthorized - valid JWT token required"
// @Failure 404 {object} dto.ErrorResponse "Video not found with provided ID"
// @Failure 500 {object} dto.ErrorResponse "Internal server error while fetching video"
// @Security BearerAuth
// @Router /api/videos/{id} [get]
func (c *controller) GetByID(ctx *gin.Context) entity.Video {
	id := ctx.Param("id")
	video, err := c.videoService.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return entity.Video{}
	}
	return *video
}

// Update godoc
// @Summary Update a video
// @Description Update an existing video's information by its ID. All fields in the video object can be updated. Requires JWT authentication.
// @Tags Videos
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token" default(Bearer <token>)
// @Param id path string true "Video UUID" format(uuid)
// @Param video body entity.Video true "Updated video object with new information"
// @Success 200 {object} entity.Video "Successfully updated video"
// @Failure 400 {object} dto.ValidationErrorResponse "Invalid request format or validation errors"
// @Failure 401 {object} dto.ErrorResponse "Unauthorized - valid JWT token required"
// @Failure 404 {object} dto.ErrorResponse "Video not found with provided ID"
// @Failure 500 {object} dto.ErrorResponse "Internal server error while updating video"
// @Security BearerAuth
// @Router /api/videos/{id} [put]
func (c *controller) Update(ctx *gin.Context) entity.Video {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ValidationErrorResponse{Errors: utils.FormatValidationError(err)})
		return entity.Video{}
	}
	id := ctx.Param("id")
	video.ID = utils.ParseUUID(id)
	err = validate.Struct(video)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ValidationErrorResponse{Errors: utils.FormatValidationError(err)})
		return entity.Video{}
	}
	updatedVideo, err := c.videoService.Update(video)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return entity.Video{}
	}
	return updatedVideo
}

// Delete godoc
// @Summary Delete a video
// @Description Permanently delete a video by its ID. This action cannot be undone. Requires JWT authentication.
// @Tags Videos
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token" default(Bearer <token>)
// @Param id path string true "Video UUID" format(uuid)
// @Success 200 {object} dto.MessageResponse "Video successfully deleted"
// @Failure 401 {object} dto.ErrorResponse "Unauthorized - valid JWT token required"
// @Failure 404 {object} dto.ErrorResponse "Video not found with provided ID"
// @Failure 500 {object} dto.ErrorResponse "Internal server error while deleting video"
// @Security BearerAuth
// @Router /api/videos/{id} [delete]
func (c *controller) Delete(ctx *gin.Context) error {
	id := ctx.Param("id")
	err := c.videoService.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return err
	}
	return nil
}
