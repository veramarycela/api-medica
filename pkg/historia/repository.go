package historia

import "context"

// Repository handle the CRUD operations with Users.
type Repository interface {
	GetAll(ctx context.Context) ([]Historia, error)
	GetOne(ctx context.Context, id uint) (Historia, error)
	GetByHiatorianombre(ctx context.Context, id int) (Historia, error)
	Create(ctx context.Context, historia *Historia) error
	Update(ctx context.Context, id uint, historia Historia) error
	Delete(ctx context.Context, id uint) error
}
