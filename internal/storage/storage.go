package storage

// https://open.spotify.com/track/4aMVT4VVMrbw3OakN1P22R?si=a2cfd97c14694623

import "github.com/htmluz/worklog/internal/domain"

type Storage interface {
	Load() (*domain.Store, error)
	Save(store *domain.Store) error
}
