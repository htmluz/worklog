package storage

// https://open.spotify.com/track/6wJslEwvjmwfu7gthHc4cs?si=83ac5a964fa14eb2

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"

	"github.com/htmluz/worklog/internal/domain"
)

type JSONStorage struct {
	path string
}

func NewJSONStorage() (*JSONStorage, error) {
	dir, err := getWorkLogDir()
	if err != nil {
		return nil, err
	}

	return &JSONStorage{
		path: filepath.Join(dir, "data.json"),
	}, nil
}

func (s *JSONStorage) Load() (*domain.Store, error) {
	data, err := os.ReadFile(s.path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return domain.NewStore(), nil
		}
		return nil, err
	}

	var store domain.Store
	if err := json.Unmarshal(data, &store); err != nil {
		return nil, err
	}

	if store.Tasks == nil {
		store.Tasks = make(map[string]*domain.Task)
	}
	if store.Windows == nil {
		store.Windows = make(map[string]*domain.Window)
	}

	return &store, nil
}

func (s *JSONStorage) Save(store *domain.Store) error {
	data, err := json.MarshalIndent(store, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.path, data, 0644)
}

func getWorkLogDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	dir := filepath.Join(home, ".worklog")

	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}
	return dir, nil
}
