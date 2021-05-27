package repositories_test

import (
	"encoder/application/repositories"
	"encoder/domain"
	"encoder/framework/database"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestJobRepositoryDbInsert(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	//Defining repositories
	repoVideo := repositories.VideoRepositoryDb{Db: db}
	repoJob := repositories.JobRepositoryDb{Db: db}

	//Creating new video
	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()
	video, err := repoVideo.Insert(video)
	require.Nil(t, err)
	video, err = repoVideo.Find(video.ID)
	require.Nil(t, err)

	job, err := domain.NewJob("OutPut Test", "pending", video)
	require.Nil(t, err)
	require.Equal(t, video, job.Video)

	j, err := repoJob.Insert(job)
	require.Nil(t, err)

	j, err = repoJob.Find(j.ID)
	require.NotEmpty(t, j.ID)
	require.Nil(t, err)
	require.Equal(t, j.ID, job.ID)
	require.Equal(t, video.ID, j.Video.ID)
}

func TestJobRepositoryDbUpdate(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()
	repoVideo := repositories.VideoRepositoryDb{Db: db}
	video, err := repoVideo.Insert(video)
	require.Nil(t, err)
	video, err = repoVideo.Find(video.ID)
	require.Nil(t, err)

	job, err := domain.NewJob("OutPut Test", "pending", video)
	require.Nil(t, err)
	require.Equal(t, video, job.Video)
	repoJob := repositories.JobRepositoryDb{Db: db}
	j, err := repoJob.Insert(job)
	require.Nil(t, err)

	//Update Status
	j.Status = "Complete"
	repoJob.Update(j)

	j, err = repoJob.Find(j.ID)
	require.NotEmpty(t, j.ID)
	require.Nil(t, err)
	require.Equal(t, j.Status, job.Status)
}
