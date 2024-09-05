package repository

import (
	"database/sql"
	"sync"
)

type BDRepository struct {
	db *sql.DB
	mu sync.Mutex
}

func NewBdRepository(db *sql.DB) *BDRepository {
	return &BDRepository{
		db: db,
	}
}

type BDRepositoryChallenge struct {
	db *sql.DB
	mu sync.Mutex
}

func NewBdRepositoryChallenge(db *sql.DB) *BDRepositoryChallenge {
	return &BDRepositoryChallenge{
		db: db,
	}
}
