package medico

import "context"

// Repository handle the CRUD operations with Medico.
type Repository interface {
	GetAll(ctx context.Context) ([]Medico, error)
	GetOne(ctx context.Context, id uint) (Medico, error)
	GetByMediconombre(ctx context.Context, nombre string) (Medico, error)
	Create(ctx context.Context, medico *Medico) error
	Update(ctx context.Context, id uint, medico Medico) error
	Delete(ctx context.Context, id uint) error
}
