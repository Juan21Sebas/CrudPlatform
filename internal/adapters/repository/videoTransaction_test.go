package repository

import (
	model "CrudPlatform/internal/core/domain/repository/model/videos"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBDRepositoryVideo(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := &BDRepositoryVideo{db: db}
	ctx := &gin.Context{}

	t.Run("CreateVideo", func(t *testing.T) {
		video := &model.Videos{
			Title:       "Test Video",
			Description: "This is a test video",
		}

		mock.ExpectExec("INSERT INTO videos").
			WithArgs(sqlmock.AnyArg(), video.Title, video.Description, sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnResult(sqlmock.NewResult(1, 1))

		id, err := repo.CreateVideo(ctx, video)
		assert.NoError(t, err)
		assert.NotEmpty(t, id)
	})

	t.Run("CreateVideo_ExecError", func(t *testing.T) {
		video := &model.Videos{
			Title:       "Test Video",
			Description: "This is a test video",
		}

		mock.ExpectExec("INSERT INTO videos").
			WithArgs(sqlmock.AnyArg(), video.Title, video.Description, sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnError(fmt.Errorf("exec error"))

		id, err := repo.CreateVideo(ctx, video)
		assert.Error(t, err)
		assert.Empty(t, id)
		assert.Contains(t, err.Error(), "error executing statement: exec error")
	})

	t.Run("SelectVideo", func(t *testing.T) {
		request := &model.GetVideo{ID: "123"}
		rows := sqlmock.NewRows([]string{"title", "description", "created_at", "updated_at"}).
			AddRow("Test Video", "This is a test video", time.Now(), time.Now())

		mock.ExpectQuery("SELECT (.+) FROM videos WHERE id = \\$1").
			WithArgs(request.ID).
			WillReturnRows(rows)

		video, err := repo.SelectVideo(ctx, request)
		assert.NoError(t, err)
		assert.NotNil(t, video)
		assert.Equal(t, "Test Video", video.Title)
	})

	t.Run("SelectVideo_NotFound", func(t *testing.T) {
		request := &model.GetVideo{ID: "999"}

		mock.ExpectQuery("SELECT (.+) FROM videos WHERE id = \\$1").
			WithArgs(request.ID).
			WillReturnError(sql.ErrNoRows)

		video, err := repo.SelectVideo(ctx, request)
		assert.Error(t, err)
		assert.Nil(t, video)
		assert.Contains(t, err.Error(), "video with id 999 not found")
	})

	t.Run("SelectVideo_ScanError", func(t *testing.T) {
		request := &model.GetVideo{ID: "123"}

		rows := sqlmock.NewRows([]string{"title", "description", "created_at", "updated_at", "extra_column"}).
			AddRow("Test Video", "Test Description", time.Now(), time.Now(), "extra data")

		mock.ExpectQuery("SELECT (.+) FROM videos WHERE id = \\$1").
			WithArgs(request.ID).
			WillReturnRows(rows)

		video, err := repo.SelectVideo(ctx, request)
		assert.Error(t, err)
		assert.Nil(t, video)
		assert.Contains(t, err.Error(), "error scanning video row")
	})

	t.Run("UpdateVideo", func(t *testing.T) {
		request := &model.UpdateVideo{
			ID:          "123",
			Title:       "Updated Video",
			Description: "This is an updated test video",
		}

		mock.ExpectExec("UPDATE videos SET").
			WithArgs(request.Title, request.Description, sqlmock.AnyArg(), request.ID).
			WillReturnResult(sqlmock.NewResult(0, 1))

		rows := sqlmock.NewRows([]string{"title", "description", "updated_at"}).
			AddRow(request.Title, request.Description, time.Now())

		mock.ExpectQuery("SELECT (.+) FROM videos WHERE id = \\$1").
			WithArgs(request.ID).
			WillReturnRows(rows)

		video, err := repo.UpdateVideo(ctx, request)
		assert.NoError(t, err)
		assert.NotNil(t, video)
		assert.Equal(t, request.Title, video.Title)
	})

	t.Run("UpdateVideo_ExecError", func(t *testing.T) {
		request := &model.UpdateVideo{
			ID:          "123",
			Title:       "Updated Video",
			Description: "This is an updated test video",
		}

		mock.ExpectExec("UPDATE videos SET").
			WithArgs(request.Title, request.Description, sqlmock.AnyArg(), request.ID).
			WillReturnError(fmt.Errorf("exec error"))

		video, err := repo.UpdateVideo(ctx, request)
		assert.Error(t, err)
		assert.Nil(t, video)
		assert.Contains(t, err.Error(), "error executing update")
	})

	t.Run("UpdateVideo_NotFound", func(t *testing.T) {
		request := &model.UpdateVideo{
			ID:          "123",
			Title:       "Updated Video",
			Description: "This is an updated test video",
		}

		mock.ExpectExec("UPDATE videos SET").
			WithArgs(request.Title, request.Description, sqlmock.AnyArg(), request.ID).
			WillReturnResult(sqlmock.NewResult(0, 1))

		mock.ExpectQuery("SELECT (.+) FROM videos WHERE id = \\$1").
			WithArgs(request.ID).
			WillReturnError(sql.ErrNoRows)

		video, err := repo.UpdateVideo(ctx, request)
		assert.Error(t, err)
		assert.Nil(t, video)
		assert.Contains(t, err.Error(), "video with id 123 not found after update")
	})

	t.Run("UpdateVideo_ScanError", func(t *testing.T) {
		request := &model.UpdateVideo{
			ID:          "123",
			Title:       "Updated Video",
			Description: "This is an updated test video",
		}

		mock.ExpectExec("UPDATE videos SET").
			WithArgs(request.Title, request.Description, sqlmock.AnyArg(), request.ID).
			WillReturnResult(sqlmock.NewResult(0, 1))

		rows := sqlmock.NewRows([]string{"title", "description", "updated_at", "extra_column"}).
			AddRow(request.Title, request.Description, time.Now(), "extra data")

		mock.ExpectQuery("SELECT (.+) FROM videos WHERE id = \\$1").
			WithArgs(request.ID).
			WillReturnRows(rows)

		video, err := repo.UpdateVideo(ctx, request)
		assert.Error(t, err)
		assert.Nil(t, video)
		assert.Contains(t, err.Error(), "error scanning updated video row")
	})

	t.Run("DeleteVideo", func(t *testing.T) {
		request := &model.DeleteVideo{ID: "123"}

		mock.ExpectExec("DELETE FROM videos WHERE id = \\$1").
			WithArgs(request.ID).
			WillReturnResult(sqlmock.NewResult(0, 1))

		err := repo.DeleteVideo(ctx, request)
		assert.NoError(t, err)
	})

	t.Run("DeleteVideo_NotFound", func(t *testing.T) {
		request := &model.DeleteVideo{ID: "999"}

		mock.ExpectExec("DELETE FROM videos WHERE id = \\$1").
			WithArgs(request.ID).
			WillReturnResult(sqlmock.NewResult(0, 0))

		err := repo.DeleteVideo(ctx, request)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "video with id 999 not found")
	})

	t.Run("DeleteVideo_ExecError", func(t *testing.T) {
		request := &model.DeleteVideo{ID: "123"}

		mock.ExpectExec("DELETE FROM videos WHERE id = \\$1").
			WithArgs(request.ID).
			WillReturnError(fmt.Errorf("exec error"))

		err := repo.DeleteVideo(ctx, request)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "exec error")
	})

	t.Run("DeleteVideo_RowsAffectedError", func(t *testing.T) {
		request := &model.DeleteVideo{ID: "123"}

		mock.ExpectExec("DELETE FROM videos WHERE id = \\$1").
			WithArgs(request.ID).
			WillReturnResult(sqlmock.NewErrorResult(fmt.Errorf("rows affected error")))

		err := repo.DeleteVideo(ctx, request)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "rows affected error")
	})
}
