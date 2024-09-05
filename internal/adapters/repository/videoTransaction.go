package repository

import (
	model "CrudPlatform/internal/core/domain/repository/model/videos"
	schema "CrudPlatform/internal/core/domain/repository/schema/videos"
	"database/sql"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (p *BDRepositoryVideo) CreateVideo(ctx *gin.Context, request *model.Videos) (string, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	id := uuid.NewString()
	now := time.Now().UTC().Format("2006-01-02 15:04:05")

	query := `
		INSERT INTO videos (id, title, description, created_at, updated_at) 
		VALUES (?, ?, ?, ?, ?)
	`
	stmt, err := p.db.Prepare(query)
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id, request.Title, request.Description, now, now)
	if err != nil {
		return "", fmt.Errorf("error executing statement: %w", err)
	}

	return id, nil
}

func (p *BDRepositoryVideo) SelectVideo(ctx *gin.Context, request *model.GetVideo) (*schema.VideosGetResponse, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	query := "SELECT title, description, created_at, updated_at FROM videos WHERE id = ?"
	row := p.db.QueryRow(query, request.ID)

	var response schema.VideosGetResponse

	err := row.Scan(&response.Title, &response.Description, &response.CreatedAt, &response.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("video with id %s not found", request.ID)
		}
		return nil, fmt.Errorf("error scanning video row: %w", err)
	}

	return &response, nil
}

func (p *BDRepositoryVideo) UpdateVideo(ctx *gin.Context, request *model.UpdateVideo) (*schema.VideosUpdateResponse, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	now := time.Now().UTC().Format("2006-01-02 15:04:05")

	query := "UPDATE videos SET title = ?, description = ?, updated_at = ? WHERE id = ?"
	stmt, err := p.db.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("error preparing update statement: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(request.Title, request.Description, now, request.ID)
	if err != nil {
		return nil, fmt.Errorf("error executing update: %w", err)
	}

	updatedQuery := "SELECT title, description, updated_at FROM videos WHERE id = ?"
	updatedRow := p.db.QueryRow(updatedQuery, request.ID)

	var response schema.VideosUpdateResponse
	err = updatedRow.Scan(&response.Title, &response.Description, &response.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("video with id %s not found after update", request.ID)
		}
		return nil, fmt.Errorf("error scanning updated video row: %w", err)
	}

	return &response, nil
}

func (p *BDRepositoryVideo) DeleteVideo(ctx *gin.Context, request *model.DeleteVideo) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	query := "DELETE FROM videos WHERE id = ?"
	stmt, err := p.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(request.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("video with id %s not found", request.ID)
	}

	return nil
}
