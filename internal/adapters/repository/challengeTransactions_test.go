package repository

import (
	"CrudPlatform/internal/core/domain/repository/model/challenges"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBDRepositoryChallenge(t *testing.T) {
	// Setup
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := &BDRepositoryChallenge{db: db}
	ctx := &gin.Context{}

	t.Run("CreateChallenge", func(t *testing.T) {
		challenge := &challenges.Challenge{
			Title:       "Test Challenge",
			Description: "This is a test challenge",
			Difficulty:  3,
		}

		mock.ExpectExec("INSERT INTO challenges").
			WithArgs(sqlmock.AnyArg(), challenge.Title, challenge.Description, challenge.Difficulty, sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnResult(sqlmock.NewResult(1, 1))

		id, err := repo.CreateChallenge(ctx, challenge)
		assert.NoError(t, err)
		assert.NotEmpty(t, id)
	})

	t.Run("CreateChallenge_ExecError", func(t *testing.T) {
		challenge := &challenges.Challenge{
			Title:       "Test Challenge",
			Description: "This is a test challenge",
			Difficulty:  3,
		}

		mock.ExpectExec("INSERT INTO challenges").
			WithArgs(sqlmock.AnyArg(), challenge.Title, challenge.Description, challenge.Difficulty, sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnError(fmt.Errorf("error de ejecución"))

		id, err := repo.CreateChallenge(ctx, challenge)
		assert.Error(t, err)
		assert.Empty(t, id)
		assert.Contains(t, err.Error(), "error executing statement: error de ejecución")
	})

	t.Run("SelectChallenge", func(t *testing.T) {
		request := &challenges.GetChallenge{ID: "123"}
		rows := sqlmock.NewRows([]string{"title", "description", "difficulty", "created_at", "updated_at"}).
			AddRow("Test Challenge", "This is a test challenge", 3, time.Now(), time.Now())

		mock.ExpectQuery("SELECT (.+) FROM challenges WHERE id = \\$1").
			WithArgs(request.ID).
			WillReturnRows(rows)

		challenge, err := repo.SelectChallenge(ctx, request)
		assert.NoError(t, err)
		assert.NotNil(t, challenge)
		assert.Equal(t, "Test Challenge", challenge.Title)
		assert.Equal(t, 3, challenge.Difficulty)
	})

	t.Run("SelectChallenge_NotFound", func(t *testing.T) {
		request := &challenges.GetChallenge{ID: "999"}

		mock.ExpectQuery("SELECT (.+) FROM challenges WHERE id = \\$1").
			WithArgs(request.ID).
			WillReturnError(sql.ErrNoRows)

		challenge, err := repo.SelectChallenge(ctx, request)
		assert.Error(t, err)
		assert.Nil(t, challenge)
		assert.Contains(t, err.Error(), "challenge with id 999 not found")
	})

	t.Run("SelectChallenge_ScanError", func(t *testing.T) {
		request := &challenges.GetChallenge{ID: "123"}

		mock.ExpectQuery("SELECT (.+) FROM challenges WHERE id = \\$1").
			WithArgs(request.ID).
			WillReturnRows(sqlmock.NewRows([]string{"title", "description", "difficulty", "created_at", "updated_at"}).
				AddRow("Test Challenge", "This is a test challenge", "no es un número", time.Now(), time.Now()))

		challenge, err := repo.SelectChallenge(ctx, request)
		assert.Error(t, err)
		assert.Nil(t, challenge)
		assert.Contains(t, err.Error(), "error scanning challenge row")
	})

	t.Run("UpdateChallenge", func(t *testing.T) {
		request := &challenges.UpdateChallenge{
			ID:          "123",
			Title:       "Updated Challenge",
			Description: "This is an updated test challenge",
			Difficulty:  4,
		}

		mock.ExpectExec("UPDATE challenges").
			WithArgs(request.Title, request.Description, request.Difficulty, sqlmock.AnyArg(), request.ID).
			WillReturnResult(sqlmock.NewResult(0, 1))

		rows := sqlmock.NewRows([]string{"title", "description", "difficulty", "updated_at"}).
			AddRow(request.Title, request.Description, request.Difficulty, time.Now())

		mock.ExpectQuery("SELECT (.+) FROM challenges WHERE id = \\$1").
			WithArgs(request.ID).
			WillReturnRows(rows)

		challenge, err := repo.UpdateChallenge(ctx, request)
		assert.NoError(t, err)
		assert.NotNil(t, challenge)
		assert.Equal(t, request.Title, challenge.Title)
		assert.Equal(t, request.Difficulty, challenge.Difficulty)
	})

	t.Run("UpdateChallenge_ExecError", func(t *testing.T) {
		request := &challenges.UpdateChallenge{
			ID:          "123",
			Title:       "Updated Challenge",
			Description: "This is an updated test challenge",
			Difficulty:  4,
		}

		mock.ExpectExec("UPDATE challenges").
			WithArgs(request.Title, request.Description, request.Difficulty, sqlmock.AnyArg(), request.ID).
			WillReturnError(fmt.Errorf("error de ejecución"))

		challenge, err := repo.UpdateChallenge(ctx, request)
		assert.Error(t, err)
		assert.Nil(t, challenge)
		assert.Contains(t, err.Error(), "error executing update")
	})

	t.Run("UpdateChallenge_NotFoundAfterUpdate", func(t *testing.T) {
		request := &challenges.UpdateChallenge{
			ID:          "123",
			Title:       "Updated Challenge",
			Description: "This is an updated test challenge",
			Difficulty:  4,
		}

		mock.ExpectExec("UPDATE challenges").
			WithArgs(request.Title, request.Description, request.Difficulty, sqlmock.AnyArg(), request.ID).
			WillReturnResult(sqlmock.NewResult(0, 1))

		mock.ExpectQuery("SELECT (.+) FROM challenges WHERE id = \\$1").
			WithArgs(request.ID).
			WillReturnError(sql.ErrNoRows)

		challenge, err := repo.UpdateChallenge(ctx, request)
		assert.Error(t, err)
		assert.Nil(t, challenge)
		assert.Contains(t, err.Error(), "challenge with id 123 not found after update")
	})

	t.Run("UpdateChallenge_ScanError", func(t *testing.T) {
		request := &challenges.UpdateChallenge{
			ID:          "123",
			Title:       "Updated Challenge",
			Description: "This is an updated test challenge",
			Difficulty:  4,
		}

		mock.ExpectExec("UPDATE challenges").
			WithArgs(request.Title, request.Description, request.Difficulty, sqlmock.AnyArg(), request.ID).
			WillReturnResult(sqlmock.NewResult(0, 1))

		mock.ExpectQuery("SELECT (.+) FROM challenges WHERE id = \\$1").
			WithArgs(request.ID).
			WillReturnRows(sqlmock.NewRows([]string{"title", "description", "difficulty", "updated_at"}).
				AddRow(request.Title, request.Description, "no es un número", time.Now()))

		challenge, err := repo.UpdateChallenge(ctx, request)
		assert.Error(t, err)
		assert.Nil(t, challenge)
		assert.Contains(t, err.Error(), "error scanning updated challenge row")
	})

	t.Run("DeleteChallenge", func(t *testing.T) {
		request := &challenges.DeleteChallenge{ID: "123"}

		mock.ExpectExec("DELETE FROM challenges").
			WithArgs(request.ID).
			WillReturnResult(sqlmock.NewResult(0, 1))

		err := repo.DeleteChallenge(ctx, request)
		assert.NoError(t, err)
	})

	t.Run("DeleteChallenge_NotFound", func(t *testing.T) {
		request := &challenges.DeleteChallenge{ID: "999"}

		mock.ExpectExec("DELETE FROM challenges").
			WithArgs(request.ID).
			WillReturnResult(sqlmock.NewResult(0, 0))

		err := repo.DeleteChallenge(ctx, request)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "challenge with id 999 not found")
	})

	t.Run("DeleteChallenge_ExecError", func(t *testing.T) {
		request := &challenges.DeleteChallenge{ID: "123"}

		mock.ExpectExec("DELETE FROM challenges").
			WithArgs(request.ID).
			WillReturnError(fmt.Errorf("error de ejecución"))

		err := repo.DeleteChallenge(ctx, request)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "error de ejecución")
	})

	t.Run("DeleteChallenge_RowsAffectedError", func(t *testing.T) {
		request := &challenges.DeleteChallenge{ID: "123"}

		mock.ExpectExec("DELETE FROM challenges").
			WithArgs(request.ID).
			WillReturnResult(sqlmock.NewErrorResult(fmt.Errorf("error de filas afectadas")))

		err := repo.DeleteChallenge(ctx, request)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "error de filas afectadas")
	})
}
