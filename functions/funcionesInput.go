package functions

//InputGlobal ... estructura global para validar los campos tipo json
type InputGlobal struct {
	DiscID      uint   `json:"iddiscapacidad"`
	CiuID       uint   `json:"idciudad"`
	CargoEmpID  uint   `json:"idcargoemp"`
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
	Foto        string `json:"foto"`
	CodigoEmp   string `json:"codigoemp"`

	Descripcion string `json:"descripcion"`
	ProID       uint   `json:"idprovincia"`
}
