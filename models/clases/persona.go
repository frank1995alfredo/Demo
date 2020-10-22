package clases

//PersonaM ...
type PersonaM struct {
	DiscID      uint   `json:"iddiscapacidad"`
	CiuID       uint   `json:"idciudad"`
	PriNombre   string `json:"primnombrepapu"`
	SegNombre   string `json:"segnombre"`
	PriApellido string `json:"priapellido"`
	SegApellido string `json:"segapellido"`
	FechaNac    string `json:"fechanac"`
	NumCedula   string `json:"numcedula"`
	Direccion   string `json:"direccion"`
	Email       string `json:"email"`
	Telefono    string `json:"telefono"`
	Genero      string `json:"genero"`
	NivelDis    string `json:"niveldis"`
}

//GetNombres ...
func (p PersonaM) GetNombres() string {
	return p.PriNombre + " " + p.SegNombre
}

//GetApellidos ...
func (p PersonaM) GetApellidos() string {
	return p.PriApellido + " " + p.SegApellido
}
