package memory

import (
	"context"

	"upcrm.com/inscritos/internal/repository"
	"upcrm.com/inscritos/pkg/model"
)

type Repository struct {
	data map[model.InscritoType]map[model.InscritoID][]model.Inscritos
}

func New() *Repository {
	return &Repository{map[model.InscritoType]map[model.InscritoID][]model.Inscritos{}}
}

func (r *Repository) Get(ctx context.Context, inscritosID model.InscritoID, inscritosType model.InscritoType) ([]model.Inscritos, error) {
	if _, ok := r.data[inscritosType]; !ok {
		return nil, repository.ErrNotFound
	}

	if inscritos, ok := r.data[inscritosType][inscritosID]; !ok || len(inscritos) == 0 {
		return nil, repository.ErrNotFound
	}
	return r.data[inscritosType][inscritosID], nil
}

func (r *Repository) Put(ctx context.Context, inscritosID model.InscritoID, inscritosType model.InscritoType, rating *model.Inscritos) error {
	if _, ok := r.data[inscritosType]; !ok {
		r.data[inscritosType] = map[model.InscritoID][]model.Inscritos{}
	}
	r.data[inscritosType][inscritosID] = append(r.data[inscritosType][inscritosID], *rating)
	return nil
}
