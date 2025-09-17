package interesados

import (
	"context"
	"errors"

	"upcrm.com/interesados/internal/repository"
	"upcrm.com/interesados/pkg/model"
)

var ErrNotFound = errors.New("Ratings not found for record")

type interesadosRepository interface {
	Get(ctx context.Context, interesadoID model.InteresadoID, interesadoType model.InteresadoType) ([]model.Interesados, error)
	Put(ctx context.Context, interesadoID model.InteresadoID, interesadoType model.InteresadoType, interesados *model.Interesados) error
}

type Controller struct {
	repo interesadosRepository
}

func New(repo interesadosRepository) *Controller {
	return &Controller{repo}
}

func (c *Controller) GetAggregatedInteresados(ctx context.Context, interesadoID model.InteresadoID, interesadoType model.InteresadoType) (float64, error) {
	interesados, err := c.repo.Get(ctx, interesadoID, interesadoType)
	if err != nil && err == repository.ErrNotFound {
		return 0, err
	} else if err != nil {
		return 0, err
	}

	sum := float64(0)
	for _, r := range interesados {
		sum += float64(r.Value)
	}

	return sum / float64(len(interesados)), nil
}

func (c *Controller) PutRating(ctx context.Context, interesadoID model.InteresadoID, interesadoType model.InteresadoType, interesados *model.Interesados) error {
	return c.repo.Put(ctx, interesadoID, interesadoType, interesados)
}
