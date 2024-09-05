package repository

import (
	model "CrudPlatform/internal/core/domain/repository/model/users"

	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateTask_Success(t *testing.T) {

	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal("Error al abrir la conexión de base de datos:", err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS tasks (id TEXT PRIMARY KEY, title TEXT, description TEXT, status TEXT)")
	if err != nil {
		t.Fatal("Error al crear la tabla de tareas:", err)
	}

	repo := NewBdRepository(db)
	c, _ := gin.CreateTestContext(nil)
	request := &model.User{
		Name:      "Título de la tarea",
		Email:     "Descripción de la tarea",
		ImagePath: "Pendiente",
	}

	response, err := repo.CreateUser(c, request)
	assert.NoError(t, err)
	assert.Contains(t, response, "Su ID es")
}

// func TestSelectTask_Success(t *testing.T) {

// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("Error creando el mock de la base de datos: %v", err)
// 	}
// 	defer db.Close()

// 	repo := NewBdRepository(db)

// 	c, _ := gin.CreateTestContext(nil)
// 	rows := sqlmock.NewRows([]string{"title", "description", "status"}).
// 		AddRow("Título de la tarea", "Descripción de la tarea", "Pendiente")
// 	mock.ExpectQuery("SELECT title, description, status FROM tasks WHERE id = ?").
// 		WithArgs("ID").
// 		WillReturnRows(rows)

// 	request := &model.GetUser{
// 		Id: "ID",
// 	}

// 	response, err := repo.SelectUser(c, request)

// 	assert.NoError(t, err)
// 	assert.Equal(t, "Título de la tarea", response.Title)
// 	assert.Equal(t, "Descripción de la tarea", response.Description)
// 	assert.Equal(t, "Pendiente", response.Status)

// 	err = mock.ExpectationsWereMet()
// 	assert.NoError(t, err)
// }

// func TestUpdateTask_Success(t *testing.T) {

// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("Error creating mock database: %v", err)
// 	}
// 	defer db.Close()

// 	repo := NewBdRepository(db)

// 	c, _ := gin.CreateTestContext(nil)
// 	mock.ExpectPrepare("UPDATE tasks SET title = \\?, description = \\?, status = \\? WHERE id = \\?").
// 		ExpectExec().
// 		WithArgs("Nuevo título", "Nueva descripción", "Nuevo estado", "ID").
// 		WillReturnResult(sqlmock.NewResult(0, 1))

// 	rows := sqlmock.NewRows([]string{"title", "description", "status"}).
// 		AddRow("Nuevo título", "Nueva descripción", "Nuevo estado")
// 	mock.ExpectQuery("SELECT title, description, status FROM tasks WHERE id = ?").
// 		WithArgs("ID").
// 		WillReturnRows(rows)

// 	request := &modelupdate.Task{
// 		Id:          "ID",
// 		Title:       "Nuevo título",
// 		Description: "Nueva descripción",
// 		Status:      "Nuevo estado",
// 	}

// 	response, err := repo.UpdateTask(c, request)

// 	assert.NoError(t, err, "UpdateTask should not return an error")
// 	assert.Equal(t, "Nuevo título", response.Title, "Title should match the updated value")
// 	assert.Equal(t, "Nueva descripción", response.Description, "Description should match the updated value")
// 	assert.Equal(t, "Nuevo estado", response.Status, "Status should match the updated value")

// 	err = mock.ExpectationsWereMet()
// 	assert.NoError(t, err, "Mock database expectations should be met")
// }

func TestDeleteTask_Success(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creando el mock de la base de datos: %v", err)
	}
	defer db.Close()

	repo := NewBdRepository(db)
	c, _ := gin.CreateTestContext(nil)

	mock.ExpectPrepare("DELETE FROM tasks WHERE id = ?").
		ExpectExec().
		WithArgs("ID").
		WillReturnResult(sqlmock.NewResult(0, 1))

	request := &model.DeleteUser{
		Id: "ID",
	}

	err = repo.DeleteUser(c, request)

	assert.NoError(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}
