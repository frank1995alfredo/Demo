package mantenimiento

//Cliente ...
type Cliente struct {
	ClienteID   uint   `json:"idcliente" gorm:"primary_key"`
	DiscID      uint   `json:"iddiscapacidad"`
	CiuID       uint   `json:"idciudad" gorm:"not null"`
	PriNombre   string `json:"primnombre" gorm:"not null"`
	SegNombre   string `json:"segnombre" gorm:"not null"`
	PriApellido string `json:"priapellido" gorm:"not null"`
	SegApellido string `json:"segapellido" gorm:"not null"`
	FechaNac    string `json:"fechanac"`
	NumCedula   string `json:"numcedula" gorm:"not null"`
	CodigoCli   string `json:"codigocli" gorm:"not null"`
	Direccion   string `json:"direccion" gorm:"not null"`
	Email       string `json:"email"`
	Telefono    string `json:"telefono"`
	Genero      string `json:"genero"`
	NivelDis    string `json:"niveldis"`
}

//Empleado ...
type Empleado struct {
	EmpleadoID  uint    `json:"idempleado" gorm:"primary_key"`
	DiscID      uint    `json:"iddiscapacidad" gorm:"not null"`
	CiuID       uint    `json:"idciudad" gorm:"not null"`
	CargoEmpID  uint    `json:"idcargoemp" gorm:"not null"`
	PriNombre   string  `json:"prinombre" gorm:"not null"`
	SegNombre   string  `json:"segnombre" gorm:"not null"`
	PriApellido string  `json:"priapellido" gorm:"not null"`
	SegApellido string  `json:"segapellido" gorm:"not null"`
	FechNac     string  `json:"fechnac"`
	NumCedula   string  `json:"numcedula" gorm:"not null"`
	CodigoEmp   string  `json:"codigoemp" gorm:"not null"`
	Direccion   string  `json:"direccion"`
	Email       string  `json:"email"`
	Telefono    string  `json:"telefono"`
	Genero      string  `json:"genero"`
	Estado      bool    `json:"estado"`
	Foto        *string `json:"foto"`
	NivelDis    string  `json:"niveldis"`
}

//Ciudad ...
type Ciudad struct {
	CiudadID    uint       `json:"idciudad" gorm:"primary_key"`
	ProID       uint       `json:"idprovincia" gorm:"not null"`
	Descripcion string     `json:"descripcion" gorm:"size:100;not null"`
	Estado      bool       `json:"estado"`
	Clientes    []Cliente  `json:"clientes"  gorm:"foreignkey:CiuID"`
	Empleados   []Empleado `json:"empleados" gorm:"foreignkey:CiuID"`
}

//Provincia ...
type Provincia struct {
	ProvinciaID uint     `json:"idprovincia" gorm:"primary_key"`
	Descripcion string   `json:"descripcion" gorm:"size:100" gorm:"not null;unique"`
	Estado      bool     `json:"estado"`
	Ciudades    []Ciudad `json:"ciudades" gorm:"foreignkey:ProID"`
}

//Discapacidad ...
type Discapacidad struct {
	DiscapacidadID uint       `json:"iddiscapacidad" gorm:"primary_key"`
	Descripcion    string     `json:"descripcion" gorm:"not null"`
	Clientes       []Cliente  `json:"clientes" gorm:"foreignkey:DiscID"`
	Empleados      []Empleado `json:"empleados" gorm:"foreignkey:DiscID"`
}

//CargoEmp ...
type CargoEmp struct {
	CargoEmpID  uint       `json:"idcargo" gorm:"primary_key"`
	Descripcion string     `json:"descripcion" gorm:"not null;unique"`
	Empleados   []Empleado `json:"empleados" gorm:"foreignkey:CargoEmpID"`
}
