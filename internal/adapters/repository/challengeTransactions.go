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
	now := time.Now().UTC()

	query := `
		INSERT INTO challenges (id, title, description, difficulty, created_at, updated_at) 
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := p.db.Exec(query, id, request.Title, request.Description, request.Difficulty, now, now)
	if err != nil {
		return "", fmt.Errorf("error executing statement: %w", err)
	}

	return id, nil
}

func (p *BDRepositoryChallenge) SelectChallenge(ctx *gin.Context, request *model.GetChallenge) (*schema.ChallengeGetResponse, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	query := "SELECT title, description, difficulty, created_at, updated_at FROM challenges WHERE id = $1"
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

	now := time.Now().UTC()

	query := "UPDATE challenges SET title = $1, description = $2, difficulty = $3, updated_at = $4 WHERE id = $5"
	_, err := p.db.Exec(query, request.Title, request.Description, request.Difficulty, now, request.ID)
	if err != nil {
		return nil, fmt.Errorf("error executing update: %w", err)
	}

	updatedQuery := "SELECT title, description, difficulty, updated_at FROM challenges WHERE id = $1"
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

	query := "DELETE FROM challenges WHERE id = $1"
	result, err := p.db.Exec(query, request.ID)
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
