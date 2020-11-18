package mantenimiento

import (
	Cliente "github.com/frank1995alfredo/api/models/clientes"
	Empleado "github.com/frank1995alfredo/api/models/empleados"
)

//Ciudad ...
type Ciudad struct {
	CiudadID    uint64              `json:"idciudad" gorm:"primary_key"`
	ProID       uint                `json:"idprovincia" gorm:"not null"`
	Descripcion string              `json:"descripcion" gorm:"size:100;not null"`
	Estado      bool                `json:"estado"`
	Clientes    []Cliente.Cliente   `json:"clientes"  gorm:"foreignkey:CiuID"`
	Empleados   []Empleado.Empleado `json:"empleados" gorm:"foreignkey:CiuID"`
}
