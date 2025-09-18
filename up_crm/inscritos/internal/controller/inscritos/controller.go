package inscritos

import (
	"context"
	"errors"

	"upcrm.com/inscritos/internal/repository"
	"upcrm.com/inscritos/pkg/model"
)

var ErrNotFound = errors.New("Ratings not found for record")

type inscritosRepository interface {
	Get(ctx context.Context, inscritosID model.InscritoID, inscritosType model.InscritoType) ([]model.Inscritos, error)
	Put(ctx context.Context, inscritosID model.InscritoID, inscritosType model.InscritoType, inscritos *model.Inscritos) error
}

type Controller struct {
	repo inscritosRepository
}

func New(repo inscritosRepository) *Controller {
	return &Controller{repo}
}

func (c *Controller) GetAggregatedInscritos(ctx context.Context, inscritosID model.InscritoID, inscritosType model.InscritoType) (float64, error) {
	inscritos, err := c.repo.Get(ctx, inscritosID, inscritosType)
	if err != nil && err == repository.ErrNotFound {
		return 0, err
	} else if err != nil {
		return 0, err
	}

	sum := float64(0)
	for _, r := range inscritos {
		sum += float64(r.Value)
	}

	return sum / float64(len(inscritos)), nil
}

func (c *Controller) PutRating(ctx context.Context, inscritosID model.InscritoID, inscritosType model.InscritoType, inscritos *model.Inscritos) error {
	return c.repo.Put(ctx, inscritosID, inscritosType, inscritos)
}
