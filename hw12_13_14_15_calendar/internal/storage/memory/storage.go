package memorystorage

import "sync"

type Storage struct {
	// TODO
	mu sync.RWMutex
}

func New() (*Storage, error) {
	return &Storage{}, nil
}

// TODO
