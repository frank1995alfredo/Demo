package mantenimiento

//Ciudad ...
type Ciudad struct {
	CiudadID    uint       `json:"idciudad" gorm:"primary_key"`
	ProID       uint       `json:"idprovincia" gorm:"not null"`
	Descripcion string     `json:"descripcion" gorm:"size:100;not null"`
	Estado      bool       `json:"estado"`
	Clientes    []Cliente  `json:"clientes"  gorm:"foreignkey:CiuID"`
	Empleados   []Empleado `json:"empleados" gorm:"foreignkey:CiuID"`
}
