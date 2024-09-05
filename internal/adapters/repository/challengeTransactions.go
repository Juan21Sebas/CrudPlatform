package repository

import (
	model "CrudPlatform/internal/core/domain/repository/model/challenges"
	schema "CrudPlatform/internal/core/domain/repository/schema/challenges"
	"database/sql"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (p *BDRepositoryChallenge) CreateChallenge(ctx *gin.Context, request *model.Challenge) (string, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	id := uuid.NewString()
	now := time.Now().UTC().Format("2006-01-02 15:04:05")

	query := `
		INSERT INTO challenges (id, title, description, difficulty, created_at, updated_at) 
		VALUES (?, ?, ?, ?, ?, ?)
	`
	stmt, err := p.db.Prepare(query)
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id, request.Title, request.Description, request.Difficulty, now, now)
	if err != nil {
		return "", fmt.Errorf("error executing statement: %w", err)
	}

	return id, nil
}

func (p *BDRepositoryChallenge) SelectChallenge(ctx *gin.Context, request *model.GetChallenge) (*schema.ChallengeGetResponse, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	query := "SELECT title, description, difficulty, created_at, updated_at FROM challenges WHERE id = ?"
	row := p.db.QueryRow(query, request.ID)

	var response schema.ChallengeGetResponse

	err := row.Scan(&response.Title, &response.Description, &response.Difficulty, &response.CreatedAt, &response.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("challenge with id %s not found", request.ID)
		}
		return nil, fmt.Errorf("error scanning challenge row: %w", err)
	}

	return &response, nil
}

func (p *BDRepositoryChallenge) UpdateChallenge(ctx *gin.Context, request *model.UpdateChallenge) (*schema.ChallengeUpdateResponse, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	now := time.Now().UTC().Format("2006-01-02 15:04:05")

	query := "UPDATE challenges SET title = ?, description = ?, difficulty = ?, updated_at = ? WHERE id = ?"
	stmt, err := p.db.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("error preparing update statement: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(request.Title, request.Description, request.Difficulty, now, request.ID)
	if err != nil {
		return nil, fmt.Errorf("error executing update: %w", err)
	}

	updatedQuery := "SELECT title, description, difficulty, updated_at FROM challenges WHERE id = ?"
	updatedRow := p.db.QueryRow(updatedQuery, request.ID)

	var response schema.ChallengeUpdateResponse
	err = updatedRow.Scan(&response.Title, &response.Description, &response.Difficulty, &response.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("challenge with id %s not found after update", request.ID)
		}
		return nil, fmt.Errorf("error scanning updated challenge row: %w", err)
	}

	return &response, nil
}

func (p *BDRepositoryChallenge) DeleteChallenge(ctx *gin.Context, request *model.DeleteChallenge) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	query := "DELETE FROM challenges WHERE id = ?"
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
		return fmt.Errorf("challenge with id %s not found", request.ID)
	}

	return nil
}
