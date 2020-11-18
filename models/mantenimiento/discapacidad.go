package mantenimiento

import (
	Cliente "github.com/frank1995alfredo/api/models/clientes"
	Empleado "github.com/frank1995alfredo/api/models/empleados"
)

//Discapacidad ...
type Discapacidad struct {
	DiscapacidadID uint                `json:"iddiscapacidad" gorm:"primary_key"`
	Descripcion    string              `json:"descripcion" gorm:"not null"`
	Clientes       []Cliente.Cliente   `json:"clientes" gorm:"foreignkey:DiscID"`
	Empleados      []Empleado.Empleado `json:"empleados" gorm:"foreignkey:DiscID"`
}
