package inputsmantenimiento

//EmpleadoInput ...
type EmpleadoInput struct {
	DiscID      uint    `json:"iddiscapacidad"`
	CiuID       uint    `json:"idciudad"`
	CargoEmpID  uint    `json:"idcargoemp"`
	PriNombre   string  `json:"prinombre"`
	SegNombre   string  `json:"segnombre"`
	PriApellido string  `json:"priapellido"`
	SegApellido string  `json:"segapellido"`
	FechNac     string  `json:"fechnac"`
	NumCedula   string  `json:"numcedula"`
	CodigoEmp   string  `json:"codigoemp"`
	Direccion   string  `json:"direccion"`
	Email       string  `json:"email"`
	Telefono    string  `json:"telefono"`
	Genero      string  `json:"genero"`
	Estado      bool    `json:"estado"`
	Foto        *string `json:"foto"`
	NivelDis    string  `json:"niveldis"`
}

//ValidarEntrada ...
func (empleado *EmpleadoInput) ValidarEntrada() bool {
	if empleado.CiuID == 0 || empleado.DiscID == 0 ||
		empleado.CargoEmpID == 0 || empleado.PriNombre == "" ||
		empleado.SegNombre == "" || empleado.PriApellido == "" ||
		empleado.SegApellido == "" || empleado.NumCedula == "" ||
		empleado.CodigoEmp == "" {
		return true
	}
	return false
}
