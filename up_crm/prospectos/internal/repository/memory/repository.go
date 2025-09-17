package memory

import (
	"context"
	"sync"

	"upcrm.com/prospectos/internal/repository"
	model "upcrm.com/prospectos/pkg"
)

type Repository struct {
	sync.RWMutex
	data map[string]*model.Prospectos
}

func New() *Repository {
	return &Repository{data: map[string]*model.Prospectos{}}
}

func (r *Repository) Get(_ context.Context, id string) (*model.Prospectos, error) {
	r.RLock()
	defer r.RUnlock()

	m, ok := r.data[id]
	if !ok {
		return nil, repository.ErrNotFound
	}

	return m, nil
}

func (r *Repository) Put(_ context.Context, id string, metadata *model.Prospectos) error {
	r.Lock()
	defer r.Unlock()
	r.data[id] = metadata
	return nil
}
