package repository

import (
	"database/sql"
	"sync"
)

type BDRepository struct {
	db *sql.DB
	mu sync.Mutex
}

type BDRepositoryChallenge struct {
	db *sql.DB
	mu sync.Mutex
}

type BDRepositoryVideo struct {
	db *sql.DB
	mu sync.Mutex
}

func NewBdRepository(db *sql.DB) *BDRepository {
	return &BDRepository{
		db: db,
	}
}

func NewBdRepositoryChallenge(db *sql.DB) *BDRepositoryChallenge {
	return &BDRepositoryChallenge{
		db: db,
	}
}

func NewBdRepositoryVideo(db *sql.DB) *BDRepositoryVideo {
	return &BDRepositoryVideo{
		db: db,
	}
}
