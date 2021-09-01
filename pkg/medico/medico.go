package medico

type Medico struct {
	ID              int    `json:"id_m,omitempty"`
	Nombre          string `json:"nombre,omitempty"`
	Apellido        string `json:"apellido,omitempty"`
	Direccion       string `json:"direccion,omitempty"`
	Telefono        string `json:"telefono,omitempty"`
	Id_especialidad string `json:"id_especialidad,omitempty"`
}
