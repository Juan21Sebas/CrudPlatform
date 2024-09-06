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

		mock.ExpectExec("INSERT INTO users").
			WithArgs(sqlmock.AnyArg(), user.Name, user.Email, user.ImagePath, sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnResult(sqlmock.NewResult(1, 1))

		id, err := repo.CreateUser(ctx, user)
		assert.NoError(t, err)
		assert.NotEmpty(t, id)
	})

	t.Run("CreateUser_ExecError", func(t *testing.T) {
		user := &users.User{
			Name:      "John Doe",
			Email:     "john@example.com",
			ImagePath: "/path/to/image.jpg",
		}

		mock.ExpectExec("INSERT INTO users").
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

		mock.ExpectQuery("SELECT (.+) FROM users WHERE id = \\$1").
			WithArgs(request.Id).
			WillReturnRows(rows)

		user, err := repo.SelectUser(ctx, request)
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, "John Doe", user.Name)
	})

	t.Run("SelectUser_NotFound", func(t *testing.T) {
		request := &users.GetUser{Id: "999"}

		mock.ExpectQuery("SELECT (.+) FROM users WHERE id = \\$1").
			WithArgs(request.Id).
			WillReturnError(sql.ErrNoRows)

		user, err := repo.SelectUser(ctx, request)
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Contains(t, err.Error(), "user with id 999 not found")
	})
	t.Run("SelectUser_ScanError", func(t *testing.T) {
		request := &users.GetUser{Id: "123"}

		// Agregamos una columna extra para provocar un error de escaneo
		rows := sqlmock.NewRows([]string{"name", "email", "image_path", "created_at", "updated_at", "extra_column"}).
			AddRow("John Doe", "john@example.com", "/path/to/image.jpg", time.Now(), time.Now(), "extra_data")

		mock.ExpectQuery("SELECT (.+) FROM users WHERE id = \\$1").
			WithArgs(request.Id).
			WillReturnRows(rows)

		user, err := repo.SelectUser(ctx, request)
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Contains(t, err.Error(), "error scanning user row")
	})

	t.Run("UpdateUser", func(t *testing.T) {
		request := &users.UpdateUser{
			Id:        "123",
			Name:      "Jane Doe",
			Email:     "jane@example.com",
			ImagePath: "/new/path/to/image.jpg",
		}

		mock.ExpectExec("UPDATE users SET").
			WithArgs(request.Name, request.Email, request.ImagePath, sqlmock.AnyArg(), request.Id).
			WillReturnResult(sqlmock.NewResult(0, 1))

		rows := sqlmock.NewRows([]string{"name", "email", "image_path", "updated_at"}).
			AddRow(request.Name, request.Email, request.ImagePath, time.Now())

		mock.ExpectQuery("SELECT (.+) FROM users WHERE id = \\$1").
			WithArgs(request.Id).
			WillReturnRows(rows)

		user, err := repo.UpdateUser(ctx, request)
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, request.Name, user.Name)
	})

	t.Run("UpdateUser_ExecError", func(t *testing.T) {
		request := &users.UpdateUser{
			Id:        "123",
			Name:      "Jane Doe",
			Email:     "jane@example.com",
			ImagePath: "/new/path/to/image.jpg",
		}

		mock.ExpectExec("UPDATE users SET").
			WithArgs(request.Name, request.Email, request.ImagePath, sqlmock.AnyArg(), request.Id).
			WillReturnError(fmt.Errorf("exec error"))

		user, err := repo.UpdateUser(ctx, request)
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Contains(t, err.Error(), "error executing update")
	})

	t.Run("UpdateUser_NotFound", func(t *testing.T) {
		request := &users.UpdateUser{
			Id:        "123",
			Name:      "Jane Doe",
			Email:     "jane@example.com",
			ImagePath: "/new/path/to/image.jpg",
		}

		mock.ExpectExec("UPDATE users SET").
			WithArgs(request.Name, request.Email, request.ImagePath, sqlmock.AnyArg(), request.Id).
			WillReturnResult(sqlmock.NewResult(0, 1))

		mock.ExpectQuery("SELECT (.+) FROM users WHERE id = \\$1").
			WithArgs(request.Id).
			WillReturnError(sql.ErrNoRows)

		user, err := repo.UpdateUser(ctx, request)
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Contains(t, err.Error(), "user with id 123 not found after update")
	})

	t.Run("UpdateUser_ScanError", func(t *testing.T) {
		request := &users.UpdateUser{
			Id:        "123",
			Name:      "Jane Doe",
			Email:     "jane@example.com",
			ImagePath: "/new/path/to/image.jpg",
		}

		mock.ExpectExec("UPDATE users SET").
			WithArgs(request.Name, request.Email, request.ImagePath, sqlmock.AnyArg(), request.Id).
			WillReturnResult(sqlmock.NewResult(0, 1))

		rows := sqlmock.NewRows([]string{"name", "email", "image_path", "updated_at", "extra_column"}).
			AddRow(request.Name, request.Email, request.ImagePath, time.Now(), "extra_data")

		mock.ExpectQuery("SELECT (.+) FROM users WHERE id = \\$1").
			WithArgs(request.Id).
			WillReturnRows(rows)

		user, err := repo.UpdateUser(ctx, request)
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Contains(t, err.Error(), "error scanning updated user row")
	})

	t.Run("DeleteUser", func(t *testing.T) {
		request := &users.DeleteUser{Id: "123"}

		mock.ExpectExec("DELETE FROM users WHERE id = \\$1").
			WithArgs(request.Id).
			WillReturnResult(sqlmock.NewResult(0, 1))

		err := repo.DeleteUser(ctx, request)
		assert.NoError(t, err)
	})

	t.Run("DeleteUser_ExecError", func(t *testing.T) {
		request := &users.DeleteUser{Id: "123"}

		mock.ExpectExec("DELETE FROM users WHERE id = \\$1").
			WithArgs(request.Id).
			WillReturnError(fmt.Errorf("exec error"))

		err := repo.DeleteUser(ctx, request)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "exec error")
	})

	t.Run("DeleteUser_NoRowsAffected", func(t *testing.T) {
		request := &users.DeleteUser{Id: "123"}

		mock.ExpectExec("DELETE FROM users WHERE id = \\$1").
			WithArgs(request.Id).
			WillReturnResult(sqlmock.NewResult(0, 0))

		err := repo.DeleteUser(ctx, request)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "no user found with id 123")
	})

	t.Run("DeleteUser_RowsAffectedError", func(t *testing.T) {
		request := &users.DeleteUser{Id: "123"}

		mock.ExpectExec("DELETE FROM users WHERE id = \\$1").
			WithArgs(request.Id).
			WillReturnResult(sqlmock.NewErrorResult(fmt.Errorf("rows affected error")))

		err := repo.DeleteUser(ctx, request)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "rows affected error")
	})
}
