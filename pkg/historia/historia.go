package historia

import "time"

type Historia struct {
	ID          string    `json:"id_h,omitempty"`
	Id_Paciente string    `json:"id_paciente,omitempty"`
	Id_Medico   int       `json:"id_medico,omitempty"`
	Fecha       time.Time `json:"fecha,omitempty"`
	Motivo      string    `json:"motivo,omitempty"`
	Diagnostico string    `json:"diagnostico,omitempty"`
	Receta      string    `json:"receta,omitempty"`
}
