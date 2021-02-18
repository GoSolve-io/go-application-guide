package bikes

import (
	"context"

	"github.com/nglogic/go-example-project/internal/app"
)

// Repository can manage bike data.
type Repository interface {
	List(context.Context) ([]app.Bike, error)
	Get(ctx context.Context, id string) (*app.Bike, error)
	Add(context.Context, app.Bike) (id string, err error)
	Update(ctx context.Context, id string, b app.Bike) error
	Delete(ctx context.Context, id string) error
}
