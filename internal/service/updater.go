package service

import (
	"context"
	"sync"

	"gophkeeper/internal/storage"
)

type UpdaterService struct {
	storage storage.Users

	ctx    context.Context
	cancel context.CancelFunc
	mu     sync.Mutex
}

func NewUpdaterService(storage storage.Users) *UpdaterService {
	us := &UpdaterService{
		storage: storage,
	}
	return us
}
