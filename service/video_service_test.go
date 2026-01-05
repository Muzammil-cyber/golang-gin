package service_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/muzammil-cyber/golang-gin/dto"
	"github.com/muzammil-cyber/golang-gin/entity"
	"github.com/muzammil-cyber/golang-gin/repository"
	"github.com/muzammil-cyber/golang-gin/service"
)

var testVideo = dto.VideoCreateRequest{
	Title:       "Test Video",
	Description: "This is a test video",
	URL:         "https://www.example.com/test-video",
	Author: entity.Person{
		Name:  "Test Author",
		Email: "example@gmail.com",
		Age:   23,
	},
}

var savedVideo entity.Video

var _ = Describe("VideoService", func() {
	var (
		videoService    service.VideoService
		videoRepository repository.VideoRepository
	)

	BeforeEach(func() {
		videoRepository = repository.NewVideoRepository()
		videoService = service.New(videoRepository)
	})

	Describe("Save", func() {
		It("should save a video successfully", func() {
			createdVideo, err := videoService.Save(testVideo)
			Expect(err).To(BeNil())
			Expect(createdVideo.ID).NotTo(BeEmpty())
			Expect(createdVideo.Title).To(Equal(testVideo.Title))
			Expect(createdVideo.Description).To(Equal(testVideo.Description))
			Expect(createdVideo.URL).To(Equal(testVideo.URL))
			Expect(createdVideo.Author.Name).To(Equal(testVideo.Author.Name))
			savedVideo = createdVideo
		})
	})

	Describe("GetAll", func() {
		It("should retrieve at least 1 video", func() {
			videos, err := videoService.GetAll()
			Expect(err).To(BeNil())
			Expect(len(videos)).To(BeNumerically(">", 0))
		})
	})

	Describe("GetByID", func() {
		It("should retrieve the saved video by ID", func() {
			video, err := videoService.GetByID(savedVideo.ID.String())
			Expect(err).To(BeNil())
			Expect(video).NotTo(BeNil())
			Expect(video.ID).To(Equal(savedVideo.ID))
			Expect(video.Title).To(Equal(savedVideo.Title))
		})
	})
	Describe("Update", func() {
		It("should update the video's title", func() {
			savedVideo.Title = "Updated Test Video"
			updatedVideo, err := videoService.Update(savedVideo)
			Expect(err).To(BeNil())
			Expect(updatedVideo.Title).To(Equal("Updated Test Video"))
		})
	})

	Describe("Delete", func() {
		It("should delete the saved video", func() {
			err := videoService.Delete(savedVideo.ID.String())
			Expect(err).To(BeNil())

			deletedVideo, err := videoService.GetByID(savedVideo.ID.String())
			Expect(err).NotTo(BeNil())
			Expect(deletedVideo).To(BeNil())
		})
	})
})
