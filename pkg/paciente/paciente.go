package paciente

import "time"

type Paciente struct {
	ID        string    `json:"id_p,omitempty"`
	Nombre    string    `json:"nombre,omitempty"`
	Apellido  string    `json:"apellido,omitempty"`
	Fecha_nac time.Time `json:"fecha_nac,omitempty"`
	Direccion string    `json:"direccion,omitempty"`
	Telefono  string    `json:"telefono,omitempty"`
}
