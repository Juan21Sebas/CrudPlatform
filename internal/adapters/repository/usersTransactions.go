package repository

import (
	model "CrudPlatform/internal/core/domain/repository/model/users"
	schema "CrudPlatform/internal/core/domain/repository/schema/users"
	"database/sql"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (p *BDRepository) CreateUser(ctx *gin.Context, request *model.User) (string, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	id := uuid.NewString()
	now := time.Now().UTC()

	query := `
		INSERT INTO users (id, name, email, image_path, created_at, updated_at) 
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := p.db.Exec(query, id, request.Name, request.Email, request.ImagePath, now, now)
	if err != nil {
		return "", fmt.Errorf("error executing statement: %w", err)
	}

	return id, nil
}

func (p *BDRepository) SelectUser(ctx *gin.Context, request *model.GetUser) (*schema.UsersGetResponse, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	query := "SELECT name, email, image_path, created_at, updated_at FROM users WHERE id = $1"
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

	now := time.Now().UTC()

	query := "UPDATE users SET name = $1, email = $2, image_path = $3, updated_at = $4 WHERE id = $5"
	_, err := p.db.Exec(query, request.Name, request.Email, request.ImagePath, now, request.Id)
	if err != nil {
		return nil, fmt.Errorf("error executing update: %w", err)
	}

	updatedQuery := "SELECT name, email, image_path, updated_at FROM users WHERE id = $1"
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

	query := "DELETE FROM users WHERE id = $1"
	result, err := p.db.Exec(query, request.Id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no user found with id %s", request.Id)
	}

	return nil
}
