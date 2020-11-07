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
	Estado      bool   `json:"estado"`
	NivelDis    string `json:"niveldis"`
}
