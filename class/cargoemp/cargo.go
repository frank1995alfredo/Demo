package cargoemp

import (
	"github.com/biezhi/gorm-paginator/pagination"
	empleado "github.com/frank1995alfredo/api/class/empleado"
	database "github.com/frank1995alfredo/api/database"
	cargos "github.com/frank1995alfredo/api/models/empleados"
)

//CargoEmp ...
type CargoEmp struct {
	cargoEmpID  uint
	descripcion string
	empleados   []empleado.Empleado
}

//ObtenerCargo ...
func (cargo *CargoEmp) ObtenerCargo(page int, limit int) *pagination.Paginator {

	var carg []cargos.CargoEmp

	db := database.DB.Find(&carg)

	paginator := pagination.Paging(&pagination.Param{
		DB:      db,
		Page:    page,
		Limit:   limit,
		OrderBy: []string{"cargo_emp_id asc"},
		ShowSQL: false,
	}, &carg)

	return paginator

}

//CargoConstructor ...
func (cargo *CargoEmp) CargoConstructor(cargoEmID uint, descripcion string) {
	cargo.cargoEmpID = cargoEmID
	cargo.descripcion = descripcion
}

//GetID ...
func (cargo *CargoEmp) GetID() uint {
	return cargo.cargoEmpID
}

//SetID ...
func (cargo *CargoEmp) SetID(id uint) {
	cargo.cargoEmpID = id
}

//GetDescripcion ...
func (cargo *CargoEmp) GetDescripcion() string {
	return cargo.descripcion
}

//SetDescripcion ...
func (cargo *CargoEmp) SetDescripcion(descripcion string) {
	cargo.descripcion = descripcion
}
