package empleados

import (
	User "github.com/frank1995alfredo/api/models/usuarios"
)

//Empleado ...
type Empleado struct {
	EmpleadoID  uint      `json:"idempleado" gorm:"primary_key"`
	DiscID      uint      `json:"iddiscapacidad" gorm:"not null"`
	CiuID       uint      `json:"idciudad" gorm:"not null"`
	CargoEmpID  uint      `json:"idcargoemp" gorm:"not null"`
	PriNombre   string    `json:"prinombre" gorm:"not null"`
	SegNombre   string    `json:"segnombre" gorm:"not null"`
	PriApellido string    `json:"priapellido" gorm:"not null"`
	SegApellido string    `json:"segapellido" gorm:"not null"`
	FechNac     string    `json:"fechnac"`
	NumCedula   string    `json:"numcedula" gorm:"not null"`
	CodigoEmp   string    `json:"codigoemp" gorm:"not null"`
	Direccion   string    `json:"direccion"`
	Email       string    `json:"email"`
	Telefono    string    `json:"telefono"`
	Genero      string    `json:"genero"`
	Estado      bool      `json:"estado"`
	Foto        *string   `json:"foto"`
	NivelDis    string    `json:"niveldis"`
	Usuario     User.User `json:"usuario" gorm:"foreignkey:EmpID"`
}
