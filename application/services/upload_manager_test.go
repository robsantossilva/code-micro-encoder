package services_test

import (
	"encoder/application/services"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

func init() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func TestVideoServiceUpload(t *testing.T) {

	// var err error

	// video, repo := prepare()

	// videoService := services.NewVideoService()
	// videoService.Video = video
	// videoService.VideoRepository = repo

	// err = videoService.Download("storage-micro-videos")
	// require.Nil(t, err)

	// err = videoService.Fragment()
	// require.Nil(t, err)

	// err = videoService.Encode()
	// require.Nil(t, err)

	// err = videoService.Finish()
	// require.Nil(t, err)

	// Upload Test
	videoUpload := services.NewVideoUpload()
	videoUpload.OutputBucket = "storage-micro-videos"
	//videoUpload.VideoPath = os.Getenv("localStoragePath") + "/" + video.ID
	videoUpload.VideoPath = os.Getenv("localStoragePath") + "/6fab542f-10ca-43a1-9f44-b0fa7c4eb27b"

	doneUpload := make(chan string)
	go videoUpload.ProcessUpload(4, doneUpload)

	result := <-doneUpload
	require.Equal(t, "Upload completed", result)

}
