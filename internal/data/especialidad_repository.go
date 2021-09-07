package data

import (
	"context"

	"github.com/veramarycela/api-medica/pkg/especialidad"
)

type EspecialidadRepository struct {
	Data *Data
}

func (ur *EspecialidadRepository) GetAll(ctx context.Context) ([]especialidad.Especialidad, error) {
	q := `
    SELECT id_e, nombre
        FROM especialidad;
    `

	rows, err := ur.Data.DB.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var especialidads []especialidad.Especialidad
	for rows.Next() {
		var e especialidad.Especialidad
		rows.Scan(&e.ID, &e.Nombre)
		especialidads = append(especialidads, e)
	}

	return especialidads, nil
}

func (ur *EspecialidadRepository) GetOne(ctx context.Context, id string) (especialidad.Especialidad, error) {
	q := `
    SELECT id_e, nombre
        FROM especialidad
        WHERE id_e = $1;
    `

	row := ur.Data.DB.QueryRowContext(ctx, q, id)

	var e especialidad.Especialidad
	err := row.Scan(&e.ID, &e.Nombre)
	if err != nil {
		return especialidad.Especialidad{}, err
	}

	return e, nil
}

func (ur *EspecialidadRepository) GetByEspecialidadnombre(ctx context.Context, nombre string) (especialidad.Especialidad, error) {
	q := `
	SELECT id_e, nombre
		FROM especialidad
        WHERE nombre = $1;
    `

	row := ur.Data.DB.QueryRowContext(ctx, q, nombre)

	var e especialidad.Especialidad
	err := row.Scan(&e.ID, &e.Nombre)
	if err != nil {
		return especialidad.Especialidad{}, err
	}

	return e, nil
}

func (ur *EspecialidadRepository) Create(ctx context.Context, e *especialidad.Especialidad) error {
	q := `
    INSERT INTO especialidad (nombre)
        VALUES ($1)
        RETURNING id_e;
    `

	row := ur.Data.DB.QueryRowContext(
		ctx, q, e.Nombre,
	)

	err := row.Scan(&e.ID)
	if err != nil {
		return err
	}

	return nil
}

func (ur *EspecialidadRepository) Update(ctx context.Context, id string, e especialidad.Especialidad) error {
	q := `
    UPDATE especialidad set nombre=$1
        WHERE id_e=$2;
    `

	stmt, err := ur.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx, e.Nombre,
	)
	if err != nil {
		return err
	}

	return nil
}

func (ur *EspecialidadRepository) Delete(ctx context.Context, id string) error {
	q := `DELETE FROM especialidad WHERE id_e=$1;`

	stmt, err := ur.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
