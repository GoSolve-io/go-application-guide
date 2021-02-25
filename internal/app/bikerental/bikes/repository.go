package bikes

import (
	"context"

	"github.com/nglogic/go-example-project/internal/app/bikerental"
)

// Repository can manage bike data.
type Repository interface {
	List(context.Context) ([]bikerental.Bike, error)
	Get(ctx context.Context, id string) (*bikerental.Bike, error)
	Add(context.Context, bikerental.Bike) (id string, err error)
	Update(ctx context.Context, id string, b bikerental.Bike) error
	Delete(ctx context.Context, id string) error
}
