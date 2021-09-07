package especialidad

import "context"

// Repository handle the CRUD operations with especialidad
type Repository interface {
	GetAll(ctx context.Context) ([]Especialidad, error)
	GetOne(ctx context.Context, id uint) (Especialidad, error)
	GetByEspecialidadnombre(ctx context.Context, nombre string) (Especialidad, error)
	Create(ctx context.Context, especialidad *Especialidad) error
	Update(ctx context.Context, id uint, especialidad Especialidad) error
	Delete(ctx context.Context, id uint) error
}
