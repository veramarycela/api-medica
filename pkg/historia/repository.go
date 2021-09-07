package historia

import "context"

// Repository handle the CRUD operations with Historia.
type Repository interface {
	GetAll(ctx context.Context) ([]Historia, error)
	GetOne(ctx context.Context, id string) (Historia, error)
	// GetByHistorianombre(ctx context.Context, id int) (Historia, error)
	Create(ctx context.Context, historia *Historia) error
	Update(ctx context.Context, id string, historia Historia) error
	Delete(ctx context.Context, id string) error
}
