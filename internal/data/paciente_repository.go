package data

import (
	"context"

	"github.com/veramarycela/api-medica/pkg/paciente"
)

type PacienteRepository struct {
	Data *Data
}

func (ur *PacienteRepository) GetAll(ctx context.Context) ([]paciente.Paciente, error) {
	q := `
    SELECT id_p, nombre, apellido, fecha_nac, direccion, telefono
        FROM paciente;
    `

	rows, err := ur.Data.DB.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var pacientes []paciente.Paciente
	for rows.Next() {
		var p paciente.Paciente
		rows.Scan(&p.ID, &p.Nombre, &p.Apellido, &p.Fecha_nac,
			&p.Direccion, &p.Telefono)
		pacientes = append(pacientes, p)
	}

	return pacientes, nil
}

func (ur *PacienteRepository) GetOne(ctx context.Context, id string) (paciente.Paciente, error) {
	q := `
    SELECT id_p, nombre, apellido, fecha_nac, direccion, telefono
        FROM paciente
        WHERE id_p = $1;
    `

	row := ur.Data.DB.QueryRowContext(ctx, q, id)

	var p paciente.Paciente
	err := row.Scan(&p.ID, &p.Nombre, &p.Apellido, &p.Fecha_nac,
		&p.Direccion, &p.Telefono)
	if err != nil {
		return paciente.Paciente{}, err
	}

	return p, nil
}

func (ur *PacienteRepository) GetByPacientenombre(ctx context.Context, nombre string) (paciente.Paciente, error) {
	q := `
	SELECT id_p, nombre, apellido, fecha_nac, direccion, telefono
		FROM paciente
        WHERE nombre = $1;
    `

	row := ur.Data.DB.QueryRowContext(ctx, q, nombre)

	var p paciente.Paciente
	err := row.Scan(&p.ID, &p.Nombre, &p.Apellido, &p.Fecha_nac,
		&p.Direccion, &p.Telefono)
	if err != nil {
		return paciente.Paciente{}, err
	}

	return p, nil
}

func (ur *PacienteRepository) Create(ctx context.Context, p *paciente.Paciente) error {
	q := `
    INSERT INTO paciente (nombre, apellido, fecha_nac, direccion, telefono)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id_p;
    `

	row := ur.Data.DB.QueryRowContext(
		ctx, q, p.Nombre, p.Apellido, p.Fecha_nac, p.Direccion,
		p.Telefono,
	)

	err := row.Scan(&p.ID)
	if err != nil {
		return err
	}

	return nil
}

func (ur *PacienteRepository) Update(ctx context.Context, id string, p paciente.Paciente) error {
	q := `
    UPDATE paciente set nombre=$1, apellido=$2, fecha_nac=$3, direccion=$4, telefono=$5
        WHERE id_p=$6;
    `

	stmt, err := ur.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx, p.Nombre, p.Apellido,
		p.Fecha_nac, p.Direccion, p.Telefono, id,
	)
	if err != nil {
		return err
	}

	return nil
}

func (ur *PacienteRepository) Delete(ctx context.Context, id string) error {
	q := `DELETE FROM paciente WHERE id_p=$1;`

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
