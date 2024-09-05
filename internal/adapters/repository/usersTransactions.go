package repository

import (
	model "CrudPlatform/internal/core/domain/repository/model/users"

	"database/sql"
	"fmt"
	"time"

	schema "CrudPlatform/internal/core/domain/repository/schema/users"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (p *BDRepository) CreateUser(ctx *gin.Context, request *model.User) (string, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	id := uuid.NewString()
	now := time.Now().UTC().Format("2006-01-02 15:04:05")

	query := `
		INSERT INTO users (id, name, email, image_path, created_at, updated_at) 
		VALUES (?, ?, ?, ?, ?, ?)
	`
	stmt, err := p.db.Prepare(query)
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id, request.Name, request.Email, request.ImagePath, now, now)
	if err != nil {
		return "", fmt.Errorf("error executing statement: %w", err)
	}

	return id, nil
}

func (p *BDRepository) SelectUser(ctx *gin.Context, request *model.GetUser) (*schema.UsersGetResponse, error) {

	p.mu.Lock()
	defer p.mu.Unlock()

	query := "SELECT name, email, image_path, created_at, updated_at FROM users WHERE id = ?"
	row := p.db.QueryRow(query, request.Id)

	var response schema.UsersGetResponse

	err := row.Scan(&response.Name, &response.Email, &response.ImagePath, &response.CreatedAt, &response.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user with id %s not found", request.Id)
		}
		return nil, fmt.Errorf("error scanning user row: %w", err)
	}

	return &response, nil
}

func (p *BDRepository) UpdateUser(ctx *gin.Context, request *model.UpdateUser) (*schema.UsersUpdateResponse, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	now := time.Now().UTC().Format("2006-01-02 15:04:05")

	query := "UPDATE users SET name = ?, email = ?, image_path = ?, updated_at = ? WHERE id = ?"
	stmt, err := p.db.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("error preparing update statement: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(request.Name, request.Email, request.ImagePath, now, request.Id)
	if err != nil {
		return nil, fmt.Errorf("error executing update: %w", err)
	}

	updatedQuery := "SELECT name, email, image_path, updated_at FROM users WHERE id = ?"
	updatedRow := p.db.QueryRow(updatedQuery, request.Id)

	var response schema.UsersUpdateResponse
	err = updatedRow.Scan(&response.Name, &response.Email, &response.ImagePath, &response.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user with id %s not found after update", request.Id)
		}
		return nil, fmt.Errorf("error scanning updated user row: %w", err)
	}

	return &response, nil
}

func (p *BDRepository) DeleteUser(ctx *gin.Context, request *model.DeleteUser) error {

	p.mu.Lock()
	defer p.mu.Unlock()

	query := "DELETE FROM users WHERE id = ?"
	stmt, err := p.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(request.Id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected > 0 {
		return nil
	}

	return nil
}
