package empleado

import (
	"github.com/biezhi/gorm-paginator/pagination"
	persona "github.com/frank1995alfredo/api/class/persona"
	database "github.com/frank1995alfredo/api/database"
	"github.com/frank1995alfredo/api/models/empleados"
)

//Empleado ... hace una composicion de la estructura persona
type Empleado struct {
	persona.Persona
	empleadoID string
	cargoEmpID uint
	codigoEmp  string
	foto       string
}

//ObtenerEmpleados ...
func (empleado *Empleado) ObtenerEmpleados(page int, limit int) *pagination.Paginator {

	var emple []empleados.Empleado

	db := database.DB.Where("estado=?", true).Find(&emple)

	paginator := pagination.Paging(&pagination.Param{
		DB:      db,
		Page:    page,
		Limit:   limit,
		OrderBy: []string{"empleado_id asc"},
		ShowSQL: false,
	}, &emple)

	return paginator

}

//CrearEmpleado ...
func (empleado *Empleado) CrearEmpleado(discID uint, ciuID uint,
	priNombre string, segNombre string, priApellido string,
	segApellido string, fechaNac string, numCedula string,
	direccion string, email string, telefono string,
	genero string, estado bool, nivelDis string, cargoEmpID uint,
	codigoEmp string, foto string) string {

	var emple empleados.Empleado

	empleado.PersonaConstructor("", discID, ciuID, priNombre, segNombre,
		priApellido, segApellido, fechaNac, numCedula, direccion, email, telefono,
		genero, estado, nivelDis)

	empleado.EmpleadoConstructor("", cargoEmpID, codigoEmp, foto)

	empleado.SetDiscID(discID)
	empleado.SetCiuID(ciuID)
	empleado.SetPriNombre(priNombre)
	empleado.SetSegNombre(segNombre)
	empleado.SetPriApellido(priApellido)
	empleado.SetSegApellido(segApellido)
	empleado.SetFechaNac(fechaNac)
	empleado.SetNumCedula(numCedula)
	empleado.SetDireccion(direccion)
	empleado.SetEmail(email)
	empleado.SetTelefono(telefono)
	empleado.SetGenero(genero)
	empleado.SetEstado(estado)
	empleado.SetNivelDisacapacidad(nivelDis)
	empleado.SetCargoEmpID(cargoEmpID)
	empleado.SetCodigoEmp(codigoEmp)
	empleado.SetFoto(foto)

	discapacidadID := empleado.GetDiscID()
	ciudadID := empleado.GetCiuID()
	prinombre := empleado.GetPriNombre()
	segnombre := empleado.GetSegNombre()
	priapellido := empleado.GetPriApellido()
	segapellido := empleado.GetSegApellido()
	fechanac := empleado.GetFechaNac()
	numcedula := empleado.GetNumCedula()
	direcc := empleado.GetDireccion()
	emai := empleado.GetEmail()
	telef := empleado.GetTelefono()
	gene := empleado.GetGenero()
	niveldis := empleado.GetNivelDisacapacidad()
	cargo := empleado.GetCargoEmpID()
	codigoemp := empleado.GetCodigoEmp()
	fot := empleado.GetFoto()

	if empleado.validarEntrada(ciudadID, discapacidadID, cargo, prinombre, segnombre, priapellido,
		segapellido, numcedula, codigoemp) {
		return "Por favor ingrese los datos que son obligatorios."
	}

	if err := database.DB.Where("num_cedula=?", numCedula).First(&emple).Error; err == nil {
		return "Ya existe este empleado con el mismo número de cédula, ingrese otro."
	}

	datos := empleados.Empleado{DiscID: discapacidadID, CiuID: ciudadID, CargoEmpID: cargo,
		PriNombre: prinombre, SegNombre: segnombre, PriApellido: priapellido, SegApellido: segapellido,
		FechNac: fechanac, NumCedula: numcedula, CodigoEmp: codigoemp, Direccion: direcc,
		Email: emai, Telefono: telef, Genero: gene, Estado: true, Foto: fot, NivelDis: niveldis}

	//inicio de la transaccion
	tx := database.DB.Begin()
	err := tx.Create(&datos).Error
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
	//fin de la transaccion

	return "Datos ingresados correctamente."

}

//BuscarEmpleado ...
func (empleado *Empleado) BuscarEmpleado(valor string) ([]empleados.Empleado, string) {

	var emp []empleados.Empleado

	empleado.PersonaConstructor("", 0, 0, "", "", valor, valor, "", valor, "", "", "", "", false, "")

	empleado.SetNumCedula(valor)
	empleado.SetPriApellido(valor)
	empleado.SetSegApellido(valor)

	cedula := empleado.GetNumCedula()
	priApellido := empleado.GetPriApellido()
	segApellido := empleado.GetSegNombre()

	if err := database.DB.Where("num_cedula=?", cedula).
		Or("pri_apellido=?", priApellido).
		Or("seg_apellido=?", segApellido).
		First(&emp).Error; err != nil {
		return nil, "No existe este empleado."
	}

	return emp, ""

}

