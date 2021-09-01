// pkg/user/repository.go
package paciente

import "context"

// Repository handle the CRUD operations with Users.
type Repository interface {
	GetAll(ctx context.Context) ([]Paciente, error)
	GetOne(ctx context.Context, id string) (Paciente, error)
	GetByPacientenombre(ctx context.Context, nombre string) (Paciente, error)
	Create(ctx context.Context, paciente *Paciente) error
	Update(ctx context.Context, id string, paciente Paciente) error
	Delete(ctx context.Context, id string) error
}
