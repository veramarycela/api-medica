package data

import (
	"context"

	"github.com/veramarycela/api-medica/pkg/medico"
)

type MedicoRepository struct {
	Data *Data
}

func (ur *MedicoRepository) GetAll(ctx context.Context) ([]medico.Medico, error) {
	q := `
    SELECT id_m, nombre, apellido, direccion, telefono
        FROM medico;
    `

	rows, err := ur.Data.DB.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var medicos []medico.Medico
	for rows.Next() {
		var m medico.Medico
		rows.Scan(&m.ID, &m.Nombre, &m.Apellido,
			&m.Direccion, &m.Telefono)
		medicos = append(medicos, m)
	}

	return medicos, nil
}

func (ur *MedicoRepository) GetOne(ctx context.Context, id uint) (medico.Medico, error) {
	q := `
    SELECT id_m, nombre, apellido, direccion, telefono
        FROM medico
        WHERE id_m = $1;
    `

	row := ur.Data.DB.QueryRowContext(ctx, q, id)

	var m medico.Medico
	err := row.Scan(&m.ID, &m.Nombre, &m.Apellido,
		&m.Direccion, &m.Telefono)
	if err != nil {
		return medico.Medico{}, err
	}

	return m, nil
}

func (ur *MedicoRepository) GetByMediconombre(ctx context.Context, nombre string) (medico.Medico, error) {
	q := `
	SELECT id_m, nombre, apellido, direccion, telefono
		FROM medico
        WHERE nombre = $1;
    `

	row := ur.Data.DB.QueryRowContext(ctx, q, nombre)

	var m medico.Medico
	err := row.Scan(&m.ID, &m.Nombre, &m.Apellido,
		&m.Direccion, &m.Telefono)
	if err != nil {
		return medico.Medico{}, err
	}

	return m, nil
}

func (ur *MedicoRepository) Create(ctx context.Context, m *medico.Medico) error {
	q := `
    INSERT INTO medico (nombre, apellido, direccion, telefono)
        VALUES ($1, $2, $3, $4)
        RETURNING id_m;
    `

	row := ur.Data.DB.QueryRowContext(
		ctx, q, m.Nombre, m.Apellido, m.Direccion,
		m.Telefono,
	)

	err := row.Scan(&m.ID)
	if err != nil {
		return err
	}

	return nil
}

func (ur *MedicoRepository) Update(ctx context.Context, id uint, m medico.Medico) error {
	q := `
    UPDATE medico set nombre=$1, apellido=$2, direccion=$4, telefono=$5
        WHERE id_m=$6;
    `

	stmt, err := ur.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx, m.Nombre, m.Apellido,
		m.Direccion, m.Telefono, id,
	)
	if err != nil {
		return err
	}

	return nil
}

func (ur *MedicoRepository) Delete(ctx context.Context, id uint) error {
	q := `DELETE FROM medico WHERE id_m=$1;`

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
