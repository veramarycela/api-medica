package data

import (
	"context"

	"github.com/veramarycela/api-medica/pkg/historia"
)

type HistoriaRepository struct {
	Data *Data
}

func (ur *HistoriaRepository) GetAll(ctx context.Context) ([]historia.Historia, error) {
	q := `
    SELECT id_h, id_paciente, id_medico, fecha, motivo, diagnostico, receta
        FROM historia;
    `

	rows, err := ur.Data.DB.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var historias []historia.Historia
	for rows.Next() {
		var h historia.Historia
		rows.Scan(&h.ID, &h.Id_Paciente, &h.Id_Medico, &h.Fecha,
			&h.Motivo, &h.Diagnostico, &h.Receta)
		historias = append(historias, h)
	}

	return historias, nil
}

func (ur *HistoriaRepository) GetOne(ctx context.Context, id uint) (historia.Historia, error) {
	q := `
    SELECT id_h, id_paciente, id_medico, fecha, motivo, diagnostico, receta
        FROM historia
        WHERE id_h = $1;
    `

	row := ur.Data.DB.QueryRowContext(ctx, q, id)

	var h historia.Historia
	err := row.Scan(&h.ID, &h.Id_Paciente, &h.Id_Medico, &h.Fecha,
		&h.Motivo, &h.Diagnostico, &h.Receta)
	if err != nil {
		return historia.Historia{}, err
	}

	return h, nil
}

func (ur *HistoriaRepository) GetByHistorianame(ctx context.Context, nombre string) (historia.Historia, error) {
	q := `
	SELECT id_h, id_paciente, id_medico, fecha, motivo, diagnostico, receta
		FROM paciente
        WHERE nombre = $1;
    `

	row := ur.Data.DB.QueryRowContext(ctx, q, nombre)

	var h historia.Historia
	err := row.Scan(&h.ID, &h.Id_Paciente, &h.Id_Medico, &h.Fecha,
		&h.Motivo, &h.Diagnostico, &h.Receta)
	if err != nil {
		return historia.Historia{}, err
	}

	return h, nil
}

func (ur *HistoriaRepository) Create(ctx context.Context, h *historia.Historia) error {
	q := `
    INSERT INTO historia (id_paciente, id_medico, fecha, motivo, diagnostico, receta)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
        RETURNING id_h;
    `

	row := ur.Data.DB.QueryRowContext(
		ctx, q, h.ID, h.Id_Paciente, h.Id_Medico, h.Fecha,
		h.Motivo, h.Diagnostico, h.Receta,
	)

	err := row.Scan(&h.ID)
	if err != nil {
		return err
	}

	return nil
}

func (ur *HistoriaRepository) Update(ctx context.Context, id uint, h historia.Historia) error {
	q := `
    UPDATE historia set id_paciente=$1, id_medico=$2, fecha=$3, motivo=$4, diagnostico=$6, receta=$6
	WHERE id_h=$8;
    `

	stmt, err := ur.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx, h.ID, h.Id_Paciente, h.Id_Medico, h.Fecha,
		h.Motivo, h.Diagnostico, h.Receta, id,
	)
	if err != nil {
		return err
	}

	return nil
}

func (ur *HistoriaRepository) Delete(ctx context.Context, id uint) error {
	q := `DELETE FROM historia WHERE id_h=$1;`

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
