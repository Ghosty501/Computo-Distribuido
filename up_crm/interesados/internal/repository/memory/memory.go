package memory

import (
	"context"

	"upcrm.com/interesados/internal/repository"
	"upcrm.com/interesados/pkg/model"
)

type Repository struct {
	data map[model.InteresadoType]map[model.InteresadoID][]model.Interesados
}

func New() *Repository {
	return &Repository{map[model.InteresadoType]map[model.InteresadoID][]model.Interesados{}}
}

func (r *Repository) Get(ctx context.Context, interesadoID model.InteresadoID, interesadoType model.InteresadoType) ([]model.Interesados, error) {
	if _, ok := r.data[interesadoType]; !ok {
		return nil, repository.ErrNotFound
	}

	if interesados, ok := r.data[interesadoType][interesadoID]; !ok || len(interesados) == 0 {
		return nil, repository.ErrNotFound
	}
	return r.data[interesadoType][interesadoID], nil
}

func (r *Repository) Put(ctx context.Context, interesadoID model.InteresadoID, interesadoType model.InteresadoType, rating *model.Interesados) error {
	if _, ok := r.data[interesadoType]; !ok {
		r.data[interesadoType] = map[model.InteresadoID][]model.Interesados{}
	}
	r.data[interesadoType][interesadoID] = append(r.data[interesadoType][interesadoID], *rating)
	return nil
}
