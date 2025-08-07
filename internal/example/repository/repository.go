package repository

import (
	"context"

	"github.com/charmingruby/pack/internal/example/model"
)

type ExampleRepo interface {
	FindByID(ctx context.Context, id string) (model.Example, error)
	Create(ctx context.Context, example model.Example) error
	Save(ctx context.Context, example model.Example) error
}
