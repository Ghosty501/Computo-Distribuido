package prospectos

import (
	"context"
	"errors"

	model "upcrm.com/prospectos/pkg"
)

var ErrNotFound = errors.New("Not found")

type ProspectosRepository interface {
	Get(ctx context.Context, id string) (*model.Prospectos, error)
}

type Controller struct {
	repo ProspectosRepository
}

func New(repo ProspectosRepository) *Controller {
	return &Controller{repo}
}

func (c *Controller) Get(ctx context.Context, id string) (*model.Prospectos, error) {
	res, err := c.repo.Get(ctx, id)

	if err != nil {
		return nil, ErrNotFound
	}

	return res, err
}
