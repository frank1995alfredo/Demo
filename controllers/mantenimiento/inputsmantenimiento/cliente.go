package inputsmantenimiento

//ClienteInput ...
type ClienteInput struct {
	DiscID      uint   `json:"iddiscapacidad"`
	CiuID       uint   `json:"idciudad"`
	PriNombre   string `json:"primnombre"`
	SegNombre   string `json:"segnombre"`
	PriApellido string `json:"priapellido"`
	SegApellido string `json:"segapellido"`
	FechaNac    string `json:"fechanac"`
	NumCedula   string `json:"numcedula"`
	CodigoCli   string `json:"codigocli"`
	Direccion   string `json:"direccion"`
	Email       string `json:"email"`
	Telefono    string `json:"telefono"`
	Genero      string `json:"genero"`
	Estado      bool   `json:"estado"`
	NivelDis    string `json:"niveldis"`
}

//ValidarEntrada ...
func (cliente *ClienteInput) ValidarEntrada() bool {
	if cliente.CiuID == 0 || cliente.DiscID == 0 ||
		cliente.PriNombre == "" || cliente.SegNombre == "" ||
		cliente.PriApellido == "" ||
		cliente.SegApellido == "" || cliente.NumCedula == "" ||
		cliente.CodigoCli == "" || cliente.Direccion == "" {
		return true
	}
	return false
}
