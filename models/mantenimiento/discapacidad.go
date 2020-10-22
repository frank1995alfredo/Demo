package mantenimiento

//Discapacidad ...
type Discapacidad struct {
	DiscapacidadID uint       `json:"iddiscapacidad" gorm:"primary_key"`
	Descripcion    string     `json:"descripcion" gorm:"not null"`
	Clientes       []Cliente  `json:"clientes" gorm:"foreignkey:DiscID"`
	Empleados      []Empleado `json:"empleados" gorm:"foreignkey:DiscID"`
}
