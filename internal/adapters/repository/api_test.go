package repository

import (
	"sync"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewBdRepository(t *testing.T) {
	db, _, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := NewBdRepository(db)

	assert.NotNil(t, repo)
	assert.Equal(t, db, repo.db)
}

func TestNewBdRepositoryChallenge(t *testing.T) {
	db, _, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := NewBdRepositoryChallenge(db)

	assert.NotNil(t, repo)
	assert.Equal(t, db, repo.db)
}

func TestNewBdRepositoryVideo(t *testing.T) {
	db, _, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := NewBdRepositoryVideo(db)

	assert.NotNil(t, repo)
	assert.Equal(t, db, repo.db)
}

func TestRepositoryStructures(t *testing.T) {
	db, _, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	t.Run("BDRepository Structure", func(t *testing.T) {
		repo := &BDRepository{db: db}
		assert.IsType(t, &sync.Mutex{}, &repo.mu)
		assert.Equal(t, db, repo.db)
	})

	t.Run("BDRepositoryChallenge Structure", func(t *testing.T) {
		repo := &BDRepositoryChallenge{db: db}
		assert.IsType(t, &sync.Mutex{}, &repo.mu)
		assert.Equal(t, db, repo.db)
	})

	t.Run("BDRepositoryVideo Structure", func(t *testing.T) {
		repo := &BDRepositoryVideo{db: db}
		assert.IsType(t, &sync.Mutex{}, &repo.mu)
		assert.Equal(t, db, repo.db)
	})
}
