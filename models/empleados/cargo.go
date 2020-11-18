package empleados

//CargoEmp ...
type CargoEmp struct {
	CargoEmpID  uint       `json:"idcargo" gorm:"primary_key"`
	Descripcion string     `json:"descripcion" gorm:"not null;unique"`
	Empleados   []Empleado `json:"empleados" gorm:"foreignkey:CargoEmpID"`
}
