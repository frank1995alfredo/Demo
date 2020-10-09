package inputs

//ProvinciaInput ...
type ProvinciaInput struct {
	Descripcion string `json:"descripcion"`
	Estado      bool   `json:"estado"`
}

//CiudadInput ...
type CiudadInput struct {
	ProID       uint   `json:"idprovincia"`
	Descripcion string `json:"descripcion"`
	Estado      bool   `json:"estado"`
}

//DiscapacidadInput ...
type DiscapacidadInput struct {
	Descripcion string `json:"descripcion"`
}

//CargoInput ...
type CargoInput struct {
	Descripcion string `json:"descripcion"`
}

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
	NivelDis    string `json:"niveldis"`
}
