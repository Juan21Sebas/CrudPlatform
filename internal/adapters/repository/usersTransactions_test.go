package repository

import (
	"CrudPlatform/internal/core/domain/repository/model/users"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBDRepository(t *testing.T) {
	// Setup
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := &BDRepository{db: db}
	ctx := &gin.Context{}

	t.Run("CreateUser", func(t *testing.T) {
		user := &users.User{
			Name:      "John Doe",
			Email:     "john@example.com",
			ImagePath: "/path/to/image.jpg",
		}

		mock.ExpectPrepare("INSERT INTO users").
			ExpectExec().
			WithArgs(sqlmock.AnyArg(), user.Name, user.Email, user.ImagePath, sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnResult(sqlmock.NewResult(1, 1))

		id, err := repo.CreateUser(ctx, user)
		assert.NoError(t, err)
		assert.NotEmpty(t, id)
	})

	t.Run("CreateUser_PrepareError", func(t *testing.T) {
		user := &users.User{
			Name:      "John Doe",
			Email:     "john@example.com",
			ImagePath: "/path/to/image.jpg",
		}

		mock.ExpectPrepare("INSERT INTO users").WillReturnError(fmt.Errorf("prepare error"))

		id, err := repo.CreateUser(ctx, user)
		assert.Error(t, err)
		assert.Empty(t, id)
		assert.Contains(t, err.Error(), "prepare error")
	})

	t.Run("CreateUser_ExecError", func(t *testing.T) {
		user := &users.User{
			Name:      "John Doe",
			Email:     "john@example.com",
			ImagePath: "/path/to/image.jpg",
		}

		mock.ExpectPrepare("INSERT INTO users").
			ExpectExec().
			WithArgs(sqlmock.AnyArg(), user.Name, user.Email, user.ImagePath, sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnError(fmt.Errorf("exec error"))

		id, err := repo.CreateUser(ctx, user)
		assert.Error(t, err)
		assert.Empty(t, id)
		assert.Contains(t, err.Error(), "error executing statement: exec error")
	})

	t.Run("SelectUser", func(t *testing.T) {
		request := &users.GetUser{Id: "123"}
		rows := sqlmock.NewRows([]string{"name", "email", "image_path", "created_at", "updated_at"}).
			AddRow("John Doe", "john@example.com", "/path/to/image.jpg", time.Now(), time.Now())

		mock.ExpectQuery("SELECT (.+) FROM users WHERE id = ?").
			WithArgs(request.Id).
			WillReturnRows(rows)

		user, err := repo.SelectUser(ctx, request)
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, "John Doe", user.Name)
	})

	t.Run("SelectUser_NotFound", func(t *testing.T) {
		request := &users.GetUser{Id: "999"}

		mock.ExpectQuery("SELECT (.+) FROM users WHERE id = ?").
			WithArgs(request.Id).
			WillReturnError(sql.ErrNoRows)

		video, err := repo.SelectUser(ctx, request)
		assert.Error(t, err)
		assert.Nil(t, video)
		assert.Contains(t, err.Error(), "user with id 999 not found")
	})

	t.Run("SelectUser_ScanError_ExtraField", func(t *testing.T) {
		request := &users.GetUser{Id: "123"}

		rows := sqlmock.NewRows([]string{"name", "email", "image_path", "created_at", "updated_at", "extra_field"}).
			AddRow("John Doe", "john@example.com", "/path/to/image.jpg", time.Now(), time.Now(), "extra data")

		mock.ExpectQuery("SELECT name, email, image_path, created_at, updated_at FROM users WHERE id = ?").
			WithArgs(request.Id).
			WillReturnRows(rows)

		user, err := repo.SelectUser(ctx, request)

		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Contains(t, err.Error(), "error scanning user row")
		assert.Contains(t, err.Error(), "scanning")
	})

	t.Run("UpdateUser", func(t *testing.T) {
		request := &users.UpdateUser{
			Id:        "123",
			Name:      "Jane Doe",
			Email:     "jane@example.com",
			ImagePath: "/new/path/to/image.jpg",
		}

		mock.ExpectPrepare("UPDATE users").
			ExpectExec().
			WithArgs(request.Name, request.Email, request.ImagePath, sqlmock.AnyArg(), request.Id).
			WillReturnResult(sqlmock.NewResult(0, 1))

		rows := sqlmock.NewRows([]string{"name", "email", "image_path", "updated_at"}).
			AddRow(request.Name, request.Email, request.ImagePath, time.Now())

		mock.ExpectQuery("SELECT (.+) FROM users WHERE id = ?").
			WithArgs(request.Id).
			WillReturnRows(rows)

		user, err := repo.UpdateUser(ctx, request)
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, request.Name, user.Name)
	})

	t.Run("UpdateUser_PrepareError", func(t *testing.T) {
		request := &users.UpdateUser{
			Id:        "123",
			Name:      "Jane Doe",
			Email:     "jane@example.com",
			ImagePath: "/new/path/to/image.jpg",
		}

		mock.ExpectPrepare("UPDATE users").WillReturnError(fmt.Errorf("prepare error"))

		video, err := repo.UpdateUser(ctx, request)
		assert.Error(t, err)
		assert.Nil(t, video)
		assert.Contains(t, err.Error(), "error preparing update statement")
	})

	t.Run("UpdateUser_ExecError", func(t *testing.T) {
		request := &users.UpdateUser{
			Id:        "123",
			Name:      "Jane Doe",
			Email:     "jane@example.com",
			ImagePath: "/new/path/to/image.jpg",
		}

		mock.ExpectPrepare("UPDATE users").
			ExpectExec().
			WithArgs(request.Name, request.Email, request.ImagePath, sqlmock.AnyArg(), request.Id).
			WillReturnError(fmt.Errorf("error de ejecución"))

		user, err := repo.UpdateUser(ctx, request)
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Contains(t, err.Error(), "error executing update")
	})

	t.Run("UpdateUser_ScanError", func(t *testing.T) {
		request := &users.UpdateUser{
			Id:        "123",
			Name:      "Jane Doe",
			Email:     "jane@example.com",
			ImagePath: "/new/path/to/image.jpg",
		}

		mock.ExpectPrepare("UPDATE users").
			ExpectExec().
			WithArgs(request.Name, request.Email, request.ImagePath, sqlmock.AnyArg(), request.Id).
			WillReturnResult(sqlmock.NewResult(0, 1))

		mock.ExpectQuery("SELECT name, email, image_path, updated_at FROM users WHERE id = ?").
			WithArgs(request.Id).
			WillReturnError(sql.ErrNoRows)

		user, err := repo.UpdateUser(ctx, request)
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Contains(t, err.Error(), "user with id 123 not found after update")
	})

	t.Run("UpdateUser_OtherScanError", func(t *testing.T) {
		request := &users.UpdateUser{
			Id:        "123",
			Name:      "Jane Doe",
			Email:     "jane@example.com",
			ImagePath: "/new/path/to/image.jpg",
		}

		mock.ExpectPrepare("UPDATE users").
			ExpectExec().
			WithArgs(request.Name, request.Email, request.ImagePath, sqlmock.AnyArg(), request.Id).
			WillReturnResult(sqlmock.NewResult(0, 1))

		mock.ExpectQuery("SELECT name, email, image_path, updated_at FROM users WHERE id = ?").
			WithArgs(request.Id).
			WillReturnError(fmt.Errorf("error de escaneo"))

		user, err := repo.UpdateUser(ctx, request)
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Contains(t, err.Error(), "error scanning updated user row")
	})

	t.Run("DeleteUser", func(t *testing.T) {
		request := &users.DeleteUser{Id: "123"}

		mock.ExpectPrepare("DELETE FROM users").
			ExpectExec().
			WithArgs(request.Id).
			WillReturnResult(sqlmock.NewResult(0, 1))

		err := repo.DeleteUser(ctx, request)
		assert.NoError(t, err)
	})

	t.Run("DeleteUser_PrepareError", func(t *testing.T) {
		request := &users.DeleteUser{Id: "123"}

		mock.ExpectPrepare("DELETE FROM users").WillReturnError(fmt.Errorf("error de preparación"))

		err := repo.DeleteUser(ctx, request)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "error de preparación")
	})

	t.Run("DeleteUser_ExecError", func(t *testing.T) {
		request := &users.DeleteUser{Id: "123"}

		mock.ExpectPrepare("DELETE FROM users").
			ExpectExec().
			WithArgs(request.Id).
			WillReturnError(fmt.Errorf("error de ejecución"))

		err := repo.DeleteUser(ctx, request)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "error de ejecución")
	})

	t.Run("DeleteUser_NoRowsAffected", func(t *testing.T) {
		request := &users.DeleteUser{Id: "123"}

		mock.ExpectPrepare("DELETE FROM users").
			ExpectExec().
			WithArgs(request.Id).
			WillReturnResult(sqlmock.NewResult(0, 0))

		err := repo.DeleteUser(ctx, request)
		assert.NoError(t, err)
	})

	t.Run("DeleteUser_RowsAffectedError", func(t *testing.T) {
		request := &users.DeleteUser{Id: "123"}

		mock.ExpectPrepare("DELETE FROM users").
			ExpectExec().
			WithArgs(request.Id).
			WillReturnResult(sqlmock.NewErrorResult(fmt.Errorf("error de filas afectadas")))

		err := repo.DeleteUser(ctx, request)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "error de filas afectadas")
	})

}
