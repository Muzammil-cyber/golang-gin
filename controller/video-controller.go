package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
func (c *controller) Save(ctx *gin.Context) {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}
	err = validate.Struct(video)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": utils.FormatValidationError(err)})
		return
	}
	savedVideo, err := c.videoService.Save(video)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
		return
	}
	ctx.JSON(200, savedVideo)
}

func (c *controller) GetAll(ctx *gin.Context) []entity.Video {
	videos, err := c.videoService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
		return nil
	}
	return videos
}

func (c *controller) ShowAll(ctx *gin.Context) {
	videos, err := c.videoService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
		return
	}
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"title":  "Video List",
		"videos": videos,
	})
}

func (c *controller) GetByID(ctx *gin.Context) entity.Video {
	id := ctx.Param("id")
	video, err := c.videoService.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
		return entity.Video{}
	}
	return *video
}

func (c *controller) Update(ctx *gin.Context) entity.Video {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return entity.Video{}
	}
	id := ctx.Param("id")
	video.ID = utils.ParseUUID(id)
	err = validate.Struct(video)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": utils.FormatValidationError(err)})
		return entity.Video{}
	}
	updatedVideo, err := c.videoService.Update(video)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
		return entity.Video{}
	}
	return updatedVideo
}

func (c *controller) Delete(ctx *gin.Context) error {
	id := ctx.Param("id")
	err := c.videoService.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
		return err
	}
	return nil
}