//ActualizarEmpleado ...
func (empleado *Empleado) ActualizarEmpleado(ID string, discID uint, ciuID uint,
	priNombre string, segNombre string, priApellido string,
	segApellido string, fechaNac string, numCedula string,
	direccion string, email string, telefono string,
	genero string, estado bool, nivelDis string, cargoEmp uint,
	codigoEmp string, foto string) string {

	var emple empleados.Empleado

	empleado.PersonaConstructor("", discID, ciuID, priNombre, segNombre,
		priApellido, segApellido, fechaNac, numCedula, direccion, email, telefono,
		genero, estado, nivelDis)
	empleado.EmpleadoConstructor("", cargoEmp, codigoEmp, foto)

	empleado.SetDiscID(discID)
	empleado.SetCiuID(ciuID)
	empleado.SetPriNombre(priNombre)
	empleado.SetSegNombre(segNombre)
	empleado.SetPriApellido(priApellido)
	empleado.SetSegApellido(segApellido)
	empleado.SetFechaNac(fechaNac)
	empleado.SetNumCedula(numCedula)
	empleado.SetDireccion(direccion)
	empleado.SetEmail(email)
	empleado.SetTelefono(telefono)
	empleado.SetGenero(genero)
	empleado.SetEstado(estado)
	empleado.SetNivelDisacapacidad(nivelDis)
	empleado.SetCargoEmpID(cargoEmp)
	empleado.SetCodigoEmp(codigoEmp)
	empleado.SetFoto(foto)

	discapacidadID := empleado.GetDiscID()
	ciudadID := empleado.GetCiuID()
	prinombre := empleado.GetPriNombre()
	segnombre := empleado.GetSegNombre()
	priapellido := empleado.GetPriApellido()
	segapellido := empleado.GetSegApellido()
	fechanac := empleado.GetFechaNac()
	numcedula := empleado.GetNumCedula()
	direcc := empleado.GetDireccion()
	emai := empleado.GetEmail()
	telef := empleado.GetTelefono()
	gene := empleado.GetGenero()
	estad := empleado.GetEstado()
	niveldis := empleado.GetNivelDisacapacidad()
	cargo := empleado.GetCargoEmpID()
	codigoemp := empleado.GetCodigoEmp()
	fot := empleado.GetFoto()

	if empleado.validarEntrada(ciudadID, discapacidadID, cargo, prinombre, segnombre, priapellido,
		segapellido, numcedula, codigoemp) {
		return "Por favor ingrese los datos que son obligatorios."
	}

	datos := empleados.Empleado{DiscID: discapacidadID, CiuID: ciudadID, CargoEmpID: cargo,
		PriNombre: prinombre, SegNombre: segnombre, PriApellido: priapellido, SegApellido: segapellido,
		FechNac: fechanac, NumCedula: numcedula, CodigoEmp: codigoemp, Direccion: direcc,
		Email: emai, Telefono: telef, Genero: gene, Estado: estad, Foto: fot, NivelDis: niveldis}

	//inicio de la transaccion
	tx := database.DB.Begin()
	err := tx.Model(&emple).Where("empleado_id=?", ID).Update(datos).Error
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
	//fin de la transaccion

	return "Registro modificado correctamente."

}

//EliminarEmpleado ...
func (empleado *Empleado) EliminarEmpleado(ID string) string {
	var emp empleados.Empleado

	empleado.EmpleadoConstructor(ID, 0, "", "")

	empleado.SetEmpID(ID)
	id := empleado.GetEmpID()

	//inicio de la transaccion
	tx := database.DB.Begin()
	err := tx.Model(&emp).Where("empleado_id=?", id).Update("estado", false).Error
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
	//fin de la transaccion

	return "Registro eliminado correctamente."
}

//EmpleadoConstructor ... constructor de la estructura empleado
func (empleado *Empleado) EmpleadoConstructor(empleadoID string, cargoEmpID uint, codigoEmp string, foto string) {
	empleado.cargoEmpID = cargoEmpID
	empleado.codigoEmp = codigoEmp
	empleado.foto = foto
}

//GetEmpID ...
func (empleado *Empleado) GetEmpID() string {
	return empleado.empleadoID
}

//SetEmpID ...
func (empleado *Empleado) SetEmpID(id string) {
	empleado.empleadoID = id
}

//GetCargoEmpID ...
func (empleado *Empleado) GetCargoEmpID() uint {
	return empleado.cargoEmpID
}

//SetCargoEmpID ...
func (empleado *Empleado) SetCargoEmpID(cargo uint) {
	empleado.cargoEmpID = cargo
}

//GetCodigoEmp ...
func (empleado *Empleado) GetCodigoEmp() string {
	return empleado.codigoEmp
}

//SetCodigoEmp ...
func (empleado *Empleado) SetCodigoEmp(codigo string) {
	empleado.codigoEmp = codigo
}

//GetFoto ...
func (empleado *Empleado) GetFoto() string {
	return empleado.foto
}

//SetFoto ...
func (empleado *Empleado) SetFoto(foto string) {
	empleado.foto = foto
}

func (empleado *Empleado) validarEntrada(ciudadID uint, discapacidadID,
	cargoEmpID uint, priNombre string, segNombre string, priApellido string,
	segApellido string, numCedula string, codigoEmp string) bool {

	if ciudadID == 0 || discapacidadID == 0 || cargoEmpID == 0 || priNombre == "" ||
		segNombre == "" || priApellido == "" || segApellido == "" || numCedula == "" ||
		codigoEmp == "" {
		return true
	}

	return false

}
